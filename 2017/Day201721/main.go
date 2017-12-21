package Day201721

import (
	"bufio"
	"fmt"
	"strings"
)

type dayEntry bool

var Entry dayEntry

type enchanceRule struct {
	inputPattern  []string
	outputPattern []string
}

func getMatchBook(inputData string) []enchanceRule {
	matchBook := make([]enchanceRule, 0)
	scanner := bufio.NewScanner(strings.NewReader(inputData))
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, " => ")

		inputs := strings.Split(parts[0], "/")
		outputs := strings.Split(parts[1], "/")

		matchBook = append(matchBook, enchanceRule{inputs, outputs})
	}

	return matchBook
}

func cmpCells(source []string, cmp []string) bool {
	for index := range source {
		if source[index] != cmp[index] {
			return false
		}
	}
	return true
}

func printGrid(input []string) {
	for _, val := range input {
		fmt.Println(val)
	}
	fmt.Println("")
}

func loopCells(input []string, rules []enchanceRule) []string {
	cellSize := 0

	if len(input[0])%2 == 0 {
		cellSize = 2
	} else if len(input[0])%3 == 0 {
		cellSize = 3
	}

	if cellSize == 0 {
		return input
	}

	newSize := (len(input) / cellSize) * (cellSize + 1)
	newGrid := make([]string, newSize)

	for y := 0; y*cellSize < len(input); y++ {
		for x := 0; x*cellSize < len(input); x++ {
			cell := getCell(cellSize, input, x*cellSize, y*cellSize)

			newCell := getEnhancementRule(cell, rules)

			newY := y * (cellSize + 1)
			newGrid[newY] += newCell[0]
			newGrid[newY+1] += newCell[1]
			newGrid[newY+2] += newCell[2]

			if cellSize == 3 {
				newGrid[newY+3] += newCell[3]
			}
		}
	}

	return newGrid
}

func getEnhancementRule(input []string, rules []enchanceRule) []string {
	for _, value := range rules {
		if cmpCells(input, value.inputPattern) {
			return value.outputPattern
		}

		if cmpCells(rotate90(input), value.inputPattern) {
			return value.outputPattern
		}

		if cmpCells(rotate90(rotate90(input)), value.inputPattern) {
			return value.outputPattern
		}

		if cmpCells(rotate90(rotate90(rotate90(input))), value.inputPattern) {
			return value.outputPattern
		}

		if cmpCells(flipX(input), value.inputPattern) {
			return value.outputPattern
		}

		if cmpCells(flipX(rotate90(input)), value.inputPattern) {
			return value.outputPattern
		}

		if cmpCells(flipX(rotate90(rotate90(input))), value.inputPattern) {
			return value.outputPattern
		}

		if cmpCells(flipX(rotate90(rotate90(rotate90(input)))), value.inputPattern) {
			return value.outputPattern
		}

		if cmpCells(flipY(input), value.inputPattern) {
			return value.outputPattern
		}

		if cmpCells(flipY(rotate90(input)), value.inputPattern) {
			return value.outputPattern
		}

		if cmpCells(flipY(rotate90(rotate90(input))), value.inputPattern) {
			return value.outputPattern
		}

		if cmpCells(flipY(rotate90(rotate90(rotate90(input)))), value.inputPattern) {
			return value.outputPattern
		}
	}

	return []string{"NOT FOUND"}
}

func countLights(input []string) int {
	lights := 0
	for _, line := range input {
		for _, char := range line {
			if string(char) == "#" {
				lights++
			}
		}
	}
	return lights
}

func startGrid() []string {
	return []string{".#.", "..#", "###"}
}

func flipY(input []string) []string {
	if len(input[0]) == 2 {
		return []string{input[1], input[0]}
	}

	return []string{input[2], input[1], input[0]}
}

func flipX(input []string) []string {
	if len(input[0]) == 2 {
		bytes0 := []byte{input[0][1], input[0][0]}
		bytes1 := []byte{input[1][1], input[1][0]}

		return []string{string(bytes0), string(bytes1)}
	}

	bytes0 := []byte{input[0][2], input[0][1], input[0][0]}
	bytes1 := []byte{input[1][2], input[1][1], input[1][0]}
	bytes2 := []byte{input[2][2], input[2][1], input[2][0]}
	return []string{string(bytes0), string(bytes1), string(bytes2)}
}

func rotate90(input []string) []string {
	if len(input[0]) == 2 {
		bytes0 := []byte{input[1][0], input[0][0]}
		bytes1 := []byte{input[1][1], input[0][1]}

		return []string{string(bytes0), string(bytes1)}
	}

	bytes0 := []byte{input[2][0], input[1][0], input[0][0]}
	bytes1 := []byte{input[2][1], input[1][1], input[0][1]}
	bytes2 := []byte{input[2][2], input[1][2], input[0][2]}
	return []string{string(bytes0), string(bytes1), string(bytes2)}
}

func getCell(cellSize int, input []string, startX int, startY int) []string {
	cell := make([]string, 0)

	for y := startY; y < startY+cellSize; y++ {

		line := ""
		for x := startX; x < startX+cellSize; x++ {
			line += string(input[y][x])
		}

		cell = append(cell, line)
	}

	return cell
}

func (td dayEntry) Describe() (int, int, string) {
	return 2017, 21, "Fractal Art"
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	grid, rules := startGrid(), getMatchBook(inputData)
	grid = loopCells(grid, rules)
	grid = loopCells(grid, rules)
	grid = loopCells(grid, rules)
	grid = loopCells(grid, rules)
	grid = loopCells(grid, rules)

	lights := countLights(grid)

	return fmt.Sprintf("%v", lights)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	grid, rules := startGrid(), getMatchBook(inputData)
	grid = loopCells(grid, rules)
	updateChan <- []string{"1"}
	grid = loopCells(grid, rules)
	updateChan <- []string{"2"}
	grid = loopCells(grid, rules)
	updateChan <- []string{"3"}
	grid = loopCells(grid, rules)
	updateChan <- []string{"4"}
	grid = loopCells(grid, rules)
	updateChan <- []string{"5"}
	grid = loopCells(grid, rules)
	updateChan <- []string{"6"}
	grid = loopCells(grid, rules)
	updateChan <- []string{"7"}
	grid = loopCells(grid, rules)
	updateChan <- []string{"8"}
	grid = loopCells(grid, rules)
	updateChan <- []string{"9"}
	grid = loopCells(grid, rules)
	updateChan <- []string{"10"}
	grid = loopCells(grid, rules)
	updateChan <- []string{"11"}
	grid = loopCells(grid, rules)
	updateChan <- []string{"12"}
	grid = loopCells(grid, rules)
	updateChan <- []string{"13"}
	grid = loopCells(grid, rules)
	updateChan <- []string{"14"}
	grid = loopCells(grid, rules)
	updateChan <- []string{"15"}
	grid = loopCells(grid, rules)
	updateChan <- []string{"16"}
	grid = loopCells(grid, rules)
	updateChan <- []string{"17"}
	grid = loopCells(grid, rules)
	updateChan <- []string{"18"}
	lights := countLights(grid)

	return fmt.Sprintf("%v", lights)
}
