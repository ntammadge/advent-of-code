package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func day1(fileBytes []byte, part int) {
	fileContent := string(fileBytes)
	lines := strings.Split(fileContent, "\n")

	// Convert word-form numbers to number characters and preserve the beginning and final characters for part 2 only
	if part == 2 {
		for i := 0; i < len(lines); i++ {
			lines[i] = strings.ReplaceAll(lines[i], "one", "o1e")
			lines[i] = strings.ReplaceAll(lines[i], "two", "t2o")
			lines[i] = strings.ReplaceAll(lines[i], "three", "t3e")
			lines[i] = strings.ReplaceAll(lines[i], "four", "f4r")
			lines[i] = strings.ReplaceAll(lines[i], "five", "f5e")
			lines[i] = strings.ReplaceAll(lines[i], "six", "s6x")
			lines[i] = strings.ReplaceAll(lines[i], "seven", "s7n")
			lines[i] = strings.ReplaceAll(lines[i], "eight", "e8t")
			lines[i] = strings.ReplaceAll(lines[i], "nine", "n9e")
		}
	}

	calibrationValue := 0

	digitRegexp := regexp.MustCompile("\\d")
	for i := 0; i < len(lines); i++ {
		matches := digitRegexp.FindAllString(lines[i], -1)
		if matches == nil {
			panic("match error")
		}

		tens, _ := strconv.Atoi(matches[0])
		ones, _ := strconv.Atoi(matches[len(matches)-1])

		calibrationValue += tens*10 + ones
	}

	fmt.Printf("Calibration total %v\n", calibrationValue)
}
