package Day201511

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	Expect(increment("aaz")).Should(Equal("aba"))
	Expect(increment("")).Should(Equal("a"))
	Expect(increment("a")).Should(Equal("b"))
	Expect(increment("z")).Should(Equal("aa"))
	Expect(increment("azzz")).Should(Equal("baaa"))
	Expect(increment("vzbxkghb")).Should(Equal("vzbxkghc"))

	Expect(validChars("john")).Should(Equal(false))
	Expect(validChars("jhn")).Should(Equal(true))
	Expect(validChars("jhnl")).Should(Equal(false))
	Expect(validChars("jhni")).Should(Equal(false))

	Expect(hasTwoPairs("jhni")).Should(Equal(false))
	Expect(hasTwoPairs("jhniaa")).Should(Equal(false))
	Expect(hasTwoPairs("jhniaaa")).Should(Equal(false))
	Expect(hasTwoPairs("jhniaaaa")).Should(Equal(false))
	Expect(hasTwoPairs("jhniaaaaa")).Should(Equal(false))
	Expect(hasTwoPairs("jhniaabb")).Should(Equal(true))
	Expect(hasTwoPairs("bbjhnibb")).Should(Equal(false))
	Expect(hasTwoPairs("zzjhnibb")).Should(Equal(true))

	Expect(hasStraight("")).Should(Equal(false))
	Expect(hasStraight("ab")).Should(Equal(false))
	Expect(hasStraight("abc")).Should(Equal(true))
	Expect(hasStraight("xxabcx")).Should(Equal(true))
	Expect(hasStraight("aaxyz")).Should(Equal(true))
	Expect(hasStraight("aaacd")).Should(Equal(false))
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
