package Day201924

var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string, int) {
	return 2019, 24, "Planet of Discord"
}

func (td dayEntry) PuzzleInput() string {
	return `#..##
#.#..
#...#
##..#
#..##`
}
