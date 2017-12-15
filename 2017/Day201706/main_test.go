package Day201706

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	Expect(registerList{0, 2, 7, 0}.distributeRegister(2)).Should(Equal(registerList{2, 4, 1, 2}))
	Expect(registerList{2, 4, 1, 2}.distributeRegister(1)).Should(Equal(registerList{3, 1, 2, 3}))

	Expect(registerList{3, 9, 9, 6, 5, 3, 4, 2, 6, 6, 2, 3, 16, 14, 5, 0}.getHighestIndex()).Should(Equal(12))

	previousStates := stateList{registerList{0, 0, 0, 0}, registerList{1, 1, 1, 1}}

	Expect(previousStates.isStatePreviouslySeen(registerList{1, 1, 1, 1})).Should(Equal(1))
	Expect(previousStates.isStatePreviouslySeen(registerList{1, 2, 1, 1})).Should(Equal(-1))

	cycles, _ := distributionCycle([]int{0, 2, 7, 0})
	Expect(cycles).Should(Equal(5))
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
