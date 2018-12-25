package Day201508

import (
	"fmt"
	"strings"
)

// Entry holds wraps the data and runner interfaces for this puzzle
var Entry dayEntry

type dayEntry bool

func escapeString(input string) string {
	output := make([]rune, 0)

	for i := 0; i < len(input); i++ {
		if input[i] == 92 {
			if input[i+1] == 34 || input[i+1] == 92 {
				output = append(output, rune(input[i+1]))
				i++
			} else if input[i+1] == 120 {
				output = append(output, rune(48))
				i += 3
			}
		} else {
			output = append(output, rune(input[i]))
		}
	}

	return string(output[1 : len(output)-1])
}

func encodeString(input string) string {
	output := make([]rune, 0)

	for i := 0; i < len(input); i++ {
		if input[i] == 92 {
			output = append(output, 92)
			output = append(output, 92)
		} else if input[i] == 34 {
			output = append(output, 92)
			output = append(output, 34)
		} else {
			output = append(output, rune(input[i]))
		}
	}

	return `"` + string(output) + `"`
}

func (td dayEntry) Describe() (int, int, string) {
	return 2015, 8, "Matchsticks"
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	literalLength, actualChars := 0, 0
	for _, data := range strings.Split(inputData, "\n") {
		literalLength += len(data)
		actualChars += len(escapeString(data))
	}
	return fmt.Sprintf("%v", literalLength-actualChars)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	literalLength, encodedChars := 0, 0
	for _, data := range strings.Split(inputData, "\n") {
		literalLength += len(data)
		encodedChars += len(encodeString(data))
	}
	return fmt.Sprintf("%v", encodedChars-literalLength)
}
