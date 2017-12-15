package Day201701

import (
	"strconv"

	"github.com/pezza/AoC2017/Common"
)

type dayEntry bool

var Entry dayEntry

func (td dayEntry) Describe() (int, int, string) {
	return 2017, 01, "Inverse Captcha"
}

func (td dayEntry) PartOne(inputData string) (string, error) {
	total := 0
	loopSize := len(inputData)

	for index := range inputData {
		cmp := common.GetWrappedIndex(index, loopSize, 1)

		if inputData[index] == inputData[cmp] {
			intVal, _ := strconv.Atoi(string(inputData[cmp]))
			total += intVal
		}

	}
	return strconv.Itoa(total), nil
}

func (td dayEntry) PartTwo(inputData string) (string, error) {
	total := 0
	loopSize := len(inputData)

	for index := range inputData {
		cmp := common.GetWrappedIndex(index, loopSize, loopSize/2)

		if inputData[index] == inputData[cmp] {
			intVal, _ := strconv.Atoi(string(inputData[cmp]))
			total += intVal
		}

	}
	return strconv.Itoa(total), nil
}
