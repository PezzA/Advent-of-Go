package Day201623

import (
	"fmt"

	"github.com/pezza/advent-of-code/puzzle-support/assembunny"
)

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	program := assembunny.ParseProgram(inputData)
	registers := assembunny.RunProgram(program, assembunny.NewRegisterSet(7, 0, 0, 0))
	return fmt.Sprintf("%v", registers["a"])
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	program := assembunny.ParseProgram(inputData)
	registers := assembunny.RunProgram(program, assembunny.NewRegisterSet(12, 0, 0, 0))
	return fmt.Sprintf("%v", registers["a"])
}
