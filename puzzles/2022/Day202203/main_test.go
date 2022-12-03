package Day202203

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_ReadData(t *testing.T) {
	RegisterTestingT(t)

	data := getData(Entry.PuzzleInput())

	Expect(data[0]).Should(Equal(rucksack("FqdWDFppHWhmwwzdjvjTRTznjdMv")))
}

func Test_GetCommonItem(t *testing.T) {
	RegisterTestingT(t)

	r := rucksack("vJrwpWtwJgWrhcsFMMfFFhFp")

	Expect(r.getCommonItem()).Should(Equal(rune(112)))

}

func Test_GetPriority(t *testing.T) {
	RegisterTestingT(t)

	Expect(getItemPriority(rune(97))).Should(Equal(1))
	Expect(getItemPriority(rune(65))).Should(Equal(27))
}

func Test_SplitInHalf(t *testing.T) {
	RegisterTestingT(t)

	cmp1, cmp2 := splitInHalf("vJrwpWtwJgWrhcsFMMfFFhFp")

	Expect(cmp1).Should(Equal("vJrwpWtwJgWr"))
	Expect(cmp2).Should(Equal("hcsFMMfFFhFp"))
}

func Test_GetBadge(t *testing.T) {
	RegisterTestingT(t)

	data := getData(`vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg`)

	Expect(getCommonItemType(data[0], data[1], data[2])).Should(Equal(rune(114)))
}

func Test_PartOneCases(t *testing.T) {
	RegisterTestingT(t)

	data := getData(`vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`)

	Expect(getItemPriority(data[0].getCommonItem())).Should(Equal(16))
}

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)
	// not 8111
	Expect(Entry.PartOne(Entry.PuzzleInput(), nil)).Should(Equal("8243"))
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

	Expect(Entry.PartTwo(Entry.PuzzleInput(), nil)).Should(Equal("2631"))
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
