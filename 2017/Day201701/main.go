package Day201701

import (
	"strconv"

	"github.com/pezza/advent-of-code/common"
)

type dayEntry bool

var Entry dayEntry

func (td dayEntry) Describe() (int, int, string) {
	return 2017, 1, "Inverse Captcha"
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	total := 0
	loopSize := len(inputData)

	for index := range inputData {
		cmp := common.GetWrappedIndex(index, loopSize, 1)

		if inputData[index] == inputData[cmp] {
			intVal, _ := strconv.Atoi(string(inputData[cmp]))
			total += intVal
		}

	}
	return strconv.Itoa(total)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	total := 0
	loopSize := len(inputData)

	for index := range inputData {
		cmp := common.GetWrappedIndex(index, loopSize, loopSize/2)

		if inputData[index] == inputData[cmp] {
			intVal, _ := strconv.Atoi(string(inputData[cmp]))
			total += intVal
		}

	}
	return strconv.Itoa(total)
}
