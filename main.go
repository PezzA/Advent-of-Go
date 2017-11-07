package main

import (
	"fmt"
	"os"
	"strconv"
)

func outputUseage() {
	fmt.Println("	USEAGE: advent-of-go <day>")
	fmt.Println("		<day> = day number of puzzle to run.  e.g. 2")
	fmt.Println("")
	fmt.Println("	Flags")
	fmt.Println("		--test : run the day with test data")
}

func main() {
	data := [...]dayData{testDayEntry}
	runners := [...]dayRunner{testDayEntry}

	if len(os.Args) < 2 {
		fmt.Println("Not all parameters specified")
		fmt.Println("")
		outputUseage()
		return
	}

	day, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Println("Called with Invalid parameters")
		fmt.Println("")
		outputUseage()
		return
	}

	runTest := false
	for _, val := range os.Args {
		if val == "--test" {
			runTest = true
		}
	}

	dayFound := false
	for _, runner := range runners {
		if runner.Day() == day {
			dayFound = true
		}
	}

	if !dayFound {
		fmt.Println("Day specified has not been implemented yet")
		return
	}

	runner(data[day], runners[day], runTest)
}

func runner(data dayData, runner dayRunner, runTest bool) {
	inputData := data.GetData()

	if runTest {
		inputData = data.GetTestData()
	}

	fmt.Println(runner.Heading())
	part1Output, part1Err := runner.PartOne(inputData)

	if part1Err != nil {
		fmt.Println("    Part 1: ERROR", part1Err)
		return
	}

	fmt.Println("    Part 1:", part1Output)

	part2Output, part2Err := runner.PartTwo(inputData)

	if part2Err != nil {
		fmt.Println("    Part 2: ERROR", part2Err)
		return
	}

	fmt.Println("    Part 2:", part2Output)
}
