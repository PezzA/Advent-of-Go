package main

type puzzlePart func(string, chan []string) (string, error)

type dailyPuzzle interface {
	Describe() (int, int, string)
	PartOne(string, chan []string) (string, error)
	PartTwo(string, chan []string) (string, error)
	PuzzleInput() string
}
