package Day201716

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	pl := progList("abcde")

	pl = pl.spin(1)
	Expect(pl).Should(Equal(progList("eabcd")))
	pl = pl.exchangePos(3, 4)
	Expect(pl).Should(Equal(progList("eabdc")))
	pl = pl.exchangeProg("e", "b")
	Expect(pl).Should(Equal(progList("baedc")))

	pl = progList("abcde")
	Expect(pl.doDance(getInstructions("s1,x3/4,pe/b"))).Should(Equal(progList("baedc")))

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
