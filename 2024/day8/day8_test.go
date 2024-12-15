package day8_test

import (
	"adventOfCode/main/2024/day8"
	"adventOfCode/main/utils"
	"testing"

	"gotest.tools/v3/assert"
)

func TestPart1(t *testing.T) {
	result := Part1()

	assert.Equal(t, result, 14)
}

func Part1() int {
	lines := utils.GetDataSplittedInLines("day8/day8")

	gridMap := day8.GetAntennaAltPositions(lines)

	return day8.GetNumberOfAntinodes(gridMap, lines)
}
