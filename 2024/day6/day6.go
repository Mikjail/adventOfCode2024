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

	_, initialPosition := getPositions(path)

	res, _ := dfs(path, guardDirections, initialPosition.Row, initialPosition.Col)

	return res
}

func getPath() [][]string {
	data := utils.GetDataSplittedInLines("day6/day6")
	path := utils.SplitStringIntoArray(data)
	return path
}

func dfs(path [][]string, guardDirections GuardDirections, row int, col int) (result int, isLoop bool) {
	ROWS := len(path)
	COLS := len(path[0])
	dir := guardDirections[path[row][col]]
	isLoop = false
	visitedDirection := make(map[string]bool)
	visitedStep := make(map[string]bool)
	for {
		nextRow := row + dir.RowChange
		nextCol := col + dir.ColChange

		Pathkey := fmt.Sprintf("%d,%d,%s", nextRow, nextCol, dir.NewDirection)
		stepKey := fmt.Sprintf("%d,%d", row, col)

		if !visitedStep[stepKey] {
			result++
		}

		visitedStep[stepKey] = true

		if visitedDirection[Pathkey] {
			isLoop = true
			return
		}

		visitedDirection[Pathkey] = true

		if nextRow < 0 || nextRow >= ROWS || nextCol < 0 || nextCol >= COLS {
			return
		}

		if path[nextRow][nextCol] == "#" {
			dir = guardDirections[dir.NewDirection]
			continue
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

	obstaclePostitions, initialPosition := getPositions(path)

	for _, obstacle := range obstaclePostitions {
		path[obstacle.Row][obstacle.Col] = "#"
		_, isLoop := dfs(path, guardDirections, initialPosition.Row, initialPosition.Col)
		if isLoop {
			result++
		}
		path[obstacle.Row][obstacle.Col] = "."
	}

	return result
}

type Position struct {
	Row int
	Col int
}

func getPositions(path [][]string) (obstaclePositions []Position, initialPosition Position) {
	ROWS := len(path)
	COLS := len(path[0])
	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLS; j++ {
			if path[i][j] == "." {
				obstaclePositions = append(obstaclePositions, Position{i, j})
			}
			if path[i][j] == "^" {
				initialPosition = Position{i, j}
			}
		}
	}
	return
}
