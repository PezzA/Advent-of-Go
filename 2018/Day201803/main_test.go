package Day201803

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	claims := getData(`#1 @ 1,3: 4x4
	#2 @ 3,1: 4x4
	#3 @ 5,5: 2x2`)

	//claims := getData(Entry.PuzzleInput())
	claimed, tags := applyClaims(claims)

	overlaps := countOverlaps(claimed)

	nonOverlapped := getNonOverlappedClaim(tags, claims)

	Expect(overlaps).Should(Equal(4))
	Expect(nonOverlapped).Should(Equal("#3"))
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
