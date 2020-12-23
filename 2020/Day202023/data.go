package Day202023

var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string) {
	return 2020, 23, "Crab Cups"
}

func (td dayEntry) PuzzleInput() string {
	return `685974213`
}

func (td dayEntry) PuzzleSpec() string {
	return ``
}
