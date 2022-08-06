package Day201523

import (
	"fmt"
	"strconv"
	"strings"
)

// Entry holds wraps the data and runner interfaces for this puzzle
var Entry dayEntry

type dayEntry bool

type instruction struct {
	command  string
	register string
	offset   int
}

type registers map[string]int

func getData(input string) []instruction {
	ins := make([]instruction, 0)
	for _, line := range strings.Split(input, "\n") {
		bits := strings.Fields(line)

		switch bits[0] {
		case "hlf", "tpl", "inc":
			ins = append(ins, instruction{bits[0], bits[1], 0})
		case "jmp":
			offSet, _ := strconv.Atoi(bits[1])
			ins = append(ins, instruction{bits[0], "", offSet})
		case "jie", "jio":
			offSet, _ := strconv.Atoi(bits[2])
			reg := strings.Replace(bits[1], ",", "", -1)
			ins = append(ins, instruction{bits[0], reg, offSet})
		}
	}
	return ins
}

func (r registers) getRegister(reg string) int {
	if val, ok := r[reg]; ok {
		return val
	}
	return 0
}

func runProgram(program []instruction, reg registers) registers {
	if reg == nil {
		reg = make(registers)
	}

	caret := 0

	for caret < len(program) {
		ins := program[caret]

		switch ins.command {
		case "hlf":
			reg[ins.register] = reg.getRegister(ins.register) / 2
			caret++
		case "tpl":
			reg[ins.register] = reg.getRegister(ins.register) * 3
			caret++
		case "inc":
			reg[ins.register] = reg.getRegister(ins.register) + 1
			caret++
		case "jmp":
			caret += ins.offset
		case "jie":
			if reg.getRegister(ins.register)%2 == 0 {
				caret += ins.offset
			} else {
				caret++
			}
		case "jio":
			if reg.getRegister(ins.register) == 1 {
				caret += ins.offset
			} else {
				caret++
			}
		}
	}

	return reg
}

func (td dayEntry) Describe() (int, int, string, int) {
	return 2015, 23, "Opening the Turing Lock", 2
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	regList := make(registers)

	regList = runProgram(getData(inputData), regList)

	return fmt.Sprintf("%v", regList["b"])
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	regList := make(registers)

	regList["a"] = 1
	regList = runProgram(getData(inputData), regList)

	return fmt.Sprintf("%v", regList["b"])
}
