package Day201506

import (
	"bufio"
	"regexp"
	"strconv"
	"strings"
)

type dayEntry bool

var Entry dayEntry

func (td dayEntry) Describe() (int, int, string) {
	return 2015, 6, "Probably a Fire Hazard"
}

type instruction struct {
	ins string
	x1  int
	y1  int
	x2  int
	y2  int
}

type grid [][]int

var instructionRegex = regexp.MustCompile(`(toggle|turn on|turn off) ([0-9]+),([0-9]+) through ([0-9]+),([0-9]+)`)

func getInstructions(input string) []instruction {
	ins := make([]instruction, 0)

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()

		matches := instructionRegex.FindStringSubmatch(line)

		x1, _ := strconv.Atoi(matches[2])
		y1, _ := strconv.Atoi(matches[3])
		x2, _ := strconv.Atoi(matches[4])
		y2, _ := strconv.Atoi(matches[5])

		ins = append(ins, instruction{matches[1], x1, y1, x2, y2})
	}

	return ins
}

func getGrid(x int, y int) grid {
	grid := make(grid, y)

	for index := range grid {
		grid[index] = make([]int, x)
		for subindex := range grid[index] {
			grid[index][subindex] = 0
		}
	}
	return grid
}

func (g grid) turnOff(x1 int, y1 int, x2 int, y2 int) grid {
	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			g[y][x] = 0

		}
	}
	return g
}

func (g grid) turnOn(x1 int, y1 int, x2 int, y2 int) grid {
	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			g[y][x] = 1

		}
	}
	return g
}

func (g grid) toggle(x1 int, y1 int, x2 int, y2 int) grid {
	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			if g[y][x] == 0 {
				g[y][x] = 1
			} else {
				g[y][x] = 0
			}

		}
	}
	return g
}

func (g grid) shift(x1 int, y1 int, x2 int, y2 int, shift int) grid {
	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			g[y][x] += shift

			if g[y][x] < 0 {
				g[y][x] = 0
			}
		}
	}
	return g
}

func (g grid) countLights() int {
	lights := 0
	for _, line := range g {
		for _, light := range line {
			lights += light
		}
	}
	return lights
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	grid, instructions := getGrid(1000, 1000), getInstructions(inputData)

	for _, ins := range instructions {
		switch ins.ins {
		case "toggle":
			grid = grid.toggle(ins.x1, ins.y1, ins.x2, ins.y2)
		case "turn on":
			grid = grid.turnOn(ins.x1, ins.y1, ins.x2, ins.y2)
		case "turn off":
			grid = grid.turnOff(ins.x1, ins.y1, ins.x2, ins.y2)
		}
	}

	return strconv.Itoa(grid.countLights())
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	grid, instructions := getGrid(1000, 1000), getInstructions(inputData)

	for _, ins := range instructions {
		switch ins.ins {
		case "toggle":
			grid = grid.shift(ins.x1, ins.y1, ins.x2, ins.y2, 2)
		case "turn on":
			grid = grid.shift(ins.x1, ins.y1, ins.x2, ins.y2, 1)
		case "turn off":
			grid = grid.shift(ins.x1, ins.y1, ins.x2, ins.y2, -1)
		}
	}

	return strconv.Itoa(grid.countLights())
}
