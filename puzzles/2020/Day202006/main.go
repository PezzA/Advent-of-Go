package Day202006

import (
	"fmt"
	"strings"
)

type group []string

func getData(input string) []group {
	data := make([]group, 0)

	current := make(group, 0)
	for _, line := range strings.Split(input, "\n") {

		if line == "" {
			data = append(data, current)
			current = make(group, 0)
			continue
		}
		current = append(current, line)
	}
	data = append(data, current)

	return data
}

func (g group) getDistinctYes() int {
	questions := make(map[rune]int, 0)

	for _, passenger := range g {
		for _, question := range passenger {

			if _, ok := questions[question]; ok {
				questions[question]++
			} else {
				questions[question] = 1
			}
		}
	}

	return len(questions)
}

func (g group) getCollectiveYes() int {
	questions := make(map[rune]int, 0)

	for _, passenger := range g {
		for _, question := range passenger {

			if _, ok := questions[question]; ok {
				questions[question]++
			} else {
				questions[question] = 1
			}
		}
	}

	collective := 0

	for _, v := range questions {
		if v == len(g) {
			collective++
		}
	}

	return collective
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	total := 0

	for _, g := range getData(inputData) {
		total += g.getDistinctYes()
	}

	return fmt.Sprintf("%v", total)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	total := 0

	for _, g := range getData(inputData) {
		total += g.getCollectiveYes()
	}

	return fmt.Sprintf("%v", total)
}
