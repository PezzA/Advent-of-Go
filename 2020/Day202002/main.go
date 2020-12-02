package Day202002

import (
	"fmt"
	"strings"
)

type passwordSpec struct {
	letter         string
	minOccurrences int
	maxOccurrences int
	password       string
}

func getData(input string) []passwordSpec {
	data := make([]passwordSpec, 0)
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		var min int
		var max int
		var letter string
		var pass string

		fmt.Sscanf(line, "%d-%d %v %v", &min, &max, &letter, &pass)

		data = append(data, passwordSpec{
			letter:         letter[:1],
			minOccurrences: min,
			maxOccurrences: max,
			password:       pass,
		})
	}

	return data
}

func isPasswordValid(spec passwordSpec) bool {
	charMap := make(map[string]int, 0)

	for _, c := range spec.password {
		cmp := string(c)
		if _, ok := charMap[cmp]; ok {
			charMap[cmp]++
		} else {
			charMap[cmp] = 1
		}
	}

	if val, ok := charMap[spec.letter]; !ok {
		return false
	} else {
		if val < spec.minOccurrences || val > spec.maxOccurrences {
			return false
		}
	}

	return true
}

func isPositionValid(spec passwordSpec) bool {
	posA := spec.password[spec.minOccurrences-1 : spec.minOccurrences]
	posB := spec.password[spec.maxOccurrences-1 : spec.maxOccurrences]

	if posA == posB {
		return false
	}

	return posA == spec.letter || posB == spec.letter
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	valid := 0
	data := getData(Entry.PuzzleInput())

	for _, spec := range data {
		if isPasswordValid(spec) {
			valid++
		}
	}

	return fmt.Sprintf("%v", valid)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	valid := 0
	data := getData(Entry.PuzzleInput())

	for _, spec := range data {
		if isPositionValid(spec) {
			valid++
		}
	}

	return fmt.Sprintf("%v", valid)
}
