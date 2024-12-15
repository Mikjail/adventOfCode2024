package day8

import "adventOfCode/main/utils"

type Antenna struct {
	freq string
	row  int
	col  int
}

type AntennasMap map[string][]Antenna

type AntinodeTuple struct {
	Row int
	Col int
}

type Antinodes map[AntinodeTuple]struct{}

func Part1() int {
	lines := utils.GetDataSplittedInLines("day8/day8")

	gridMap := getAntennaPositions(lines)

	antinodes := getNumberOfAntinodes(gridMap, lines)

	return len(antinodes)
}

func Part2() int {
	lines := utils.GetDataSplittedInLines("day8/day8")

	gridMap := getAntennaPositions(lines)

	antiNodeCoord := getAtniNodesPart2(gridMap, lines)

	return len(antiNodeCoord)

}

func getNumberOfAntinodes(gridMap AntennasMap, lines []string) Antinodes {
	ROWS := len(lines)
	COLS := len(lines[0])
	antinodes := make(Antinodes)
	for _, antennas := range gridMap {
		for i := 0; i < len(antennas); i++ {
			for j := 0; j < len(antennas); j++ {
				if antennas[i] == antennas[j] {
					continue

				}
				difRow := antennas[j].row - antennas[i].row
				difCol := antennas[j].col - antennas[i].col
				antiNodeRow := antennas[i].row - difRow
				antiNodeCol := antennas[i].col - difCol
				if antiNodeRow >= 0 && antiNodeRow < ROWS && antiNodeCol >= 0 && antiNodeCol < COLS {
					antiNode := AntinodeTuple{Row: antiNodeRow, Col: antiNodeCol}
					antinodes[antiNode] = struct{}{}
				}

			}
		}
	}
	return antinodes
}

func getAtniNodesPart2(gridMap AntennasMap, lines []string) Antinodes {
	ROWS := len(lines)
	COLS := len(lines[0])
	antinodes := make(Antinodes)
	for _, antennas := range gridMap {
		for i := 0; i < len(antennas); i++ {
			for j := 0; j < len(antennas); j++ {
				if antennas[i] == antennas[j] {
					continue

				}
				difRow := antennas[j].row - antennas[i].row
				difCol := antennas[j].col - antennas[i].col
				totalRow := difRow
				totalCol := difCol
				inBoard := true
				for inBoard {
					antiNodeRow := antennas[i].row - totalRow
					antiNodeCol := antennas[i].col - totalCol
					if antiNodeRow >= 0 && antiNodeRow < ROWS && antiNodeCol >= 0 && antiNodeCol < COLS {
						antiNode := AntinodeTuple{Row: antiNodeRow, Col: antiNodeCol}
						antinodes[antiNode] = struct{}{}
						totalRow += difRow
						totalCol += difCol
					} else {
						inBoard = false
					}
				}
				antiNode := AntinodeTuple{Row: antennas[j].row, Col: antennas[j].col}
				antinodes[antiNode] = struct{}{}
			}
		}
	}
	return antinodes
}
func getAntennaPositions(lines []string) (gridMap AntennasMap) {
	ROWS := len(lines)
	COLS := len(lines[0])
	gridMap = make(AntennasMap)
	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLS; j++ {
			if lines[i][j] != '.' {
				freq := string(lines[i][j])
				gridMap[freq] = append(gridMap[freq], Antenna{freq: freq, row: i, col: j})
			}
		}
	}
	return
}
