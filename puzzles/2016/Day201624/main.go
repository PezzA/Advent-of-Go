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
			cellPoint := common.NewPoint(x, y)

			switch position {
			case '#':
				mazeData[cellPoint] = WALL
			case '.':
				mazeData[cellPoint] = SPACE
			default:
				location, _ := strconv.Atoi(string(position))
				if wallsAsSpace {
					mazeData[cellPoint] = SPACE
				} else {
					mazeData[cellPoint] = location
				}
				locations[location] = cellPoint
			}
		}
	}

	return mazeData, locations
}

func checkTestPoint(newPoint common.Point, target common.Point, mm mazeMap) int {
	// are we the targer
	if newPoint.Equals(target) {
		return 1
	}

	if visited.c

	// can we move up?
	if mm[newPoint] != WALL {
		return 2
	}



	return 0
}

func walk(depth int, loc common.Point, target common.Point, mm mazeMap, visited []common.Point) int {

	visited = append(visited, loc)

	for _, cardinal := range common.Cardinal4 {
		testPoint := loc.Add(cardinal)

		move := checkTestPoint(testPoint, target, visited, mm)

	}
	// testup
	newPoint := common.NewPoint(loc.X, loc.Y-1)

	// is up the target?
	if loc.Equals(target) {
		return depth + 1
	}

	// can we move up?
	if mm[newPoint] != WALL {
		return walk(depth+1, newPoint, target, mm)
	}

}

func distance(start common.Point, dest common.Point, mm mazeMap) int {

}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	return fmt.Sprintf("%v", " -- Not Yet Implemented --")
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	return fmt.Sprintf("%v", " -- Not Yet Implemented --")
}
