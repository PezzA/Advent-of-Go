package Day201905

import (
	"testing"

	. "github.com/onsi/gomega"
	"github.com/pezza/advent-of-code/2019/intcode"
)

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	vm := intcode.New(Entry.PuzzleInput())
	results := vm.RunProgram(nil, []int64{1}, nil, nil)
	Expect(results[len(results)-1]).Should(Equal(int64(10987514)))
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)
	vm := intcode.New(Entry.PuzzleInput())
	results := vm.RunProgram(nil, []int64{5}, nil, nil)
	Expect(results[0]).Should(Equal(int64(14195011)))
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
