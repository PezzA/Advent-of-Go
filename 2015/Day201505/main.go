package Day201505

import (
	"bufio"
	"strconv"
	"strings"
)

// Entry holds wraps the data and runner interfaces for this puzzle
var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string) {
	return 2015, 5, "--- (2015) Day 5: Doesn't He Have Intern-Elves For This? ---"
}

var vowels = "aeiou"
var invalidMatches map[string]bool

func init() {
	invalidMatches = make(map[string]bool, 0)
	invalidMatches["ab"] = true
	invalidMatches["cd"] = true
	invalidMatches["pq"] = true
	invalidMatches["xy"] = true
}

func sufficentVowels(input string) bool {
	count := 0

	for _, char := range input {
		if strings.ContainsRune(vowels, char) {
			count++
		}
	}

	return count >= 3
}

func hasRepeat(input string) bool {
	for i := 0; i < len(input)-1; i++ {
		if input[i] == input[i+1] {
			return true
		}
	}
	return false
}

func hasBumpy(input string) bool {
	for i := 0; i < len(input)-2; i++ {
		if input[i] == input[i+2] {
			return true
		}
	}
	return false
}

func noInvalidMatches(input string) bool {
	for i := 0; i < len(input)-1; i++ {
		if _, ok := invalidMatches[input[i:i+2]]; ok {
			return false
		}
	}
	return true
}

func hasRepeatNoOverlap(input string) bool {
	bits, span := make(map[string]int, 0), ""

	for i := 0; i < len(input)-1; i++ {

		if span == input[i:i+2] {
			span = ""
			continue
		}
		span = input[i : i+2]

		if _, ok := bits[span]; ok {
			bits[span]++
		} else {
			bits[span] = 1
		}
	}

	for _, val := range bits {
		if val > 1 {
			return true
		}
	}

	return false
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) (string, error) {
	total := 0
	scanner := bufio.NewScanner(strings.NewReader(inputData))
	for scanner.Scan() {
		testText := scanner.Text()
		if sufficentVowels(testText) && noInvalidMatches(testText) && hasRepeat(testText) {
			total++
		}
	}
	return strconv.Itoa(total), nil
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) (string, error) {
	total := 0
	scanner := bufio.NewScanner(strings.NewReader(inputData))
	for scanner.Scan() {
		testText := scanner.Text()
		if hasBumpy(testText) && hasRepeatNoOverlap(testText) {
			total++
		}
	}
	return strconv.Itoa(total), nil
}
