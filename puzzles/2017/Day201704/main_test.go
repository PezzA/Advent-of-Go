package Day201704

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)
	Expect(noDupedWord("aa bb cc dd ee")).Should(Equal(true))
	Expect(noDupedWord("aa bb cc dd aa")).Should(Equal(false))
	Expect(noDupedWord("aa bb cc dd aaa")).Should(Equal(true))
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)
	Expect(noAnagram("abcde fghij")).Should(Equal(true))
	Expect(noAnagram("abcde xyz ecdab")).Should(Equal(false))
	Expect(noAnagram("a ab abc abd abf abj")).Should(Equal(true))
	Expect(noAnagram("iiii oiii ooii oooi oooo")).Should(Equal(true))
	Expect(noAnagram("oiii ioii iioi iiio")).Should(Equal(false))
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
