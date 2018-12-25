package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
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

	web := flag.Bool("web", false, "Run Aoc Web-Server")
	year := flag.Int("year", 0, "Year of puzzle")
	day := flag.Int("day", 0, "Day of puzzle")
	flag.Parse()

	if *web {
		runWeb(*year, *day)
		return
	}

	runCli(*year, *day)
}

func runCli(year int, day int) {
	if puzzle, err := getPuzzle(day, year); err != nil {
		cli.OutputUseage(err)
		return
	} else {
		runner(puzzle)
	}
}

func runWeb(year int, day int) {
	http.HandleFunc("/", handler)
	http.HandleFunc("/puzzles", puzzleHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func puzzleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(`<ul>`)
	for _, puzzle := range puzzles {
		year, day, name := puzzle.Describe()

	}
	fmt.Println(`</ul>`)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
