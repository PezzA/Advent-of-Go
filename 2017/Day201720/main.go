package Day201720

type dayEntry bool

var Entry dayEntry

func (td dayEntry) Describe() (int, int, string) {
	return 2017, 20, "Getting the boilerplate in place"
}

func (td dayEntry) PartOne(inputData string) (string, error) {
	return " -- Not Yet Implemented --", nil
}

func (td dayEntry) PartTwo(inputData string) (string, error) {
	return " -- Not Yet Implemented --", nil
}

func (td dayEntry) Day() int {
	return 20
}

func (td dayEntry) Year() int {
	return 2017
}

func (td dayEntry) PuzzleInput() string {
	return "Actual Data"
}
