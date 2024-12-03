package day1

import (
	util "adventOfCode/main/utils"
	"fmt"
	"math"
	"sort"
	"strings"
)

type Locations struct {
	left  []int
	right []int
}

func FirstPart(locations Locations) int {
	locations.sortLists()

	return locations.getDistance()
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

func GetLocations(fileName string) (result Locations, err error) {
	data, err := util.GetData(fileName)

	if err != nil {
		return Locations{}, fmt.Errorf("error getting the data: %s", err)
	}

	lines := strings.Split(data, "\n")

	for i := 0; i < len(lines); i++ {
		parts := strings.Fields(lines[i])
		result.left = append(result.left, util.ParseStringToNum(parts[0]))
		result.right = append(result.right, util.ParseStringToNum(parts[1]))
	}

	return
}

func (locations *Locations) sortLists() {
	sort.Ints(locations.left)

	sort.Ints(locations.right)
}

func (locations *Locations) getDistance() (total int) {
	n := len(locations.left)

	for i := 0; i < n; i++ {
		diff := math.Abs(float64(locations.left[i]) - float64(locations.right[i]))
		total += int(diff)
	}

	return
}
