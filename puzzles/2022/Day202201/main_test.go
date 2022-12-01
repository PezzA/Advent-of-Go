package Day202201

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_ReadData(t *testing.T) {
	RegisterTestingT(t)

	elves := loadData(Entry.PuzzleInput())

	Expect(len(elves[0])).Should(Equal(15))
	Expect(elves[0][0]).Should(Equal(5229))
	Expect(elves[0][14]).Should(Equal(4504))

	Expect(elves[len(elves)-1][0]).Should(Equal(6500))
	Expect(elves[len(elves)-1][6]).Should(Equal(10283))
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
