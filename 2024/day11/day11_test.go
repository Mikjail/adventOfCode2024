package day11_test

import (
	"adventOfCode/main/2024/day11"
	"testing"

	"gotest.tools/v3/assert"
)

func TestPart1(t *testing.T) {
	data := day11.Part1()

	assert.Equal(t, data, 197357)
}
