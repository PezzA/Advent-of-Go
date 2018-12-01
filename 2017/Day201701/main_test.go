package Day201701

import (
	"testing"

	. "github.com/onsi/gomega"
	common "github.com/pezza/AoC2017/Common"
)

func Test_Stuff(t *testing.T) {
	RegisterTestingT(t)
	Expect(common.GetWrappedIndex(0, 10, 0)).Should(Equal(0))
	Expect(common.GetWrappedIndex(0, 10, 10)).Should(Equal(0))
	Expect(common.GetWrappedIndex(0, 10, 1)).Should(Equal(1))
	Expect(common.GetWrappedIndex(0, 10, 15)).Should(Equal(5))
}

func Benchmark_PartOne(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Entry.PartOne(Entry.PuzzleInput())
	}
}

func Benchmark_PartTwo(b *testing.B) {
	data := Entry.PuzzleInput()
	for n := 0; n < b.N; n++ {
		Entry.PartTwo(data)
	}
}
