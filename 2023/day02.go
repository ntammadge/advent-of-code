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

	gameIdSum := 0
	redsPattern := regexp.MustCompile("(\\d+) red")
	greensPattern := regexp.MustCompile("(\\d+) green")
	bluesPattern := regexp.MustCompile("(\\d+) blue")

	for i := 0; i < len(lines); i++ {
		gameData := strings.Split(lines[i], ": ")[1]
		subsets := strings.Split(gameData, "; ")
		invalid := false

		for _, subset := range subsets {
			colors := strings.Split(subset, ", ")
			for _, color := range colors {
				if redsMatch := redsPattern.FindStringSubmatch(color); redsMatch != nil {
					if redCount, _ := strconv.Atoi(redsMatch[1]); redCount > maxReds {
						invalid = true
						break
					}
				} else if greensMatch := greensPattern.FindStringSubmatch(color); greensMatch != nil {
					if greenCount, _ := strconv.Atoi(greensMatch[1]); greenCount > maxGreens {
						invalid = true
						break
					}
				} else if bluesMatch := bluesPattern.FindStringSubmatch(color); bluesMatch != nil {
					if blueCount, _ := strconv.Atoi(bluesMatch[1]); blueCount > maxBlues {
						invalid = true
						break
					}
				}
			}

			if invalid {
				break
			}
		}

		if !invalid {
			gameIdSum += i + 1
		}
	}

	fmt.Printf("Sum of invalid game ids: %v\n", gameIdSum)
}
