package day2_test

import (
	"adventOfCode/main/2024/day2"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	data := day2.GetLocations("day2/day2")

	result := day2.FirstPart(data)

	assert.Equal(t, 680, result)

}

func TestPart2(t *testing.T) {
	data := day2.GetLocations("day2/day2")

	result := day2.SecondPart(data)

	assert.Equal(t, 710, result)

}
