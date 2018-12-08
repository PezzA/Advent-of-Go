package Day201808

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

var testData = "2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2"

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	vals := getData(Entry.PuzzleInput())
	node, _ := getNode(vals)
	mc := countMetaData(node)

	fmt.Println(mc)

	fmt.Println(advancedCount(node))
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
