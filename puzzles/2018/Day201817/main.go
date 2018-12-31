package Day201817

import (
	"fmt"
	"strings"

	"github.com/pezza/advent-of-code/puzzles/Common"
)

var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string) {
	return 2018, 17, "Reservoir Research"
}

type BoreHoleLog [][]common.Point

func (bhl BoreHoleLog) GetBoundaries() (common.Point, common.Point) {
	minx, miny, maxx, maxy := -1, -1, -1, -1
	for _, line := range bhl {
		for _, pos := range line {
			if minx == -1 || pos.X < minx {
				minx = pos.X
			}
			if maxx == -1 || pos.X > maxx {
				maxx = pos.X
			}
			if miny == -1 || pos.Y < miny {
				miny = pos.Y
			}
			if maxy == -1 || pos.Y > maxy {
				maxy = pos.Y
			}
		}
	}
	return common.Point{X: minx, Y: miny}, common.Point{X: maxx, Y: maxy}
}

func getData(input string) BoreHoleLog {
	bhl := make(BoreHoleLog, 0)
	for _, line := range strings.Split(input, "\n") {
		cl := make([]common.Point, 0)
		if strings.HasPrefix(line, "x") {
			x, miny, maxy := 0, 0, 0
			fmt.Sscanf(line, "x=%v, y=%v..%v", &x, &miny, &maxy)
			for i := miny; i <= maxy; i++ {
				cl = append(cl, common.Point{X: x, Y: i})
			}

			bhl = append(bhl, cl)
			continue
		}

		y, minx, maxx := 0, 0, 0
		fmt.Sscanf(line, "y=%v, x=%v..%v", &y, &minx, &maxx)
		for i := minx; i <= maxx; i++ {
			cl = append(cl, common.Point{X: i, Y: y})
		}
		bhl = append(bhl, cl)
	}
	return bhl
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	return fmt.Sprintf(" -- Not Yet Implemented --")
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	return fmt.Sprintf(" -- Not Yet Implemented --")
}

func (td dayEntry) Page() int { return 45 }

func (td dayEntry) Start(inputData string) interface{} {
	if inputData == "" {
		inputData = Entry.PuzzleInput()
	}
	return getData(inputData)
}

func (td dayEntry) Prev() interface{} {
	return 0
}

func (td dayEntry) Next() interface{} {
	return 0
}
func (td dayEntry) End() interface{} { return "" }
