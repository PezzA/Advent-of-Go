package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"flag"

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

	web := flag.Bool("web", false, "run AoC webserver")
	flag.Parse()
	if *web {
		log.Fatal(http.ListenAndServe(":8000", http.FileServer(http.Dir("..\\advent-of-wasm\\Content"))))
	}

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
