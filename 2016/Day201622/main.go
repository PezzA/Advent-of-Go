package Day201622

import "regexp"

// Entry holds wraps the data and runner interfaces for this puzzle
var Entry testDay

type testDay bool

var regex = "/dev/grid/node-x([0-9]+)-y([0-9]+)     ([0-9]+)T   ([0-9]+)T    ([0-9]+)T   ([0-9]+)%"
var instructionRegex = regexp.MustCompile(regex)

func (td testDay) PartOne(inputData string) (string, error) {
	return " -- Not Yet Implemented --", nil
}

func (td testDay) PartTwo(inputData string) (string, error) {
	return " -- Not Yet Implemented --", nil
}

func (td testDay) Day() int {
	return 22
}

func (td testDay) Year() int {
	return 2016
}

func (td testDay) Title() string {
	return "Grid Computing"
}
