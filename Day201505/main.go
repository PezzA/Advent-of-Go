package Day201505

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
	return 201505
}

func (td testDay) Heading() string {
	return "--- (2015) Day 5: Doesn't He Have Intern-Elves For This? ---"
}

func (td testDay) GetTestData() ([]string, []string) {
	return []string{
			"ugknbfddgicrmopn",
			"aaa",
			"jchzalrnumimnmhp",
			"haegwjzuvuyypxyu",
			"dvszwmarrgswjxmb",
		},
		[]string{}
}
