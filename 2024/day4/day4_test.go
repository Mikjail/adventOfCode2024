package day4_test

import (
	"adventOfCode/main/2024/day4"
	"testing"

	"gotest.tools/v3/assert"
)

func TestPart1(t *testing.T) {
	result := Part1()

	assert.Equal(t, result, 18)
}

func Part1() int {
	words := day4.GetAllLetters()

	return day4.SearchWord(words)
}
