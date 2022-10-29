package chronalcompiler

func RunProgram(program []Instruction, regCount int, insPointerRegister int, registers RegisterSet) RegisterSet {
	opCodes := GetOpCodes()

	insPointer := 0

	boundInsPointer := insPointerRegister >= 0

	if registers == nil {
		registers = NewRegisterSet(regCount)
	}

	for _, ins := range program {

		if ins.OpCode == "ip" {
			continue
		}

		if boundInsPointer {
			registers[insPointerRegister] = insPointer
		}

		registers[ins.C] = opCodes[ins.OpCode].Process(registers, ins.A, ins.B)

		if boundInsPointer {
			insPointer = registers[insPointerRegister]
		}

		insPointer++

		// if pointer is outside of program bounds, exist
		if insPointer < 0 || insPointer >= len(program) {
			break
		}
	}

	return registers
}
