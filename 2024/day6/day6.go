package day6

import (
	"adventOfCode/main/utils"
	"fmt"
)

type GuardDirections map[string]Direction

type Direction struct {
	RowChange    int
	ColChange    int
	NewDirection string
}

func Part1() int {
	path := getPath()
	guardDirections := GuardDirections{
		"^": {-1, 0, ">"},
		">": {0, 1, "V"},
		"V": {1, 0, "<"},
		"<": {0, -1, "^"},
	}

	ROWS := len(path)
	COLS := len(path[0])
	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLS; j++ {
			if path[i][j] == "^" {
				res, _ := dfs(path, guardDirections, i, j, ROWS, COLS)
				return res
			}
		}
	}

	return 0
}

func getPath() [][]string {
	data := utils.GetDataSplittedInLines("day6/day6")
	path := utils.SplitStringIntoArray(data)
	return path
}

func dfs(path [][]string, guardDirections GuardDirections, row int, col int, ROWS int, COLS int) (result int, isLoop bool) {
	dir := guardDirections[path[row][col]]
	result = 1
	isLoop = false
	visited := make(map[string]bool) // tracking visited path
	for {
		nextRow := row + dir.RowChange
		nextCol := col + dir.ColChange

		key := fmt.Sprintf("%d,%d,%s", nextRow, nextCol, dir.NewDirection)

		if visited[key] {
			isLoop = true
			return
		}

		visited[key] = true

		if nextRow < 0 || nextRow >= ROWS || nextCol < 0 || nextCol >= COLS {
			result++
			return
		}

		if path[nextRow][nextCol] == "#" {
			dir = guardDirections[dir.NewDirection]
			continue
		}

		if path[row][col] == "." {
			result++
		}

		row = nextRow
		col = nextCol
	}
}

func Part2() (result int) {
	path := getPath()
	guardDirections := GuardDirections{
		"^": {-1, 0, ">"},
		">": {0, 1, "V"},
		"V": {1, 0, "<"},
		"<": {0, -1, "^"},
	}
	ROWS := len(path)
	COLS := len(path[0])

	obstaclePostitions := getObstaclePositions(path)

	for _, obstacle := range obstaclePostitions {
		for i := 0; i < ROWS; i++ {
			for j := 0; j < COLS; j++ {
				path[obstacle.Row][obstacle.Col] = "#"
				if path[i][j] == "^" {
					_, isLoop := dfs(path, guardDirections, i, j, ROWS, COLS)
					if isLoop {
						result++
					}
				}
				path[obstacle.Row][obstacle.Col] = "."
			}
		}

	}

	return result
}

type ObstaclePositions struct {
	Row int
	Col int
}

func getObstaclePositions(path [][]string) (positions []ObstaclePositions) {
	ROWS := len(path)
	COLS := len(path[0])
	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLS; j++ {
			if path[i][j] == "." {
				positions = append(positions, ObstaclePositions{i, j})
			}
		}
	}
	return positions
}
