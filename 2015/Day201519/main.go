package Day201519

import (
	"fmt"
	"strings"

	"github.com/pezza/advent-of-code/Common"
)

var Entry dayEntry

type replacement struct {
	key         string
	replacement string
}

type replacements []replacement

func getData(input string) (replacements, string) {
	replacementList, medicineMolecule := make(replacements, 0), ""

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

type dayEntry bool

func (td dayEntry) Describe() (int, int, string) {
	return 2015, 19, "Medicine for Rudolph"
}

func getCombinations() []string {

}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	replacementList, molecule := getData(inputData)
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

	return fmt.Sprintf("%v", len(molecules))
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	return fmt.Sprintf(" -- Not Yet Implemented --")
}
