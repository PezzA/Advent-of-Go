package Day201519

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/pezza/advent-of-code/puzzles/common"
)

var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string, int) {
	return 2015, 19, "Medicine for Rudolph", 1
}

type replacement struct {
	key         string
	replacement string
	keyRegex    *regexp.Regexp
	repRegex    *regexp.Regexp
}

func getData(input string) ([]replacement, string) {
	replacementList, medicineMolecule := make([]replacement, 0), ""

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		if strings.Contains(line, "=>") {
			bits := strings.Fields(line)
			replacementList = append(replacementList, replacement{bits[0], bits[2], regexp.MustCompile(bits[0]), regexp.MustCompile(bits[2])})
			continue
		}

		medicineMolecule = line
	}

	return replacementList, medicineMolecule
}

func stripKey(key string, reps []replacement) []replacement {
	newReps := []replacement{}

	for index := range reps {
		if reps[index].key != key {
			newReps = append(newReps, reps[index])
		}
	}

	return newReps
}

func getAllDistinctReplacements(reps []replacement, molecule string) []string {
	mutations := []string{}

	for _, rep := range reps {
		for _, variant := range getReplacements(rep.keyRegex, rep.replacement, molecule) {
			if !common.Contains(mutations, variant) {
				mutations = append(mutations, variant)
			}
		}

	}
	return mutations
}

func getAllDistinctReverseReplacements(reps []replacement, molecule string) []string {
	mutations := []string{}

	for _, rep := range reps {
		for _, variant := range getReplacements(rep.repRegex, rep.key, molecule) {
			if !isSeen(variant) {
				mutations = append(mutations, variant)
			}
		}

	}
	return mutations
}

func getReplacements(search *regexp.Regexp, replacement, molecule string) []string {
	reps := []string{}
	for _, rep := range search.FindAllStringIndex(molecule, -1) {
		incoming := molecule[:rep[0]] + replacement + molecule[rep[1]:]
		if !common.Contains(reps, incoming) {
			reps = append(reps, incoming)
		}
	}

	return reps
}

var seen map[int][]string = make(map[int][]string, 0)

func isSeen(molecule string) bool {
	l := len(molecule)

	if list, ok := seen[l]; ok {
		for _, item := range list {
			if item == molecule {
				return true
			}
		}
	}
	return false
}

func addIfNotExists(molecule string) {
	if isSeen(molecule) {
		return
	}

	l := len(molecule)

	if list, ok := seen[l]; ok {
		seen[l] = append(list, molecule)
	} else {
		totSize := 0

		for _, v := range seen {
			totSize += len(v)
		}
		log.Println(totSize)
		seen[l] = []string{molecule}
	}
}

func moveToTargets(targets []string, molecule string, history []string, repList []replacement, depth int) ([]string, bool) {
	log.Println(molecule)
	if common.Contains(targets, molecule) {
		return history, true
	}

	newRepList := getAllDistinctReverseReplacements(repList, molecule)

	if len(newRepList) == 0 {
		addIfNotExists(molecule)
		return history, false
	}

	for _, rep := range newRepList {
		val, success := []string{}, false
		if !isSeen(rep) {
			val, success = moveToTargets(targets, rep, append(history, molecule), repList, depth+1)

			if success {
				return val, success
			}

			for _, item := range val {
				addIfNotExists(item)
			}
		}

	}

	return history, false
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	replacementList, molecule := getData(Entry.PuzzleInput())
	return fmt.Sprintf("%v", len(getAllDistinctReplacements(replacementList, molecule)))

}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	return fmt.Sprintf(" -- Not Yet Implemented --")
}
