package Day201816

import (
	"fmt"
	"log"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/pezza/advent-of-code/puzzle-support/chronalcompiler"
)

func Test_GetData(t *testing.T) {
	RegisterTestingT(t)

	_, codes := getData(Entry.PuzzleInput())

	Expect(codes[0]).Should(Equal([]int{7, 3, 2, 0}))

	log.Println(codes[len(codes)-1])
	Expect(codes[len(codes)-1]).Should(Equal([]int{1, 3, 0, 0}))

}
func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

}

func Test_ResolveCodeMap(t *testing.T) {
	RegisterTestingT(t)

	tests, _ := getData(Entry.PuzzleInput())

	results := determineCodeMap(chronalcompiler.GetOpCodes(), tests)

	for i := 0; i < len(results); i++ {
		fmt.Println(i, results[i])
	}
}

func Test_RunProgram(t *testing.T) {
	RegisterTestingT(t)

	tests, program := getData(Entry.PuzzleInput())
	opCodes, regSet := chronalcompiler.GetOpCodes(), chronalcompiler.NewRegisterSet(4)

	insLookup := determineCodeMap(opCodes, tests)
	opCodes = chronalcompiler.GetOpCodes()
	for _, ins := range program {
		instruction := insLookup[ins[0]]
		val := opCodes[instruction].Process(regSet, ins[1], ins[2])
		regSet[ins[3]] = val
	}

	Expect(regSet[0]).ShouldNot(Equal(617))
	fmt.Println(regSet)
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
