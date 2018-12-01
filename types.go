package main

type puzzlePart func(string, chan []string) string

type contenter interface {
	Describe() (int, int, string)
	PuzzleInput() string
}

type dailyPuzzle interface {
	contenter
	PartOne(string, chan []string) string
	PartTwo(string, chan []string) string
}

type visualiser interface {
	contenter
	Visualise(chan<- []string)
}
