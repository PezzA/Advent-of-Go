package Day202015

import (
	"fmt"
	"strconv"
	"strings"
)

func getData(input string) []int {
	lines := strings.Split(input, ",")
	data := make([]int, len(lines))

	for i := range lines {
		val, _ := strconv.Atoi(lines[i])

		data[i] = val
	}

	return data
}

func getNextNumber(seq []int, records map[int][]int, iteration int) ([]int, map[int][]int) {
	lastSpoken := seq[len(seq)-1]

	lastSeen := records[lastSpoken]

	if len(lastSeen) == 1 {
		seq = append(seq, 0)
		records[0] = append(records[0], iteration)
		return seq, records
	}

	previouslySpoken := lastSeen[len(lastSeen)-1]
	previouslySpokenBeforeThen := lastSeen[len(lastSeen)-2]

	newNumber := previouslySpoken - previouslySpokenBeforeThen
	seq = append(seq, newNumber)

	if _, ok := records[newNumber]; ok {
		records[newNumber] = append(records[newNumber], iteration)
	} else {
		records[newNumber] = []int{iteration}
	}

	return seq, records
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	data := getData(inputData)
	seens := make(map[int][]int, 0)

	for i := range data {
		seens[data[i]] = []int{i + 1}
	}

	turn := 2020
	turnScore := 0
	for i := len(data) + 1; i <= turn; i++ {
		data, seens = getNextNumber(data, seens, i)

		if i == turn {
			turnScore = data[len(data)-1]
		}
	}

	return fmt.Sprintf("%v", turnScore)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	data := getData(inputData)
	seens := make(map[int][]int, 0)

	for i := range data {
		seens[data[i]] = []int{i + 1}
	}

	turn := 30000000
	turnScore := 0
	for i := len(data) + 1; i <= turn; i++ {
		data, seens = getNextNumber(data, seens, i)

		if i == turn {
			turnScore = data[len(data)-1]
		}
	}

	return fmt.Sprintf("%v", turnScore)
}
