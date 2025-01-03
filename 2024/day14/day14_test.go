package day14_test

import (
	"adventOfCode/main/2024/day14"
	"testing"

	"gotest.tools/v3/assert"
)

func TestPart1(t *testing.T) {
	result := day14.Part1()

	assert.Equal(t, result, 217328832)

}

func TestPart2(t *testing.T) {
	day14.Part2()

	assert.Equal(t, 7412, 7412)
}
