package Day202024

import (
	"fmt"
	"math"
	"regexp"
	"strings"

	"github.com/pezza/advent-of-code/common"
)

var pathSplitter = regexp.MustCompile(`(w)|(e)|(ne)|(nw)|(se)|(sw)`)

var hexMove = map[string]common.Point{
	"e":  {1, 0},
	"w":  {-1, 0},
	"ne": {1, -1},
	"nw": {0, -1},
	"se": {0, 1},
	"sw": {-1, 1},
}

type floor map[common.Point]bool

func getData(input string) [][]string {
	lines := strings.Split(input, "\n")
	data := make([][]string, len(lines))

	for i := range lines {
		data[i] = pathSplitter.FindAllString(lines[i], -1)
	}

	return data
}

func (f floor) flipTile(insList []string) {
	currPoint := common.Point{}
	for i := range insList {
		currPoint = currPoint.Add(hexMove[insList[i]])
	}
	if _, ok := f[currPoint]; ok {
		f[currPoint] = !f[currPoint]
	} else {
		f[currPoint] = true
	}
}

func (f floor) flipAllTiles(data [][]string) {
	for i := range data {
		f.flipTile(data[i])
	}
}

func (f floor) adjacentCount(p common.Point) int {
	tot := 0
	for _, v := range hexMove {
		testPoint := p.Add(v)

		if val, ok := f[testPoint]; ok {
			if val {
				tot++
			}
		}
	}

	return tot
}

func (f floor) automata() floor {
	topLeft, bottomRight := f.getBounds()

	newFloor := make(floor, 0)

	for x := topLeft.X - 2; x < bottomRight.X+2; x++ {
		for y := topLeft.Y - 2; y < bottomRight.Y+2; y++ {
			cp := common.Point{X: x, Y: y}
			if _, ok := f[cp]; !ok {
				f[cp] = false
			}

			n := f.adjacentCount(common.Point{X: x, Y: y})

			if f[cp] && (n == 0 || n > 2) {
				newFloor[cp] = false
				continue
			}

			if !f[cp] && n == 2 {
				newFloor[cp] = true
				continue
			}
			newFloor[cp] = f[cp]
		}
	}

	return newFloor
}

func (f floor) getBounds() (common.Point, common.Point) {
	minX, minY := math.MaxInt32, math.MaxInt32
	maxX, maxY := math.MinInt32, math.MinInt32

	for k := range f {
		if k.X > maxX {
			maxX = k.X
		}

		if k.X < minX {
			minX = k.X
		}

		if k.Y > maxY {
			maxY = k.Y
		}

		if k.Y < minY {
			minY = k.Y
		}
	}

	return common.Point{X: minX, Y: minY}, common.Point{X: maxX, Y: maxY}
}

func (f floor) countBlackTiles() int {
	tot := 0
	for _, v := range f {
		if v {
			tot++
		}
	}
	return tot
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	data := getData(inputData)

	f := make(floor, 0)

	f.flipAllTiles(data)

	return fmt.Sprintf("%v", f.countBlackTiles())
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	data := getData(inputData)
	f := make(floor, 0)

	f.flipAllTiles(data)

	for i := 0; i < 100; i++ {
		f = f.automata()
	}

	return fmt.Sprintf("%v", f.countBlackTiles())
}
