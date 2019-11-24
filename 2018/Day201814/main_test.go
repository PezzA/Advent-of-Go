package Day201814

import (
	"container/list"
	"testing"

	. "github.com/onsi/gomega"
)

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	scores := list.New()
	elf1 := scores.PushBack(3)
	elf2 := scores.PushBack(7)

	elf1, elf2, _ = getRecipes(elf1, elf2, scores, "")

	Expect(elf1.Value.(int)).Should(Equal(3))
	Expect(elf2.Value.(int)).Should(Equal(7))

	elf1, elf2, _ = getRecipes(elf1, elf2, scores, "")

	Expect(elf1.Value.(int)).Should(Equal(1))
	Expect(elf2.Value.(int)).Should(Equal(0))

	scores = list.New()
	elf1 = scores.PushBack(3)
	elf2 = scores.PushBack(7)

}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)
	Entry.PartTwo(Entry.PuzzleInput(), nil)

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
