package day7

import (
	"adventOfCode/main/utils"
	"strings"
)

type Equations struct {
	Numbers []int
	Target  int
}

func Part1() int {
	lines := utils.GetDataSplittedInLines("day7/day7")

	equations := getEquations(lines)

	return getCombinationResult(equations)

}

func getCombinationResult(equations []Equations) (result int) {
	for _, equation := range equations {
		if isValidBackTracking(0, 0, equation.Numbers, equation.Target) {
			result += equation.Target
		}
	}
	return
}

func isValidBackTracking(index int, current int, comb []int, target int) bool {
	if index == len(comb) {
		return current == target
	}

	if isValidBackTracking(index+1, current+comb[index], comb, target) {
		return true
	}

	if isValidBackTracking(index+1, current*comb[index], comb, target) {
		return true
	}

	return false
}

func getEquations(lines []string) (equations []Equations) {
	for _, line := range lines {
		data := strings.Split(line, ":")
		result := utils.ParseStringToNum(data[0])
		trimmedData := strings.TrimSpace(data[1])
		numbers := getNumbers(strings.Split(trimmedData, " "))
		equations = append(equations, Equations{Target: result, Numbers: numbers})
	}
	return
}

func getNumbers(numbers []string) (result []int) {

	for _, number := range numbers {
		result = append(result, utils.ParseStringToNum(number))
	}
	return
}
