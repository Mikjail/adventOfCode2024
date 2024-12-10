package day4_test

import (
	"adventOfCode/main/2024/day4"
	"testing"

	"gotest.tools/v3/assert"
)

func TestPart1(t *testing.T) {
	result := day4.Part1()

	assert.Equal(t, result, 2517)
}

func TestPart2(t *testing.T) {
	result := day4.Part2()

	assert.Equal(t, result, 1960)
}
