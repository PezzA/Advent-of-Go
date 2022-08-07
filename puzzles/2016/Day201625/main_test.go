package Day201625

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/pezza/advent-of-code/puzzle-support/assembunny"
)

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	program := assembunny.ParseProgram(Entry.PuzzleInput())

	test := "0101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101"
	found := false
	index := 0
	for !found {

		registerSet := assembunny.NewRegisterSet(index, 0, 0, 0)

		outputChan := make(chan int)

		go assembunny.RunProgram(program, registerSet, outputChan, 100)

		testString := ""
		for elem := range outputChan {
			testString += fmt.Sprint(elem)
		}

		fmt.Println("Index:", index, "Result", testString)

		if testString == test {
			break
		}
		index++
	}

	fmt.Println(index)
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
