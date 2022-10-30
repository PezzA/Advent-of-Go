package chronalcompiler

import "fmt"

func RunProgram(program []Instruction, regCount int, registers RegisterSet, output chan []string) RegisterSet {
	opCodes := GetOpCodes()

	hasBoundRegister, insPointer, boundRegister := false, 0, 0

	if registers == nil {
		registers = NewRegisterSet(regCount)
	}

	if program[0].OpCode == "#ip" {
		boundRegister = program[0].A
		hasBoundRegister = true
		program = program[1:]
	}

	for {
		ins := program[insPointer]
		if hasBoundRegister {
			registers[boundRegister] = insPointer
		}

		registers[ins.C] = opCodes[ins.OpCode].Process(registers, ins.A, ins.B)

		if hasBoundRegister {
			insPointer = registers[boundRegister]
		}

		insPointer++

		if output != nil {
			output <- []string{fmt.Sprintf("-- A:%v B:%v C:%v D:%v E:%v JUMP:%v --", registers[0], registers[1], registers[2], registers[3], registers[4], registers[5])}
		}

		// if pointer is outside of program bounds, exist
		if insPointer < 0 || insPointer >= len(program) {
			break
		}
	}

	return registers
}
