package main

import (
	"fmt"
	"time"

	tm "github.com/buger/goterm"
)

func runner(puzzle dailyPuzzle, mode int) {
	inputData := puzzle.PuzzleInput()
	year, day, title := puzzle.Describe()

	part1Answer, part2Answer := "", ""

	part1Chan := make(chan string, 0)
	part2Chan := make(chan string, 0)

	if mode == 0 || mode == 1 {
		go doPartOne(puzzle, inputData, part1Chan)
	}

	if mode == 0 || mode == 2 {
		go doPartTwo(puzzle, inputData, part2Chan)
	}

	for part1Answer != "" && part2Answer != "" {
		select {
		case result := <-part1Chan:
			part1Answer = result
		case result := <-part2Chan:
			part2Answer = result
		default:
			tm.Println(fmt.Sprintf("--- %v Day %v : %v ---", tm.Color(fmt.Sprintf("%v", year), tm.YELLOW), day, tm.Bold(tm.Color(title, tm.WHITE))))

			if part1Answer == "" {
				tm.Println("    Part 1 : *")
			} else {
				tm.Println("    Part 1 :", part1Answer)
			}

			if part2Answer == "" {
				tm.Println("    Part 2 : *")
			} else {
				tm.Println("    Part 2 :", part2Answer)
			}
			tm.Flush()
		}
	}
}

func doPartOne(puzzle dailyPuzzle, inputData string, response chan string) {
	part1Output, _ := puzzle.PartOne(inputData)
	response <- part1Output
}

func doPartTwo(puzzle dailyPuzzle, inputData string, response chan string) {
	part2Output, _ := puzzle.PartTwo(inputData)
	response <- part2Output
}

func printTime() {
	for {
		tm.MoveCursorUp(0)
		tm.ResetLine("")
		tm.Print("Current Time:", time.Now().Format(time.RFC1123))

		tm.Flush() // Call it every time at the end of rendering
		time.Sleep(time.Second)
	}
}
