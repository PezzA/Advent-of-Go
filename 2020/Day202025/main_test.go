package Day202025

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

func Test_ReadData(t *testing.T) {
	RegisterTestingT(t)

}

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	cardPublicKey, doorPublicKey := getData(Entry.PuzzleInput())
	val := 1

	cardLoopSize := 0
	for i := 0; true; i++ {
		val = transForm(val, 7)
		if cardPublicKey == val {
			cardLoopSize = i
			break
		}
	}

	encVal := 1
	for i := 0; i <= cardLoopSize; i++ {
		encVal = transForm(encVal, doorPublicKey)
	}

	fmt.Println(encVal)

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
