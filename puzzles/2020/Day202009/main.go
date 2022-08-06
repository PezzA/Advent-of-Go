package Day202009

import (
	"fmt"
	"strconv"
	"strings"
)

func getData(input string) []int {
	lines := strings.Split(input, "\n")
	data := make([]int, len(lines))

	for i := range lines {
		val, _ := strconv.Atoi(lines[i])
		data[i] = val
	}

	return data
}

func getPreamble(index int, values []int, preamble int) []int {
	return values[index-preamble : index]
}

func checkValidIndex(val int, preamble []int) bool {
	for i := range preamble {
		for j := range preamble {
			if preamble[i]+preamble[j] == val && preamble[i] != preamble[j] {
				return true
			}
		}
	}
	return false
}

func checkSequence(values []int, preambleLength int) int {
	for i := preambleLength; i < len(values); i++ {
		pa := getPreamble(i, values, preambleLength)

		if !checkValidIndex(values[i], pa) {
			return values[i]
		}
	}

	return -1
}

func getWeaknessList(value int, cypher []int) []int {
	for i := range cypher {
		total := 0
		for j := i; j < len(cypher); j++ {
			total += cypher[j]

			if total == value {
				return cypher[i : j+1]
			}
		}

	}

	return []int{}
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	val := checkSequence(getData(inputData), 25)
	return fmt.Sprintf("%v", val)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	data := getData(inputData)
	val := checkSequence(data, 25)
	wl := getWeaknessList(val, data)

	min, max := -1, 0

	for i := range wl {
		if wl[i] < min || min == -1 {
			min = wl[i]
		}

		if wl[i] > max {
			max = wl[i]
		}
	}

	return fmt.Sprintf("%v", min+max)
}
