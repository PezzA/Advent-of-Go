package Day201811

import (
	"fmt"

	"strconv"

	"github.com/pezza/advent-of-code/Common"
)

var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string) {
	return 2018, 11, "Chronal Charge "
}

func getData(input string) int {
	val, _ := strconv.Atoi(input)

	return val
}
func getPowerLevel(point common.Point, serial int) int {
	rackId := point.X + 10

	pl := rackId * point.Y

	pl = pl + serial

	pl = pl * rackId

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
	maxPower := -1
	var powerPoint common.Point
	for x := 1; x <= 298; x++ {
		for y := 1; y <= 298; y++ {
			testPoint := common.Point{x, y}
			squarePower := getPowerLevel(common.Point{x, y}, gridSerial)
			squarePower += getPowerLevel(common.Point{x + 1, y}, gridSerial)
			squarePower += getPowerLevel(common.Point{x + 2, y}, gridSerial)
			squarePower += getPowerLevel(common.Point{x, y + 1}, gridSerial)
			squarePower += getPowerLevel(common.Point{x + 1, y + 1}, gridSerial)
			squarePower += getPowerLevel(common.Point{x + 2, y + 1}, gridSerial)
			squarePower += getPowerLevel(common.Point{x, y + 2}, gridSerial)
			squarePower += getPowerLevel(common.Point{x + 1, y + 2}, gridSerial)
			squarePower += getPowerLevel(common.Point{x + 2, y + 2}, gridSerial)

			if maxPower == -1 || squarePower > maxPower {
				maxPower = squarePower
				powerPoint = testPoint
			}
		}
	}

	return fmt.Sprintf("%v,%v", powerPoint.X, powerPoint.Y)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {

	//	gridSerial := getData(Entry.PuzzleInput())
	//maxPower := -1
	var powerPoint common.Point
	squareSize := -1
	for z := 1; z <= 300; z++ {
		count := 0
		for x := 1; x <= z; x++ {
			for y := 1; y <= z; y++ {
				count++

			}
		}
		fmt.Println(z, count)
	}
	return fmt.Sprintf("%v,%v,%v", powerPoint.X, powerPoint.Y, squareSize)
}
