package Day201725

import (
	"fmt"
	"regexp"
	"strconv"
)

type dayEntry bool

var Entry dayEntry

func (td dayEntry) Describe() (int, int, string, int) {
	return 2017, 25, "The Halting Problem", 2
}

var controlRegex = regexp.MustCompile(`Begin in state ([A-Z]).\nPerform a diagnostic checksum after ([0-9]*) steps.`)
var commandRegex = regexp.MustCompile(`In state ([A-Z]*):\n(\tIf the current value is ([0-9])?:\n\t- Write the value ([0-9]?).\n\t- Move one slot to the (right|left).\n\t- Continue with state ([A-Z]?).\n\tIf the current value is ([0-9])?:\n\t- Write the value ([0-9]?).\n\t- Move one slot to the (right|left).\n\t- Continue with state ([A-Z]?).)`)

type instruction struct {
	ZeroWriteValue      int
	ZeroMoveInstruction string
	ZeroContinueState   string
	OneWriteValue       int
	OneMoveInstruction  string
	OneContinueState    string
}

type tape map[int]int

func getTuringData(input string) (string, int, map[string]instruction) {
	controlMatch := controlRegex.FindStringSubmatch(input)
	commandMatches := commandRegex.FindAllStringSubmatch(input, -1)

	steps, _ := strconv.Atoi(controlMatch[2])
	states := make(map[string]instruction, 0)

	for _, val := range commandMatches {
		zeroWrite, _ := strconv.Atoi(val[4])
		oneWrite, _ := strconv.Atoi(val[8])

		states[val[1]] = instruction{zeroWrite, val[5], val[6], oneWrite, val[9], val[10]}
	}

	return controlMatch[1], steps, states
}

func runStep(state instruction, cursor int, tape tape) (string, int, tape) {
	if _, ok := tape[cursor]; !ok {
		tape[cursor] = 0
	}

	var newstate = ""
	if tape[cursor] == 0 {
		tape[cursor] = state.ZeroWriteValue
		if state.ZeroMoveInstruction == "left" {
			cursor--
		} else {
			cursor++
		}
		newstate = state.ZeroContinueState
	} else {
		tape[cursor] = state.OneWriteValue
		if state.OneMoveInstruction == "left" {
			cursor--
		} else {
			cursor++
		}
		newstate = state.OneContinueState
	}

	return newstate, cursor, tape
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {

	state, steps, commands := getTuringData(inputData)

	cursor := 0
	tape := make(map[int]int, 0)
	for i := 0; i < steps; i++ {
		state, cursor, tape = runStep(commands[state], cursor, tape)
	}

	checksum := 0
	for _, v := range tape {
		checksum += v
	}

	return fmt.Sprintf("%v", checksum)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	return fmt.Sprintf("Reboot Complete!")
}
