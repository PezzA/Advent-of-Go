package Day201525

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	Expect(fillMatrix(10)).Should(Equal(30943339))
	Expect(getFirstColIndex(4)).Should(Equal(10))
	Expect(getColVal(4, 1)).Should(Equal(10))
	Expect(getColVal(4, 3)).Should(Equal(19))
	Expect(getColVal(6, 1)).Should(Equal(21))

	Expect(fillMatrix(getColVal(4, 1))).Should(Equal(30943339))
	Expect(fillMatrix(getColVal(4, 2))).Should(Equal(7726640))
	Expect(fillMatrix(getColVal(4, 3))).Should(Equal(7981243))
	Expect(fillMatrix(getColVal(4, 4))).Should(Equal(9380097))
	Expect(fillMatrix(getColVal(4, 5))).Should(Equal(6899651))
	Expect(fillMatrix(getColVal(4, 6))).Should(Equal(24659492))
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
