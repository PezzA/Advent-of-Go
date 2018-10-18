package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/pezza/advent-of-code/cli"
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

	runner(puzzle)
}
