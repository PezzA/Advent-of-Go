package Day202003

import (
	"fmt"
	"strings"

	"github.com/pezza/advent-of-code/puzzles/common"
)

func getData(input string) [][]bool {
	lines := strings.Split(input, "\n")
	var grid = make([][]bool, len(lines))

	for l, line := range lines {
		gridRow := make([]bool, len(line))

		for i, r := range line {
			gridRow[i] = r == 35
		}

		grid[l] = gridRow
	}
	return grid
}

func raceToTheBottom(grid [][]bool, increment int, skipAlternate bool) int {
	hits, pos := 0, 0

	for i, row := range grid {
		if skipAlternate && i%2 == 1 {
			continue
		}

		if row[common.GetWrappedIndex(pos, len(row), 0)] {
			hits++
		}
		pos += increment
	}

	return hits
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	data := getData(inputData)

	return fmt.Sprintf("%v", raceToTheBottom(data, 3, false))
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	data := getData(inputData)

	total := raceToTheBottom(data, 1, false) *
		raceToTheBottom(data, 3, false) *
		raceToTheBottom(data, 5, false) *
		raceToTheBottom(data, 7, false) *
		raceToTheBottom(data, 1, true)

	return fmt.Sprintf("%v", total)
}
