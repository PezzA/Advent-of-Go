package main

import (
	"flag"
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

	visualisePtr := flag.Bool("vis", false, "Perform visualisation for this puzzle (if supported)")
	flag.Parse()

	year, day, error := cli.CheckParams(flag.Args())

	if error != nil {
		cli.OutputUseage(error)
		return
	}

	if !*visualisePtr {
		if puzzle, err := getPuzzle(day, year); err != nil {
			cli.OutputUseage(error)

		} else {
			runner(puzzle)
		}

		return
	}

	if puzzle, err := getVisualiser(day, year); err != nil {
		cli.OutputUseage(error)
		return
	} else {
		visualise(puzzle)
	}
}
