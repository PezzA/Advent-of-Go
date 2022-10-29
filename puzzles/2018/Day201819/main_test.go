package Day201819

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/pezza/advent-of-code/puzzle-support/chronalcompiler"
)

func Test_GetData(t *testing.T) {
	RegisterTestingT(t)

	ipBinding, insList := getData(Entry.PuzzleInput())

	Expect(ipBinding).Should(Equal(5))
	Expect(insList[0]).Should(Equal(chronalcompiler.Instruction{OpCode: "addi", A: 5, B: 16, C: 5}))
	Expect(insList[len(insList)-1]).Should(Equal(chronalcompiler.Instruction{OpCode: "seti", A: 0, B: 0, C: 5}))
}

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	ipBinding, program := getData(`#ip 0
seti 5 0 1
seti 6 0 2
addi 0 1 0
addr 1 2 3
setr 1 0 0
seti 8 0 4
seti 9 0 5`)

	regSet := chronalcompiler.RunProgram(program, 6, ipBinding, nil)

	fmt.Println(regSet[0])
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
