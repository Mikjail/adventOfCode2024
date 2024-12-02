package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Data struct {
	Array1 []int `json:"array1"`
	Array2 []int `json:"array2"`
}

func main() {
	data, err := getData("data/day1.txt")

	if err != nil {
		fmt.Println("Error: ", err)
	}

	firstPart(data)

	locationIdMap := map[int]int{}

	for _, num := range data.Array1 {
		_, exist := locationIdMap[num]
		if !exist {
			locationIdMap[num] = 0
		}
	}

	for _, num := range data.Array2 {
		_, exist := locationIdMap[num]
		if exist {
			locationIdMap[num] += 1
		}
	}

	total := 0
	for _, num := range data.Array1 {
		nTimes, exist := locationIdMap[num]
		if exist {
			total += nTimes * num
		}
	}

	fmt.Printf("%d", total)

}

func firstPart(data Data) {

	sort.Ints(data.Array1)

	sort.Ints(data.Array2)

	n := len(data.Array1)

	total := 0.0
	for i := 0; i < n; i++ {
		diff := math.Abs(float64(data.Array1[i]) - float64(data.Array2[i]))
		total += diff
	}

	fmt.Printf("%f\n", total)
}

func getData(filename string) (Data, error) {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("Error:", err)
		return Data{}, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	var array1, array2 []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		if len(parts) != 2 {
			return Data{}, fmt.Errorf("invalid line format: %s", line)
		}

		val1, err1 := strconv.Atoi(parts[0])
		val2, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			return Data{}, fmt.Errorf("error converting string to int: %s", parts)
		}

		array1 = append(array1, val1)
		array2 = append(array2, val2)
	}

	if err := scanner.Err(); err != nil {
		return Data{}, fmt.Errorf("error reading file: %w", err)
	}

	return Data{Array1: array1, Array2: array2}, nil
}
