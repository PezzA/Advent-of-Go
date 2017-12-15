package Day201611

func PuzzleInput() ([]string, []string) {
	return []string{`The first floor contains a hydrogen-compatible microchip and a lithium-compatible microchip.
	The second floor contains a hydrogen generator.
	The third floor contains a lithium generator.
	The fourth floor contains nothing relevant.`}, []string{}
}

// PuzzleInput returns the main puzzle input for day 11
func (d dayEntry) PuzzleInput() string {
	return `The first floor contains a promethium generator and a promethium-compatible microchip.
The second floor contains a cobalt generator, a curium generator, a ruthenium generator, and a plutonium generator.
The third floor contains a cobalt-compatible microchip, a curium-compatible microchip, a ruthenium-compatible microchip, and a plutonium-compatible microchip.
The fourth floor contains nothing relevant.`
}
