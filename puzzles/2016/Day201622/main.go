package Day201622

import (
	"fmt"
	"regexp"
)

// Entry holds wraps the data and runner interfaces for this puzzle
var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string, int) {
	return 2016, 22, "Grid Computing", 6
}

var regex = "/dev/grid/node-x([0-9]+)-y([0-9]+)     ([0-9]+)T   ([0-9]+)T    ([0-9]+)T   ([0-9]+)%"
var instructionRegex = regexp.MustCompile(regex)

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	return fmt.Sprintf(" -- Not Yet Implemented --")
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	return fmt.Sprintf(" -- Not Yet Implemented --")
}
