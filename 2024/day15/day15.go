package day15

import (
	"adventOfCode/main/utils"
	"strconv"
	"strings"
)

type Position struct {
	Row int
	Col int
}

type MapWithPositions map[Position]bool

func Part1() int {
	lines := utils.GetData("day15/day15")

	dataSplitted := strings.Split(lines, "\n\n")
	pathMap := strings.Split(dataSplitted[0], "\n")
	directionLines := strings.Split(dataSplitted[1], "\n")

	robotPosition, boxesPosition, obstaclePosition, directions := getBoxPositionsAndDirections(pathMap, directionLines)
	ROWS := len(pathMap)
	COLS := len(pathMap[0])
	boxPositions := getSumOFAllBoxesGPS(robotPosition, boxesPosition, obstaclePosition, directions, ROWS, COLS)
	return getGPSCoordinatesSum(boxPositions)
}

func getGPSCoordinatesSum(boxPositions MapWithPositions) (result int) {
	for position, exist := range boxPositions {
		if exist {
			result += (100 * position.Row) + position.Col
		}
	}
	return result
}

func getSumOFAllBoxesGPS(robotPosition Position, boxPositions MapWithPositions, obstaclePosition MapWithPositions, directions [][]string, ROWS, COLS int) MapWithPositions {
	directionMap := map[string][]int{
		"^": {-1, 0},
		">": {0, 1},
		"v": {1, 0},
		"<": {0, -1},
	}

	for _, direction := range directions {
		move := direction[0]
		moves := utils.ParseStringToNum(direction[1])
		row := robotPosition.Row
		col := robotPosition.Col
		for moves > 0 {
			dir := directionMap[move]
			nextRow := row + dir[0]
			nextCol := col + dir[1]

			if nextRow < 1 || nextRow > ROWS || nextCol < 1 || nextCol > COLS || obstaclePosition[Position{nextRow, nextCol}] {
				break
			}

			if boxPositions[Position{nextRow, nextCol}] {
				boxNextRow := nextRow
				boxNextCol := nextCol

				for boxPositions[Position{boxNextRow, boxNextCol}] {
					boxNextRow = boxNextRow + dir[0]
					boxNextCol = boxNextCol + dir[1]
					if boxNextRow < 1 || boxNextRow > ROWS || boxNextCol < 1 || boxNextCol > COLS || obstaclePosition[Position{boxNextRow, boxNextCol}] {
						break
					}
				}

				if obstaclePosition[Position{boxNextRow, boxNextCol}] {
					break
				}
				boxPositions[Position{nextRow, nextCol}] = false
				boxPositions[Position{boxNextRow, boxNextCol}] = true

				if boxNextRow == ROWS || boxNextCol == COLS {
					row = nextRow
					col = nextCol
					break
				}
			}
			row = nextRow
			col = nextCol
			moves--
		}
		robotPosition.Row = row
		robotPosition.Col = col
	}

	return boxPositions
}

func getBoxPositionsAndDirections(pathMap []string, directionLines []string) (Position, MapWithPositions, MapWithPositions, [][]string) {
	boxesPosition := map[Position]bool{}
	robotPosition := Position{}
	obstaclePosition := map[Position]bool{}
	for i := 0; i < len(pathMap); i++ {
		for j := 0; j < len(pathMap[0]); j++ {
			if pathMap[i][j] == '#' {
				obstaclePosition[Position{Row: i, Col: j}] = true
			}
			if pathMap[i][j] == '@' {
				robotPosition = Position{Row: i, Col: j}
			}
			if pathMap[i][j] == 'O' {
				boxesPosition[Position{Row: i, Col: j}] = true
			}
		}
	}

	directions := [][]string{}
	direction := string(directionLines[0][0])
	count := 0

	for _, line := range directionLines {
		for i := 0; i < len(line); i++ {
			if direction != string(line[i]) {
				directions = append(directions, []string{direction, strconv.Itoa(count)})
				direction = string(line[i])
				count = 0
			}
			count++
		}
	}

	if count > 0 {
		directions = append(directions, []string{direction, strconv.Itoa(count)})
	}

	return robotPosition, boxesPosition, obstaclePosition, directions
}
