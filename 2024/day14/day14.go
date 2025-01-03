package day14

import (
	"adventOfCode/main/utils"
	"strings"
)

type Direction struct {
	Col int
	Row int
}

type Robot struct {
	Position Direction
	Steps    Direction
}

type Boundaries struct {
	COLS int
	ROWS int
}

type Quadrants struct {
	TopLeft     int
	TopRight    int
	BottomLeft  int
	BottomRight int
}

func Part1() int {
	lines := utils.GetDataSplittedInLines("day14/day14")

	robots := getRobots(lines)

	boundaries := Boundaries{COLS: 101, ROWS: 103}

	quadrants := getQuadrants(robots, boundaries)

	result := getFactor(quadrants)

	return result
}

func getFactor(quadrants *Quadrants) int {
	return quadrants.TopLeft * quadrants.TopRight * quadrants.BottomRight * quadrants.BottomLeft
}

func getQuadrants(robots []Robot, boundaries Boundaries) *Quadrants {
	quadrants := &Quadrants{}
	for j := range robots {
		robot := &robots[j]
		totalStepsCol := robot.Steps.Col * 100
		if totalStepsCol > 0 {
			robot.Position.Col = (robot.Position.Col + totalStepsCol) % boundaries.COLS
		} else {
			robot.Position.Col = (robot.Position.Col + totalStepsCol) % boundaries.COLS
			if robot.Position.Col < 0 {
				robot.Position.Col += boundaries.COLS
			}
		}

		totalStepsRow := robot.Steps.Row * 100
		if totalStepsRow > 0 {
			robot.Position.Row = (robot.Position.Row + totalStepsRow) % boundaries.ROWS
		} else {
			robot.Position.Row = (robot.Position.Row + totalStepsRow) % boundaries.ROWS
			if robot.Position.Row < 0 {
				robot.Position.Row += boundaries.ROWS
			}
		}

		//check  top right
		middleCol := boundaries.COLS / 2
		middleRow := boundaries.ROWS / 2
		if robot.Position.Col != middleCol && robot.Position.Row != middleRow {
			setQudrant(robot, middleCol, middleRow, quadrants)
		}
	}
	return quadrants
}

func setQudrant(robot *Robot, middleCol int, middleRow int, qudrants *Quadrants) {
	if robot.Position.Col < middleCol && robot.Position.Row < middleRow {
		qudrants.TopLeft++
	}
	if robot.Position.Col > middleCol && robot.Position.Row < middleRow {
		qudrants.TopRight++
	}
	if robot.Position.Col > middleCol && robot.Position.Row > middleRow {
		qudrants.BottomRight++
	}
	if robot.Position.Col < middleCol && robot.Position.Row > middleRow {
		qudrants.BottomLeft++
	}
}

func getRobots(lines []string) (robots []Robot) {
	for _, line := range lines {
		posStr := strings.Split(strings.Split(line, "p=")[1], ",")
		position := getDirection(posStr)
		velStr := strings.Split(strings.Split(line, "v=")[1], ",")
		steps := getDirection(velStr)
		robots = append(robots, Robot{position, steps})
	}

	return
}

func getDirection(posStr []string) Direction {
	return Direction{utils.ParseStringToNum(posStr[0]), utils.ParseStringToNum(strings.Split(posStr[1], " ")[0])}
}
