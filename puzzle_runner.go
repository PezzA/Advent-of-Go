package main

import (
	"fmt"
	"time"

	"github.com/pezza/advent-of-code/puzzles"
)

type partResult struct {
	answered bool
	result   string
	time     time.Duration
}

func runner(puzzle puzzles.DailyPuzzle) {
	hideCursor()

	defer showCursor()

	inputData := puzzle.PuzzleInput()
	year, day, title := puzzle.Describe()

	part1Answer, part2Answer := partResult{false, "", 0}, partResult{false, "", 0}
	part1Update, part2Update := []string{"*"}, []string{"*"}

	part1Chan := make(chan partResult)
	part2Chan := make(chan partResult)

	part1UpdateChan := make(chan []string, 10)
	part2UpdateChan := make(chan []string, 10)

	heading := getHeader(year, day, title)

	drawFrame(part1Answer.result, part1Update, part1Answer.time, part2Answer.result, part2Update, part2Answer.time, heading)

	go doPart(puzzle.PartOne, inputData, part1Chan, part1UpdateChan)
	go doPart(puzzle.PartTwo, inputData, part2Chan, part2UpdateChan)

	complete := false
	outputAnswer := ""
	for !complete {
		select {
		case result := <-part1Chan:
			part1Answer = result
		case result := <-part2Chan:
			part2Answer = result
		case update := <-part1UpdateChan:
			part1Update = update
		case update := <-part2UpdateChan:
			part2Update = update
		default:
			newFrame(4 + len(part1Update) + len(part1Update))
			outputAnswer = part1Answer.result

			if year == 2018 && day == 10 {
				outputAnswer = ""
			}
			drawFrame(outputAnswer, part1Update, part1Answer.time, part2Answer.result, part2Update, part2Answer.time, heading)

			if part1Answer.answered && part2Answer.answered {
				complete = true
			}

			time.Sleep(17 * time.Millisecond)
		}
	}

	if year == 2018 && day == 10 {
		fmt.Println(part1Answer.result)
	}
}

func doPart(fn puzzles.PuzzlePart, inputData string, response chan partResult, updateChan chan []string) {
	start := time.Now()
	output := fn(inputData, updateChan)
	end := time.Now()
	elapsed := end.Sub(start)
	response <- partResult{true, output, elapsed}
}
