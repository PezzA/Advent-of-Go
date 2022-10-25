package Day201713

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)
	Expect(getPositionForTime(0, 4)).Should(Equal(0))
	Expect(getPositionForTime(1, 4)).Should(Equal(1))
	Expect(getPositionForTime(2, 4)).Should(Equal(2))
	Expect(getPositionForTime(3, 4)).Should(Equal(3))
	Expect(getPositionForTime(4, 4)).Should(Equal(2))
	Expect(getPositionForTime(5, 4)).Should(Equal(1))
	Expect(getPositionForTime(6, 4)).Should(Equal(0))
	Expect(getPositionForTime(7, 4)).Should(Equal(1))
	Expect(getPositionForTime(8, 4)).Should(Equal(2))
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)
}

func Benchmark_PartOne(b *testing.B) {
	data := Entry.PuzzleInput()
	for n := 0; n < b.N; n++ {
		Entry.PartOne(data, nil)
	}
}

func Benchmark_PartTwo(b *testing.B) {
	data := Entry.PuzzleInput()
	for n := 0; n < b.N; n++ {
		Entry.PartTwo(data, nil)
	}
}
