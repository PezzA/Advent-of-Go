package Day202003

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

func Test_ReadData(t *testing.T) {
	RegisterTestingT(t)

	data := getData(Entry.PuzzleInput())

	fmt.Println(data)

}

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	data := getData(Entry.PuzzleInput())

	fmt.Println(raceToTheBottom(data, 3, false))
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

	data := getData(Entry.PuzzleInput())

	total := raceToTheBottom(data, 1, false) *
		raceToTheBottom(data, 3, false) *
		raceToTheBottom(data, 5, false) *
		raceToTheBottom(data, 7, false) *
		raceToTheBottom(data, 1, true)

	fmt.Println(total)
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
