package day7

import (
	"adventOfCode/main/utils"
	"math"
	"strings"
)

type Equations struct {
	Numbers []int
	Target  int
}

type OperationsMap map[string]bool

func Part1() int {
	operationMap := OperationsMap{
		"+": true,
		"*": true,
	}
	lines := utils.GetDataSplittedInLines("day7/day7")

	equations := getEquations(lines)

	return getCombinationResult(equations, operationMap)

}

func Part2() int {
	operationMap := OperationsMap{
		"+":  true,
		"*":  true,
		"||": true,
	}
	lines := utils.GetDataSplittedInLines("day7/day7")

	equations := getEquations(lines)

	return getCombinationResult(equations, operationMap)
}

func getCombinationResult(equations []Equations, operationsMap OperationsMap) (result int) {
	for _, equation := range equations {
		if isValidBackTracking(0, 0, equation.Numbers, equation.Target, operationsMap) {
			result += equation.Target
		}
	}
	return
}

func isValidBackTracking(index int, current int, comb []int, target int, operationsMap OperationsMap) bool {

	if index == len(comb) {
		return current == target
	}

	if operationsMap["+"] && isValidBackTracking(index+1, current+comb[index], comb, target, operationsMap) {
		return true
	}

	if operationsMap["*"] && isValidBackTracking(index+1, current*comb[index], comb, target, operationsMap) {
		return true
	}

	if operationsMap["||"] {
		n := utils.CountDigits(comb[index])
		total := int((math.Pow(10, n) * float64(current)) + float64(comb[index]))
		if isValidBackTracking(index+1, total, comb, target, operationsMap) {
			return true
		}
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
