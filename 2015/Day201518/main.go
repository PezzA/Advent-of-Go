package Day201518

import (
	"fmt"
	"strings"

	"github.com/pezza/advent-of-code/Common"
)

// Entry holds wraps the data and runner interfaces for this puzzle
var Entry dayEntry

type dayEntry bool

type display [][]int

func (d display) getCellValue(p common.Point) int {
	if p.X < 0 || p.Y < 0 || p.X > len(d)-1 || p.Y > len(d)-1 {
		return 0
	}
	return d[p.Y][p.X]
}

func (d display) getNeighbourCount(p common.Point) int {
	nc := 0

	nc += d.getCellValue(common.Point{p.X, p.Y + 1})
	nc += d.getCellValue(common.Point{p.X + 1, p.Y + 1})
	nc += d.getCellValue(common.Point{p.X + 1, p.Y})
	nc += d.getCellValue(common.Point{p.X + 1, p.Y - 1})
	nc += d.getCellValue(common.Point{p.X, p.Y - 1})
	nc += d.getCellValue(common.Point{p.X - 1, p.Y - 1})
	nc += d.getCellValue(common.Point{p.X, p.Y - 1})
	nc += d.getCellValue(common.Point{p.X + 1, p.Y - 1})

	return nc
}

func getData(input string) display {
	startingScreen := make(display, 0)
	for _, line := range strings.Split(input, "\n") {
		screenLine := make([]int, 0)
		for _, char := range line {

			val := 0
			if char == 35 {
				val = 1
			}

			screenLine = append(screenLine, val)
		}

		startingScreen = append(startingScreen, screenLine)
	}

	return startingScreen
}

func (td dayEntry) Describe() (int, int, string) {
	return 2015, 18, "Like a GIF For Your Yard"
}
func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	return fmt.Sprintf(" -- Not Yet Implemented --")
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	return fmt.Sprintf(" -- Not Yet Implemented --")
}
