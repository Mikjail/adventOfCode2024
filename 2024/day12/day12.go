package day12

import "adventOfCode/main/utils"

type Area struct {
	Row int
	Col int
}

func Part1() int {
	data := utils.GetDataSplittedInLines("day12/day12")
	arr := utils.ParseArrayStringIntoMatrix(data)
	ROWS := len(arr)
	COLS := len(arr[0])

	visited := make(map[Area]bool)
	result := 0
	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLS; j++ {
			_, exist := visited[Area{Row: i, Col: j}]
			if exist {
				continue
			}
			// check if the area is already visited
			areaCount, perimeter := newFunction(i, j, arr, visited)
			result += areaCount * perimeter
		}
	}

	return result
}

func newFunction(i int, j int, arr [][]string, visited map[Area]bool) (int, int) {
	// DFS through area until char changes
	directions := [][]int{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}
	ROWS := len(arr)
	COLS := len(arr[0])
	char := arr[i][j]
	areaCount := 1
	perimeter := 0
	stack := []Area{{
		Row: i, Col: j,
	}}
	for len(stack) > 0 {
		lenArea := len(stack) - 1
		area := stack[lenArea]
		stack = stack[:lenArea]
		visited[area] = true
		for _, dir := range directions {
			row := area.Row + dir[0]
			col := area.Col + dir[1]

			if row < 0 || row >= ROWS || col < 0 || col >= COLS || arr[row][col] != char {
				perimeter++
				continue
			}
			if !visited[Area{Row: row, Col: col}] {
				visited[Area{Row: row, Col: col}] = true
				areaCount++
				stack = append(stack, Area{Row: row, Col: col})
			}
		}
	}
	return areaCount, perimeter
}
