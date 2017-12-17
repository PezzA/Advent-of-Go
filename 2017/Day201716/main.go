package Day201716

import (
	"regexp"
	"strconv"
	"strings"
)

type dayEntry bool

var Entry dayEntry

func (td dayEntry) Describe() (int, int, string) {
	return 2017, 16, "Permutation Promenade"
}

var instructionRegex = regexp.MustCompile("((x)([0-9]+)\\/([0-9]+)|(s)([0-9]+)|(p)([a-z])\\/([a-z]))")

type progList string

type instruction struct {
	letter   string
	operand1 string
	operand2 string
}

func getInstructions(input string) []instruction {
	bits, instructions := strings.Split(input, ","), make([]instruction, 0)

	for _, val := range bits {
		match := instructionRegex.FindStringSubmatch(val)

		if match[2] == "x" {
			instructions = append(instructions, instruction{match[2], match[3], match[4]})
		}

		if match[5] == "s" {
			instructions = append(instructions, instruction{match[5], match[6], ""})
		}

		if match[7] == "p" {
			instructions = append(instructions, instruction{match[7], match[8], match[9]})
		}
	}

	return instructions
}

func (pl progList) spin(spinAmount int) progList {
	return pl[len(pl)-spinAmount:] + pl[:len(pl)-spinAmount]
}

func (pl progList) exchangePos(posA int, posB int) progList {
	newInput := []byte(pl)
	newInput[posA] = pl[posB]
	newInput[posB] = pl[posA]
	return progList(newInput)
}

func (pl progList) exchangeProg(progA string, progB string) progList {
	indexA := strings.Index(string(pl), progA)
	indexB := strings.Index(string(pl), progB)
	return pl.exchangePos(indexA, indexB)
}

func (pl progList) doDance(ins []instruction) progList {
	for _, val := range ins {
		switch val.letter {
		case "s":
			spinAmount, _ := strconv.Atoi(val.operand1)
			pl = pl.spin(spinAmount)
		case "p":
			pl = pl.exchangeProg(val.operand1, val.operand2)
		case "x":
			posA, _ := strconv.Atoi(val.operand1)
			posB, _ := strconv.Atoi(val.operand2)
			pl = pl.exchangePos(posA, posB)
		}
	}

	return pl
}

var masterProgList = progList("abcdefghijklmnop")

func (td dayEntry) PartOne(inputData string) (string, error) {
	programs, instructions := masterProgList, getInstructions(inputData)
	return string(programs.doDance(instructions)), nil
}

func (td dayEntry) PartTwo(inputData string) (string, error) {
	dance, instructions := masterProgList, getInstructions(inputData)

	loopIndex := 0
	for {
		dance = dance.doDance(instructions)

		if dance == masterProgList {
			break
		}
		loopIndex++
	}

	actual := 1000000000 % (loopIndex + 1)

	dance = masterProgList
	for i := 0; i < actual; i++ {
		dance = dance.doDance(instructions)
	}

	return string(dance), nil
}
