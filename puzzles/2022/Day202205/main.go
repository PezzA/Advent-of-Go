package Day202205

import (
	"fmt"
	"strings"

	"github.com/pezza/advent-of-code/puzzles/common"
)

type instruction struct {
	source      int
	destination int
	count       int
}

func getData(input string) (map[int]*common.Stack, []instruction) {
	lines := strings.Split(input, "\n")

	loadingInstructions := false
	stackCount := -1

	stacks := make(map[int]*common.Stack, 0)
	ins := make([]instruction, 0)

	for _, line := range lines {
		if line == "" {
			loadingInstructions = true
			continue
		}

		if !loadingInstructions {
			if stackCount == -1 {
				// cater for no space after the last stack by adding one to the length
				stackCount = (len(line) + 1) / 4

				// init all the stacks
				for i := 0; i < stackCount; i++ {
					stacks[i+1] = common.NewStack()
				}
			}

			for i := 0; i < stackCount; i++ {

				index := (i * 4) + 1
				if line[index] >= 65 && line[index] <= 90 {
					stacks[i+1].AddAtBottom(string(line[index]))
				}
			}
			continue
		}

		count, source, destination := 0, 0, 0
		fmt.Sscanf(line, "move %d from %d to %d", &count, &source, &destination)
		ins = append(ins, instruction{source, destination, count})
	}

	return stacks, ins
}

func processInstruction(ins instruction, stacks map[int]*common.Stack, maintainOrder bool) {

	tempStack := common.NewStack()

	for i := 0; i < ins.count; i++ {
		crate, err := stacks[ins.source].Pop()

		if maintainOrder {
			tempStack.Push(crate)
			continue
		}

		if err == nil {
			stacks[ins.destination].Push(crate)
		}
	}

	if maintainOrder {

		for {

			crate, err := tempStack.Pop()

			if err != nil {
				break
			}

			stacks[ins.destination].Push(crate)
		}
	}
}

func processInstructionList(insList []instruction, stacks map[int]*common.Stack, maintainOrder bool) {
	for _, ins := range insList {
		processInstruction(ins, stacks, maintainOrder)
	}
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	stacks, instructions := getData(inputData)

	processInstructionList(instructions, stacks, false)

	tops := ""
	for i := 1; i <= 9; i++ {
		tops += stacks[i].Peek()
	}

	return fmt.Sprint(tops)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	stacks, instructions := getData(inputData)

	processInstructionList(instructions, stacks, true)

	tops := ""
	for i := 1; i <= 9; i++ {
		tops += stacks[i].Peek()
	}

	return fmt.Sprint(tops)
}
