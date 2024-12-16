package day9

import (
	"adventOfCode/main/utils"
	"strconv"
	"strings"
)

func Part1() int {
	data := utils.GetData("day9/day9")

	mappedDisk := getMapDisk(data)

	filesMoved := moveFilesFromDisk(mappedDisk)

	return getFileSystemCheckSum(filesMoved)
}

func getFileSystemCheckSum(filesMoved []string) (result int) {
	for i := 0; i < len(filesMoved); i++ {
		if filesMoved[i] != "." {
			number := utils.ParseStringToNum(string(filesMoved[i]))
			result += (number * i)
		}
	}
	return
}

func moveFilesFromDisk(mappedDisk []string) []string {
	left := 0
	right := len(mappedDisk) - 1

	for left < right {
		if mappedDisk[right][0] == '.' {
			right--
			continue
		}
		if mappedDisk[left][0] == '.' {
			mappedDisk[left], mappedDisk[right] = mappedDisk[right], mappedDisk[left]
			right--
		}
		left++
	}

	return mappedDisk
}

func getMapDisk(data string) (diskMapped []string) {
	fileIndex := 0
	isFile := true
	for i := 0; i < len(data); i++ {
		times := utils.ParseStringToNum(string(data[i]))
		if isFile {
			files := getFiles(strconv.Itoa(fileIndex), times)
			diskMapped = append(diskMapped, files...)
			isFile = false
			fileIndex++
		} else {
			freeSpaces := strings.Split(strings.Repeat(".", times), "")
			diskMapped = append(diskMapped, freeSpaces...)
			isFile = true
		}
	}

	return
}

func getFiles(number string, times int) (files []string) {
	for i := 0; i < times; i++ {
		files = append(files, number)
	}
	return
}
