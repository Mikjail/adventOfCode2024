package day1

import (
	util "adventOfCode/main/utils"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Locations struct {
	left  []int
	right []int
}

func FirstPart(data Locations) float64 {

	sort.Ints(data.left)

	sort.Ints(data.right)

	n := len(data.left)

	total := 0.0
	for i := 0; i < n; i++ {
		diff := math.Abs(float64(data.left[i]) - float64(data.right[i]))
		total += diff
	}

	return total
}

func SecondPart(data Locations) int {
	locationIdMap := map[int]int{}

	for _, num := range data.left {
		_, exist := locationIdMap[num]
		if !exist {
			locationIdMap[num] = 0
		}
	}

	for _, num := range data.right {
		_, exist := locationIdMap[num]
		if exist {
			locationIdMap[num] += 1
		}
	}

	total := 0
	for _, num := range data.left {
		nTimes, exist := locationIdMap[num]
		if exist {
			total += nTimes * num
		}
	}

	return total
}

func GetLocations(fileName string) (Locations, error) {
	data, err := util.GetData(fileName)

	if err != nil {
		return Locations{}, fmt.Errorf("error getting the data: %s", err)
	}

	lines := strings.Split(data, "\n")

	left := []int{}
	right := []int{}

	for i := 0; i < len(lines); i++ {
		parts := strings.Fields(lines[i])
		leftVal, err := strconv.Atoi(parts[0])
		if err != nil {
			return Locations{}, fmt.Errorf("error parsing left value: %s", err)
		}
		left = append(left, leftVal)
		rightVal, err := strconv.Atoi(parts[1])
		if err != nil {
			return Locations{}, fmt.Errorf("error parsing right value: %s", err)
		}
		right = append(right, rightVal)
	}

	return Locations{left: left, right: right}, nil
}
