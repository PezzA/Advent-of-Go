package Day201813

import (
	"fmt"
	"sort"
	"strings"

	"github.com/pezza/advent-of-code/puzzles/common"
)

var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string, int) {
	return 2018, 13, "Mine Cart Madness", 2
}

type cart struct {
	id        int
	direction int
	x         int
	y         int
	state     int
}

func getData(input string) ([]string, []cart) {
	lines := make([]string, 0)
	carts := make([]cart, 0)

	for y, line := range strings.Split(input, "\n") {
		lines = append(lines, line)

		for x, char := range line {
			if char == 60 {
				carts = append(carts, cart{len(carts), 270, x, y, 0})
			}
			if char == 62 {
				carts = append(carts, cart{len(carts), 90, x, y, 0})
			}
			if char == 118 {
				carts = append(carts, cart{len(carts), 180, x, y, 0})
			}
			if char == 94 {
				carts = append(carts, cart{len(carts), 0, x, y, 0})
			}
		}
	}

	return lines, carts
}

func (c cart) turnRight() cart {
	c.direction += 90
	if c.direction == 360 {
		c.direction = 0
	}
	return c
}

func (c cart) turnLeft() cart {
	c.direction -= 90
	if c.direction == -90 {
		c.direction = 270
	}
	return c
}

func moveCart(c cart, grid []string, carts []cart) (cart, cart, bool) {
	// get the location of where it would move next
	switch c.direction {
	case 0:
		c.y--
	case 180:
		c.y++
	case 270:
		c.x--
	case 90:
		c.x++
	}

	// See if this will collide with another cart
	for _, testCart := range carts {
		if testCart.x == c.x && testCart.y == c.y {
			return c, testCart, true
		}
	}

	// see if we need to change direction
	switch grid[c.y][c.x] {
	case 47: // "/"
		switch c.direction {
		case 0:
			c = c.turnRight()
		case 90:
			c = c.turnLeft()
		case 180:
			c = c.turnRight()
		case 270:
			c = c.turnLeft()
		}

	case 92: // "\"
		switch c.direction {
		case 0:
			c = c.turnLeft()
		case 90:
			c = c.turnRight()
		case 180:
			c = c.turnLeft()
		case 270:
			c = c.turnRight()
		}
	case 43: // "+"
		switch c.state {
		case 0:
			c.state = 1
			c = c.turnLeft()
		case 1:
			c.state = 2
		case 2:
			c.state = 0
			c = c.turnRight()
		}
	}

	return c, c, false
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	grid, carts := getData(Entry.PuzzleInput())

	crashed := false
	var point common.Point
	for {
		for index := range carts {

			carts[index], _, crashed = moveCart(carts[index], grid, carts)
			if crashed {
				point = common.Point{X: carts[index].x, Y: carts[index].y}
				break
			}
		}

		// sort the carts into x y order
		sort.Slice(carts, func(i, j int) bool {
			if carts[i].y == carts[j].y {
				return carts[i].x < carts[j].x
			}
			return carts[i].y < carts[j].y
		})

		if crashed {
			break
		}
	}

	return fmt.Sprintf("%v,%v", point.X, point.Y)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	grid, carts := getData(Entry.PuzzleInput())

	for {
		crashIds := make([]int, 0)
		for index := range carts {
			crashed := false
			var otherCrashedCart cart

			carts[index], otherCrashedCart, crashed = moveCart(carts[index], grid, carts)

			if crashed {
				crashIds = append(crashIds, carts[index].id)
				crashIds = append(crashIds, otherCrashedCart.id)
			}
		}

		if len(crashIds) > 0 {
			for _, val := range crashIds {
				foundIndex := 0
				for searchIndex := range carts {
					if carts[searchIndex].id == val {
						foundIndex = searchIndex
						break
					}
				}
				carts = append(carts[:foundIndex], carts[foundIndex+1:]...)
			}
		}

		if len(carts) == 1 {
			break
		}

		// sort the carts into x y order
		sort.Slice(carts, func(i, j int) bool {
			if carts[i].y == carts[j].y {
				return carts[i].x < carts[j].x
			}
			return carts[i].y < carts[j].y
		})
	}

	return fmt.Sprintf("%v,%v", carts[0].x, carts[0].y)
}
