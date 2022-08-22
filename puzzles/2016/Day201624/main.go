package Day201624

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pezza/advent-of-code/puzzles/common"
)

const WALL int = -2
const SPACE int = -1

type mazeMap = map[common.Point]int

func getMaze(input string, wallsAsSpace bool) (mazeMap, map[int]common.Point) {
	mazeData := make(mazeMap)
	locations := make(map[int]common.Point)

	lines := strings.Split(input, "\n")

	for y, line := range lines {
		for x, position := range line {

			switch position {
			case '#':
				mazeData[common.Point{X: x, Y: y}] = WALL
			case '.':
				mazeData[common.Point{X: x, Y: y}] = SPACE
			default:
				location, _ := strconv.Atoi(string(position))
				if wallsAsSpace {
					mazeData[common.Point{X: x, Y: y}] = SPACE
				} else {
					mazeData[common.Point{X: x, Y: y}] = location
				}

				locations[location] = common.Point{X: x, Y: y}
			}

		}
	}

	return mazeData, locations
}

func walk(depth int, loc common.Point, target common.Point, mm mazeMap) int {

	// is up the target?
	if  loc.X == target.X && loc.Y-1 == target.Y {
		return depth + 1
	}



	if testup === target

}

func distance(start common.Point, dest common.Point, mm mazeMap) int {

}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	return fmt.Sprintf("%v", " -- Not Yet Implemented --")
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	return fmt.Sprintf("%v", " -- Not Yet Implemented --")
}
