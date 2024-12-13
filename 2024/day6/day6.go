package day6

import "adventOfCode/main/utils"

type Direction struct {
	RowChange    int
	ColChange    int
	NewDirection string
}

func Part1() int {
	data := utils.GetDataSplittedInLines("day6/day6")
	path := utils.SplitStringIntoArray(data)
	guardDirections := map[string]Direction{
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
				return dfs(path, guardDirections, i, j, ROWS, COLS)
			}
		}
	}

	return 0
}

func dfs(path [][]string, guardDirections map[string]Direction, row int, col int, ROWS int, COLS int) (result int) {
	dir := guardDirections[path[row][col]]
	result = 1
	for {
		nextRow := row + dir.RowChange
		nextCol := col + dir.ColChange
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
		path[row][col] = "X"
		row = nextRow
		col = nextCol
	}

}
