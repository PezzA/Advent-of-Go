package main

import "fmt"

func runner(puzzle dailyPuzzle, mode int) {
	inputData := puzzle.PuzzleInput()

	year, day, title := puzzle.Describe()
	fmt.Println(fmt.Sprintf("--- (%v) Day %v : %v ---", year, day, title))

	if mode == 0 || mode == 1 {
		part1Output, part1Err := puzzle.PartOne(inputData)

		if part1Err != nil {
			fmt.Println("    Part 1: ERROR", part1Err)
			return
		}

		fmt.Println("    Part 1:", part1Output)
	}

	if mode == 0 || mode == 2 {
		part2Output, part2Err := puzzle.PartTwo(inputData)

		if part2Err != nil {
			fmt.Println("    Part 2: ERROR", part2Err)
			return
		}
		fmt.Println("    Part 2:", part2Output)
	}
}
