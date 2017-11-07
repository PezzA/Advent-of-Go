package main

var testDayEntry testDay

type testDay bool

func (td testDay) PartOne(inputData string) (string, error) {
	return "Part One running " + inputData, nil
}

func (td testDay) PartTwo(inputData string) (string, error) {
	return "Part Two running " + inputData, nil
}

func (td testDay) Day() int {
	return 0
}

func (td testDay) Heading() string {
	return "Test Day : Getting the boilerplate in place"
}

func (td testDay) GetTestData() string {
	return "Test Data"
}

func (td testDay) GetData() string {
	return "Actual Data"
}
