package util

import (
	"fmt"
	"os"
)

func GetData(fileName string) (string, error) {
	filePath := fmt.Sprintf("../data/%s.txt", fileName)
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}
	return string(data), nil
}
