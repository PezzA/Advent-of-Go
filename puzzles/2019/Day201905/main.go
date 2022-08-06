package Day201905

import (
	"fmt"

	"github.com/pezza/advent-of-code/puzzles/2019/intcode"
)

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	vm := intcode.New(inputData)
	testResults := vm.RunProgram(nil, []int64{1}, nil, nil, nil)
	return fmt.Sprintf("%v", testResults[len(testResults)-1])
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	vm := intcode.New(inputData)
	testResults := vm.RunProgram(nil, []int64{5}, nil, nil, nil)
	return fmt.Sprintf("%v", testResults[0])
}
