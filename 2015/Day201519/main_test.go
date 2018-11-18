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

	//replacementList, molecule := getData(Entry.PuzzleInput())
	replacementList, molecule := getTestData()
	molecules := make([]string, 0)

	for _, r := range replacementList {

		foundIndex, currentPosition := 0, 0

		for {
			foundIndex = strings.Index(molecule[currentPosition:], r.key)
			if foundIndex == -1 {
				break
			}

			adjustedIndex := currentPosition + foundIndex

			newMolecule := fmt.Sprintf("%v%v%v",
				molecule[:adjustedIndex],
				r.replacement,
				molecule[adjustedIndex+len(r.key):])

			currentPosition = adjustedIndex + len(r.key)

			if !common.Contains(molecules, newMolecule) {
				molecules = append(molecules, newMolecule)
			}
		}
	}
	fmt.Println(len(molecules))
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

	//	replacementList, targetMolecule := getData(`e => H
	//e => O
	//H => HO
	//H => OH
	//O => HH
	//
	//HOH`)

	replacementList, targetMolecule := getData(Entry.PuzzleInput())

	list := replacementList.getStarterList()

	found, count := false, 0
	for !found {
		fmt.Println(count, len(list), len(seenMap))
		list, found = doPass(list, replacementList, targetMolecule)
		count++

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
