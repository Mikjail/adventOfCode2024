package day4

import (
	"adventOfCode/main/utils"
	"strings"
)

func GetAllLetters() (result []string) {
	data := utils.GetData("day4/example")

	lines := strings.Split(data, "\n")

	result = append(result, lines...)
	return
}

func SearchWord(lines []string) int {
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
