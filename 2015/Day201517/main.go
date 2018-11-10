package Day201517

import (
	"fmt"
	"strconv"
	"strings"
)

// Entry holds wraps the data and runner interfaces for this puzzle
var Entry dayEntry

type dayEntry bool

func getData(input string) []int {
	jars := make([]int, 0)

	for _, val := range strings.Split(input, "\n") {
		jarSize, _ := strconv.Atoi(val)

		jars = append(jars, jarSize)
	}
	return jars
}

func sumInts(ints []int) int {
	result := 0
	for _, val := range ints {
		result += val
	}
	return result
}
func combinator(remaining []int, combo []int, depth int, matchChan chan<- []int) {
	if depth == 0 {
		defer close(matchChan)
	}

	for i := 0; i < len(remaining); i++ {
		newCombo := append(combo, remaining[i])

		newRemain := append(remaining[:i], remaining[i+1:]...)

		if len(newRemain) > 0 {
			combinator(newRemain, newCombo, depth+1, matchChan)
		}
	}
}

func (td dayEntry) Describe() (int, int, string) {
	return 2015, 17, "No Such Thing as Too Much"
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	return fmt.Sprintf(" -- Not Yet Implemented --")
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	return fmt.Sprintf(" -- Not Yet Implemented --")
}
