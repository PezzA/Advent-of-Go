package TestDay

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
	return 0
}

func (td testDay) Heading() string {
	return "--- Test Day : Getting the boilerplate in place ---"
}

func (td testDay) GetTestData() ([]string, []string) {
	return []string{"Day 1 test case 1", "Day 1 test case 2"},
		[]string{"Day 2 test case 1", "Day 2 test case 2"}
}

func (td testDay) GetData() string {
	return "Actual Data"
}
