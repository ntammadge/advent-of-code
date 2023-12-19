package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const one string = "one"
const two string = "two"
const three string = "three"
const four string = "four"
const five string = "five"
const six string = "six"
const seven string = "seven"
const eight string = "eight"
const nine string = "nine"

func main() {
	partPointer := flag.Int("part", 1, "which part of the day's puzzle is solved")
	inputPointer := flag.String("input", "", "file to pull input data from")
	flag.Parse()

	if *inputPointer == "" {
		panic("Input file required (/input=filePath)")
	}
	if *partPointer != 1 && *partPointer != 2 {
		panic("Part must be 1 (default) or 2")
	}

	fileBytes, err := os.ReadFile(*inputPointer)
	if err != nil {
		panic("ah")
	}

	fileContent := string(fileBytes)
	lines := strings.Split(fileContent, "\n")

	// Convert word-form numbers to number characters and preserve the beginning and final characters for part 2 only
	if *partPointer == 2 {
		numberWordMap := map[string]string{
			one:   "o1e",
			two:   "t2o",
			three: "t3e",
			four:  "f4r",
			five:  "f5e",
			six:   "s6x",
			seven: "s7n",
			eight: "e8t",
			nine:  "n9e"}

		for i := 0; i < len(lines); i++ {
			lines[i] = strings.ReplaceAll(lines[i], one, numberWordMap[one])
			lines[i] = strings.ReplaceAll(lines[i], two, numberWordMap[two])
			lines[i] = strings.ReplaceAll(lines[i], three, numberWordMap[three])
			lines[i] = strings.ReplaceAll(lines[i], four, numberWordMap[four])
			lines[i] = strings.ReplaceAll(lines[i], five, numberWordMap[five])
			lines[i] = strings.ReplaceAll(lines[i], six, numberWordMap[six])
			lines[i] = strings.ReplaceAll(lines[i], seven, numberWordMap[seven])
			lines[i] = strings.ReplaceAll(lines[i], eight, numberWordMap[eight])
			lines[i] = strings.ReplaceAll(lines[i], nine, numberWordMap[nine])
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
