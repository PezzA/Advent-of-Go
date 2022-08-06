package Day202015

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_ReadData(t *testing.T) {
	RegisterTestingT(t)

	data := getData(Entry.PuzzleInput())

	Expect(data[0]).Should(Equal(13))
	Expect(data[len(data)-1]).Should(Equal(8))

}

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	data := getData(Entry.PuzzleInput())
	seens := make(map[int][]int, 0)

	for i := range data {
		seens[data[i]] = []int{i + 1}
	}

	turn2020 := 0
	for i := len(data) + 1; i <= 30000000; i++ {
		data, seens = getNextNumber(data, seens, i)

		if i == 30000000 {
			turn2020 = data[len(data)-1]
		}
	}

	Expect(turn2020).Should(Equal(436))
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
