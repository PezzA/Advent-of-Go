package Day202001

import (
	"fmt"
	"strconv"
	"strings"
)

// Assumption - lines will always be integers.
func getData(input string) []int {
	data := make([]int, 0)

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		val, _ := strconv.Atoi(line)
		data = append(data, val)
	}

	return data
}

// Assumption - Value will be found
func findSumParts(input []int) (int, int) {
	for _, cmp1 := range input {
		for _, cmp2 := range input {
			if cmp1+cmp2 == 2020 {
				return cmp1, cmp2
			}
		}
	}
	return 0, 0
}

// Assumption - Value will be found
func findSumPartsTrio(input []int) (int, int, int) {
	for _, cmp1 := range input {
		for _, cmp2 := range input {
			for _, cmp3 := range input {
				if cmp1+cmp2+cmp3 == 2020 {
					return cmp1, cmp2, cmp3
				}
			}
		}
	}
	return 0, 0, 0
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	cmp1, cmp2 := findSumParts(getData(inputData))
	return fmt.Sprintf("%v", cmp1*cmp2)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	cmp1, cmp2, cmp3 := findSumPartsTrio(getData(inputData))
	return fmt.Sprintf("%v", cmp1*cmp2*cmp3)
}
