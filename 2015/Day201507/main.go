package Day201507

import (
	"bufio"
	"regexp"
	"strings"
)

// Entry holds wraps the data and runner interfaces for this puzzle
var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string) {
	return 2015, 7, "Some Assembly Required"
}

var instructionRegex = regexp.MustCompile(`(\S*)*( )*(NOT|RSHIFT|LSHIFT|OR|AND)*( )*(\S*) -> (\S*)`)

type instruction struct {
	command  string
	operand1 string
	operand2 string
	output   string
}

func getInstructions(input string) []instruction {

	ins := make([]instruction, 0)
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()

		matches := instructionRegex.FindStringSubmatch(line)

		switch matches[3] {
		case "RSHIFT":
		case "LSHIFT":
		case "OR":
		case "AND":
		case "":
			if matches[1] == "NOT" {
				// not operation
			} else {
				// assignment
			}

		}

		ins = append(ins, instruction{})
	}

	return ins
}

func (td dayEntry) PartOne(inputData string) (string, error) {
	return " -- Not Yet Implemented --", nil
}

func (td dayEntry) PartTwo(inputData string) (string, error) {
	return " -- Not Yet Implemented --", nil
}

func (td dayEntry) PuzzleInput() string {
	return "Actual Data"
}
