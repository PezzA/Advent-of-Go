package Day202017

var Entry dayEntry

func (td dayEntry) Describe() (int, int, string) {
	return 2020, 17, "Conway Cubes"
}

func (td dayEntry) PuzzleInput() string {
	return `.##...#.
.#.###..
..##.#.#
##...#.#
#..#...#
#..###..
.##.####
..#####.`
}

func (td dayEntry) PuzzleSpec() string {
	return ``
}
