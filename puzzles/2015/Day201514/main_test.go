package Day201514

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	comet := reindeer{"comet", physicality{14, 10, 127}, state{0, true, 10}}
	dancer := reindeer{"dancer", physicality{16, 11, 162}, state{0, true, 11}}

	for i := 0; i < 1000; i++ {
		comet.tick()
		dancer.tick()
	}

	fmt.Println(comet)
	fmt.Println(dancer)
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
