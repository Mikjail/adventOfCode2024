package day8_test

import (
	"adventOfCode/main/2024/day8"
	"testing"

	"gotest.tools/v3/assert"
)

func TestPart1(t *testing.T) {
	result := day8.Part1()

	assert.Equal(t, result, 332)
}

func TestPart2(t *testing.T) {
	result := day8.Part2()

	assert.Equal(t, result, 1174)
}
