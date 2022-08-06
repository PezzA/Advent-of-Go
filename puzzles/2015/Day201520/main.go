package Day201520

import (
	"fmt"
	"strconv"
)

// Entry holds wraps the data and runner interfaces for this puzzle
var Entry dayEntry

type dayEntry bool

func getData(input string) int {
	val, _ := strconv.Atoi(input)
	return val
}

func (td dayEntry) Describe() (int, int, string, int) {
	return 2015, 20, "Infinite Elves and Infinite Houses", 2
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	size := getData(inputData)

	houses := make([]int, size)
	for i := 1; i < (size / 10); i++ {
		multi := i * 10
		for j := i; j < (size / 10); j += i {
			houses[j] += multi
		}
	}

	house := 0
	for index, val := range houses {
		if val > size {
			house = index
			break
		}
	}

	return fmt.Sprintf("%v", house)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	size := getData(inputData)

	houses := make([]int, size)
	for i := 1; i < (size); i++ {
		multi := i * 11
		count := 0
		for j := i; j < (size); j += i {
			houses[j] += multi
			count++
			if count == 50 {
				break
			}
		}
	}

	house := 0
	for index, val := range houses {
		if val > size {
			house = index
			break
		}
	}

	return fmt.Sprintf("%v", house)
}
