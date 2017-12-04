package Day201701

import (
	"strconv"
)

// Entry holds wraps the data and runner interfaces for this puzzle
var Entry testDay

type testDay bool

func (td testDay) PartOne(inputData string) (string, error) {
	total := 0
	loopSize := len(inputData)

	for index := range inputData {
		cmp := getWrappedIndex(index, loopSize, 1)

		if inputData[index] == inputData[cmp] {
			intVal, _ := strconv.Atoi(string(inputData[cmp]))
			total += intVal
		}

	}
	return strconv.Itoa(total), nil
}

// getWrappedIndex will take
func getWrappedIndex(currentPos int, loopsize int, modifier int) int {
	rawIndex := currentPos + modifier

	if rawIndex >= (loopsize) {
		rawIndex -= loopsize
	}

	return rawIndex
}

func (td testDay) PartTwo(inputData string) (string, error) {
	total := 0
	loopSize := len(inputData)

	for index := range inputData {
		cmp := getWrappedIndex(index, loopSize, loopSize/2)

		if inputData[index] == inputData[cmp] {
			intVal, _ := strconv.Atoi(string(inputData[cmp]))
			total += intVal
		}

	}
	return strconv.Itoa(total), nil
}

func (td testDay) Day() int {
	return 01
}

func (td testDay) Year() int {
	return 2017
}

func (td testDay) Title() string {
	return "Inverse Captcha"
}
