package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	partPointer := flag.Int("part", 1, "which part of the day's puzzle is solved")
	inputPointer := flag.String("input", "", "file to pull input data from")
	flag.Parse()

	if *inputPointer == "" {
		panic("Input file required (/input=filePath)")
	}

	fileBytes, err := os.ReadFile(*inputPointer)
	if err != nil {
		panic("ah")
	}

	fileContent := string(fileBytes)
	fileLines := strings.Split(fileContent, "\n")

	calibrationValue := 0

	if *partPointer == 1 {
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
	}

	fmt.Printf("Calibration total %v\n", calibrationValue)
}
