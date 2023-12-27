package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func day2(inputData string, part int) {
	maxReds := 12
	maxGreens := 13
	maxBlues := 14

	lines := strings.Split(inputData, "\n")

	gameSum := 0
	redsPattern := regexp.MustCompile("(\\d+) red")
	greensPattern := regexp.MustCompile("(\\d+) green")
	bluesPattern := regexp.MustCompile("(\\d+) blue")

	for i := 0; i < len(lines); i++ {
		gameData := strings.Split(lines[i], ": ")[1]
		subsets := strings.Split(gameData, "; ")
		invalid := false
		// Part 2 only
		// Can I make this cleaner?
		minRed := 0
		minGreen := 0
		minBlue := 0

		for _, subset := range subsets {
			redsMatch := redsPattern.FindStringSubmatch(subset)
			greensMatch := greensPattern.FindStringSubmatch(subset)
			bluesMatch := bluesPattern.FindStringSubmatch(subset)

			redCount := 0
			if redsMatch != nil {
				redCount, _ = strconv.Atoi((redsMatch[1]))
			}
			blueCount := 0
			if bluesMatch != nil {
				blueCount, _ = strconv.Atoi((bluesMatch[1]))
			}
			greenCount := 0
			if greensMatch != nil {
				greenCount, _ = strconv.Atoi((greensMatch[1]))
			}

			if part == 1 {
				if redCount > maxReds || blueCount > maxBlues || greenCount > maxGreens {
					invalid = true
					break
				}
			} else {
				if redCount > minRed {
					minRed = redCount
				}
				if greenCount > minGreen {
					minGreen = greenCount
				}
				if blueCount > minBlue {
					minBlue = blueCount
				}
			}
		}

		if part == 2 {
			gameSum += minRed * minGreen * minBlue
		} else if !invalid {
			gameSum += i + 1
		}
	}

	fmt.Printf("Sum from games: %v\n", gameSum)
}
