package Day201904

var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string, int) {
	return 2019, 4, "Secure Container", 0
}

func (td dayEntry) PuzzleInput() string {
	return `271973-785961`
}
