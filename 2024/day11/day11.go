package day11

import (
	"adventOfCode/main/utils"
	"fmt"
)

func Part1() int {
	data := utils.GetData("day11/day11")
	arr := utils.ParseStringToArr(data)

	count := blinkTimes(25, arr)

	return count

}

func blinkTimes(blinktingTimes int, arr []string) int {
	for i := 0; i < blinktingTimes; i++ {
		stones := []string{}
		for _, stone := range arr {
			modifiedStone := modifyStone(stone)
			stones = append(stones, modifiedStone...)
		}
		arr = stones
	}

	return len(arr)
}

func modifyStone(stone string) []string {
	if stone == "0" {
		return []string{"1"}
	}
	//even digits: split into two
	if len(stone)%2 == 0 {
		middle := len(stone) / 2
		firstHalf := stone[:middle]
		secondHalf := stone[middle:]
		//if the lead are 0, then split into only one 0 on the right
		countLeading0 := countLeadingZeroes(secondHalf)
		secondHalf = secondHalf[countLeading0:]
		if secondHalf == "" {
			secondHalf = "0"
		}
		return []string{firstHalf, secondHalf}
	}
	// if previous doesnt happen then we multiply by 2024
	number := utils.ParseStringToNum(stone)
	number *= 2024
	return []string{fmt.Sprint(number)}
}

func countLeadingZeroes(secondHalf string) (countLeading0 int) {
	n := len(secondHalf)
	if n > 1 && secondHalf[0] == '0' {

		for j := 0; j < n; j++ {
			if secondHalf[j] != '0' {
				break
			}
			countLeading0++
		}
	}
	return countLeading0
}
