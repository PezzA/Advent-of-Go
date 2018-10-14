package Day201509

import (
	"testing"

	"fmt"

	. "github.com/onsi/gomega"
)

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)
	destinations := getData(Entry.PuzzleInput()).extractDestinations()

	resultChan := make(chan []string)

	go permutation(destinations, make([]string, 0), resultChan)

	i := 0
	for range resultChan {
		i++
		if i == factorial(len(destinations)) {
			close(resultChan)
		}
	}

	fmt.Println("done!")

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
