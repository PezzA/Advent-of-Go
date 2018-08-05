package Day201507

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Entry holds wraps the data and runner interfaces for this puzzle
var Entry dayEntry

type dayEntry bool

type wireList map[string]uint16

type gateList []*gate

func (td dayEntry) Describe() (int, int, string) {
	return 2015, 7, "Some Assembly Required"
}

var instructionRegex = regexp.MustCompile(`((([a-z0-9]*) (RSHIFT|OR|AND|LSHIFT)|NOT) )*([a-z0-9]*) -> ([a-z0-9]*)`)

type gate struct {
	command  string
	operand1 string
	operand2 string
	output   string
	running  bool
}

func (w wireList) signalState(testWire string) (bool, uint16) {
	if val, ok := w[testWire]; ok {
		return true, val
	}
	return false, 0
}

func isOperandConstant(operand string) bool {
	if _, err := strconv.Atoi(operand); err == nil {
		return true
	}
	return false
}

func getGateOperand(operand string, wires wireList) (bool, uint16) {
	if isOperandConstant(operand) {
		cons, _ := strconv.ParseUint(operand, 10, 16)
		return true, uint16(cons)
	} else {
		active, val := wires.signalState(operand)
		if active {
			return true, val
		}
	}
	return false, 0
}

func runCircuitLoop(gates gateList, wires wireList) (bool, wireList, gateList) {
	hasUpdate := false

	for _, g := range gates {
		if !g.running {
			hasUpdate = true
			switch g.command {
			case "ASSIGN":
				hasOp1Signal, op1Signal := getGateOperand(g.operand1, wires)

				// only assign if the output wire does not already have a value, part 2 gotcha!
				hasOutputSignal, _ := wires.signalState(g.output)

				if hasOutputSignal {
					g.running = true
					continue
				}

				if hasOp1Signal {
					wires[g.output] = op1Signal
					g.running = true
				}

			case "AND", "OR", "LSHIFT", "RSHIFT":
				hasOp1Signal, op1Signal := getGateOperand(g.operand1, wires)
				hasOp2Signal, op2Signal := getGateOperand(g.operand2, wires)

				if hasOp1Signal && hasOp2Signal {
					switch g.command {
					case "AND":
						wires[g.output] = op1Signal & op2Signal
					case "OR":
						wires[g.output] = op1Signal | op2Signal
					case "LSHIFT":
						wires[g.output] = op1Signal << op2Signal
					case "RSHIFT":
						wires[g.output] = op1Signal >> op2Signal
					}

					g.running = true
				}

			case "NOT":
				hasOp1Signal, op1Signal := getGateOperand(g.operand1, wires)

				if hasOp1Signal {
					wires[g.output] = ^op1Signal
					g.running = true
				}
			}
		}
	}

	return hasUpdate, wires, gates
}

func runCircuit(gates gateList, wires wireList) wireList {
	if wires == nil {
		wires = make(wireList, 0)
	}

	var hasUpdates bool

	for {
		hasUpdates, wires, gates = runCircuitLoop(gates, wires)
		if !hasUpdates {
			break
		}
	}

	return wires
}

func getInstructions(input string) gateList {
	insList, scanner := make(gateList, 0), bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		matches := instructionRegex.FindStringSubmatch(scanner.Text())

		var ins *gate

		switch matches[4] {
		case "RSHIFT", "LSHIFT", "OR", "AND":
			ins = &gate{
				command:  matches[4],
				operand1: matches[3],
				operand2: matches[5],
				output:   matches[6],
				running:  false,
			}
		case "":
			if matches[2] == "NOT" {
				ins = &gate{
					command:  matches[2],
					operand1: matches[5],
					operand2: "",
					output:   matches[6],
					running:  false,
				}
			} else {
				ins = &gate{
					command:  "ASSIGN",
					operand1: matches[5],
					operand2: "",
					output:   matches[6],
					running:  false,
				}
			}
		}
		insList = append(insList, ins)
	}
	return insList
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	wires := runCircuit(getInstructions(inputData), nil)
	return fmt.Sprintf("%v", wires["a"])
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	wires := runCircuit(getInstructions(inputData), nil)

	newWires := make(wireList, 0)
	newWires["b"] = wires["a"]

	updatedWires := runCircuit(getInstructions(inputData), newWires)
	return fmt.Sprintf("%v", updatedWires["a"])
}
