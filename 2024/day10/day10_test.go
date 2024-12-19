package day10

import (
	"adventOfCode/main/utils"
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
)

func TestPart1(t *testing.T) {
	result := Part1()

	assert.Equal(t, result, 512)

}

func Part1() int {
	lines := utils.GetDataSplittedInLines("day10/day10")
	arr := utils.SplitStringIntoArray(lines)
	ROWS := len(arr)
	COLS := len(arr[0])
	count := 0
	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLS; j++ {
			if arr[i][j] == "0" {
				score := dfs(arr, i, j)
				count += score
			}
		}
	}

	return count
}

type Path struct {
	Row int
	Col int
}

func dfs(arr [][]string, row, col int) int {
	ROWS := len(arr)
	COLS := len(arr[0])
	directions := [][]int{
		{0, 1},
		{1, 0},
		{-1, 0},
		{0, -1},
	}

	stack := []Path{
		{Row: row, Col: col},
	}

	visited := make(map[string]bool)
	score := 0
	for len(stack) != 0 {
		stackLen := len(stack)
		path := stack[stackLen-1]
		stack = stack[:stackLen-1]

		key := fmt.Sprintf("%d,%d", path.Row, path.Col)
		if visited[key] {
			continue
		}
		visited[key] = true

		if arr[path.Row][path.Col] == "9" {
			score++
		}

		for _, direction := range directions {
			r := path.Row + direction[0]
			c := path.Col + direction[1]

			if r < 0 || r >= ROWS || c < 0 || c >= COLS {
				continue
			}

			number := utils.ParseStringToNum(arr[path.Row][path.Col])
			nextNumber := utils.ParseStringToNum(arr[r][c])

			if nextNumber != number+1 {
				continue
			}

			stack = append(stack, Path{Row: r, Col: c})
		}
	}

	return score
}
