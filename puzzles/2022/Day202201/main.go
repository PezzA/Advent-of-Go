package Day202201

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func loadData(input string) [][]int {
	lines := strings.Split(input, "\n")

	elves := make([][]int, 0)
	current := make([]int, 0)

	for _, line := range lines {
		if line == "" {
			elves = append(elves, current)
			current = make([]int, 0)
			continue
		}

		val, _ := strconv.Atoi(line)

		current = append(current, val)
	}

	elves = append(elves, current)
	return elves
}

func sum(input []int) int {
	retval := 0
	for _, v := range input {
		retval += v
	}

	return retval
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	elves := loadData(inputData)

	maxCal := 0

	for _, elf := range elves {
		calCount := sum(elf)

		if calCount > maxCal {
			maxCal = calCount
		}
	}

	return fmt.Sprint(maxCal)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	elves := loadData(inputData)

	totals := make([]int, 0)

	for _, elf := range elves {
		totals = append(totals, sum(elf))
	}

	sort.Slice(totals, func(i, j int) bool {
		return totals[i] > totals[j]
	})

	grandTotal := totals[0] + totals[1] + totals[2]
	return fmt.Sprint(grandTotal)
}
