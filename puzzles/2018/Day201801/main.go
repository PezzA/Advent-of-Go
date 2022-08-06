package Day201801

import (
	"fmt"
	"strconv"
	"strings"
)

// Entry holds wraps the data and runner interfaces for this puzzle
var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string, int) {
	return 2018, 1, "Chronal Calibration", 0
}

func getData(input string) []int {
	vals := make([]int, 0)

	for _, val := range strings.Split(input, "\n") {
		testVal, _ := strconv.Atoi(val)
		vals = append(vals, testVal)
	}

	return vals
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	retVal := 0
	for _, value := range getData(inputData) {
		retVal += value
	}
	return fmt.Sprintf("%v", retVal)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	changes := getData(inputData)
	freqs := make(map[int]bool)

	retVal := 0
	found := false

	for !found {
		for _, value := range changes {
			retVal += value

			if _, ok := freqs[retVal]; ok {
				found = true
				break
			}

			freqs[retVal] = true
		}
	}
	return fmt.Sprintf("%v", retVal)
}
