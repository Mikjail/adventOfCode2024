package day11

import (
	"adventOfCode/main/utils"
	"fmt"
	"strings"
)

type StonesMap map[string]int

func Part1(n int) int {
	data := utils.GetData("day11/day11")
	arr := utils.ParseStringToArr(data)

	count := blinkTimes(n, arr)

	return count

}

func blinkTimes(blinktingTimes int, arr []string) int {
	stoneNumberMap := make(StonesMap)
	for _, char := range arr {
		stoneNumberMap[char] += 1
	}

	for i := 0; i < blinktingTimes; i++ {
		newStones := make(StonesMap)
		for key := range stoneNumberMap {
			modifyStone(key, newStones, stoneNumberMap)
		}
		stoneNumberMap = newStones
	}

	result := 0
	for _, val := range stoneNumberMap {
		result += val
	}

	return result
}

func modifyStone(stone string, newStones StonesMap, memoStones StonesMap) {
	if stone == "0" {
		newStones["1"] += memoStones[stone]
		return
	}
	//even digits: split into two
	if len(stone)%2 == 0 {
		middle := len(stone) / 2
		firstHalf := stone[:middle]
		secondHalf := stone[middle:]
		//if the lead are 0, then split into only one 0 on the right
		secondHalf = strings.TrimLeft(secondHalf, "0")
		if secondHalf == "" {
			secondHalf = "0"
		}
		newStones[firstHalf] += memoStones[stone]
		newStones[secondHalf] += memoStones[stone]
		return
	}
	// if previous doesnt happen then we multiply by 2024
	number := utils.ParseStringToNum(stone)
	number *= 2024
	newStones[fmt.Sprint(number)] += memoStones[stone]
	return
}
