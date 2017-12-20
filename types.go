package main

type puzzlePart func(string, chan []string) string

type dailyPuzzle interface {
	Describe() (int, int, string)
	PartOne(string, chan []string) string
	PartTwo(string, chan []string) string
	PuzzleInput() string
}
