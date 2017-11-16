package main

import "fmt"

func testRunner(data dayData, runner dayRunner) {
	day1Tests, day2Tests := data.GetTestData()

	fmt.Println(runner.Heading(), "** TEST MODE **")

	fmt.Println("Part 1:")
	for index, test := range day1Tests {
		testOutput, testError := runner.PartOne(test)

		if testError != nil {
			fmt.Println("Error in test", index, testError)
			continue
		}

		fmt.Println(fmt.Sprintf("[%v] %v -> %v", index, test, testOutput))
	}

	fmt.Println("Part 2:")
	for index, test := range day2Tests {
		testOutput, testError := runner.PartTwo(test)

		if testError != nil {
			fmt.Println("Error in test", index, testError)
			continue
		}

		fmt.Println(fmt.Sprintf("[%v] %v -> %v", index, test, testOutput))
	}
}

func runner(data dayData, runner dayRunner) {
	inputData := data.GetData()

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
