package intcode

import (
	"fmt"
	"strconv"
	"sync"
)

type OpCode struct {
	codeType   int
	firstMode  int
	secondMode int
	thirdMode  int
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
func setValue(parameter int64, codes []int64, mmap map[int64]int64, val int64) {
	if parameter > int64(len(codes))-1 {
		mapPos := parameter - int64(len(codes))
		mmap[mapPos] = val
		return
	}

	codes[parameter] = val
}

func resolveValue(mode int, parameter int64, codes []int64, rb int64, mmap map[int64]int64) int64 {
	switch mode {
	case 0:
		if parameter > int64(len(codes))-1 {
			mapPos := parameter - int64(len(codes))
			return mmap[mapPos]
		} else {
			return codes[parameter]
		}

	case 1:
		return parameter
	case 2:
		if rb+parameter > int64(len(codes))-1 {
			mapPos := (rb + parameter) - int64(len(codes))
			return mmap[mapPos]
		} else {
			return codes[rb+parameter]
		}

	default:
		return -1
	}
}

// RunProgram runs an intcode program
func RunProgram(opcodes []int64, inputs []int64, debug bool, inputChan chan int64, outputChan chan int64, name string, wg *sync.WaitGroup) []int64 {
	if wg != nil {
		defer wg.Done()
	}
	//defer close(inputChan)
	relativeBase, inputPosition, position, outputs := int64(0), 0, int64(0), []int64{}

	memory := make(map[int64]int64, 0)

	for {
		op := parseOpCode(opcodes[position])

		if op.codeType == 99 {

			if debug {
				fmt.Println("exited")
			}

			break

		} else if op.codeType == 1 {

			param1, param2, param3 := opcodes[position+1], opcodes[position+2], opcodes[position+3]
			resParam1, resParam2 := resolveValue(op.firstMode, param1, opcodes, relativeBase, memory), resolveValue(op.secondMode, param2, opcodes, relativeBase, memory)

			if op.thirdMode == 2 {
				param3 += relativeBase
			}

			setValue(param3, opcodes, memory, resParam1+resParam2)

			position += 4

		} else if op.codeType == 2 {

			param1, param2, param3 := opcodes[position+1], opcodes[position+2], opcodes[position+3]
			resParam1, resParam2 := resolveValue(op.firstMode, param1, opcodes, relativeBase, memory), resolveValue(op.secondMode, param2, opcodes, relativeBase, memory)

			if op.thirdMode == 2 {
				param3 += relativeBase
			}

			setValue(param3, opcodes, memory, resParam1*resParam2)

			position += 4

		} else if op.codeType == 3 {

			param1 := opcodes[position+1]

			if inputChan != nil {
				fmt.Println(name, "waiting on input")
				setValue(param1, opcodes, memory, <-inputChan)
			} else {
				if op.firstMode == 2 {
					param1 += relativeBase
				}
				setValue(param1, opcodes, memory, inputs[inputPosition])
			}

			if debug {
				fmt.Printf("Storing input %v into %v\n", inputs, param1)
			}

			inputPosition++
			position += 2

		} else if op.codeType == 4 {
			param1 := opcodes[position+1]

			if outputChan != nil {
				fmt.Println(name, "waiting on output")
				outputChan <- resolveValue(op.firstMode, param1, opcodes, relativeBase, memory)
			}

			outputs = append(outputs, resolveValue(op.firstMode, param1, opcodes, relativeBase, memory))

			if debug {
				fmt.Println("===============================")
				fmt.Println(resolveValue(op.firstMode, param1, opcodes, relativeBase, memory))
				fmt.Println()
				fmt.Println()
			}

			position += 2

		} else if op.codeType == 5 {

			param1, param2 := opcodes[position+1], opcodes[position+2]

			if debug {
				fmt.Printf("jump-if-true is %v true? : ", resolveValue(op.firstMode, param1, opcodes, relativeBase, memory))
			}

			if resolveValue(op.firstMode, param1, opcodes, relativeBase, memory) != 0 {
				if debug {
					fmt.Printf("yes. ")
				}
				position = resolveValue(op.secondMode, param2, opcodes, relativeBase, memory)
			} else {
				if debug {
					fmt.Printf("no. ")
				}
				position += 3
			}

			if debug {
				fmt.Printf("Jumping to %v\n", position)
			}

		} else if op.codeType == 6 {

			param1, param2 := opcodes[position+1], opcodes[position+2]

			if debug {
				fmt.Printf("jump-if-false, is %v false? : ", resolveValue(op.firstMode, param1, opcodes, relativeBase, memory))
			}

			if resolveValue(op.firstMode, param1, opcodes, relativeBase, memory) == 0 {
				if debug {
					fmt.Printf("yes. Jumping to %v\n", position)
				}
				position = resolveValue(op.secondMode, param2, opcodes, relativeBase, memory)
			} else {
				if debug {
					fmt.Printf("no. incrementing by 3")
				}
				position += 3
			}

		} else if op.codeType == 7 {

			param1, param2, param3 := opcodes[position+1], opcodes[position+2], opcodes[position+3]

			if op.thirdMode == 2 {
				param3 += relativeBase
			}

			if debug {
				fmt.Printf("Less-than.  is %v less then %v", resolveValue(op.firstMode, param1, opcodes, relativeBase, memory), resolveValue(op.secondMode, param2, opcodes, relativeBase, memory))
			}

			if resolveValue(op.firstMode, param1, opcodes, relativeBase, memory) < resolveValue(op.secondMode, param2, opcodes, relativeBase, memory) {
				if debug {
					fmt.Printf(": yes. : Setting 1 to %v", param3)
				}
				setValue(param3, opcodes, memory, 1)

			} else {
				if debug {
					fmt.Printf(": no. : Setting 0 to %v", param3)
				}
				setValue(param3, opcodes, memory, 0)
			}

			if debug {
				fmt.Println()
			}
			position += 4
		} else if op.codeType == 8 {

			param1, param2, param3 := opcodes[position+1], opcodes[position+2], opcodes[position+3]

			if op.thirdMode == 2 {
				param3 += relativeBase
			}

			if debug {
				fmt.Printf("equal-to.  is %v equal to %v", resolveValue(op.firstMode, param1, opcodes, relativeBase, memory), resolveValue(op.secondMode, param2, opcodes, relativeBase, memory))
			}

			if resolveValue(op.firstMode, param1, opcodes, relativeBase, memory) == resolveValue(op.secondMode, param2, opcodes, relativeBase, memory) {
				if debug {
					fmt.Printf(": yes. : Setting 1 to %v", param3)
				}

				setValue(param3, opcodes, memory, 1)

			} else {
				if debug {
					fmt.Printf(": no. : Setting 0 to %v", param3)
				}

				setValue(param3, opcodes, memory, 0)
			}

			if debug {
				fmt.Println()
			}

			position += 4
		} else if op.codeType == 9 {
			param1 := opcodes[position+1]
			relativeBase += resolveValue(op.firstMode, param1, opcodes, relativeBase, memory)
			position += 2
		}

	}

	return outputs
}
