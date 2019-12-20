package Day201920

import (
	"fmt"
	"strings"

	"github.com/pezza/advent-of-code/common"
)

type maze [][]rune

type portal struct {
	pos  common.Point
	name string
}

const floor = 46
const wall = 35

func printMaze(in maze) {
	for _, line := range in {
		for _, cell := range line {
			if cell == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print(string(cell))
			}
		}
		fmt.Println()
	}
}
func getData(input string) (maze, []portal, common.Point, common.Point) {
	lines := strings.Split(input, "\n")
	newMaze := make(maze, len(lines))

	for y, line := range lines {
		newMaze[y] = make([]rune, len(line))
		for x, char := range line {
			if char == wall || char == floor {
				newMaze[y][x] = char
			}
		}
	}

	portals := make([]portal, 0)

	for y := range newMaze {
		for x := range newMaze[y] {
			if newMaze[y][x] == floor {
				if newMaze[y-1][x] == 0 {
					portals = append(portals, portal{
						pos: common.Point{X: x, Y: y},
						name: fmt.Sprintf("%v%v",
							string(lines[y-2][x]),
							string(lines[y-1][x])),
					})
				}

				if newMaze[y+1][x] == 0 {
					portals = append(portals, portal{
						pos: common.Point{X: x, Y: y},
						name: fmt.Sprintf("%v%v",
							string(lines[y+1][x]),
							string(lines[y+2][x])),
					})
				}
				if newMaze[y][x-1] == 0 {
					portals = append(portals, portal{
						pos: common.Point{X: x, Y: y},
						name: fmt.Sprintf("%v%v",
							string(lines[y][x-2]),
							string(lines[y][x-1])),
					})
				}
				if newMaze[y][x+1] == 0 {
					portals = append(portals, portal{
						pos: common.Point{X: x, Y: y},
						name: fmt.Sprintf("%v%v",
							string(lines[y][x+1]),
							string(lines[y][x+2])),
					})
				}
			}
		}
	}

	var start, end common.Point
	var actualPortals []portal

	for _, p := range portals {
		if p.name == "AA" {
			start = p.pos
			continue
		}

		if p.name == "ZZ" {
			end = p.pos
			continue
		}

		actualPortals = append(actualPortals, p)
	}
	return newMaze, actualPortals, start, end
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	return fmt.Sprintf("%v", " -- Not Yet Implemented --")
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	return fmt.Sprintf("%v", " -- Not Yet Implemented --")
}
