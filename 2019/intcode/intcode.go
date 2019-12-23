package intcode

import (
	"log"
	"strconv"
	"strings"
)

type OpCode struct {
	codeType   int
	firstMode  int
	secondMode int
	thirdMode  int
}

type IntCode struct {
	program       string
	codes         []int64
	memory        map[int64]int64
	position      int64
	inputPosition int
	relativeBase  int64
	requestInput  chan bool
}

func New(input string) IntCode {
	return IntCode{
		program:       input,
		codes:         nil,
		memory:        make(map[int64]int64, 0),
		position:      0,
		inputPosition: 0,
		relativeBase:  0,
		requestInput:  nil,
	}
}

func getListIntData(input string) []int64 {
	retval := []int64{}

	for _, i := range strings.Split(input, ",") {
		newVal, _ := strconv.ParseInt(i, 10, 64)

		retval = append(retval, newVal)
	}

	return retval
}

func parseOpCode(input int64) OpCode {
	str := strconv.FormatInt(input, 10)

	code, first, second, third := 0, 0, 0, 0

	switch len(str) {
	case 1:
		fallthrough
	case 2:
		code, _ = strconv.Atoi(str)
	case 3:
		code, _ = strconv.Atoi(str[1:])
		first, _ = strconv.Atoi(string(str[0]))
	case 4:
		code, _ = strconv.Atoi(str[2:])
		first, _ = strconv.Atoi(string(str[1]))
		second, _ = strconv.Atoi(string(str[0]))
	case 5:
		code, _ = strconv.Atoi(str[3:])
		first, _ = strconv.Atoi(string(str[2]))
		second, _ = strconv.Atoi(string(str[1]))
		third, _ = strconv.Atoi(string(str[0]))
	}

	return OpCode{code, first, second, third}
}

func (ic *IntCode) GetValue(pos int64) int64 {
	if int64(len(ic.codes)+1) >= pos {
		return ic.codes[0]
	}
	return 0
}

func (ic *IntCode) setValue(mode int, parameter int64, val int64) {
	if mode == 2 {
		parameter += ic.relativeBase
	}

	if parameter > int64(len(ic.codes))-1 {
		mapPos := parameter - int64(len(ic.codes))
		ic.memory[mapPos] = val
		return
	}

	ic.codes[parameter] = val
}

func (ic *IntCode) resolveValue(mode int, parameter int64) int64 {
	if mode == 1 {
		return parameter
	}

	if mode == 2 {
		parameter += ic.relativeBase
	}

	if parameter > int64(len(ic.codes))-1 {
		mapPos := parameter - int64(len(ic.codes))
		return ic.memory[mapPos]
	}

	return ic.codes[parameter]
}

func (ic *IntCode) SetMemory(pos, val int64) {
	ic.memory[pos] = val
}

// RunProgram runs an intcode program
func (ic *IntCode) RunProgram(init map[int64]int64, inputs []int64, inputChan chan int64, outputChan chan int64, requestChan chan bool) []int64 {
	outputs := []int64{}

	ic.codes = getListIntData(ic.program)

	if init != nil {
		for k, v := range init {
			ic.codes[k] = v
		}
	}

	ic.position = 0
	ic.relativeBase = 0
	ic.inputPosition = 0
	ic.requestInput = requestChan

	for {
		op := parseOpCode(ic.codes[ic.position])

		switch op.codeType {
		case 1:
			param1, param2, param3 := ic.codes[ic.position+1], ic.codes[ic.position+2], ic.codes[ic.position+3]
			resParam1, resParam2 := ic.resolveValue(op.firstMode, param1), ic.resolveValue(op.secondMode, param2)
			ic.setValue(op.thirdMode, param3, resParam1+resParam2)
			ic.position += 4
		case 2:
			param1, param2, param3 := ic.codes[ic.position+1], ic.codes[ic.position+2], ic.codes[ic.position+3]
			resParam1, resParam2 := ic.resolveValue(op.firstMode, param1), ic.resolveValue(op.secondMode, param2)
			ic.setValue(op.thirdMode, param3, resParam1*resParam2)
			ic.position += 4
		case 3:
			param1 := ic.codes[ic.position+1]

			if ic.requestInput != nil {
				ic.requestInput <- true
			}

			if inputChan != nil {
				val := int64(-1)
				select {
				case nval := <-inputChan:
					val = nval
				default:
				}
				ic.setValue(op.firstMode, param1, val)

			} else {
				ic.setValue(op.firstMode, param1, inputs[ic.inputPosition])
			}

			ic.inputPosition++
			ic.position += 2
		case 4:
			param1 := ic.codes[ic.position+1]

			if outputChan != nil {
				outputChan <- ic.resolveValue(op.firstMode, param1)
			}

			outputs = append(outputs, ic.resolveValue(op.firstMode, param1))
			ic.position += 2
		case 5:
			param1, param2 := ic.codes[ic.position+1], ic.codes[ic.position+2]

			if ic.resolveValue(op.firstMode, param1) != 0 {
				ic.position = ic.resolveValue(op.secondMode, param2)
			} else {
				ic.position += 3
			}
		case 6:
			param1, param2 := ic.codes[ic.position+1], ic.codes[ic.position+2]

			if ic.resolveValue(op.firstMode, param1) == 0 {
				ic.position = ic.resolveValue(op.secondMode, param2)
			} else {
				ic.position += 3
			}
		case 7:
			param1, param2, param3 := ic.codes[ic.position+1], ic.codes[ic.position+2], ic.codes[ic.position+3]

			if ic.resolveValue(op.firstMode, param1) < ic.resolveValue(op.secondMode, param2) {
				ic.setValue(op.thirdMode, param3, 1)
			} else {
				ic.setValue(op.thirdMode, param3, 0)
			}

			ic.position += 4
		case 8:
			param1, param2, param3 := ic.codes[ic.position+1], ic.codes[ic.position+2], ic.codes[ic.position+3]

			if ic.resolveValue(op.firstMode, param1) == ic.resolveValue(op.secondMode, param2) {
				ic.setValue(op.thirdMode, param3, 1)

			} else {
				ic.setValue(op.thirdMode, param3, 0)
			}

			ic.position += 4
		case 9:
			param1 := ic.codes[ic.position+1]
			ic.relativeBase += ic.resolveValue(op.firstMode, param1)
			ic.position += 2
		case 99:
			return outputs
		default:
			log.Fatalln("Got unknown command")
		}
	}
}
