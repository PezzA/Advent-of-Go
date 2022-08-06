package Day202022

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

func Test_ReadData(t *testing.T) {
	RegisterTestingT(t)

	d1, d2 := getData(Entry.PuzzleInput())

	Expect(len(d1)).Should(Equal(len(d2)))
}

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	d1, d2 := getData(Entry.PuzzleInput())

	d1, d2 = runGame(d1, d2)

	score := 0
	if len(d1) > 0 {
		score = d1.score()
	} else {
		score = d2.score()
	}

	fmt.Println(score)
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

	d1, d2 := getData(Entry.PuzzleInput())

	d1, d2 = runRecursiveGame(d1, d2, []state{})

	fmt.Println(d1)
	fmt.Println(d2)
	score := 0
	if len(d1) > 0 {
		score = d1.score()
	} else {
		score = d2.score()
	}

	fmt.Println(score)

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
