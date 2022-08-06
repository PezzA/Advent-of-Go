package Day202112

import "fmt"

var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string, int) {
	return 2021, 12, "Getting the boilerplate in place", 3
}

func (td dayEntry) PuzzleInput() string {
	return ``
}

func (td dayEntry) PuzzleSpec() string {
	return ``
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	return fmt.Sprintf("%v", " -- Not Yet Implemented --")
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	return fmt.Sprintf("%v", " -- Not Yet Implemented --")
}
