package main

type dailyPuzzle interface {
	Describe() (int, int, string)
	PartOne(string) (string, error)
	PartTwo(string) (string, error)
	PuzzleInput() string
}
