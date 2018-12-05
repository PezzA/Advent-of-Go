package Day201805

import (
	"fmt"
	"strings"
)

// Entry holds wraps the data and runner interfaces for this puzzle
var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string) {
	return 2018, 5, "Alchemical Reduction"
}

func reactPolymer(input string) (string, bool) {
	for i := 0; i < len(input)-1; i++ {
		if input[i]-input[i+1] == 32 || input[i+1]-input[i] == 32 {
			return input[:i] + input[i+2:], true
		}
	}

	return input, false
}

func fullyReactPolymer(input string) string {
	reduction := true
	for reduction {
		input, reduction = reactPolymer(input)
	}

	return input
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	return fmt.Sprintf("%v", len(fullyReactPolymer(inputData)))
}

func removeUnit(unit rune, input string) string {
	input = strings.Replace(input, string(unit), "", -1)
	input = strings.Replace(input, string(unit-32), "", -1)
	return input
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	replacements := "abcdefghijklmnopqrstuvwxyz"

	inputData = fullyReactPolymer(inputData)
	min := -1
	for _, char := range replacements {
		newPolymer := removeUnit(char, inputData)

		newPolymer = fullyReactPolymer(newPolymer)

		if min == -1 || len(newPolymer) < min {
			min = len(newPolymer)
		}
	}
	// not 4, 1
	return fmt.Sprintf("%v", min)
}
