package Day201505

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

	Expect(hasRepeatNoOverlap("xyxy")).Should(Equal(true))
	Expect(hasRepeatNoOverlap("aaa")).Should(Equal(false))
	Expect(hasRepeatNoOverlap("aaaa")).Should(Equal(true))
	Expect(hasBumpy("aaa")).Should(Equal(true))
	Expect(hasBumpy("xyx")).Should(Equal(true))

	Expect(hasRepeatNoOverlap("qjhvhtzxzqqjkmpb")).Should(Equal(true))

	Expect(hasBumpy("qjhvhtzxzqqjkmpb") && hasRepeatNoOverlap("qjhvhtzxzqqjkmpb")).Should(Equal(true))
}

func Benchmark_BenchPartOne(b *testing.B) {
	data := Entry.PuzzleInput()
	for n := 0; n < b.N; n++ {
		Entry.PartOne(data)
	}
}

func Benchmark_BenchPartTwo(b *testing.B) {
	data := Entry.PuzzleInput()
	for n := 0; n < b.N; n++ {
		Entry.PartTwo(data)
	}
}
