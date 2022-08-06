package Day202015

var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string, int) {
	return 2020, 15, "Rambunctious Recitation", 2
}

func (td dayEntry) PuzzleInput() string {
	return `13,0,10,12,1,5,8`
}

func (td dayEntry) PuzzleSpec() string {
	return ``
}
