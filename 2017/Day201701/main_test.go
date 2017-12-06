package Day201701

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_Stuff(t *testing.T) {
	RegisterTestingT(t)
	Expect(getWrappedIndex(0, 10, 0)).Should(Equal(0))
	Expect(getWrappedIndex(0, 10, 10)).Should(Equal(0))
	Expect(getWrappedIndex(0, 10, 1)).Should(Equal(1))
	Expect(getWrappedIndex(0, 10, 15)).Should(Equal(5))
}

func Benchmark_PartOne(b *testing.B) {
	data := Entry.GetData()
	for n := 0; n < b.N; n++ {
		Entry.PartOne(data)
	}
}

func Benchmark_PartTwo(b *testing.B) {
	data := Entry.GetData()
	for n := 0; n < b.N; n++ {
		Entry.PartTwo(data)
	}
}
