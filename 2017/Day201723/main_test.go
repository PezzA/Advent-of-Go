package Day201723

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	/*
		ins := getInstructions(Entry.PuzzleInput())

		c := chip{
			registers:      make(registers, 0),
			index:          0,
			mulInvocations: 0,
		}

		fmt.Println(runProgram(ins, c, nil))
	*/
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

	//ins := getInstructions(Entry.PuzzleInput())

	c := chip{
		registers:      make(registers, 0),
		index:          0,
		mulInvocations: 0,
	}

	c.registers["a"] = 1

	//fmt.Println(runProgram(ins, c, nil))

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
