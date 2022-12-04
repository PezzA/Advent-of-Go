package Day202205

import (
	"testing"

	. "github.com/onsi/gomega"
	"github.com/pezza/advent-of-code/puzzles/common"
)

func Test_ReadData(t *testing.T) {
	RegisterTestingT(t)

}

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	Expect(Entry.PartOne(Entry.PuzzleInput(), nil)).Should(Equal(common.NOT_IMPLEMENTED))
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

	Expect(Entry.PartTwo(Entry.PuzzleInput(), nil)).Should(Equal(common.NOT_IMPLEMENTED))
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
