package Day201716

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	Expect(spin("abcde", 1)).Should(Equal("eabcd"))
	Expect(exchangePos("eabcd", 3, 4)).Should(Equal("eabdc"))
	Expect(exchangeProg("e", "b", "eabdc")).Should(Equal("baedc"))
	Expect(doDance("abcde", getInstructions("s1,x3/4,pe/b"))).Should(Equal("baedc"))

}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

}

func Benchmark_PartOne(b *testing.B) {
	data := Entry.PuzzleInput()
	for n := 0; n < b.N; n++ {
		Entry.PartOne(data)
	}
}

func Benchmark_PartTwo(b *testing.B) {
	data := Entry.PuzzleInput()
	for n := 0; n < b.N; n++ {
		Entry.PartTwo(data)
	}
}
