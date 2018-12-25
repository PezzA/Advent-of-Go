package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/pezza/advent-of-code/cli"
	"github.com/pezza/advent-of-code/puzzles"
)

func main() {
	// Handle re-showing cursor on Ctrl+C
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		cli.Interrupted()
		os.Exit(1)
	}()

	year, day, err := cli.CheckParams()

	if err != nil {
		cli.OutputUseage(err)
		return
	}

	if puzzle, err := puzzles.GetPuzzle(day, year); err != nil {
		cli.OutputUseage(err)

	} else {
		runner(puzzle)
	}

	return
}
