package Day201705

import (
	"bufio"
	"strconv"
	"strings"
)

// Entry holds wraps the data and runner interfaces for this puzzle
var Entry testDay

type testDay bool

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

func (td testDay) PartOne(inputData string) (string, error) {
	insList := getInsList(inputData)

	exit := doJumps(insList)

	return strconv.Itoa(exit), nil
}

func (td testDay) PartTwo(inputData string) (string, error) {
	insList := getInsList(inputData)

	exit := doModifiedJumps(insList)

	return strconv.Itoa(exit), nil
}

func (td testDay) Day() int {
	return 5
}

func (td testDay) Year() int {
	return 2017
}

func (td testDay) Title() string {
	return "A Maze of Twisty Trampolines, All Alike"
}
