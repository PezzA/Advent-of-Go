package main

import (
	"time"

	"github.com/pezza/Aoc2017/cli"
)

type partResult struct {
	answered bool
	result   string
	time     time.Duration
}

func runner(puzzle dailyPuzzle, mode int) {
	inputData := puzzle.PuzzleInput()
	year, day, title := puzzle.Describe()

	part1Answer, part2Answer := partResult{false, "", 0}, partResult{false, "", 0}
	part1Update, part2Update := []string{"*"}, []string{"*"}

	part1Chan := make(chan partResult, 0)
	part2Chan := make(chan partResult, 0)

	part1UpdateChan := make(chan []string, 10)
	part2UpdateChan := make(chan []string, 10)

	heading := cli.GetHeader(year, day, title)

	cli.DrawFrame(part1Answer.result, part1Update, part1Answer.time, part2Answer.result, part2Update, part2Answer.time, heading)

	if mode == 0 || mode == 1 {
		go doPart(puzzle.PartOne, inputData, part1Chan, part1UpdateChan)
	}

	if mode == 0 || mode == 2 {
		go doPart(puzzle.PartTwo, inputData, part2Chan, part2UpdateChan)
	}

	complete := false
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
			cli.DrawFrame(part1Answer.result, part1Update, part1Answer.time, part2Answer.result, part2Update, part2Answer.time, heading)

			if part1Answer.answered && part2Answer.answered {
				complete = true
			}

			time.Sleep(100 * time.Millisecond)
		}
	}
}

func doPart(fn puzzlePart, inputData string, response chan partResult, updateChan chan []string) {
	start := time.Now()
	part1Output, _ := fn(inputData, updateChan)
	end := time.Now()
	elapsed := end.Sub(start)
	response <- partResult{true, part1Output, elapsed}
}
