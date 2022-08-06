package Day201511

import (
	"strings"
)

// Entry holds wraps the data and runner interfaces for this puzzle
var Entry dayEntry

type dayEntry bool

func updateString(input string, index int, newChar rune) string {
	/*var r strings.Builder
	if index > 0 {
		r.WriteString(input[0:index])
		r.WriteRune(newChar)
		r.WriteString(input[index+1:])
	} else {
		r.WriteRune(newChar)
		r.WriteString(input[index+1:])
	}*/
	return ""
}

func increment(input string) string {
	if input == "" {
		return "a"
	}

	// cache
	inputLen := len(input)

	// for each char in the string starting at the right
	for index := range input {
		currIndex := (inputLen - index) - 1
		char := input[currIndex]

		if char < 122 {
			return updateString(input, currIndex, rune(char+1))
		} else if char == 122 {
			input = updateString(input, currIndex, 97)
		}
	}

	return "a" + input
}

func hasStraight(input string) bool {
	for index, _ := range input {
		if index == len(input)-2 {
			return false
		}

		if input[index+2] == input[index]+2 &&
			input[index+1] == input[index]+1 {
			return true
		}
	}
	return false
}

func validChars(input string) bool {
	return !strings.ContainsAny(input, "iol")
}

func hasTwoPairs(input string) bool {
	hits := make(map[string]bool)

	var testChar rune
	for _, char := range input {
		if testChar != char {
			testChar = char
			continue
		}

		hits[string(char)] = true
		testChar = 0
	}
	return len(hits) == 2
}

func (td dayEntry) Describe() (int, int, string, int) {
	return 2015, 11, "Corporate Policy", 4
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	seq := Entry.PuzzleInput()

	for {
		seq = increment(seq)

		if hasTwoPairs(seq) && hasStraight(seq) && validChars(seq) {
			break
		}
	}

	return seq
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	seq := Entry.PartOne(Entry.PuzzleInput(), updateChan)

	for {
		seq = increment(seq)

		if hasTwoPairs(seq) && hasStraight(seq) && validChars(seq) {
			break
		}
	}

	return seq
}
