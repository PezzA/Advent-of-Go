package dayEntry

var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string) {
	return 2019, 11, "Getting the boilerplate in place"
}

func (td dayEntry) PuzzleInput() string {
	return ``
}
