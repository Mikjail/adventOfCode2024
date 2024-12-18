package day9_test

import (
	"adventOfCode/main/2024/day9"
	"testing"

	"gotest.tools/v3/assert"
)

func TestPart1(t *testing.T) {
	data := day9.Part1()

	assert.Equal(t, data, 6341711060162)
}

func TestPart2(t *testing.T) {
	data := day9.Part2()

	assert.Equal(t, data, 6377400869326)
}
