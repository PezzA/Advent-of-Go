package Day201902

import (
	"fmt"
	"strconv"
	"strings"
)

var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string) {
	return 2019, 2, "1202 Program Alarm"
}

func getListIntData(input string) []int {
	retval := []int{}

	for _, i := range strings.Split(input, ",") {
		newVal, _ := strconv.Atoi(i)

		retval = append(retval, newVal)
	}

	return retval
}

func runProgram(opcodes []int) []int {
	position := 0

	for {
		if opcodes[position] == 99 {
			break
		} else if opcodes[position] == 1 {
			opcodes[opcodes[position+3]] = opcodes[opcodes[position+1]] + opcodes[opcodes[position+2]]
		} else if opcodes[position] == 2 {
			opcodes[opcodes[position+3]] = opcodes[opcodes[position+1]] * opcodes[opcodes[position+2]]
		}
		position += 4
	}

	return opcodes
}
func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	codes := getListIntData(inputData)

	codes[1] = 12
	codes[2] = 2

	return fmt.Sprintf("%v", runProgram(codes)[0])
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	done := false
	actNoun, actVerb := 0, 0
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			codes := getListIntData(inputData)
			codes[1] = noun
			codes[2] = verb

			if runProgram(codes)[0] == 19690720 {
				actNoun = noun
				actVerb = verb
				done = true
			}

			if done {
				break
			}
		}
		if done {
			break
		}
	}

	return fmt.Sprintf("%v", 100*actNoun+actVerb)
}
