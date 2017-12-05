package Day201705

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	Expect(doJumps([]int{0, 3, 0, 1, -3})).Should(Equal(5))
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

	Expect(doModifiedJumps([]int{0, 3, 0, 1, -3})).Should(Equal(10))
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
