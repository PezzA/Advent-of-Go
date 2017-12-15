package Day201715

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	Expect(nextFactor(16807, 65, 0)).Should(Equal(int64(1092455)))

	fmt.Println(lowest16bits(1092455))

}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)
	Expect(nextFactor(16807, 65, 4)).Should(Equal(int64(1352636452)))
	Expect(nextFactor(48271, 8921, 8)).Should(Equal(int64(1233683848)))
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
