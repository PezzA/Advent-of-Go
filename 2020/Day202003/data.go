package Day202003

var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string) {
	return 2020, 3, "Getting the boilerplate in place"
}

func (td dayEntry) PuzzleInput() string {
	return ``
}
