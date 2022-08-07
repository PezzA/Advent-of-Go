package Day201612

import (
	"fmt"
	"strconv"
	"strings"
)

type registerSet = map[string]int

type instruction struct {
	instruction string
	value1      int
	value2      int
	register1   string
	register2   string
}

func getData(input string) []instruction {
	lines := strings.Split(input, "\n")
	instructions := make([]instruction, len(lines))

	for index, line := range lines {
		bits := strings.Split(line, " ")

		switch bits[0] {
		case "cpy":
			val, err := strconv.Atoi(bits[1])

			if err != nil {
				instructions[index] = instruction{instruction: "cpyr", register1: bits[1], register2: bits[2]}
			} else {
				instructions[index] = instruction{instruction: "cpyv", value1: val, register2: bits[2]}
			}

		case "jnz":
			val, _ := strconv.Atoi(bits[2])
			source, err := strconv.Atoi(bits[1])

			if err != nil {
				instructions[index] = instruction{instruction: "jnzr", register1: bits[1], value2: val}
			} else {
				instructions[index] = instruction{instruction: "jnzv", value1: source, value2: val}
			}

		default:
			instructions[index] = instruction{instruction: bits[0], register1: bits[1]}
		}
	}

	return instructions
}

func runProgram(program []instruction, regSet registerSet) registerSet {
	index := 0

	for index < len(program) {

		currIns := program[index]

		switch currIns.instruction {
		case "inc":
			regSet[currIns.register1]++
		case "dec":
			regSet[currIns.register1]--
		case "cpyv":
			regSet[currIns.register2] = currIns.value1
		case "cpyr":
			regSet[currIns.register2] = regSet[currIns.register1]
		case "jnzr":
			if regSet[currIns.register1] != 0 {
				index += currIns.value2
				continue
			}
		case "jnzv":
			if currIns.value1 != 0 {
				index += currIns.value2
				continue
			}
		}

		index++
	}

	return regSet

}

func makeNewRegisterSet() registerSet {
	regSet := make(registerSet)

	regSet["a"] = 0
	regSet["b"] = 0
	regSet["c"] = 0
	regSet["d"] = 0

	return regSet
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	program := getData(inputData)

	registers := runProgram(program, makeNewRegisterSet())

	return fmt.Sprintf("%v", registers["a"])
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	program := getData(inputData)

	registerSet := makeNewRegisterSet()

	registerSet["c"] = 1

	registers := runProgram(program, registerSet)

	return fmt.Sprintf("%v", registers["a"])
}
