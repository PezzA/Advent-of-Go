package Day201510

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)
	Expect(lookSay("1")).Should(Equal("11"))
	Expect(lookSay("11")).Should(Equal("21"))
	Expect(lookSay("21")).Should(Equal("1211"))
	Expect(lookSay("1211")).Should(Equal("111221"))
	Expect(lookSay("111221")).Should(Equal("312211"))
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
