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
	return 201622
}

func (td testDay) Heading() string {
	return "--- (2016) Day 22: Grid Computing ---"
}

func (td testDay) GetTestData() ([]string, []string) {
	return []string{},
		[]string{}
}
