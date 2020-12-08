package Day202008

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

func Test_ReadData(t *testing.T) {
	RegisterTestingT(t)

	data := getData(Entry.PuzzleInput())

	Expect(data[0].command).Should(Equal("acc"))
	Expect(data[0].modifier).Should(Equal(-1))
}

func Test_RunProgram(t *testing.T) {
	RegisterTestingT(t)

	data := getData(`nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`)

	acc, _ := runProgram(data)

	Expect(acc).Should(Equal(5))
}
func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)
	data := getData(Entry.PuzzleInput())

	term, acc := runProgram(data)

	Expect(term).Should(Equal(false))
	Expect(acc).Should(Equal(2014))
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)
	data := getData(Entry.PuzzleInput())

	jumps := getCommands("jmp", data)
	noOps := getCommands("nop", data)

	termAcc := 0
	for _, jmp := range jumps {

		testProg := deepCopy(data)

		testProg[jmp].command = "nop"

		term, acc := runProgram(testProg)

		if term {
			termAcc = acc
		}

	}

	if termAcc == 0 {
		fmt.Println("NO TERMACC")
	} else {
		fmt.Println("TERMACC", termAcc)
		return
	}

	for _, nop := range noOps {
		testProg := deepCopy(data)

		testProg[nop].command = "jmp"

		term, acc := runProgram(testProg)

		if term {
			termAcc = acc
		}
	}

	if termAcc == 0 {
		fmt.Println("NO TERMACC")
	} else {
		fmt.Println("TERMACC", termAcc)
	}

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
