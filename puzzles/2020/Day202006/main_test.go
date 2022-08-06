package Day202006

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

func Test_ReadData(t *testing.T) {
	RegisterTestingT(t)

	data := getData(Entry.PuzzleInput())

	Expect(len(data[0])).Should(Equal(5))

	Expect(len(data[len(data)-1])).Should(Equal(4))
}

func Test_DistinctYes(t *testing.T) {
	RegisterTestingT(t)

	data := getData(`abc

a
b
c

ab
ac

a
a
a
a

b`)

	for _, g := range data {
		fmt.Println(g.getDistinctYes())
	}
}

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	total := 0

	for _, g := range getData(Entry.PuzzleInput()) {
		total += g.getDistinctYes()
	}

	fmt.Println(total)
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

	total := 0

	for _, g := range getData(Entry.PuzzleInput()) {
		total += g.getCollectiveYes()
	}

	fmt.Println(total)
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
