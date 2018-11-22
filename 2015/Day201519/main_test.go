package Day201519

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

func getTestData() (replacements, string) {
	replacementList := make(replacements, 0)
	replacementList = append(replacementList, replacement{"H", "HO"})
	replacementList = append(replacementList, replacement{"H", "OH"})
	replacementList = append(replacementList, replacement{"O", "HH"})

	return replacementList, "HOHOHO"
}

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	//replacementList, molecule := getData(Entry.PuzzleInput())
	replacementList, molecule := getTestData()
	molecules := getCombinations(molecule, replacementList)
	fmt.Println(len(molecules))
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)
	replacementList, molecule := getTestData()

	for {
		molecules := getCombinations(molecule, replacementList)
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
