package Day201916

import (
	"fmt"
	"strings"
	"testing"

	. "github.com/onsi/gomega"
)

func Test_ReadData(t *testing.T) {
	RegisterTestingT(t)

	data := getData(Entry.PuzzleInput())

	Expect(data[0]).Should(Equal(5))
	Expect(data[len(data)-1]).Should(Equal(1))
}

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)
	Expect(basePattern(1)).Should(Equal([]int{0, 1, 0, -1}))
	Expect(basePattern(2)).Should(Equal([]int{0, 0, 1, 1, 0, 0, -1, -1}))

	result := applyPhaseShift(getData("12345678"))
	Expect(result).Should(Equal([]int{4, 8, 2, 2, 6, 1, 5, 8}))

	result = applyPhaseShift(result)
	Expect(result).Should(Equal([]int{3, 4, 0, 4, 0, 4, 3, 8}))

	seq := getData("80871224585914546619083218645595")
	for i := 0; i < 100; i++ {
		seq = applyPhaseShift(seq)
	}
	Expect(seq[:8]).Should(Equal([]int{2, 4, 1, 7, 6, 1, 7, 6}))

	p1test := getData(Entry.PuzzleInput())
	for i := 0; i < 100; i++ {
		p1test = applyPhaseShift(p1test)
	}

	fmt.Println(p1test[:8])
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

	p1test := getData(strings.Repeat(Entry.PuzzleInput(), 10000))
	for i := 0; i < 100; i++ {
		fmt.Println(i)
		p1test = applyPhaseShift(p1test)
	}

	fmt.Println(p1test[:8])
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
