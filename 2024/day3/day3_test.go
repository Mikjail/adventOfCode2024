package day3_test

import (
	"adventOfCode/main/2024/day3"
	"testing"

	"gotest.tools/v3/assert"
)

func TestPart1(t *testing.T) {
	data := day3.GetComputerInstructions()

	result := day3.Part1(data)

	assert.Equal(t, result, 164730528)

}
