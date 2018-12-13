package Day201813

import (
	"fmt"
	"strings"
)

var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string) {
	return 2018, 13, "Mine Cart Madness"
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

func moveCart(c cart, grid []string, carts []cart) (cart, bool) {
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
			return c, true
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

	return c, false
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	return fmt.Sprintf(" -- Not Yet Implemented --")
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	return fmt.Sprintf(" -- Not Yet Implemented --")
}
