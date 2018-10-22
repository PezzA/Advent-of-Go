package Day201518

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
	common "github.com/pezza/advent-of-code/Common"
)

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	screen := getData(Entry.PuzzleInput())

	fmt.Println(screen.getNeighbourCount(common.Point{0, 0}))
	fmt.Println(screen.getNeighbourCount(common.Point{1, 1}))
	fmt.Println(screen.getNeighbourCount(common.Point{2, 2}))
	fmt.Println(screen.getNeighbourCount(common.Point{3, 3}))

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
