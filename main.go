package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/pezza/advent-of-code/puzzles"
)

func main() {
	// Handle re-showing cursor on Ctrl+C
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		interrupted()
		os.Exit(1)
	}()

	year, day, err := checkParams()

	if err != nil {
		outputUseage(err)
		return
	}

	if puzzle, err := puzzles.GetPuzzle(day, year); err != nil {

		outputUseage(err)

	} else {
		runner(puzzle)
	}
}
