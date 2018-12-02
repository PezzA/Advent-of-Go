package Day201802

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	counts := getCounts("abcdef")

	Expect(hasSameLetterCount(2, counts)).Should(Equal(false))
	Expect(hasSameLetterCount(3, counts)).Should(Equal(false))

	counts = getCounts("bababc")
	Expect(hasSameLetterCount(2, counts)).Should(Equal(true))
	Expect(hasSameLetterCount(3, counts)).Should(Equal(true))

	counts = getCounts("abcccd")
	Expect(hasSameLetterCount(2, counts)).Should(Equal(false))
	Expect(hasSameLetterCount(3, counts)).Should(Equal(true))

	counts = getCounts("abcdee")
	Expect(hasSameLetterCount(2, counts)).Should(Equal(true))
	Expect(hasSameLetterCount(3, counts)).Should(Equal(false))
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
