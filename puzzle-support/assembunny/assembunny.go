package assembunny

import (
	"strconv"
	"strings"
)

func ParseProgram(input string) Program {
	lines := strings.Split(input, "\n")
	instructions := make(Program, len(lines))

	for index, line := range lines {
		bits := strings.Split(line, " ")

		instructions[index] = Instruction{instruction: bits[0], arg1: getArgument(bits[1])}

		if len(bits) == 3 {
			instructions[index].arg2 = getArgument(bits[2])
		}
	}

	return instructions
}

func getArgument(input string) Argument {
	arg, err := strconv.Atoi(input)

	if err != nil {
		return Argument{isRegister: true, register: input}
	}

	return Argument{isRegister: false, value: arg}
}

func RunProgram(prog Program, regSet RegisterSet) RegisterSet {
	index := 0

	for index < len(prog) {
		currIns := prog[index]

		switch currIns.instruction {
		case "inc":
			if currIns.arg1.isRegister {
				regSet[currIns.arg1.register]++
			}
		case "dec":
			if currIns.arg1.isRegister {
				regSet[currIns.arg1.register]--
			}
		case "cpy":
			if !currIns.arg2.isRegister {
				index++
				continue
			}

			if currIns.arg1.isRegister {
				regSet[currIns.arg2.register] = regSet[currIns.arg1.register]
			} else {
				regSet[currIns.arg2.register] = currIns.arg1.value
			}
		case "jnz":
			if currIns.arg1.isRegister {
				if regSet[currIns.arg1.register] == 0 {
					index++
					continue
				}
			} else {
				if currIns.arg1.value == 0 {
					index++
					continue
				}
			}

			if currIns.arg2.isRegister {
				index += regSet[currIns.arg2.register]
			} else {
				index += currIns.arg2.value
			}
			continue
		case "tgl":

			tglIndex := index
			if currIns.arg1.isRegister {
				tglIndex += regSet[currIns.arg1.register]
			} else {
				tglIndex += currIns.arg2.value
			}

			if tglIndex >= len(prog) {
				index++
				continue
			}

			tglIns := prog[tglIndex]

			switch tglIns.instruction {
			case "inc":
				prog[tglIndex].instruction = "dec"
			case "tgl":
				fallthrough
			case "dec":
				prog[tglIndex].instruction = "inc"
			case "cpy":
				prog[tglIndex].instruction = "jnz"
			case "jnz":
				prog[tglIndex].instruction = "cpy"
			}

		}

		index++
	}

	return regSet

}

func NewRegisterSet(a int, b int, c int, d int) RegisterSet {
	regSet := make(RegisterSet)

	regSet["a"] = a
	regSet["b"] = b
	regSet["c"] = c
	regSet["d"] = d

	return regSet
}
