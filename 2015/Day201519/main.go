package Day201519

import (
	"fmt"
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

type replacements []replacement

type result struct {
	replacement string
	match       string
}

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

func getReverseReplacementList(molecule string, replacementList replacements) []result {
	molecules := make([]result, 0)

	for _, r := range replacementList {

		if r.key == "e" {
			continue
		}

		foundIndex, currentPosition := 0, 0

		for {
			foundIndex = strings.Index(molecule[currentPosition:], r.replacement)
			if foundIndex == -1 {
				break
			}

			adjustedIndex := currentPosition + foundIndex

			newMolecule := fmt.Sprintf("%v%v%v",
				molecule[:adjustedIndex],
				r.key,
				molecule[adjustedIndex+len(r.replacement):])

			currentPosition = adjustedIndex + len(r.replacement)

			newList := make([]string, len(molecules))

			for index, val := range molecules {
				newList[index] = val.replacement
			}

			if !common.Contains(newList, newMolecule) {
				molecules = append(molecules, result{newMolecule, r.replacement})
			}
		}
	}

	return molecules

}

var seenList = []string{}
var bounces = 0

var seenMap map[string]bool

func init() {
	seenMap = make(map[string]bool)
}
func makeMol2(currentMolecule string, targetMolecule string, rl replacements, depth int) int {

	if _, ok := seenMap[currentMolecule]; ok {
		return 0
	}

	seenMap[currentMolecule] = true

	if len(seenMap)%100000 == 0 {
		fmt.Println(targetMolecule, currentMolecule, len(seenMap))
	}

	// is it a match
	if currentMolecule == targetMolecule {
		return depth
	}

	// the the list of possible replacements
	replacements := getReverseReplacementList(currentMolecule, rl)

	// if we dont have any more replacements, return unsuccessful
	if len(replacements) == 0 {
		return 0
	}

	for _, newMolecule := range replacements {
		depth = makeMol2(newMolecule.replacement, targetMolecule, rl, depth+1)

		if depth != 0 {
			return depth
		}
	}

	return 0
}

func makeMolecule(currentMolecule string, targetMolecule string, rl replacements, depth int) int {

	if common.Contains(seenList, currentMolecule) {
		return -1
	}

	seenList = append(seenList, currentMolecule)

	replacements := getReverseReplacementList(currentMolecule, rl)
	depth++

	for _, newMolecule := range replacements {
		if len(newMolecule.replacement) == 0 {
			continue
		}

		if newMolecule.replacement == targetMolecule {
			return depth
		}

		depth = makeMolecule(newMolecule.replacement, targetMolecule, rl, depth+1)
	}

	return depth
}

func getReplacementList(molecule string, replacementList replacements) []result {
	molecules := make([]result, 0)

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

			newList := make([]string, len(molecules))

			for index, val := range molecules {
				newList[index] = val.replacement
			}
			if _, ok := seenMap[newMolecule]; !ok {
				seenMap[newMolecule] = true
				molecules = append(molecules, result{newMolecule, r.replacement})
			}
		}
	}

	return molecules
}

func doPass(list []string, rl replacements, target string) ([]string, bool) {

	newList := make([]string, 0)

	for _, item := range list {
		for _, newItem := range getReplacementList(item, rl) {
			if newItem.replacement == target {
				return []string{newItem.replacement}, true
			}

			if !common.Contains(newList, newItem.replacement) {
				newList = append(newList, newItem.replacement)
			}
		}
	}

	return newList, false

}

func (r replacements) getStarterList() []string {
	rl := make([]string, 0)
	for _, replacement := range r {
		if replacement.key == "e" {
			rl = append(rl, replacement.replacement)
		}
	}
	return rl

}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	replacementList, molecule := getData(inputData)

	return fmt.Sprintf("%v", len(getReplacementList(molecule, replacementList)))

}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {

	return fmt.Sprintf(" -- Not Yet Implemented --")
}
