package Day201916

import (
	"fmt"
	"strconv"

	"github.com/pezza/advent-of-code/common"
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

func applyPhaseShift(input []int) []int {
	var newPhaseCode []int
	for i := 1; i <= len(input); i++ {
		mask := basePattern(i)
		total := 0

		maskPos := 1

		for y := 0; y < len(input); y++ {
			total += input[y] * mask[maskPos]

			if maskPos == len(mask)-1 {
				maskPos = 0
			} else {
				maskPos++
			}

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
