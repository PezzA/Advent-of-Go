package Day202025

var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string, int) {
	return 2020, 25, "Combo Breaker"
}

func (td dayEntry) PuzzleInput() string {
	return `1717001
523731`
}

func (td dayEntry) PuzzleSpec() string {
	return ``
}
