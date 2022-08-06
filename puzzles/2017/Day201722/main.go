package Day201722

import (
	"bufio"
	"fmt"
	"strings"
)

func (td dayEntry) Describe() (int, int, string, int) {
	return 2017, 22, "Sporifica Virus", 2
}

type dayEntry bool

var Entry dayEntry

type grid map[point]int

type point struct {
	x int
	y int
}

func translate(orientation int, shift int, startPoint point) point {
	switch orientation {
	case 0:
		return point{startPoint.x, startPoint.y - shift}
	case 90:
		return point{startPoint.x + shift, startPoint.y}
	case 180:
		return point{startPoint.x, startPoint.y + shift}
	case 270:
		return point{startPoint.x - shift, startPoint.y}
	}

	return point{0, 0}
}

func orientate(direction string, currentOrientation int) int {
	switch direction {
	case "L":
		if currentOrientation == 0 {
			return 270
		}
		return currentOrientation - 90
	case "R":
		if currentOrientation == 270 {
			return 0
		}
		return currentOrientation + 90
	default:
		return 0
	}
}

func getGrid(inputData string) (grid, int) {
	lines := make([]string, 0)
	scanner := bufio.NewScanner(strings.NewReader(inputData))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	offset := len(lines) / 2

	grid := make(grid, 0)
	for y, line := range lines {
		for x, char := range line {
			if char == '#' {
				py := y - offset
				px := x - offset
				grid[point{px, py}] = 2
			}
		}
	}

	return grid, len(lines)
}

func doWalk(bursts int, plane grid) (grid, int) {
	orientation, position, infectedCells := 0, point{0, 0}, 0

	for i := 0; i < bursts; i++ {
		if _, ok := plane[position]; ok {
			orientation = orientate("R", orientation)
			delete(plane, position)
		} else {
			orientation = orientate("L", orientation)
			plane[position] = 2
			infectedCells++
		}

		position = translate(orientation, 1, position)

		//printGrid(plane, 9, position)
	}
	return plane, infectedCells
}

// clean - [empty] // weakened == 1, infected == 2, flagged ==3
func doResistiveWalk(bursts int, plane grid) (grid, int) {
	orientation, position, infectedCells := 0, point{0, 0}, 0

	for i := 0; i < bursts; i++ {
		if val, ok := plane[position]; ok {
			switch val {
			case 1:
				plane[position] = 2
				infectedCells++
			case 2:
				orientation = orientate("R", orientation)
				plane[position] = 3
			case 3:
				orientation = orientate("L", orientation)
				orientation = orientate("L", orientation)
				delete(plane, position)
			}
		} else {
			orientation = orientate("L", orientation)
			plane[position] = 1

		}
		position = translate(orientation, 1, position)
		//	printGrid(plane, 12, position)
	}
	return plane, infectedCells
}

func printGrid(plane grid, size int, position point) {
	lines := make([]string, size)

	offSet := size / 2

	for index := range lines {
		lines[index] = strings.Repeat(".", size)
	}

	for key, value := range plane {
		y, x := key.y+offSet, key.x+offSet
		switch value {
		case 1:
			lines[y] = lines[y][:x] + "W" + lines[y][x+1:]
		case 2:
			lines[y] = lines[y][:x] + "#" + lines[y][x+1:]
		case 3:
			lines[y] = lines[y][:x] + "F" + lines[y][x+1:]
		}

	}

	posY, posX := position.y+offSet, position.x+offSet

	lines[posY] = lines[posY][:posX] + "X" + lines[posY][posX+1:]

	for index := range lines {
		fmt.Println(lines[index])
	}

	fmt.Println()
}

// 5258
func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	grid, _ := getGrid(inputData)
	var infectedCells int

	grid, infectedCells = doWalk(10000, grid)

	return fmt.Sprintf("%v", infectedCells)
}

// 2512719
func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	grid, _ := getGrid(inputData)
	var infectedCells int

	grid, infectedCells = doResistiveWalk(10000000, grid)

	return fmt.Sprintf("%v", infectedCells)
}
