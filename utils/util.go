package utils

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func GetData(fileName string) string {
	filePath := fmt.Sprintf("../../data/%s.txt", fileName)
	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(fmt.Errorf("error reading file: %w", err))
	}
	return string(data)
}

func GetDataSplittedInLines(fileName string) []string {
	data := GetData(fileName)
	return strings.Split(data, "\n")
}

func ParseStringToNum(str string) int {
	num, _ := strconv.Atoi(strings.Trim(str, " "))
	return num
}

func GetAbsolute(num1 int, num2 int) int {
	return int(math.Abs(float64(num1) - float64(num2)))
}

func SplitStringIntoArray(data []string) (result [][]string) {

	for _, val := range data {
		result = append(result, strings.Split(val, ""))
	}
	return result
}
