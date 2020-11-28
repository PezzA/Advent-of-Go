package main

type PuzzlePart func(string, chan []string) string

type Contenter interface {
	Describe() (int, int, string)
	PuzzleInput() string
}

type DailyPuzzle interface {
	Contenter
	PartOne(string, chan []string) string
	PartTwo(string, chan []string) string
}

type Pager interface {
	Page() int
	Start() string
	Prev() string
	Next() string
	End() string
}
