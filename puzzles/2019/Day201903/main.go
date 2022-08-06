package Day201903

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pezza/advent-of-code/puzzles/common"
)

var Entry dayEntry

type dayEntry bool

type timedPlot struct {
	colour string
	time   int
}

type grid map[common.Point][]timedPlot

func getData(input string) ([]string, []string) {
	wires := strings.Split(input, "\n")
	return strings.Split(wires[0], ","), strings.Split(wires[1], ",")
}

func getVector(direction string) common.Point {
	switch direction {
	case "U":
		return common.Point{X: 0, Y: 1}
	case "D":
		return common.Point{X: 0, Y: -1}
	case "L":
		return common.Point{X: -1, Y: 0}
	case "R":
		return common.Point{X: 1, Y: 0}
	default:
		return common.Point{X: 0, Y: 0}
	}
}

func (g grid) addPathToGrid(start common.Point, direction string, colour string, startTime int) (common.Point, int) {
	v := getVector(string(direction[0]))

	mag, _ := strconv.Atoi(string(direction[1:]))

	for i := 0; i < mag; i++ {
		start = start.Add(v)
		startTime++
		if _, ok := g[start]; ok {
			g[start] = append(g[start], timedPlot{colour, startTime})
		} else {
			g[start] = []timedPlot{timedPlot{colour, startTime}}
		}
	}

	return start, startTime
}

func runWires(green []string, red []string) (int, int) {
	grid := make(grid, 0)

	start, startTime := common.Point{X: 0, Y: 0}, 0
	for _, d := range green {
		start, startTime = grid.addPathToGrid(start, d, "green", startTime)
	}

	start, startTime = common.Point{X: 0, Y: 0}, 0
	for _, d := range red {
		start, startTime = grid.addPathToGrid(start, d, "red", startTime)
	}

	closet, soonest := -1, -1
	for k, v := range grid {
		if len(v) > 1 {

			if containsWire("red", v) && containsWire("green", v) {
				dist := k.GetMDistanceOrigin()
				if closet == -1 || dist < closet {
					closet = dist
				}

				time := getFirstPlot("red", v) + getFirstPlot("green", v)
				if soonest == -1 || time < soonest {
					soonest = time
				}
			}
		}
	}

	return closet, soonest
}

// it's going to be the first one right? assume will always work
func getFirstPlot(colour string, l []timedPlot) int {
	for index := range l {
		if l[index].colour == colour {
			return l[index].time
		}
	}

	return 0
}

func containsWire(colour string, l []timedPlot) bool {
	for index := range l {
		if l[index].colour == colour {
			return true
		}
	}
	return false
}

func (td dayEntry) Describe() (int, int, string) {
	return 2019, 3, "Crossed Wires"
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	green, red := getData(inputData)
	closest, _ := runWires(green, red)
	return fmt.Sprintf("%v", closest)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	green, red := getData(inputData)
	_, soonest := runWires(green, red)
	return fmt.Sprintf("%v", soonest)
}
