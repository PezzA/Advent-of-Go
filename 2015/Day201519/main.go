package Day201519

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/pezza/advent-of-code/common"
)

var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string) {
	return 2015, 19, "Medicine for Rudolph"
}

type replacement struct {
	key         string
	replacement string
}

func getData(input string) ([]replacement, string) {
	replacementList, medicineMolecule := make([]replacement, 0), ""

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		if strings.Contains(line, "=>") {
			bits := strings.Fields(line)
			replacementList = append(replacementList, replacement{bits[0], bits[2]})
			continue
		}

		medicineMolecule = line
	}

	return replacementList, medicineMolecule
}

func getAllDistinctReplacements(reps []replacement, molecule string) []string {
	mutations := []string{}

	for _, rep := range reps {
		for _, variant := range getReplacements(rep, molecule) {

			if !common.Contains(mutations, variant) {
				mutations = append(mutations, variant)
			}
		}

	}
	return mutations
}

func getReplacements(r replacement, molecule string) []string {
	re := regexp.MustCompile(r.key)
	reps := []string{}

	for _, rep := range re.FindAllStringIndex(molecule, -1) {
		reps = append(reps, molecule[:rep[0]]+r.replacement+molecule[rep[1]:])
	}

	return reps
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {

	replacementList, molecule := getData(Entry.PuzzleInput())

	list := getAllDistinctReplacements(replacementList, molecule)

	return fmt.Sprintf("%v", len(list))

}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	return fmt.Sprintf(" -- Not Yet Implemented --")
}
