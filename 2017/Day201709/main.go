package Day201709

// Entry holds wraps the data and runner interfaces for this puzzle
var Entry testDay

type testDay bool

type node struct {
	score int
	nodes []node
}

func (td testDay) PartOne(inputData string) (string, error) {
	return " -- Not Yet Implemented --", nil
}

func (td testDay) PartTwo(inputData string) (string, error) {
	return " -- Not Yet Implemented --", nil
}

func (td testDay) Day() int {
	return 9
}

func (td testDay) Year() int {
	return 2017
}

func (td testDay) Title() string {
	return "Stream Processing"
}
