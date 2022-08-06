package Day201705

import (
	"bufio"
	"strconv"
	"strings"
)

// Entry holds wraps the data and runner interfaces for this puzzle
var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string) {
	return 2017, 5, "A Maze of Twisty Trampolines, All Alike"
}

func doModifiedJumps(ins []int) int {
	index := 0
	jumps := 0

	for index < len(ins) {
		thisIns := ins[index]

		if thisIns >= 3 {
			ins[index]--
		} else {
			ins[index]++
		}

		index += thisIns
		jumps++
	}

	return jumps
}

func doJumps(ins []int) int {
	index := 0
	jumps := 0

	for index < len(ins) {
		thisIns := ins[index]
		ins[index]++
		index += thisIns
		jumps++
	}

	return jumps
}

func getInsList(input string) []int {
	ins := make([]int, 0)
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		newIns, _ := strconv.Atoi(scanner.Text())

		ins = append(ins, newIns)
	}

	return ins
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	insList := getInsList(inputData)

	exit := doJumps(insList)

	return strconv.Itoa(exit)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	insList := getInsList(inputData)

	exit := doModifiedJumps(insList)

	return strconv.Itoa(exit)
}
