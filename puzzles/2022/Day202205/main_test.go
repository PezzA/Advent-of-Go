package Day202205

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_ReadData(t *testing.T) {
	RegisterTestingT(t)

	stacks, ins := getData(Entry.PuzzleInput())

	Expect(len(stacks)).Should(Equal(9))
	Expect(len(ins)).Should(Equal(504))

	Expect(stacks[1].Peek()).Should(Equal("Q"))
	Expect(stacks[2].Peek()).Should(Equal("G"))
	Expect(stacks[3].Peek()).Should(Equal("B"))
	Expect(stacks[4].Peek()).Should(Equal("N"))
	Expect(stacks[5].Peek()).Should(Equal("F"))
	Expect(stacks[6].Peek()).Should(Equal("J"))
	Expect(stacks[7].Peek()).Should(Equal("V"))
	Expect(stacks[8].Peek()).Should(Equal("N"))
	Expect(stacks[9].Peek()).Should(Equal("M"))

	Expect(ins[0]).Should(Equal(instruction{6, 2, 3}))
	Expect(ins[len(ins)-1]).Should(Equal(instruction{2, 6, 5}))
}

func Test_Process(t *testing.T) {
	RegisterTestingT(t)

	stacks, ins := getData(`    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`)

	processInstructionList(ins, stacks, false)

	Expect(stacks[1].Peek()).Should(Equal("C"))
	Expect(stacks[2].Peek()).Should(Equal("M"))
	Expect(stacks[3].Peek()).Should(Equal("Z"))

}
func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	Expect(Entry.PartOne(Entry.PuzzleInput(), nil)).Should(Equal("FJSRQCFTN"))
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

	Expect(Entry.PartTwo(Entry.PuzzleInput(), nil)).Should(Equal("CJVLJQPHS"))
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
