package Day201625

import (
	"fmt"

	"github.com/pezza/advent-of-code/puzzle-support/assembunny"
)

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	program := assembunny.ParseProgram(Entry.PuzzleInput())

	test := "0101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101"
	found := false
	index := 0
	for !found {

		registerSet := assembunny.NewRegisterSet(index, 0, 0, 0)

		outputChan := make(chan int)

		go assembunny.RunProgram(program, registerSet, outputChan, 100)

		testString := ""
		for elem := range outputChan {
			testString += fmt.Sprint(elem)
		}

		if testString == test {
			break
		}
		index++
	}

	return fmt.Sprint(index)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	return "ho ho ho, merry xmas"
}
