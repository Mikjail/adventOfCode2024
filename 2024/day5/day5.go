package day5

import (
	"adventOfCode/main/utils"
	"strings"
)

type Rules = map[string]map[string]struct{}

func Part1() int {
	data := utils.GetData("day5/day5")

	newResult := strings.Split(data, "\n\n")

	rulesMap := getRules(newResult[0])

	return getLinesWithRightOrders(rulesMap, newResult[1])
}

func getLinesWithRightOrders(ruleMap Rules, data string) (totalOfMiddle int) {
	lines := strings.Split(data, "\n")
	for _, line := range lines {
		numbers := strings.Split(line, ",")
		n := len(numbers)
		if isValidOrder(n, ruleMap, numbers) {
			middle := n / 2
			totalOfMiddle += utils.ParseStringToNum(numbers[middle])
		}
	}
	return
}

func isValidOrder(n int, ruleMap map[string]map[string]struct{}, numbers []string) bool {
	for i := 1; i < n; i++ {
		pageMap, exist := ruleMap[numbers[i]]
		if !exist {
			return false
		}
		_, exist = pageMap[numbers[i-1]]
		if !exist {
			return false
		}
	}
	return true
}

func getRules(data string) Rules {
	rules := strings.Split(data, "\n")
	orderMap := map[string]map[string]struct{}{}
	for _, order := range rules {
		if order == " " {
			break
		}
		orderList := strings.Split(order, "|")
		key := orderList[1]
		value := orderList[0]

		if _, exist := orderMap[key]; !exist {

			orderMap[key] = map[string]struct{}{}
		}

		orderMap[key][value] = struct{}{}
	}
	return orderMap
}
