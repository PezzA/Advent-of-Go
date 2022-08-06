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

func deepCopy(input []int) []int {
	result := make([]int, len(input))

	for index, val := range input {
		result[index] = val
	}
	return result
}

func sum(input []int) int {
	retval := 0

	for _, val := range input {
		retval += val
	}
	return retval
}

func combinator(remaining []int, combo []int, matchChan chan<- []int) {
	if len(combo) == 0 {
		defer close(matchChan)
	}
	for index, nextItem := range remaining {
		newCombo := deepCopy(append(combo, nextItem))
		newRemain := deepCopy(remaining[index+1:])

		if sum(newCombo) == 150 {
			matchChan <- newCombo
		}

		combinator(newRemain, newCombo, matchChan)
	}
}

func (td dayEntry) Describe() (int, int, string, int) {
	return 2015, 17, "No Such Thing as Too Much", 2
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	matchChan := make(chan []int)

	go combinator(getData(Entry.PuzzleInput()), []int{}, matchChan)

	count := 0
	for range matchChan {
		count++
	}

	return fmt.Sprintf("%v", count)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	matchChan := make(chan []int)

	go combinator(getData(Entry.PuzzleInput()), []int{}, matchChan)

	count := -1
	var targetSet []int
	for match := range matchChan {
		if len(match) < count || count == -1 {
			count = len(match)
			targetSet = match
		}
	}

	targetLen := len(targetSet)

	matchChan = make(chan []int)
	go combinator(getData(Entry.PuzzleInput()), []int{}, matchChan)

	finalCount := 0
	for match := range matchChan {
		if len(match) == targetLen {
			finalCount++
		}
	}

	return fmt.Sprintf("%v", finalCount)
}
