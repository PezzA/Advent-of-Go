package Day201816

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
	cc "github.com/pezza/advent-of-code/puzzles/2018/chronalcompiler"
)

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

}

func Test_ResolveCodeMap(t *testing.T) {
	RegisterTestingT(t)

	tests, _ := getData(Entry.PuzzleInput())

	results := determineCodeMap(cc.GetOpCodes(), tests)

	for i := 0; i < len(results); i++ {
		fmt.Println(i, results[i])
	}
}

func Test_RunProgram(t *testing.T) {
	RegisterTestingT(t)

	tests, program := getData(Entry.PuzzleInput())
	opCodes, regSet := cc.GetOpCodes(), cc.NewRegisterSet()

	insLookup := determineCodeMap(opCodes, tests)
	opCodes = cc.GetOpCodes()
	for _, ins := range program {
		instruction := insLookup[ins[0]]
		regSet = opCodes[instruction].Process(regSet, ins[1], ins[2], ins[3])
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
