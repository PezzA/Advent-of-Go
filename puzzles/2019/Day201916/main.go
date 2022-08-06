package Day201916

import (
	"fmt"
	"strconv"

	"github.com/pezza/advent-of-code/puzzles/common"
)

func getData(input string) []int {
	var vals []int

	for _, char := range input {
		val, _ := strconv.Atoi(string(char))
		vals = append(vals, val)
	}
	return vals
}

func basePattern(position int) []int {
	var vals []int

	for i := 0; i < position; i++ {
		vals = append(vals, 0)
	}

	for i := 0; i < position; i++ {
		vals = append(vals, 1)
	}

	for i := 0; i < position; i++ {
		vals = append(vals, 0)
	}

	for i := 0; i < position; i++ {
		vals = append(vals, -1)
	}
	return vals
}

func basePatternIndex(position int, index int) int {
	patternLen := position * 4 // size of repeating pattern
	pos := index % patternLen  // index within the pattern

	if pos < position {
		return 0
	} else if pos < position*2 {
		return 1
	} else if pos < position*3 {
		return 0
	}

	return -1
}

func applyPhaseShift(input []int) []int {
	var newPhaseCode []int
	for i := 1; i <= len(input); i++ {

		total := 0

		for y := 0; y < len(input); y++ {
			total += input[y] * basePatternIndex(i, y+1)
		}

		newPhaseCode = append(newPhaseCode, common.Abs(total%10))
	}

	return newPhaseCode
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	seq := getData(Entry.PuzzleInput())
	for i := 0; i < 100; i++ {
		seq = applyPhaseShift(seq)
	}

	return fmt.Sprintf("%v", seq)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	return fmt.Sprintf("%v", " -- Not Yet Implemented --")
}
