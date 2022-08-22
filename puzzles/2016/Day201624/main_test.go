package Day201624

import (
	"testing"

	. "github.com/onsi/gomega"
	"github.com/pezza/advent-of-code/puzzles/common"
)

func Test_ReadData(t *testing.T) {
	RegisterTestingT(t)

	_, locations := getMaze(`###########
#0.1.....2#
#.#######.#
#4.......3#
###########`, true)

	Expect(locations[0]).To(Equal(common.Point{X: 1, Y: 1}))
	Expect(locations[1]).To(Equal(common.Point{X: 3, Y: 1}))
	Expect(locations[2]).To(Equal(common.Point{X: 9, Y: 1}))
	Expect(locations[3]).To(Equal(common.Point{X: 9, Y: 3}))
	Expect(locations[4]).To(Equal(common.Point{X: 1, Y: 3}))
}

func Test_PartOne(t *testing.T) {

	// get the maze and all the points.  Figure out the distance between
	// each point to each point

	// work out the shortest path to visit all of them
	RegisterTestingT(t)

}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

}

func Benchmark_BenchPartOne(b *testing.B) {
	data := Entry.PuzzleInput()
	for n := 0; n < b.N; n++ {
		Entry.PartOne(data, nil)
	}
}

func Benchmark_BenchPartTwo(b *testing.B) {
	data := Entry.PuzzleInput()
	for n := 0; n < b.N; n++ {
		Entry.PartTwo(data, nil)
	}
}
