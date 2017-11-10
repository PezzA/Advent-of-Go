package main

import "fmt"

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
