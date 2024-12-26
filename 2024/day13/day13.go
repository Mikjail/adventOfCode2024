package day13

import (
	"adventOfCode/main/utils"
	"fmt"
	"strings"
)

type Axes struct {
	X int
	Y int
}

type Game struct {
	ButtonA Axes
	ButtonB Axes
	Target  Axes
}

type CountedButtons struct {
	CountA int
	CountB int
}

func Part1() int {
	lines := utils.GetDataSplittedInLines("day13/day13")
	games := getGames(lines)
	result := 0

	for _, game := range games {
		isValid, countedButtons := backTracking(game, 100, 100)
		if isValid {
			result += countTokens(countedButtons)
		}
	}

	return result
}

func countTokens(countedButtons CountedButtons) int {
	return (countedButtons.CountA * 3) + (countedButtons.CountB * 1)
}

func backTracking(game Game, countA, countB int) (bool, CountedButtons) {
	stack := []CountedButtons{{
		countA, countB,
	}}
	visited := make(map[string]bool)

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		key := fmt.Sprintf("%d,%d", current.CountA, current.CountB)
		if visited[key] {
			continue
		}
		visited[key] = true

		buttonAX := game.ButtonA.X * current.CountA
		buttonBX := game.ButtonB.X * current.CountB
		buttonAY := game.ButtonA.Y * current.CountA
		buttonBY := game.ButtonB.Y * current.CountB

		if (buttonAX+buttonBX) < game.Target.X || (buttonAY+buttonBY) < game.Target.Y {
			continue
		}

		if game.Target.X == (buttonAX+buttonBX) && game.Target.Y == (buttonAY+buttonBY) {
			return true, CountedButtons{current.CountA, current.CountB}
		}

		if current.CountA > 0 {
			stack = append(stack, CountedButtons{current.CountA - 1, current.CountB})
		}

		if current.CountB > 0 {
			stack = append(stack, CountedButtons{current.CountA, current.CountB - 1})
		}
	}

	return false, CountedButtons{}
}

func getGames(lines []string) []Game {
	games := []Game{}
	for i := 0; i < len(lines); i += 4 {
		aX := utils.ParseStringToNum(strings.Split(strings.Split(lines[i], "X+")[1], ",")[0])
		aY := utils.ParseStringToNum(strings.TrimSpace(strings.Split(lines[i], "Y+")[1]))
		bX := utils.ParseStringToNum(strings.Split(strings.Split(lines[i+1], "X+")[1], ",")[0])
		bY := utils.ParseStringToNum(strings.TrimSpace(strings.Split(lines[i+1], "Y+")[1]))
		targetX := utils.ParseStringToNum(strings.Split(strings.Split(lines[i+2], "X=")[1], ",")[0])
		targetB := utils.ParseStringToNum(strings.TrimSpace(strings.Split(lines[i+2], "Y=")[1]))
		game := Game{Axes{aX, aY}, Axes{bX, bY}, Axes{targetX, targetB}}
		games = append(games, game)
	}
	return games
}
