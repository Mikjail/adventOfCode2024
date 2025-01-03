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

func PrintGridInFile(grid [][]string, stepNumber int) {
	file, err := os.OpenFile("grid_output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	fmt.Fprintf(file, "stepNumber %d\n", stepNumber)
	for _, row := range grid {
		for _, col := range row {
			if col == "" {
				fmt.Fprint(file, " ")
			} else {
				fmt.Fprint(file, col)
			}
		}
		fmt.Fprintln(file)
	}
}

func GetDataSplittedInLines(fileName string) []string {
	data := GetData(fileName)
	return strings.Split(data, "\n")
}

func ParseStringToNum(str string) int {
	num, _ := strconv.Atoi(strings.Trim(str, " "))
	return num
}

func ParseStringToArr(str string) []string {
	return strings.Split(str, " ")
}

func CountDigits(n int) (count float64) {
	if n == 0 {
		return 1
	}
	for n != 0 {
		n /= 10
		count++
	}
	return
}

func GetAbsolute(num1 int, num2 int) int {
	return int(math.Abs(float64(num1) - float64(num2)))
}

func ParseArrayStringIntoMatrix(data []string) (result [][]string) {

	for _, val := range data {
		result = append(result, strings.Split(val, ""))
	}
	return result
}
