package Day201811

import (
	"fmt"

	"strconv"

	"github.com/pezza/advent-of-code/puzzles/common"
)

// Entry is the type that implements the runner interface for this puzzle
var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string, int) {
	return 2018, 11, "Chronal Charge "
}

func getData(input string) int {
	val, _ := strconv.Atoi(input)
	return val
}

func getPowerLevel(point common.Point, serial int) int {
	rackID := point.X + 10

	pl := rackID * point.Y

	pl = pl + serial

	pl = pl * rackID

	hundreds := 0

	if pl > 99 {
		plString := strconv.Itoa(pl)

		charIndex := len(plString) - 3
		hundredString := string(plString[charIndex])

		hundreds, _ = strconv.Atoi(hundredString)
	}

	return hundreds - 5
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	gridSerial := getData(Entry.PuzzleInput())

	computedGrid := make([][]int, 300)

	for index := range computedGrid {
		computedGrid[index] = make([]int, 300)
	}

	for index := range computedGrid {
		for subIndex := range computedGrid[index] {
			computedGrid[index][subIndex] = getPowerLevel(common.Point{X: index + 1, Y: subIndex + 1}, gridSerial)
		}
	}

	maxPower := -1
	var powerPoint common.Point
	for x := 1; x <= 298; x++ {
		for y := 1; y <= 298; y++ {
			testPoint := common.Point{X: x, Y: y}
			squarePower := computedGrid[x-1][y-1]
			squarePower += computedGrid[x][y-1]
			squarePower += computedGrid[x+1][y-1]
			squarePower += computedGrid[x-1][y]
			squarePower += computedGrid[x][y]
			squarePower += computedGrid[x+1][y]
			squarePower += computedGrid[x-1][y+1]
			squarePower += computedGrid[x][y+1]
			squarePower += computedGrid[x+1][y+1]

			if maxPower == -1 || squarePower > maxPower {
				maxPower = squarePower
				powerPoint = testPoint
			}
		}
	}

	return fmt.Sprintf("%v,%v", powerPoint.X, powerPoint.Y)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	gridSerial := getData(Entry.PuzzleInput())
	//gridSerial := 18
	maxPower := -1
	var powerPoint common.Point
	squareSize := -1

	computedGrid := make([][]int, 300)

	for index := range computedGrid {
		computedGrid[index] = make([]int, 300)
	}

	for index := range computedGrid {
		for subIndex := range computedGrid[index] {
			computedGrid[index][subIndex] = getPowerLevel(common.Point{X: index + 1, Y: subIndex + 1}, gridSerial)
		}
	}
	//for each size
	for z := 1; z <= 300; z++ {
		count := 0
		offSet := 301 - z

		// for each instance of that size
		for x := 1; x <= offSet; x++ {
			for y := 1; y <= offSet; y++ {
				count++
				squareCount := 0

				// for each cell in that size
				for zx := x; zx < x+z; zx++ {
					for zy := y; zy < y+z; zy++ {
						squareCount += computedGrid[zx-1][zy-1]
					}
				}

				if maxPower == -1 || squareCount > maxPower {
					maxPower = squareCount
					powerPoint = common.Point{X: x, Y: y}
					squareSize = z
				}
			}
		}
	}

	return fmt.Sprintf("%v,%v,%v", powerPoint.X, powerPoint.Y, squareSize)
}
