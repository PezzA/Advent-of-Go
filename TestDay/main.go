package dayEntry

// Entry holds wraps the data and runner interfaces for this puzzle
var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string) {
	return 0, 0, "Getting the boilerplate in place"
}
func (td dayEntry) PartOne(inputData string) (string, error) {
	return " -- Not Yet Implemented --", nil
}

func (td dayEntry) PartTwo(inputData string) (string, error) {
	return " -- Not Yet Implemented --", nil
}

func (td dayEntry) PuzzleInput() string {
	return "Actual Data"
}
