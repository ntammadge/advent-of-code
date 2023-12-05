package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fileBytes, err := os.ReadFile("./inputs/day01.txt")
	if err != nil {
		panic("ah")
	}

	fileContent := string(fileBytes)
	fileLines := strings.Split(fileContent, "\n")

	calibrationValue := 0
	digitRegexp := regexp.MustCompile("\\d")
	for i := 0; i < len(fileLines); i++ {
		matches := digitRegexp.FindAllString(fileLines[i], -1)
		if matches == nil {
			panic("match error")
		}
		tens, e := strconv.Atoi(matches[0])
		ones, e := strconv.Atoi(matches[len(matches)-1])

		if e != nil {
			// Regex digit matching won't give a non-digit match
			panic("impossible")
		}

		calibrationValue += tens*10 + ones
	}

	fmt.Printf("Calibration total %v\n", calibrationValue)
}
