package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func day2(inputData string, part int) {
	lines := strings.Split(inputData, "\n")

	gameSum := 0

	for i := 0; i < len(lines); i++ {
		if part == 2 {
			gameSum += p2Eval(lines[i])
		} else if p1Eval(lines[i]) {
			gameSum += i + 1
		}
	}

	fmt.Printf("Sum from games: %v\n", gameSum)
}

func p1Eval(line string) (isInvalid bool) {
	maxAllowedReds := 12
	maxAllowedGreens := 13
	maxAllowedBlues := 14

	gameData := strings.Split(line, ": ")[1]
	subsets := strings.Split(gameData, "; ")

	for _, subset := range subsets {
		redCount := getRedCount(subset)
		blueCount := getBlueCount(subset)
		greenCount := getGreenCount(subset)

		if redCount > maxAllowedReds || blueCount > maxAllowedBlues || greenCount > maxAllowedGreens {
			return true
		}
	}
	return false
}

func p2Eval(line string) (setPower int) {
	maxReds := 0
	maxBlues := 0
	maxGreens := 0

	gameData := strings.Split(line, ": ")[1]
	subsets := strings.Split(gameData, "; ")

	for _, subset := range subsets {
		redCount := getRedCount(subset)
		blueCount := getBlueCount(subset)
		greenCount := getGreenCount(subset)

		if redCount > maxReds {
			maxReds = redCount
		}
		if blueCount > maxBlues {
			maxBlues = blueCount
		}
		if greenCount > maxGreens {
			maxGreens = greenCount
		}
	}
	return maxReds * maxBlues * maxGreens
}

func getRedCount(subset string) (redCount int) {
	redsPattern := regexp.MustCompile("(\\d+) red")
	redsMatch := redsPattern.FindStringSubmatch(subset)

	redCount = 0
	if redsMatch != nil {
		redCount, _ = strconv.Atoi(redsMatch[1])
	}
	return redCount
}

func getGreenCount(subset string) (greenCount int) {
	greensPattern := regexp.MustCompile("(\\d+) green")
	greensMatch := greensPattern.FindStringSubmatch(subset)

	greenCount = 0
	if greensMatch != nil {
		greenCount, _ = strconv.Atoi(greensMatch[1])
	}
	return greenCount
}

func getBlueCount(subset string) (blueCount int) {
	bluesPattern := regexp.MustCompile("(\\d+) blue")
	bluesMatch := bluesPattern.FindStringSubmatch(subset)

	blueCount = 0
	if bluesMatch != nil {
		blueCount, _ = strconv.Atoi(bluesMatch[1])
	}
	return blueCount
}
