package Day202008

import (
	"fmt"
	"strconv"
	"strings"
)

type instruction struct {
	command  string
	modifier int
}

func deepCopy(ins []instruction) []instruction {
	newIns := make([]instruction, len(ins))

	for i := range ins {
		newIns[i] = instruction{
			command:  ins[i].command,
			modifier: ins[i].modifier,
		}
	}

	return newIns
}

func getData(input string) []instruction {
	lines := strings.Split(input, "\n")
	data := make([]instruction, len(lines))

	for i := range lines {
		bits := strings.Split(lines[i], " ")

		mod, _ := strconv.Atoi(bits[1])
		data[i] = instruction{
			command:  bits[0],
			modifier: mod,
		}
	}
	return data
}

func getCommands(command string, ins []instruction) []int {
	indexes := make([]int, 0)

	for i := range ins {
		if ins[i].command == command {
			indexes = append(indexes, i)
		}
	}

	return indexes
}

func runProgram(ins []instruction) (bool, int) {
	acc, currIndex := 0, 0
	visitedLines := make(map[int]bool, 0)

	terminated := false
	for {

		if currIndex == len(ins) {
			terminated = true
			break
		}

		if _, ok := visitedLines[currIndex]; ok {
			break
		} else {
			visitedLines[currIndex] = true
		}

		if ins[currIndex].command == "nop" {
			currIndex++
			continue
		} else if ins[currIndex].command == "acc" {
			acc += ins[currIndex].modifier
			currIndex++
			continue
		} else if ins[currIndex].command == "jmp" {
			currIndex += ins[currIndex].modifier
			continue
		}
	}

	return terminated, acc
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	data := getData(inputData)
	_, acc := runProgram(data)
	return fmt.Sprintf("%v", acc)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	data := getData(inputData)

	jumps := getCommands("jmp", data)
	noOps := getCommands("nop", data)

	termAcc := 0

	for _, jmp := range jumps {
		testProg := deepCopy(data)
		testProg[jmp].command = "nop"

		term, acc := runProgram(testProg)

		if term {
			termAcc = acc
		}

	}

	if termAcc != 0 {
		for _, nop := range noOps {
			testProg := deepCopy(data)
			testProg[nop].command = "jmp"

			term, acc := runProgram(testProg)

			if term {
				termAcc = acc
			}
		}

	}

	return fmt.Sprintf("%v", termAcc)
}
