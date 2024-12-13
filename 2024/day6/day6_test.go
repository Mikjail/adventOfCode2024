package day6_test

import (
	"adventOfCode/main/2024/day6"
	"testing"

	"gotest.tools/v3/assert"
)

func TestPart1(t *testing.T) {
	data := day6.Part1()

	assert.Equal(t, data, 41)
}

func TestPart2(t *testing.T) {
	data := day6.Part2()

	assert.Equal(t, data, 1530)
}
