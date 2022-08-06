package dayEntry

var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string, int) {
	return 2020, 11, "Getting the boilerplate in place"
}

func (td dayEntry) PuzzleInput() string {
	return ``
}

func (td dayEntry) PuzzleSpec() string {
	return ``
}
