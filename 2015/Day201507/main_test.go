package Day201507

import (
	"testing"

	. "github.com/onsi/gomega"
)

var testInstructions = getInstructions(`123 -> x
456 -> y
x AND y -> d
x OR y -> e
x LSHIFT 2 -> f
y RSHIFT 2 -> g
NOT x -> h
NOT y -> i`)

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	// test that bitwise operations are returning the expected results!
	a, b := uint16(123), uint16(456)
	Expect(a & b).Should(Equal(uint16(72)))
	Expect(a | b).Should(Equal(uint16(507)))
	Expect(a << 2).Should(Equal(uint16(492)))
	Expect(b >> 2).Should(Equal(uint16(114)))
	Expect(^a).Should(Equal(uint16(65412)))

	Expect(*testInstructions[0]).Should(Equal(gate{"ASSIGN", "123", "", "x", false}))

	gates := testInstructions

	wires := make(wireList, 0)

	var todo bool
	todo, wires, gates = runCircuitLoop(gates, wires)
	Expect(todo).Should(Equal(true))

	todo, wires, gates = runCircuitLoop(gates, wires)
	Expect(todo).Should(Equal(false))

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
