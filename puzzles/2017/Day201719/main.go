package Day201719

import (
	"bufio"
	"strconv"
	"strings"
)

type dayEntry bool

var Entry dayEntry

func (td dayEntry) Describe() (int, int, string) {
	return 2017, 19, "A Series of Tubes"
}

func findEntryPoint(input string) point {
	return point{strings.Index(input, "|"), 0}
}

type point struct {
	x int
	y int
}

func inBounds(cp point, maxX int, maxY int) bool {
	return cp.x > 0 && cp.x <= maxX && cp.y > 0 && cp.y <= maxY
}

func walkMaze(maze []string) (string, int) {
	cp := findEntryPoint(maze[0])
	orientation := 180
	seenLetters := ""

	maxX := len(maze[0]) - 1
	maxY := len(maze) - 1

	steps := 0

	atEnd := false
	for !atEnd {
		cp = translate(orientation, 1, cp)
		steps++

		cell := maze[cp.y][cp.x]

		switch cell {
		case 43:
			// think about changing direction
			if orientation == 180 || orientation == 0 {
				right := cellToMyRight(cp, orientation)
				if inBounds(right, maxX, maxY) {
					if maze[right.y][right.x] != 32 && maze[right.y][right.x] != 124 {
						orientation = orientate("R", orientation)
						break
					}
				}

				left := cellToMyLeft(cp, orientation)
				if inBounds(left, maxX, maxY) {

					if maze[left.y][left.x] != 32 && maze[left.y][left.x] != 124 {
						orientation = orientate("L", orientation)
						break
					}
				}
			}

			if orientation == 90 || orientation == 270 {
				right := cellToMyRight(cp, orientation)
				if inBounds(right, maxX, maxY) {
					if maze[right.y][right.x] != 32 && maze[right.y][right.x] != 45 {
						orientation = orientate("R", orientation)
						break
					}
				}

				left := cellToMyLeft(cp, orientation)
				if inBounds(left, maxX, maxY) {
					if maze[left.y][left.x] != 32 && maze[left.y][left.x] != 45 {
						orientation = orientate("L", orientation)
						break
					}
				}
			}

		case 45:
			// ignore
		case 124:
			// ignore
		case 32:
			atEnd = true
		default:
			seenLetters += string(cell)
		}
	}

	return seenLetters, steps
}

func translate(orientation int, shift int, startPoint point) point {
	switch orientation {
	case 0:
		return point{startPoint.x, startPoint.y - shift}
	case 90:
		return point{startPoint.x - shift, startPoint.y}
	case 180:
		return point{startPoint.x, startPoint.y + shift}
	case 270:
		return point{startPoint.x + shift, startPoint.y}
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

func cellToMyLeft(startPoint point, orientation int) point {
	leftOrientation := orientate("L", orientation)
	return translate(leftOrientation, 1, startPoint)
}

func cellToMyRight(startPoint point, orientation int) point {
	leftOrientation := orientate("R", orientation)
	return translate(leftOrientation, 1, startPoint)
}

func getMaze(input string) []string {
	maze := make([]string, 0)
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()

		if strings.TrimSpace(line) != "" {
			maze = append(maze, line)
		}
	}

	return maze
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	letters, _ := walkMaze(getMaze(inputData))
	return letters
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	_, steps := walkMaze(getMaze(inputData))
	return strconv.Itoa(steps)
}
