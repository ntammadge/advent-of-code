package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	dayPointer := flag.Int("day", -1, "which day's puzzle to solve")
	partPointer := flag.Int("part", 1, "which part of the day's puzzle is solved")
	inputPointer := flag.String("input", "", "file to pull input data from")
	flag.Parse()

	if *dayPointer < 1 || *dayPointer > 25 {
		panic("Day must be provided and from 1 to 25 (inclusive)")
	}
	if *partPointer != 1 && *partPointer != 2 {
		panic("Part must be 1 (default) or 2")
	}
	if *inputPointer == "" {
		panic("Input file required (/input=filePath)")
	}

	solutionFunctions := map[string]func(inputData string, part int){
		"day1": day1,
		"day2": day2,
	}

	fileBytes, err := os.ReadFile(*inputPointer)
	if err != nil {
		panic("ah")
	}

	fileContent := string(fileBytes)

	solver := solutionFunctions["day"+fmt.Sprint(*dayPointer)]
	solver(fileContent, *partPointer)
}
