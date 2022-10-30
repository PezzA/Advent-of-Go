package Day201819

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pezza/advent-of-code/puzzle-support/chronalcompiler"
)

func getData(input string) []chronalcompiler.Instruction {
	lines := strings.Split(input, "\n")

	insList := make([]chronalcompiler.Instruction, len(lines))

	bits := strings.Split(lines[0], " ")

	ipBinding, _ := strconv.Atoi(bits[1])

	insList[0] = chronalcompiler.Instruction{OpCode: "ip", A: ipBinding, B: 0, C: 0}

	for i := 0; i < len(lines); i++ {
		ins, a, b, c := "", 0, 0, 0
		fmt.Sscanf(lines[i], "%v %d %d %d", &ins, &a, &b, &c)
		insList[i] = chronalcompiler.Instruction{OpCode: ins, A: a, B: b, C: c}
	}

	return insList

}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	program := getData(inputData)

	regSet := chronalcompiler.RunProgram(program, 6, nil, nil)

	return fmt.Sprint(regSet[0])
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	control := 10551377

	sumOfFactors := 0
	for i := 1; i <= control; i++ {
		if control%i == 0 {
			sumOfFactors += i
		}
	}
	return fmt.Sprint((sumOfFactors))
}
