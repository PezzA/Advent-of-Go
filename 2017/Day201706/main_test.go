package Day201706

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)
	Expect(distributeRegister(2, []int{0, 2, 7, 0})).Should(Equal([]int{2, 4, 1, 2}))
	Expect(distributeRegister(1, []int{2, 4, 1, 2})).Should(Equal([]int{3, 1, 2, 3}))

	Expect(getHighestIndex([]int{3, 9, 9, 6, 5, 3, 4, 2, 6, 6, 2, 3, 16, 14, 5, 0})).Should(Equal(12))

	previousStates := [][]int{[]int{0, 0, 0, 0}, []int{1, 1, 1, 1}}
	Expect(isStatePreviouslySeen([]int{1, 1, 1, 1}, previousStates)).Should(Equal(1))
	Expect(isStatePreviouslySeen([]int{1, 2, 1, 1}, previousStates)).Should(Equal(-1))

	_, cycles, _ := distributionCycle([]int{0, 2, 7, 0})
	Expect(cycles).Should(Equal(5))
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

}

func Benchmark_PartOne(b *testing.B) {
	data := Entry.GetData()
	for n := 0; n < b.N; n++ {
		Entry.PartOne(data)
	}
}

func Benchmark_PartTwo(b *testing.B) {
	data := Entry.GetData()
	for n := 0; n < b.N; n++ {
		Entry.PartTwo(data)
	}
}
