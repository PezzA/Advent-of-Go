package Day202204

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_ReadData(t *testing.T) {
	RegisterTestingT(t)

	pairs := getData(Entry.PuzzleInput())

	Expect(pairs[0].first.start).Should(Equal(54))
	Expect(pairs[0].first.end).Should(Equal(59))
	Expect(pairs[0].second.start).Should(Equal(17))
	Expect(pairs[0].second.end).Should(Equal(62))
}

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	Expect(Entry.PartOne(Entry.PuzzleInput(), nil)).Should(Equal("657"))
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

	Expect(Entry.PartTwo(Entry.PuzzleInput(), nil)).Should(Equal("938"))
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
