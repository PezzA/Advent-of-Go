package Day201507

import (
	"bufio"
	"fmt"
	"regexp"
	"strings"
)

// Entry holds wraps the data and runner interfaces for this puzzle
var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string) {
	return 2015, 7, "Some Assembly Required"
}

var instructionRegex = regexp.MustCompile(`((([a-z0-9]*) (RSHIFT|OR|AND|LSHIFT)|NOT) )*([a-z0-9]*) -> ([a-z0-9]*)`)

type instruction struct {
	command  string
	operand1 string
	operand2 string
	output   string
}

func getInstructions(input string) []instruction {
	insList, scanner := make([]instruction, 0), bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		matches := instructionRegex.FindStringSubmatch(scanner.Text())

		var ins instruction

		switch matches[4] {
		case "RSHIFT", "LSHIFT", "OR", "AND":
			ins = instruction{
				command:  matches[4],
				operand1: matches[3],
				operand2: matches[5],
				output:   matches[6],
			}
		case "":
			if matches[2] == "NOT" {
				ins = instruction{
					command:  matches[2],
					operand1: matches[5],
					operand2: "",
					output:   matches[6],
				}
			} else {
				ins = instruction{
					command:  "ASSIGN",
					operand1: matches[5],
					operand2: "",
					output:   matches[6],
				}
			}
		}
		insList = append(insList, ins)
	}
	return insList
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	return fmt.Sprintf(" -- Not Yet Implemented --")
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	return fmt.Sprintf(" -- Not Yet Implemented --")
}

func (td dayEntry) PuzzleInput() string {
	return "Actual Data"
}
