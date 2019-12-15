package Day201912

var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string) {
	return 2019, 12, "The N-Body Problem"
}

func (td dayEntry) PuzzleInput() string {
	return `<x=-16, y=15, z=-9>
<x=-14, y=5, z=4>
<x=2, y=0, z=6>
<x=-3, y=18, z=9>`
}
