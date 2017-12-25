package Day201725

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	state, steps, commands := getTuringData(`Begin in state A.
Perform a diagnostic checksum after 6 steps.

In state A:
	If the current value is 0:
	- Write the value 1.
	- Move one slot to the right.
	- Continue with state B.
	If the current value is 1:
	- Write the value 0.
	- Move one slot to the left.
	- Continue with state B.

In state B:
	If the current value is 0:
	- Write the value 1.
	- Move one slot to the left.
	- Continue with state A.
	If the current value is 1:
	- Write the value 1.
	- Move one slot to the right.
	- Continue with state A.`)

	Expect(state).Should(Equal("A"))
	Expect(steps).Should(Equal(6))
	Expect(commands["A"]).Should(Equal(instruction{1, "right", "B", 0, "left", "B"}))
	Expect(commands["B"]).Should(Equal(instruction{1, "left", "A", 1, "right", "A"}))

	cursor := 0
	tape := make(map[int]int, 0)
	for i := 0; i < steps; i++ {
		state, cursor, tape = runStep(commands[state], cursor, tape)
	}

	checksum := 0
	for _, v := range tape {
		checksum += v
	}

	fmt.Println(checksum)
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
