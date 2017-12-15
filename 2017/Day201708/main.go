package Day201708

import (
	"bufio"
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string) {
	return 2017, 8, "I Heard You Like Registers"
}

// link to regex tester : https://regex101.com/r/rodiWy/1/
var instructionRegex = regexp.MustCompile(`([a-z]*) (inc|dec) ([-]?[0-9]*) if ([a-z]*) (>|<|<=|==|>=|!=) ([-]?[0-9]*)`)

type instruction struct {
	register         string
	registerOp       string
	registerShift    int
	testRegister     string
	testRegisterOp   string
	testRegisterEval int
}

type registers map[string]int

func parseInstruction(input string) instruction {
	match := instructionRegex.FindStringSubmatch(input)

	registerShift, _ := strconv.Atoi(match[3])
	testRegisterEval, _ := strconv.Atoi(match[6])

	return instruction{
		register:         match[1],
		registerOp:       match[2],
		registerShift:    registerShift,
		testRegister:     match[4],
		testRegisterOp:   match[5],
		testRegisterEval: testRegisterEval,
	}
}

func getInstructions(input string) []instruction {
	scanner := bufio.NewScanner(strings.NewReader(input))
	instructionList := make([]instruction, 0)

	for scanner.Scan() {
		instructionList = append(instructionList, parseInstruction(scanner.Text()))
	}
	return instructionList
}

func (r registers) getRegister(reg string) int {
	if val, ok := r[reg]; ok {
		return val
	}
	return 0
}

func (i instruction) isTestValid(r registers) (bool, error) {
	testReg := r.getRegister(i.testRegister)

	switch i.testRegisterOp {
	case "<":
		return testReg < i.testRegisterEval, nil
	case ">":
		return testReg > i.testRegisterEval, nil
	case "==":
		return testReg == i.testRegisterEval, nil
	case "!=":
		return testReg != i.testRegisterEval, nil
	case ">=":
		return testReg >= i.testRegisterEval, nil
	case "<=":
		return testReg <= i.testRegisterEval, nil
	}

	return false, errors.New("did not recieved supported operation")
}

func (i instruction) runInstruction(r registers) registers {
	valid, _ := i.isTestValid(r)

	if !valid {
		return r
	}

	targetReg := r.getRegister(i.register)

	switch i.registerOp {
	case "inc":
		r[i.register] = targetReg + i.registerShift
	case "dec":
		r[i.register] = targetReg - i.registerShift
	}

	return r
}

func runProgram(r registers, instructions []instruction) (registers, int) {
	highVal := 0
	for _, instruction := range instructions {
		r = instruction.runInstruction(r)
		for _, value := range r {
			if value > highVal {
				highVal = value
			}
		}
	}
	return r, highVal
}

func (td dayEntry) PartOne(inputData string) (string, error) {
	r, _ := runProgram(make(registers, 0), getInstructions(inputData))

	test := 0
	for _, value := range r {
		if value > test {
			test = value
		}
	}
	return strconv.Itoa(test), nil
}

func (td dayEntry) PartTwo(inputData string) (string, error) {
	_, highVal := runProgram(make(registers, 0), getInstructions(inputData))
	return strconv.Itoa(highVal), nil
}
