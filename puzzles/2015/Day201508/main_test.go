package Day201508

import (
	"strings"
	"testing"

	"fmt"

	. "github.com/onsi/gomega"
)

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)
	for _, data := range strings.Split(Entry.PuzzleInput(), "\n") {
		fmt.Println(data, escapeString(data))
	}
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)
	for _, data := range strings.Split(Entry.PuzzleInput(), "\n") {
		fmt.Println(data, encodeString(data))
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
