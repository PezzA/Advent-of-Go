package Day201519

import (
	"fmt"
	"strings"
)

var Entry dayEntry

type replacements map[string]string

func getData(input string) (replacements, string) {
	replacementList, medicineMolecule := make(replacements), ""

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		if strings.Contains(line, "=>") {
			bits := strings.Fields(line)
			replacementList[bits[0]] = bits[2]
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
func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	return fmt.Sprintf(" -- Not Yet Implemented --")
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	return fmt.Sprintf(" -- Not Yet Implemented --")
}
