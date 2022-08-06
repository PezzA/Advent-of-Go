package Day201522

var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string) {
	return 2015, 22, "Wizard Simulator 20XX"
}

func (td dayEntry) PuzzleInput() string {
	return `Hit Points: 58
Damage: 9`
}
