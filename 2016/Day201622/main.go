package Day201622

import "regexp"

// Entry holds wraps the data and runner interfaces for this puzzle
var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string) {
	return 2016, 22, "Grid Computing"
}

var regex = "/dev/grid/node-x([0-9]+)-y([0-9]+)     ([0-9]+)T   ([0-9]+)T    ([0-9]+)T   ([0-9]+)%"
var instructionRegex = regexp.MustCompile(regex)

func (td dayEntry) PartOne(inputData string) (string, error) {
	return " -- Not Yet Implemented --", nil
}

func (td dayEntry) PartTwo(inputData string) (string, error) {
	return " -- Not Yet Implemented --", nil
}
