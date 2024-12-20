package day10_test

import (
	"adventOfCode/main/2024/day10"
	"testing"

	"gotest.tools/v3/assert"
)

func TestPart1(t *testing.T) {
	result := day10.Part1()

	assert.Equal(t, result, 512)

}

func TestPart2(t *testing.T) {
	result := day10.Part2()

	assert.Equal(t, result, 1045)
}
