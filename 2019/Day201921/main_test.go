package Day201921

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_ReadData(t *testing.T) {
	RegisterTestingT(t)

}

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	/*
		`NOT A J` // 1) jump if hole in front
		`NOT B T` // 2) is there a hole two spaces away
		`AND D T` //    and ground four spaces away
		`OR T J`  // jump if either is true
		`NOT C T` // 3) if there is a hole 3 spaces away
		`AND D T` //   and ground four spaces away
		`OR T J`  // jump if either is true
	*/

	Expect(Entry.PartOne(Entry.PuzzleInput(), nil)).Should(Equal("19357335"))
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

	/*
		.................
		.................
		..@..............
		#####.#.#..##.###
		   ABCDEFGHI
	*/
	Expect(Entry.PartTwo(Entry.PuzzleInput(), nil)).Should(Equal("1140147758"))

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
