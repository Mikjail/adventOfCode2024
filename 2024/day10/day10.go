package day10

import (
	"adventOfCode/main/utils"
	"fmt"
)

func Part1() int {
	lines := utils.GetDataSplittedInLines("day10/day10")
	arr := utils.SplitStringIntoArray(lines)
	ROWS := len(arr)
	COLS := len(arr[0])
	count := 0
	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLS; j++ {
			if arr[i][j] == "0" {
				score, _ := dfs(arr, i, j)
				count += score
			}
		}
	}

	return count
}

func Part2() int {
	lines := utils.GetDataSplittedInLines("day10/day10")
	arr := utils.SplitStringIntoArray(lines)

	ROWS := len(arr)
	COLS := len(arr[0])
	count := 0

	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLS; j++ {
			if arr[i][j] == "0" {
				_, visited := dfs(arr, i, j)
				score := rating(arr, i, j, visited)
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
type VisitedPath map[string]Path

func dfs(arr [][]string, row, col int) (int, VisitedPath) {
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

	visited := make(VisitedPath)
	score := 0
	for len(stack) != 0 {
		stackLen := len(stack)
		path := stack[stackLen-1]
		stack = stack[:stackLen-1]

		key := fmt.Sprintf("%d,%d", path.Row, path.Col)
		_, exist := visited[key]
		if exist {
			continue
		}

		visited[key] = Path{Row: path.Row, Col: path.Col}

		if arr[path.Row][path.Col] == "9" {
			score++
		}

		for _, direction := range directions {
			r := path.Row + direction[0]
			c := path.Col + direction[1]

			if notValid(r, ROWS, c, COLS) {
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

	return score, visited
}

func rating(arr [][]string, row, col int, visited VisitedPath) int {
	ROWS := len(arr)
	COLS := len(arr[0])
	directions := [][]int{
		{0, 1},
		{1, 0},
		{-1, 0},
		{0, -1},
	}

	rating := make(map[Path]int)
	for i := 9; i > 0; i-- {
		for _, path := range visited {
			if arr[path.Row][path.Col] != fmt.Sprint(i) {
				continue
			}

			curr := utils.ParseStringToNum(arr[path.Row][path.Col])
			if curr == 9 {
				rating[Path{Row: path.Row, Col: path.Col}] = 1
			}

			for _, dir := range directions {
				row := path.Row + dir[0]
				col := path.Col + dir[1]

				if notValid(row, ROWS, col, COLS) {
					continue
				}
				number := utils.ParseStringToNum(arr[row][col])
				if number == curr-1 {
					rating[Path{Row: row, Col: col}] += rating[Path{Row: path.Row, Col: path.Col}]
				}
			}

		}
	}

	return rating[Path{Row: row, Col: col}]
}

func notValid(r int, ROWS int, c int, COLS int) bool {
	return r < 0 || r >= ROWS || c < 0 || c >= COLS
}
