package day15_test

import (
	"adventOfCode/main/2024/day15"
	"testing"

	"gotest.tools/v3/assert"
)

func TestPart1(t *testing.T) {
	result := day15.Part1()

	assert.Equal(t, result, 1552463)
}
