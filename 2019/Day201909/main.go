package Day201909

import (
	"fmt"

	"github.com/pezza/advent-of-code/2019/intcode"
)

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	vm := intcode.New(Entry.PuzzleInput())
	results := vm.RunProgram(nil, []int64{1}, nil, nil)
	return fmt.Sprintf("%v", results[0])
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	vm := intcode.New(Entry.PuzzleInput())
	results := vm.RunProgram(nil, []int64{2}, nil, nil)
	return fmt.Sprintf("%v", results[0])
}
