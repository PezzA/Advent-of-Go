package Day201722

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	grid, _ := getGrid(`..#
#..
...`)

	var infectedCells int
	grid, infectedCells = doWalk(70, grid)
	Expect(infectedCells).Should(Equal(41))

}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)
	grid, _ := getGrid(`..#
#..
...`)

	var infectedCells int
	printGrid(grid, 12, point{0, 0})
	grid, infectedCells = doResistiveWalk(100, grid)
	Expect(infectedCells).Should(Equal(26))
}

func Benchmark_PartOne(b *testing.B) {
	data := Entry.PuzzleInput()
	for n := 0; n < b.N; n++ {
		Entry.PartOne(data, nil)
	}
}

func Benchmark_PartTwo(b *testing.B) {
	data := Entry.PuzzleInput()
	for n := 0; n < b.N; n++ {
		Entry.PartTwo(data, nil)
	}
}
