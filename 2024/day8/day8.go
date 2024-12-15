package day8

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

func GetNumberOfAntinodes(gridMap AntennasMap, lines []string) int {
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
	return len(antinodes)
}

func GetAntennaAltPositions(lines []string) (gridMap AntennasMap) {
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
