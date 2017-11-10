package main

type dayData interface {
	Day() int
	GetTestData() string
	GetData() string
}

type dayRunner interface {
	Day() int
	Heading() string
	PartOne(string) (string, error)
	PartTwo(string) (string, error)
}
