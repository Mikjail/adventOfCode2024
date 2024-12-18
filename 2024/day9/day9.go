package day9

import (
	"adventOfCode/main/utils"
	"strconv"
	"strings"
)

type ListDisk []string

type FileInfoMap map[int]FileInfo

type FileInfo struct {
	Pos  int
	Size int
}

type FreeSpace struct {
	Start int
	Size  int
}

func Part1() int {
	data := utils.GetData("day9/day9")

	mappedDisk := getMapDisk(data)

	filesMoved := moveFilesFromDisk(mappedDisk)

	return getFileSystemCheckSum(filesMoved)
}

func Part2() int {
	data := utils.GetData("day9/day9")

	fileInfo, freeSpace, fileIds, totalPositions := orderTimeAndPositionOfFilesAndFreeSpaces(data)

	fileInfoMap := moveCompleteFilesIntoBlankSpaces(fileInfo, freeSpace, fileIds)

	filesMoved := getFilesMovedIntoList(fileInfoMap, totalPositions)

	return sumFiles(filesMoved)
}

func sumFiles(filesMoved []int) (result int) {
	for i := 0; i < len(filesMoved); i++ {
		if filesMoved[i] != -1 {
			result += (filesMoved[i] * i)
		}

	}
	return
}

func getFilesMovedIntoList(fileInfoMap FileInfoMap, length int) []int {
	arrResult := make([]int, length)

	for i := 0; i < len(arrResult); i++ {
		arrResult[i] = -1
	}

	for i := 0; i < length; i++ {
		for number, file := range fileInfoMap {
			if file.Pos == i {
				for j := file.Pos; j < file.Pos+file.Size; j++ {
					arrResult[j] = number
				}
			}
		}
	}

	return arrResult
}

func moveCompleteFilesIntoBlankSpaces(fileInfos FileInfoMap, freeSpaces []FreeSpace, fileIndexId int) FileInfoMap {
	for fileIndexId > 0 {
		fileIndexId-- //pointing to the right.
		file := fileInfos[fileIndexId]
		for i, freeSpace := range freeSpaces {
			if freeSpace.Start >= file.Pos {
				freeSpaces = freeSpaces[:i]
				break
			}
			if file.Size <= freeSpace.Size {
				fileInfos[fileIndexId] = FileInfo{Pos: freeSpace.Start, Size: file.Size}
				if file.Size < freeSpace.Size {
					freeSpaces[i] = FreeSpace{Start: freeSpace.Start + file.Size, Size: freeSpace.Size - file.Size}
					break
				}
				freeSpaces = append(freeSpaces[:i], freeSpaces[i+1:]...)
				break
			}
		}
	}
	return fileInfos
}

func orderTimeAndPositionOfFilesAndFreeSpaces(data string) (FileInfoMap, []FreeSpace, int, int) {
	fileMap := FileInfoMap{}
	freeSpace := []FreeSpace{}
	lenFiles := 0
	position := 0
	for i, char := range data {
		times := utils.ParseStringToNum(string(char))
		if i%2 == 0 {
			fileMap[lenFiles] = FileInfo{Pos: position, Size: times}
			lenFiles += 1
		} else {
			freeSpace = append(freeSpace, FreeSpace{Start: position, Size: times})
		}
		position += times
	}

	return fileMap, freeSpace, lenFiles, position
}

func getFileSystemCheckSum(filesMoved ListDisk) (result int) {
	for i := 0; i < len(filesMoved); i++ {
		if filesMoved[i] != "." {
			number := utils.ParseStringToNum(string(filesMoved[i]))
			result += (number * i)
		}
	}
	return
}

func moveFilesFromDisk(mappedDisk ListDisk) ListDisk {
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

func getMapDisk(data string) (diskMapped ListDisk) {
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

func getFiles(number string, times int) (files ListDisk) {
	for i := 0; i < times; i++ {
		files = append(files, number)
	}
	return
}
