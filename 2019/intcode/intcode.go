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

func parseOpCode(input int) OpCode {
	str := strconv.Itoa(input)

	code, first, second, third := 0, 0, 0, 0

	switch len(str) {
	case 1:
		fallthrough
	case 2:
		code, _ = strconv.Atoi(str)
	case 3:
		code, _ = strconv.Atoi(str[1:])
		first = 1
	case 4:
		code, _ = strconv.Atoi(str[2:])
		first, _ = strconv.Atoi(string(str[1]))
		second = 1
	case 5:
		code, _ = strconv.Atoi(str[3:])
		first, _ = strconv.Atoi(string(str[2]))
		second, _ = strconv.Atoi(string(str[1]))
		third = 1
	}

	return OpCode{code, first, second, third}
}

func resolveValue(mode int, parameter int, codes []int) int {
	if mode == 1 {
		return parameter
	}

	return codes[parameter]
}

func RunProgram(opcodes []int, inputs []int, debug bool, inputChan chan int, outputChan chan int, name string, wg *sync.WaitGroup) []int {
	if wg != nil {
		defer wg.Done()
	}
	defer fmt.Println(name, "complete")
	//defer close(inputChan)
	inputPosition, position, outputs := 0, 0, []int{}

	for {
		op := parseOpCode(opcodes[position])

		if op.codeType == 99 {

			if debug {
				fmt.Println("exited")
			}

			break

		} else if op.codeType == 1 {

			param1, param2, param3 := opcodes[position+1], opcodes[position+2], opcodes[position+3]
			resParam1, resParam2 := resolveValue(op.firstMode, param1, opcodes), resolveValue(op.secondMode, param2, opcodes)

			before := opcodes[param3]
			opcodes[param3] = resParam1 + resParam2
			after := opcodes[param3]

			if debug {
				fmt.Printf("%v\t+\t%v\t", resParam1, resParam2)
				fmt.Println("=>", after, "  :", before)
			}

			position += 4

		} else if op.codeType == 2 {

			param1, param2, param3 := opcodes[position+1], opcodes[position+2], opcodes[position+3]
			resParam1, resParam2 := resolveValue(op.firstMode, param1, opcodes), resolveValue(op.secondMode, param2, opcodes)

			before := opcodes[param3]
			opcodes[param3] = resParam1 * resParam2
			after := opcodes[param3]

			if debug {
				fmt.Printf("%v\t*\t%v\t", resParam1, resParam2)
				fmt.Println("=>", after, "  :", before)
			}

			position += 4

		} else if op.codeType == 3 {

			param1 := opcodes[position+1]

			if inputChan != nil {
				fmt.Println(name, "waiting on input")
				opcodes[param1] = <-inputChan
			} else {
				opcodes[param1] = inputs[inputPosition]
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
				outputChan <- resolveValue(op.firstMode, param1, opcodes)
			}

			outputs = append(outputs, resolveValue(op.firstMode, param1, opcodes))

			if debug {
				fmt.Println("===============================")
				fmt.Println(resolveValue(op.firstMode, param1, opcodes))
				fmt.Println()
				fmt.Println()
			}

			position += 2

		} else if op.codeType == 5 {

			param1, param2 := opcodes[position+1], opcodes[position+2]

			if debug {
				fmt.Printf("jump-if-true is %v true? : ", resolveValue(op.firstMode, param1, opcodes))
			}

			if resolveValue(op.firstMode, param1, opcodes) != 0 {
				if debug {
					fmt.Printf("yes. ")
				}
				position = resolveValue(op.secondMode, param2, opcodes)
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
				fmt.Printf("jump-if-false, is %v false? : ", resolveValue(op.firstMode, param1, opcodes))
			}

			if resolveValue(op.firstMode, param1, opcodes) == 0 {
				if debug {
					fmt.Printf("yes. Jumping to %v\n", position)
				}
				position = resolveValue(op.secondMode, param2, opcodes)
			} else {
				if debug {
					fmt.Printf("no. incrementing by 3")
				}
				position += 3
			}

		} else if op.codeType == 7 {

			param1, param2, param3 := opcodes[position+1], opcodes[position+2], opcodes[position+3]

			if debug {
				fmt.Printf("Less-than.  is %v less then %v", resolveValue(op.firstMode, param1, opcodes), resolveValue(op.secondMode, param2, opcodes))
			}

			if resolveValue(op.firstMode, param1, opcodes) < resolveValue(op.secondMode, param2, opcodes) {
				if debug {
					fmt.Printf(": yes. : Setting 1 to %v", param3)
				}
				opcodes[param3] = 1
			} else {
				if debug {
					fmt.Printf(": no. : Setting 0 to %v", param3)
				}
				opcodes[param3] = 0
			}

			fmt.Println()
			position += 4
		} else if op.codeType == 8 {

			param1, param2, param3 := opcodes[position+1], opcodes[position+2], opcodes[position+3]

			if debug {
				fmt.Printf("equal-to.  is %v equal to %v", resolveValue(op.firstMode, param1, opcodes), resolveValue(op.secondMode, param2, opcodes))
			}

			if resolveValue(op.firstMode, param1, opcodes) == resolveValue(op.secondMode, param2, opcodes) {
				if debug {
					fmt.Printf(": yes. : Setting 1 to %v", param3)
				}
				opcodes[param3] = 1
			} else {
				if debug {
					fmt.Printf(": no. : Setting 0 to %v", param3)
				}
				opcodes[param3] = 0
			}
			fmt.Println()
			position += 4
		}

	}

	return outputs
}
