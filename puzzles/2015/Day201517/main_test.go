package Day201517

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	matchChan := make(chan []int)

	go combinator(getData(Entry.PuzzleInput()), []int{}, matchChan)

	count := 0
	for match := range matchChan {
		count++
		fmt.Println(match)
	}

	fmt.Println(count)
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)
	matchChan := make(chan []int)

	go combinator(getData(Entry.PuzzleInput()), []int{}, matchChan)

	count := -1
	var targetSet []int
	for match := range matchChan {
		if len(match) < count || count == -1 {
			count = len(match)
			targetSet = match
		}
	}

	targetLen := len(targetSet)
	fmt.Println("------------------")
	matchChan = make(chan []int)
	go combinator(getData(Entry.PuzzleInput()), []int{}, matchChan)

	for match := range matchChan {
		if len(match) == targetLen {
			fmt.Println(match)
		}

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
