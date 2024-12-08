package day3_test

import (
	"adventOfCode/main/2024/day3"
	"testing"

	"gotest.tools/v3/assert"
)

func TestPart1(t *testing.T) {
	result := day3.Part1()

	assert.Equal(t, result, 164730528)

}

func TestPart2(t *testing.T) {
	result := day3.Part2()
	assert.Equal(t, result, 70478672)
}
