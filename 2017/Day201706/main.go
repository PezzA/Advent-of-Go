package Day201706

import (
	"strconv"
	"strings"
)

var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string) {
	return 2017, 6, "Memory Reallocation"
}

type registerList []int

type stateList []registerList

func getWrappedIndex(currentPos int, loopsize int, modifier int) int {
	rawIndex := currentPos + modifier

	if rawIndex >= (loopsize) {
		rawIndex -= loopsize
	}

	return rawIndex
}

func (sl stateList) isStatePreviouslySeen(rl registerList) int {
	for index, state := range sl {
		match := true
		for index := range rl {
			if state[index] != rl[index] {
				match = false
			}
		}
		if match {
			return index
		}
	}
	return -1
}

func (sl stateList) addStateToStateList(rl registerList) stateList {
	newState := make(registerList, 0)

	for _, val := range rl {
		newState = append(newState, val)
	}

	return append(sl, newState)
}

func (rl registerList) getHighestIndex() int {
	highIndex := 0
	maxVal := rl[0]

	for index, val := range rl {
		if val > maxVal {
			highIndex = index
			maxVal = val
		}
	}

	return highIndex
}

func (rl registerList) distributeRegister(targetReg int) registerList {
	target := rl[targetReg]
	rl[targetReg] = 0
	size := len(rl)

	for i := 0; i < target; i++ {
		targetReg = getWrappedIndex(targetReg, size, 1)
		rl[targetReg]++
	}

	return rl
}

func distributionCycle(rl registerList) (int, int) {
	states := make(stateList, 0)
	cycles := 0
	firstLoopIndex := 0

	for {
		rl.distributeRegister(rl.getHighestIndex())

		cycles++

		firstLoopIndex = states.isStatePreviouslySeen(rl)

		if firstLoopIndex > -1 {
			break
		}

		states = states.addStateToStateList(rl)
	}

	return cycles, firstLoopIndex
}

func getRegisters(input string) registerList {
	bits := strings.Fields(input)
	registers := make(registerList, 0)

	for _, val := range bits {
		regVal, _ := strconv.Atoi(val)
		registers = append(registers, regVal)
	}

	return registers
}

func (td dayEntry) PartOne(inputData string) (string, error) {
	cycles, _ := distributionCycle(getRegisters(inputData))
	return strconv.Itoa(cycles), nil
}

func (td dayEntry) PartTwo(inputData string) (string, error) {
	cycles, firstLoop := distributionCycle(getRegisters(inputData))
	return strconv.Itoa(cycles - firstLoop - 1), nil
}

func (td dayEntry) PuzzleInput() string {
	return "2	8	8	5	4	2	3	1	5	5	1	2	15	13	5	14"
}
