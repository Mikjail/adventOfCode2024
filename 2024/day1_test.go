package day1_test

import (
	day1 "adventOfCode/main/2024"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	data, err := day1.GetLocations("day1")

	assert.NoError(t, err, "Expected no error, but got one")

	result := day1.FirstPart(data)

	assert.Equal(t, result, 1970720)

}

func TestPart2(t *testing.T) {
	data, err := day1.GetLocations("day1")

	assert.NoError(t, err, "Expected no error, but got one")

	result := day1.SecondPart(data)

	assert.Equal(t, result, 17191599)
}
