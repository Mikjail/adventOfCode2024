package util

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetData(fileName string) (string, error) {
	filePath := fmt.Sprintf("../data/%s.txt", fileName)
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}
	return string(data), nil
}

func ParseStringToNum(str string) int {
	num, _ := strconv.Atoi(strings.Trim(str, " "))
	return num
}
