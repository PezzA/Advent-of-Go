package Day202202

import (
	"fmt"
	"strings"
)

type round struct {
	opponentChoice string
	playerChoice   string
}

func newRound(opponent string, player string) round {
	return round{opponentChoice: opponent, playerChoice: player}
}

func (r round) shortCode() string {
	return r.opponentChoice + r.playerChoice
}

var outcomes = map[string]string{
	"AX": "Z",
	"AY": "X",
	"AZ": "Y",
	"BX": "X",
	"BY": "Y",
	"BZ": "Z",
	"CX": "Y",
	"CY": "Z",
	"CZ": "X",
}

func importData(input string) []round {
	lines := strings.Split(input, "\n")

	data := []round{}

	for _, line := range lines {
		bits := strings.Split(line, " ")
		data = append(data, newRound(bits[0], bits[1]))
	}

	return data
}

func scoreWithResult(input round) int {
	choice := outcomes[input.shortCode()]

	return scoreChoice(choice) +
		scoreOptions(input.opponentChoice, choice)
}

func scoreRound(input round) int {
	return scoreChoice(input.playerChoice) +
		scoreOptions(input.opponentChoice, input.playerChoice)
}

func scoreOptions(opponent string, player string) int {
	// A for Rock, B for Paper, and C for Scissors
	// X for Rock, Y for Paper, and Z for Scissors
	switch opponent + player {
	case "AY":
		fallthrough
	case "CX":
		fallthrough
	case "BZ":
		return 6
	case "AX":
		fallthrough
	case "BY":
		fallthrough
	case "CZ":
		return 3
	}

	return 0
}

// assumes input always valid
func scoreChoice(input string) int {
	switch input {
	case "X":
		return 1
	case "Y":
		return 2
	case "Z":
		return 3
	}
	return 0
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	predictions := importData(inputData)
	score := 0

	for _, gameRound := range predictions {
		score += scoreRound(gameRound)
	}

	return fmt.Sprint(score)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	predictions := importData(inputData)
	score := 0

	for _, gameRound := range predictions {
		score += scoreWithResult(gameRound)
	}

	return fmt.Sprint(score)
}
