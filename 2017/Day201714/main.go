package Day201714

import (
	"fmt"
	"strconv"

	Day10 "github.com/pezza/AoC2017/2017/Day201710"
)

type dayEntry bool

var Entry dayEntry

func (td dayEntry) Describe() (int, int, string) {
	return 2017, 14, "Disk Defragmentation"
}

func stringHexToBinary(input string) string {
	switch input {
	case "0":
		return "0000"
	case "1":
		return "0001"
	case "2":
		return "0010"
	case "3":
		return "0011"
	case "4":
		return "0100"
	case "5":
		return "0101"
	case "6":
		return "0110"
	case "7":
		return "0111"
	case "8":
		return "1000"
	case "9":
		return "1001"
	case "a":
		return "1010"
	case "b":
		return "1011"
	case "c":
		return "1100"
	case "d":
		return "1101"
	case "e":
		return "1110"
	case "f":
		return "1111"
	}

	return ""
}

func getLines(input string, linecount int) []string {
	lines := make([]string, 0)

	for i := 0; i < linecount; i++ {
		hashSeed := strconv.Itoa(i)
		hash := Day10.DoFullHash(input + "-" + hashSeed)
		line := ""
		for _, val := range hash {
			binary := stringHexToBinary(string(val))
			line += binary
		}
		lines = append(lines, line)
	}

	return lines
}

func countChar(input string, char rune) int {
	count := 0
	for _, val := range input {
		if val == char {
			count++
		}
	}
	return count

}

type grid [][]int
type point struct {
	x int
	y int
}

func linesToArrays(lines []string) grid {
	grid := make(grid, 0)

	for index, val := range lines {
		grid = append(grid, []int{})
		for _, bit := range val {
			intVal, _ := strconv.Atoi(string(bit))
			grid[index] = append(grid[index], intVal)
		}
	}
	return grid
}

func (g grid) getNextOpenCell() (point, bool) {
	nextPos := point{}
	found := false
	for y, line := range g {
		for x, bit := range line {
			if bit == 1 {
				found = true
				nextPos = point{x, y}
			}
			if found {
				break
			}
		}
		if found {
			break
		}
	}

	return nextPos, found
}

func (g grid) flagRegion(regionTag int, start point) grid {
	if g[start.y][start.x] == 1 {
		g[start.y][start.x] = regionTag
	} else {
		return g
	}

	if start.y > 0 {
		g.flagRegion(regionTag, point{start.x, start.y - 1})
	}

	if start.y < (len(g) - 1) {
		g.flagRegion(regionTag, point{start.x, start.y + 1})
	}

	if start.x < (len(g[start.y]) - 1) {
		g.flagRegion(regionTag, point{start.x + 1, start.y})
	}

	if start.x > 0 {
		g.flagRegion(regionTag, point{start.x - 1, start.y})
	}

	return g
}

func (g grid) PrintGrid() {
	for _, l := range g {
		for _, bit := range l {
			fmt.Print(bit)
		}
		fmt.Print("\n")
	}
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) (string, error) {
	lines := getLines(inputData, 128)

	total := 0
	for _, val := range lines {

		total += countChar(val, rune('1'))
	}

	return strconv.Itoa(total), nil
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) (string, error) {

	grid := linesToArrays(getLines(inputData, 128))

	found := true
	startRegionCount := 0
	for found {
		var point point
		point, found = grid.getNextOpenCell()

		if found {
			grid = grid.flagRegion(startRegionCount+2, point)
			startRegionCount++
		}
	}

	return strconv.Itoa(startRegionCount), nil
}

func (td dayEntry) PuzzleInput() string {
	return "ugkiagan"
}
