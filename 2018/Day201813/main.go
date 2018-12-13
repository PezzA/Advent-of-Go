package Day201813

import (
	"fmt"
	"io/ioutil"
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

	b, err := ioutil.ReadFile("grid.txt") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	input = string(b) // convert content to a 'string'

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

var testCartId = 0

func moveCart(c cart, grid []string, carts []cart) (cart, bool) {

	if c.id == testCartId {
		fmt.Printf("(%v,%v)\theading %v\t", c.x, c.y, c.direction)
	}

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

	if c.id == testCartId {
		fmt.Printf(" moves to (%v,%v)\t", c.x, c.y)
	}

	// See if this will collide with another cart
	for _, testCart := range carts {
		if testCart.x == c.x && testCart.y == c.y && c.id != testCart.id {
			fmt.Printf("%v CRASHED WITH %v !\n", c, testCart)
			return c, true
		}
	}

	if c.id == testCartId {
		fmt.Printf("finds %v ", string(grid[c.y][c.x]))
	}

	// see if we need to change direction
	switch grid[c.y][c.x] {
	case 32:
		if c.id == testCartId {
			fmt.Println("p")
		}
	case 45:
		if c.id == testCartId && (c.direction == 0 || c.direction == 180) {
			fmt.Println("wrong -")
		}
	case 124:
		if c.id == testCartId && (c.direction == 90 || c.direction == 270) {
			fmt.Println("wrong |")
		}
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
			if c.id == testCartId {
				fmt.Print("L")
			}
			c.state = 1
			c = c.turnLeft()
		case 1:
			if c.id == testCartId {
				fmt.Print("S")
			}
			c.state = 2
		case 2:
			if c.id == testCartId {
				fmt.Print("R")
			}
			c.state = 0
			c = c.turnRight()
		}

	}

	if c.id == testCartId {
		fmt.Printf("\tand is now heading %v\n", c.direction)
	}

	return c, false
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	return fmt.Sprintf(" -- Not Yet Implemented --")
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	return fmt.Sprintf(" -- Not Yet Implemented --")
}
