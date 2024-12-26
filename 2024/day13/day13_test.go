package day13_test

import (
	"adventOfCode/main/2024/day13"
	"testing"

	"gotest.tools/v3/assert"
)

func TestPart1(t *testing.T) {
	data := day13.Part1()

	assert.Equal(t, data, 28059)
}
