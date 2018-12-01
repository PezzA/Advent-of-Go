package Day201519

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

func getTestData() (replacements, string) {
	replacementList := make(replacements, 0)

	replacementList = append(replacementList, replacement{"e", "H"})
	replacementList = append(replacementList, replacement{"e", "O"})
	replacementList = append(replacementList, replacement{"H", "HO"})
	replacementList = append(replacementList, replacement{"H", "OH"})
	replacementList = append(replacementList, replacement{"O", "HH"})

	return replacementList, "HOHOHO"
}

func Test_ReverseList(t *testing.T) {
	RegisterTestingT(t)

	rlist, target := getData(Entry.PuzzleInput())
	//rlist, target := getTestData()

	starts := rlist.getStarterList()

	for _, start := range starts {
		fmt.Println(makeMol2(target, start, rlist, 0))
	}

}
func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	//replacementList, molecule := getData(Entry.PuzzleInput())
	//replacementList, molecule := getTestData()
	//molecules := getCombinations(molecule, replacementList)
	//fmt.Println(len(molecules))
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
