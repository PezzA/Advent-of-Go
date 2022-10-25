package Day201924

import (
	"fmt"
	"math"
	"strings"

	"github.com/pezza/advent-of-code/puzzles/common"
)

type hive map[int][][]bool

type mapping struct {
	cell int
}

const layerSize = 5

func getCell(x, y int) int {
	return (y * layerSize) + x
}

func countAdjacentBugs(point common.Point, bugs [][]bool) int {
	count := 0

	ubound := len(bugs) - 1

	for _, test := range common.Cardinal4 {
		testPos := point.Add(test)
		if testPos.X >= 0 && testPos.Y >= 0 && testPos.X <= ubound && testPos.Y <= ubound {
			if bugs[testPos.Y][testPos.X] {
				count++
			}
		}
	}
	return count
}

func tickAutomata(input [][]bool) [][]bool {

	newBugs := make([][]bool, len(input))

	for y, bugline := range input {
		newBugs[y] = make([]bool, len(bugline))
		for x := range bugline {
			newVal := input[y][x]
			adjBugs := countAdjacentBugs(common.Point{X: x, Y: y}, input)

			if adjBugs != 1 && input[y][x] {
				newVal = false
			}

			if adjBugs == 1 || adjBugs == 2 && !input[y][x] {
				newVal = true
			}

			newBugs[y][x] = newVal

		}
	}
	return newBugs
}

func tickRecursiveAutomata(input [][]bool) [][]bool {

	newBugs := make([][]bool, len(input))

	for y, bugline := range input {
		newBugs[y] = make([]bool, len(bugline))
		for x := range bugline {
			newVal := input[y][x]
			adjBugs := countAdjacentBugs(common.Point{X: x, Y: y}, input)

			if adjBugs != 1 && input[y][x] {
				newVal = false
			}

			if adjBugs == 1 || adjBugs == 2 && !input[y][x] {
				newVal = true
			}

			newBugs[y][x] = newVal

		}
	}
	return newBugs
}

func getData(input string) [][]bool {
	lines := strings.Split(input, "\n")
	bugs := make([][]bool, len(lines))

	for y, line := range lines {
		bugs[y] = make([]bool, len(line))
		for x, char := range line {
			if char == 35 {
				bugs[y][x] = true
			}
		}
	}
	return bugs
}

func same(src [][]bool, cmp [][]bool) bool {
	for y := range src {
		for x := range src[y] {
			if src[y][x] != cmp[y][x] {
				return false
			}
		}
	}
	return true
}

func getBioDiversity(input [][]bool) float64 {
	val, count := float64(0), float64(0)

	for y := range input {
		for x := range input[y] {
			if input[y][x] {
				val += math.Pow(2, count)
			}
			count++
		}
	}

	return val
}

func printBugs(input [][]bool) {
	for y := range input {
		for x := range input[y] {
			if input[y][x] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	return fmt.Sprintf("%v", " -- Not Yet Implemented --")
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	return fmt.Sprintf("%v", " -- Not Yet Implemented --")
}
