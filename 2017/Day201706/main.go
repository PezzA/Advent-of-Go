package Day201706

import (
	"strconv"
	"strings"
)

// Entry holds wraps the data and runner interfaces for this puzzle
var Entry testDay

type testDay bool

func getWrappedIndex(currentPos int, loopsize int, modifier int) int {
	rawIndex := currentPos + modifier

	if rawIndex >= (loopsize) {
		rawIndex -= loopsize
	}

	return rawIndex
}

func isStatePreviouslySeen(registers []int, states [][]int) int {
	for index, state := range states {
		match := true
		for index := range registers {
			if state[index] != registers[index] {
				match = false
			}
		}
		if match {
			return index
		}
	}
	return -1
}

func getHighestIndex(registers []int) int {
	highIndex := 0
	maxVal := registers[0]

	for index, val := range registers {
		if val > maxVal {
			highIndex = index
			maxVal = val
		}
	}

	return highIndex
}

func distributeRegister(targetReg int, registers []int) []int {
	target := registers[targetReg]
	registers[targetReg] = 0
	size := len(registers)

	for i := 0; i < target; i++ {
		targetReg = getWrappedIndex(targetReg, size, 1)
		registers[targetReg]++
	}

	return registers
}

func addStateToStateList(reg []int, states [][]int) [][]int {
	newState := make([]int, 0)

	for _, val := range reg {
		newState = append(newState, val)
	}

	return append(states, newState)
}

func distributionCycle(registers []int) ([]int, int, int) {
	states := make([][]int, 0)
	cycles := 0
	firstLoopIndex := 0

	for {
		// get the index of the highest value
		index := getHighestIndex(registers)

		// get it's value and distribute over registers
		registers := distributeRegister(index, registers)
		cycles++

		// see if this set of registers has been seen before
		firstLoopIndex = isStatePreviouslySeen(registers, states)

		if firstLoopIndex > -1 {
			break
		}

		// if not append to the list of seen states and update cycles
		states = addStateToStateList(registers, states)
	}

	return []int{}, cycles, firstLoopIndex
}

func getRegisters(input string) []int {
	bits := strings.Fields(input)
	registers := make([]int, 0)

	for _, val := range bits {
		regVal, _ := strconv.Atoi(val)
		registers = append(registers, regVal)
	}

	return registers
}

func (td testDay) PartOne(inputData string) (string, error) {
	registers := getRegisters(inputData)

	_, cycles, _ := distributionCycle(registers)

	return strconv.Itoa(cycles), nil
}

func (td testDay) PartTwo(inputData string) (string, error) {
	registers := getRegisters(inputData)

	_, cycles, firstLoop := distributionCycle(registers)

	return strconv.Itoa(cycles - firstLoop - 1), nil
}

func (td testDay) Day() int {
	return 6
}

func (td testDay) Year() int {
	return 2017
}

func (td testDay) Title() string {
	return "Memory Reallocation"
}

func (td testDay) GetData() string {
	return "2	8	8	5	4	2	3	1	5	5	1	2	15	13	5	14"
}
