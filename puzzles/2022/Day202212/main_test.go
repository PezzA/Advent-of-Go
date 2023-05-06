package Day202212

import (
	"testing"

	. "github.com/onsi/gomega"
	"github.com/pezza/advent-of-code/puzzles/common"
)

func Test_ReadData(t *testing.T) {
	RegisterTestingT(t)

	vista, start, end := getData(Entry.PuzzleInput())

	Expect(vista[common.Point{0, 0}]).Should(Equal(int32(97)))
	Expect(start).Should(Equal(common.Point{0, 20}))
	Expect(end).Should(Equal(common.Point{88, 20}))

}

func Test_PartOne(t *testing.T) {
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
