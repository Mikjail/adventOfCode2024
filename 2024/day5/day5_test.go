package day5_test

import (
	"adventOfCode/main/2024/day5"
	"testing"

	"gotest.tools/v3/assert"
)

func TestPart1(t *testing.T) {
	result := day5.Part1()

	assert.DeepEqual(t, result, 5732)
}
