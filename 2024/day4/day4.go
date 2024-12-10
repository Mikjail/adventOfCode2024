package day4

import (
	"adventOfCode/main/utils"
)

func Part1() int {
	words := getAllLines()

	return searchWord(words)
}

func Part2() int {
	words := getAllLines()

	return searchWordMAS(words)
}

func searchWordMAS(lines []string) int {
	count := 0
	rows := len(lines)
	cols := len(lines[0])

	var isXMAS = func(lines []string, row, col int) bool {

		if (lines[row][col] == 'S' && lines[row][col+2] == 'S') && (lines[row+2][col] == 'M' && lines[row+2][col+2] == 'M') && lines[row+1][col+1] == 'A' {
			return true
		}
		if (lines[row][col] == 'M' && lines[row][col+2] == 'M') && (lines[row+2][col] == 'S' && lines[row+2][col+2] == 'S') && lines[row+1][col+1] == 'A' {
			return true
		}
		if (lines[row][col] == 'M' && lines[row][col+2] == 'S') && (lines[row+2][col] == 'M' && lines[row+2][col+2] == 'S') && lines[row+1][col+1] == 'A' {
			return true
		}
		if (lines[row][col] == 'S' && lines[row][col+2] == 'M') && (lines[row+2][col] == 'S' && lines[row+2][col+2] == 'M') && lines[row+1][col+1] == 'A' {
			return true
		}
		return false
	}

	for i := 0; i < rows-2; i++ {
		for j := 0; j < cols-2; j++ {
			if isXMAS(lines, i, j) {
				count++
			}
		}
	}

	return count
}

func getAllLines() (result []string) {
	lines := utils.GetDataSplittedInLines("day4/day4")

	result = append(result, lines...)
	return
}

func searchWord(lines []string) int {
	ROWS := len(lines)
	COLS := len(lines[0])
	keyword := "XMAS"

	isValidDirection := func(row int, col int, dr int, dc int) bool {
		for i := 0; i < len(keyword); i++ {
			r := row + i*dr
			c := col + i*dc
			if r >= ROWS || r < 0 || c >= COLS || c < 0 || lines[r][c] != keyword[i] {
				return false
			}
		}
		return true
	}
	result := 0

	for r := 0; r < ROWS; r++ {
		for c := 0; c < COLS; c++ {
			if lines[r][c] == keyword[0] {
				directions := [][]int{
					{0, 1},   // right
					{0, -1},  // left
					{1, 0},   // down
					{-1, 0},  // up
					{1, 1},   // down-right
					{1, -1},  // down-left
					{-1, 1},  // up-right
					{-1, -1}, // up-left
				}
				for _, d := range directions {
					if isValidDirection(r, c, d[0], d[1]) {
						result += 1
					}
				}
			}
		}
	}

	return result
}
