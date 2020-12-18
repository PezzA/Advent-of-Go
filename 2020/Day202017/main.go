package Day202017

import (
	"fmt"
	"strings"
)

type dayEntry bool

type cellState string

const (
	inactive cellState = "."
	active   cellState = "#"
)

type point3D struct {
	x int
	y int
	z int
}

type universe map[point3D]cellState

var neighbours = []point3D{
	point3D{0, -1, 0},
	point3D{1, -1, 0},
	point3D{1, 0, 0},
	point3D{1, 1, 0},
	point3D{0, 1, 0},
	point3D{-1, 1, 0},
	point3D{-1, 0, 0},
	point3D{-1, -1, 0},
	point3D{0, -1, -1},
	point3D{1, -1, -1},
	point3D{1, 0, -1},
	point3D{1, 1, -1},
	point3D{0, 1, -1},
	point3D{-1, 1, -1},
	point3D{-1, 0, -1},
	point3D{-1, -1, -1},
	point3D{0, 0, -1},
	point3D{0, -1, 1},
	point3D{1, -1, 1},
	point3D{1, 0, 1},
	point3D{1, 1, 1},
	point3D{0, 1, 1},
	point3D{-1, 1, 1},
	point3D{-1, 0, 1},
	point3D{-1, -1, 1},
	point3D{0, 0, 1},
}

func (p point3D) Add(inc point3D) point3D {
	return point3D{
		x: p.x + inc.x,
		y: p.y + inc.y,
		z: p.z + inc.z,
	}
}
func getData(input string) universe {
	lines := strings.Split(input, "\n")
	universe := make(universe, 0)

	for y, line := range lines {
		for x, c := range line {
			universe[point3D{x, y, 0}] = cellState(c)
		}
	}

	return universe
}

func (u universe) neighbourCount(p point3D) int {
	count := 0

	for _, n := range neighbours {
		testPoint := p.Add(n)
		if val, ok := u[testPoint]; ok && val == active {
			count++
		}
	}
	return count
}

func (u universe) deepCopy() universe {
	newUniverse := make(universe, 0)

	for k, v := range u {
		newUniverse[k] = v
	}

	return newUniverse
}

func (u universe) cycle(xRange, yRange, zRange, iter int) universe {
	newU := u.deepCopy()
	for z := zRange - iter; z <= zRange+iter; z++ {
		for y := 0 - iter; y < yRange+iter; y++ {
			for x := 0 - iter; x < xRange+iter; x++ {

				point := point3D{x, y, z}
				state := inactive

				if val, ok := u[point]; ok {
					state = val
				}

				newU[point] = newState(state, u.neighbourCount(point))
			}
		}
	}

	return newU
}

func (u universe) getActiveCubes() int {
	tot := 0

	for _, v := range u {
		if v == active {
			tot++
		}
	}

	return tot
}

func newState(currentState cellState, neighbourCount int) cellState {
	if currentState == active {
		if neighbourCount == 2 || neighbourCount == 3 {
			return active
		} else {
			return inactive
		}

	}

	if currentState == inactive && neighbourCount == 3 {
		return active
	}

	return currentState
}

func (u universe) print(xRange, yRange, zRange, iter int) {
	for z := zRange - iter; z <= zRange+iter; z++ {
		fmt.Printf("z=%d\n", z)
		for y := 0 - iter; y < yRange+iter; y++ {
			for x := 0 - iter; x < xRange+iter; x++ {

				cell := inactive

				if val, ok := u[point3D{x, y, z}]; ok {
					cell = val
				}

				fmt.Print(cell)
			}
			fmt.Print("\n")
		}
		fmt.Println()
	}

}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	return fmt.Sprintf("%v", " -- Not Yet Implemented --")
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	return fmt.Sprintf("%v", " -- Not Yet Implemented --")
}
