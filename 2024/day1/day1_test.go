package day1_test

import (
	"adventOfCode/main/2024/day1"
	"testing"

	"gotest.tools/v3/assert"
)

func TestPart1(t *testing.T) {
	data, err := day1.GetLocations()

	assert.NilError(t, err, "Expected no error, but got one")

	result := day1.FirstPart(data)

	assert.Equal(t, 1970720, result)

}

func TestPart2(t *testing.T) {
	data, err := day1.GetLocations()

	assert.NilError(t, err, "Expected no error, but got one")

	result := day1.SecondPart(data)

	assert.Equal(t, 17191599, result)
}
