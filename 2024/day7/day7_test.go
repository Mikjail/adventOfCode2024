package day7_test

import (
	"adventOfCode/main/2024/day7"
	"testing"

	"gotest.tools/v3/assert"
)

func TestPart1(t *testing.T) {
	data := day7.Part1()

	assert.Equal(t, data, 2501605301465)
}

func TestPart2(t *testing.T) {
	data := day7.Part2()

	assert.Equal(t, data, 44841372855953)
}
