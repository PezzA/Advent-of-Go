package Day201623

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/pezza/advent-of-code/puzzle-support/assembunny"
)

func Test_ReadData(t *testing.T) {
	RegisterTestingT(t)
	data := assembunny.ParseProgram(Entry.PuzzleInput())

	fmt.Println(data)
}

func Test_PartOneSample(t *testing.T) {
	RegisterTestingT(t)

	program := assembunny.ParseProgram(`cpy 2 a
tgl a
tgl a
tgl a
cpy 1 a
dec a
dec a`)

	registerSet := make(assembunny.RegisterSet)

	registerSet["a"] = 0
	registerSet["b"] = 0
	registerSet["c"] = 0
	registerSet["d"] = 0

	registers := assembunny.RunProgram(program, registerSet)

	fmt.Println(registers["a"])
}

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	program := assembunny.ParseProgram(Entry.PuzzleInput())

	registerSet := assembunny.NewRegisterSet(7, 0, 0, 0)

	registers := assembunny.RunProgram(program, registerSet)

	fmt.Println(registers["a"])
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

	program := assembunny.ParseProgram(Entry.PuzzleInput())

	registerSet := assembunny.NewRegisterSet(12, 0, 0, 0)

	registers := assembunny.RunProgram(program, registerSet)

	fmt.Println(registers["a"])
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
