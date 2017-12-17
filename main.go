package main

import "github.com/pezza/Aoc2017/cli"

func main() {

	year, day, error := cli.CheckParams()

	if error != nil {
		cli.OutputUseage(error)
		return
	}

	puzzle, error := getPuzzle(day, year)

	if error != nil {
		cli.OutputUseage(error)
		return
	}

	runner(puzzle, 0)
}
