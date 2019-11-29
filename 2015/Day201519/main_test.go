package Day201519

import (
	"testing"

	. "github.com/onsi/gomega"
)

func getTestData() ([]replacement, string) {
	replacementList := make([]replacement, 0)

	replacementList = append(replacementList, replacement{"e", "H"})
	replacementList = append(replacementList, replacement{"e", "O"})
	replacementList = append(replacementList, replacement{"H", "HO"})
	replacementList = append(replacementList, replacement{"H", "OH"})
	replacementList = append(replacementList, replacement{"O", "HH"})

	return replacementList, "HOHOHO"
}

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	replacementList, molecule := getData(Entry.PuzzleInput())

	list := getAllDistinctReplacements(replacementList, molecule)

	Expect(len(list)).Should(Equal(509))
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
