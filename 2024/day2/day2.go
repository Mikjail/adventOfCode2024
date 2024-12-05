package day2

import (
	"adventOfCode/main/utils"
	"log"
	"strconv"
	"strings"
)

type LocationIds []int

type Locations struct {
	data []LocationIds
}

func GetLocations(fileName string) (result Locations) {
	data := utils.GetData(fileName)

	lines := strings.Split(data, "\n")

	for i := 0; i < len(lines); i++ {

		parts := strings.Fields(lines[i])
		intParts := make([]int, len(parts))

		for j, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				// Handle the error appropriately
				log.Fatalf("Error converting string to int: %v", err)
			}
			intParts[j] = num
		}
		result.data = append(result.data, intParts)
	}

	return
}

func FirstPart(locations Locations) (safe int) {
	for _, ids := range locations.data {
		if isLocationSafe(ids) {
			safe += 1
		}
	}
	return
}

func isLocationSafe(ids []int) bool {
	limit := 3
	isIncreasing := ids[0] < ids[1]
	for i := 1; i < len(ids); i++ {
		num1 := ids[i-1]
		num2 := ids[i]
		diff := utils.GetAbsolute(num1, num2)

		if diff > limit || diff == 0 {
			return false
		}
		if num1 < num2 && !isIncreasing {
			return false
		}
		if num1 > num2 && isIncreasing {
			return false
		}
	}
	return true
}

func SecondPart(locations Locations) int {
	result := 0
	for _, ids := range locations.data {
		if isLocationSafe(ids) {
			result++
			continue
		}

		for i := range ids {
			newIds := append(append(LocationIds{}, ids[:i]...), ids[i+1:]...)
			if isLocationSafe(newIds) {
				result++
				break
			}
		}
	}
	return result
}
