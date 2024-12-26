package day12

import (
	"adventOfCode/main/utils"
)

type Area struct {
	Row int
	Col int
}

type Direction struct {
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
			areaCount, perimeter := findAreaAndPerimeter(i, j, arr, visited)
			result += areaCount * perimeter
		}
	}

	return result
}

func findAreaAndPerimeter(i int, j int, arr [][]string, visited map[Area]bool) (int, int) {
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

func Part2() int {
	data := utils.GetDataSplittedInLines("day12/day12")
	grid := utils.ParseArrayStringIntoMatrix(data)

	output := 0

	directions := [][2]int{
		{-1, 0},
		{0, -1},
		{1, 0},
		{0, 1},
	}

	outerCorners := []int{
		0, 0, 0, 1, 0, 0, 1,
		2, 0, 1, 0, 2,
		1, 2, 2, 4,
	}

	checkInnerCorners := [][][]int{
		{{-1, -1}, {-1, 1}, {1, -1}, {1, 1}},
		{{1, -1}, {1, 1}},
		{{-1, 1}, {1, 1}},
		{{1, 1}},
		{{-1, -1}, {-1, 1}},
		{},
		{{-1, 1}},

		{},
		{{-1, -1}, {1, -1}},
		{{1, -1}},
		{},
		{},

		{{-1, -1}},
		{},
		{},
		{},
	}

	visited := make(map[[2]int]bool)

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if visited[[2]int{row, col}] {
				continue
			}
			plant := grid[row][col]
			area := 0
			corners := 0

			queue := [][2]int{{row, col}}
			for len(queue) > 0 {
				current := queue[0]
				queue = queue[1:]

				currentRow, currentCol := current[0], current[1]
				if visited[[2]int{currentRow, currentCol}] {
					continue
				}
				visited[[2]int{currentRow, currentCol}] = true

				area++
				cornerType := 0

				for i, dir := range directions {
					newRow := currentRow + dir[0]
					newCol := currentCol + dir[1]

					if newRow < 0 || newCol < 0 || newRow >= len(grid) || newCol >= len(grid[row]) {
						cornerType += (1 << i)
					} else if grid[newRow][newCol] != plant {
						cornerType += (1 << i)
					} else if !visited[[2]int{newRow, newCol}] {
						queue = append(queue, [2]int{newRow, newCol})
					}
				}

				outerCornerCount := outerCorners[cornerType]
				innerCornerCount := 0
				for _, corner := range checkInnerCorners[cornerType] {
					newRow := currentRow + corner[0]
					newCol := currentCol + corner[1]

					if newRow < 0 || newCol < 0 || newRow >= len(grid) || newCol >= len(grid[row]) {
						continue
					} else if grid[newRow][newCol] != plant {
						innerCornerCount++
					}
				}
				corners += outerCornerCount + innerCornerCount
			}

			price := area * corners
			output += price
		}
	}
	return output

}
