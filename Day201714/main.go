package Day201714

// Entry holds wraps the data and runner interfaces for this puzzle
var Entry testDay

type testDay bool

func (td testDay) PartOne(inputData string) (string, error) {
	return " -- Not Yet Implemented --", nil
}

func (td testDay) PartTwo(inputData string) (string, error) {
	return " -- Not Yet Implemented --", nil
}

func (td testDay) Day() int {
	return 14
}

func (td testDay) Year() int {
	return 2017
}

func (td testDay) Title() string {
	return "Getting the boilerplate in place"
}

func (td testDay) GetData() string {
	return "Actual Data"
}
