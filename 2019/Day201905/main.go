package Day201905

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pezza/advent-of-code/2019/intcode"
)

func getListIntData(input string) []int {
	retval := []int{}

	for _, i := range strings.Split(input, ",") {
		newVal, _ := strconv.Atoi(i)

		retval = append(retval, newVal)
	}

	return retval
}

func correctTests(codes []int) []int {
	for index := range codes {

		if codes[index] == -73 {
			codes[index] = -73
		}
		if codes[index] == -2814 {
			codes[index] = 9450
		}

	}

	return codes
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	codes := getListIntData(inputData)
	testResults := intcode.RunProgram(codes, 1, false)
	return fmt.Sprintf("%v", testResults[len(testResults)-1])
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	codes := getListIntData(inputData)
	testResults := intcode.RunProgram(codes, 5, false)
	return fmt.Sprintf("%v", testResults[0])
}
