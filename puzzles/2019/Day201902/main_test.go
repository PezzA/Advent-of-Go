package Day201902

import (
	"testing"

	. "github.com/onsi/gomega"
	"github.com/pezza/advent-of-code/puzzles/2019/intcode"
)

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	vc := intcode.New(Entry.PuzzleInput())

	init := make(map[int64]int64, 0)
	init[1] = 12
	init[2] = 2

	vc.RunProgram(init, []int64{}, nil, nil, nil)
	Expect(vc.GetValue(0)).Should(Equal(int64(3562672)))
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

	vc := intcode.New(Entry.PuzzleInput())

	done := false
	actNoun, actVerb := 0, 0

	for noun := 0; noun <= 100; noun++ {
		for verb := 0; verb <= 100; verb++ {

			init := make(map[int64]int64, 0)
			init[1] = int64(noun)
			init[2] = int64(verb)

			vc.RunProgram(init, []int64{}, nil, nil, nil)

			if vc.GetValue(0) == 19690720 {
				actNoun = noun
				actVerb = verb
				done = true
			}

			if done {
				break
			}
		}
		if done {
			break
		}
	}

	Expect(100*actNoun + actVerb).Should(Equal(8250))

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
