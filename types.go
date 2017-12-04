package main

type dayRunner interface {
	Day() int
	Year() int
	Title() string
	PartOne(string) (string, error)
	PartTwo(string) (string, error)
	GetData() string
}
