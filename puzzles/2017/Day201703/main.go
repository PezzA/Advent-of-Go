package Day201703

import (
	"strconv"
)

var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string, int) {
	return 2017, 3, "Spiral Memory"
}

type point struct {
	x int
	y int
}

func translate(orientation int, shift int, startPoint point) point {
	switch orientation {
	case 0:
		return point{startPoint.x, startPoint.y + shift}
	case 90:
		return point{startPoint.x + shift, startPoint.y}
	case 180:
		return point{startPoint.x, startPoint.y - shift}
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

func squareToMyLeft(startPoint point, orientation int) point {
	leftOrientation := orientate("L", orientation)
	return translate(leftOrientation, 1, startPoint)
}

func intAbs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}

func getSpiral(size int, adjacentFill bool) (map[point]int, int) {
	x := 0
	y := 0
	orientation := 180

	smap := make(map[point]int, 0)

	for i := 1; i < size; i++ {
		pos := point{x, y}

		if !adjacentFill || i == 1 {
			smap[pos] = i
		} else {
			cellValue := 0
			cellValue += getAdjacentValue(pos.x, pos.y+1, smap)
			cellValue += getAdjacentValue(pos.x+1, pos.y+1, smap)
			cellValue += getAdjacentValue(pos.x+1, pos.y, smap)
			cellValue += getAdjacentValue(pos.x+1, pos.y-1, smap)
			cellValue += getAdjacentValue(pos.x, pos.y-1, smap)
			cellValue += getAdjacentValue(pos.x-1, pos.y-1, smap)
			cellValue += getAdjacentValue(pos.x-1, pos.y, smap)
			cellValue += getAdjacentValue(pos.x-1, pos.y+1, smap)
			smap[pos] = cellValue

			if cellValue > size {
				return smap, cellValue
			}
		}

		if _, ok := smap[squareToMyLeft(pos, orientation)]; !ok {
			orientation = orientate("L", orientation)
		}

		newPos := translate(orientation, 1, pos)

		x = newPos.x
		y = newPos.y
	}

	return smap, 0
}

func getAdjacentValue(x int, y int, smap map[point]int) int {
	if value, ok := smap[point{x, y}]; ok {
		return value
	}
	return 0
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	size, _ := strconv.Atoi(inputData)

	spmap, _ := getSpiral(size+1, false)

	for key, value := range spmap {
		if value == size {
			return strconv.Itoa(intAbs(key.x) + intAbs(key.y))
		}
	}

	return "Should not have got here"
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	size, _ := strconv.Atoi(inputData)

	_, targetVal := getSpiral(size+1, true)

	return strconv.Itoa(targetVal)
}

func (td dayEntry) PuzzleInput() string {
	return "312051"
}
