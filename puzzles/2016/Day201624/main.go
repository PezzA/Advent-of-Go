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

func distance(loc common.Point, target common.Point, mm mazeMap) int {
	visited := make(map[common.Point]int)

	walk(0, loc, target, mm, visited)

	return visited[target]
}

func walk(depth int, loc common.Point, target common.Point, mm mazeMap, visited map[common.Point]int) {

	for _, cardinal := range common.Cardinal4 {
		testPoint := loc.Add(cardinal)

		// is up the target?
		if loc.Equals(target) {
			break
		}

		if mm[testPoint] != WALL {

			if val, ok := visited[testPoint]; ok {
				if val <= depth+1 {
					continue
				}
			}

			visited[testPoint] = depth + 1

			// If we havent, continue walking
			walk(depth+1, testPoint, target, mm, visited)
		}
	}
}

func getRouteLength(route []int64, distanceMap map[common.Point]int, returnToZero bool) int {
	actualRoute := append([]int64{0}, route...)
	routeLength := 0

	if returnToZero {
		actualRoute = append(actualRoute, 0)
	}

	for i := 0; i < len(actualRoute)-1; i++ {
		routeLength += distanceMap[common.NewPoint(int(actualRoute[i]), int(actualRoute[i+1]))]
	}

	return routeLength
}

func getDistances(locations map[int]common.Point, maze mazeMap) map[common.Point]int {
	distances := make(map[common.Point]int)
	for ka, va := range locations {
		for kb, vb := range locations {

			if ka == kb || va.Equals(vb) {
				continue
			}

			distances[common.NewPoint(ka, kb)] = distance(va, vb, maze)
		}
	}

	return distances
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {

	maze, locations := getMaze(inputData, true)

	distanceMap := getDistances(locations, maze)

	vals := make([]int64, 0)

	for k := range locations {
		if k == 0 {
			continue
		}
		vals = append(vals, int64(k))
	}
	perms := common.GetPermuations(vals)

	minVal := -1
	for _, perm := range perms {
		val := getRouteLength(perm, distanceMap, false)

		if minVal == -1 || val < minVal {
			minVal = val
		}

	}

	return fmt.Sprintf("%v", minVal)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	maze, locations := getMaze(inputData, true)

	distanceMap := getDistances(locations, maze)

	vals := make([]int64, 0)

	for k := range locations {
		if k == 0 {
			continue
		}
		vals = append(vals, int64(k))
	}
	perms := common.GetPermuations(vals)

	minVal := -1
	for _, perm := range perms {
		val := getRouteLength(perm, distanceMap, true)

		if minVal == -1 || val < minVal {
			minVal = val
		}

	}

	return fmt.Sprintf("%v", minVal)
}
