package day12_test

import (
	"adventOfCode/main/2024/day12"
	"testing"

	"gotest.tools/v3/assert"
)

func TestPart1(t *testing.T) {
	data := day12.Part1()

	assert.Equal(t, data, 1467094)
}
