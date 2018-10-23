package Day201519

import (
	"fmt"
	"strings"
	"testing"

	"github.com/pezza/advent-of-code/Common"

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

	replacementList, molecule := getTestData()

	fmt.Println("Medicine molecule", molecule)
	fmt.Println("Replacement List", replacementList)

	molecules := make([]string, 0)

	for _, r := range replacementList {

		foundIndex, currentPosition := 0, 0

		for {
			foundIndex = strings.Index(molecule[currentPosition:], r.key)
			if foundIndex == -1 {
				break
			}
			currentPosition += foundIndex + 1

			newMolecule := molecule[:currentPosition-1]
			newMolecule += r.replacement
			newMolecule += molecule[currentPosition:]

			if !common.Contains(molecules, newMolecule) {
				molecules = append(molecules, newMolecule)
			}
		}
	}

	fmt.Println(len(molecules))
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
