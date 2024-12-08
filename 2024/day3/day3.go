package day3

import (
	"adventOfCode/main/utils"
	"regexp"
)

func Part1(compInst [][]string) (result int) {
	for _, numbers := range compInst {
		num1 := utils.ParseStringToNum(numbers[1])
		num2 := utils.ParseStringToNum(numbers[2])
		product := num1 * num2
		result += product
	}
	return
}

func GetComputerInstructions() (compInstructions [][]string) {
	data := utils.GetData("day3/day3")

	pattern := `mul\((\d{1,3}),\s*(\d{1,3})\)`
	re := regexp.MustCompile(pattern)
	compInstructions = re.FindAllStringSubmatch(data, -1)

	return
}
