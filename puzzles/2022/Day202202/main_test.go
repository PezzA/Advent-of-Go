package Day202202

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_ReadData(t *testing.T) {
	RegisterTestingT(t)

	data := importData(Entry.PuzzleInput())

	Expect(data[0]).Should(Equal(newRound("A", "Y")))

	Expect(data[len(data)-1]).Should(Equal(newRound("A", "Y")))
}

func Test_Scoring(t *testing.T) {
	RegisterTestingT(t)
	Expect(scoreRound(newRound("A", "Y"))).Should(Equal(8))
	Expect(scoreRound(newRound("B", "X"))).Should(Equal(1))
	Expect(scoreRound(newRound("C", "Z"))).Should(Equal(6))
}

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	Expect(Entry.PartOne(Entry.PuzzleInput(), nil)).Should(Equal("15632"))
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

	Expect(Entry.PartTwo(Entry.PuzzleInput(), nil)).Should(Equal("14416"))
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
