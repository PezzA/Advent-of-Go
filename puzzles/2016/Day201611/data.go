package Day201611

// Entry holds wraps the data and runner interfaces for this puzzle
var Entry dayEntry

type dayEntry bool

func (d dayEntry) Describe() (int, int, string, int) {
	return 2016, 11, "Radioisotope Thermoelectric Generators", 5
}

// PuzzleInput returns the main puzzle input for day 11
func (d dayEntry) PuzzleInput() string {
	return `The first floor contains a promethium generator and a promethium-compatible microchip.
The second floor contains a cobalt generator, a curium generator, a ruthenium generator, and a plutonium generator.
The third floor contains a cobalt-compatible microchip, a curium-compatible microchip, a ruthenium-compatible microchip, and a plutonium-compatible microchip.
The fourth floor contains nothing relevant.`
}

/*
How to solve

from the start position

get a list of the safe


*/
