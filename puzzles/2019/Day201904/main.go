package Day201904

import (
	"fmt"
	"strconv"
	"strings"
)

func getData(input string) (int, int) {
	data := strings.Split(input, "-")

	min, _ := strconv.Atoi(data[0])
	max, _ := strconv.Atoi(data[1])

	return min, max
}

func isValidPassword(input int) bool {
	rep := strconv.Itoa(input)

	prev, hasDouble, isAscending := byte(0), false, true
	for index := range rep {
		str := string(rep[index])
		cmp := string(prev)

		if cmp == str {
			hasDouble = true
		}

		if rep[index] < prev {
			isAscending = false
		}
		prev = rep[index]
	}

	return hasDouble && isAscending
}

func getChunks(input int) []string {
	rep := strconv.Itoa(input)

	chunks := []string{}
	prev := byte(0)
	current := ""
	for index := range rep {
		str := string(rep[index])
		cmp := string(prev)

		if cmp == str {
			current += str
		} else {
			if current != "" {
				chunks = append(chunks, current)
			}
			current = str
		}

		prev = rep[index]
	}

	if current != "" {
		chunks = append(chunks, current)
	}

	return chunks
}

func isValidPassword2(input int) bool {
	chunks := getChunks(input)

	prev, hasDouble, isAscending := byte(0), false, true
	for index := range chunks {
		if len(chunks[index]) == 2 {
			hasDouble = true
		}

		if chunks[index][0] < prev {
			isAscending = false
		}
		prev = chunks[index][0]
	}

	return hasDouble && isAscending
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	min, max := getData(inputData)

	count := 0
	for i := min; i <= max; i++ {
		if isValidPassword(i) {
			count++
		}
	}

	return fmt.Sprintf("%v", count)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	min, max := getData(Entry.PuzzleInput())

	count := 0
	for i := min; i <= max; i++ {
		if isValidPassword2(i) {
			count++
		}
	}

	return fmt.Sprintf("%v", count)
}
