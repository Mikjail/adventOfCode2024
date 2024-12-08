package day3

import (
	"adventOfCode/main/utils"
	"regexp"
)

type CompInstructions = [][]string

func Part1() int {
	data := utils.GetData("day3/day3")

	compInst := GetComputerInstructions(data)

	return compute(compInst)

}

func Part2() (result int) {
	data := utils.GetData("day3/day3")

	enabledInst := GetEnabledComputerInstructions(data)

	computerInst := GetComputerInstructions(enabledInst)

	return compute(computerInst)
}

func compute(compInst CompInstructions) (result int) {
	for _, numbers := range compInst {
		num1 := utils.ParseStringToNum(numbers[1])
		num2 := utils.ParseStringToNum(numbers[2])
		product := num1 * num2
		result += product
	}
	return
}

func GetEnabledComputerInstructions(data string) (result string) {
	enabled := true
	end := len(data)
	for i := 0; i < end; i++ {
		if i+4 <= end && data[i:i+4] == "do()" {
			enabled = true
			continue
		}
		if i+7 <= end && data[i:i+7] == "don't()" {
			enabled = false
			continue
		}
		if enabled {
			result += string(data[i])
		}
	}

	return
}

func GetComputerInstructions(data string) (compInstructions CompInstructions) {
	pattern := `mul\((\d{1,3}),\s*(\d{1,3})\)`
	re := regexp.MustCompile(pattern)
	compInstructions = re.FindAllStringSubmatch(data, -1)

	return
}
