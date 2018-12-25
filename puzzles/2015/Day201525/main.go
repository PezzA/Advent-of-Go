package Day201525

import "fmt"

// Entry holds wraps the data and runner interfaces for this puzzle
var Entry dayEntry

type dayEntry bool

func fillMatrix(length int) int {
	val := 20151125

	for x := 1; x < length; x++ {
		val = nextTick(val)
	}
	return val
}

func nextTick(input int) int {
	return (input * 252533) % 33554393
}

func getFirstColIndex(col int) int {
	accum := 0
	for i := 1; i <= col; i++ {
		accum += i
	}

	return accum
}

func getColVal(colIndex int, targetRow int) int {
	accum := getFirstColIndex(colIndex)

	for i := 1; i < targetRow; i++ {
		accum += (i - 1) + colIndex

	}
	return accum
}

func (td dayEntry) Describe() (int, int, string) {
	return 2015, 25, "Let It Snow"
}

// 6591724,21653676
func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	numNeeded := getColVal(3019, 3010)
	return fmt.Sprintf("%v", fillMatrix(numNeeded))
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	return fmt.Sprintf(" -- Not Yet Implemented --")
}
