package Day201505

import (
	"bufio"
	"strconv"
	"strings"
)

// Entry holds wraps the data and runner interfaces for this puzzle
var Entry testDay

type testDay bool

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
	bits := make(map[string]int, 0)

	for i := 0; i < len(input)-1; i++ {
		if _, ok := bits[input[i:i+2]]; ok {
			bits[input[i:i+2]]++
		} else {
			bits[input[i:i+2]] = 1
		}

		if input[i] == input[i+1] {
			i++
		}
	}

	for _, val := range bits {
		if val > 1 {
			return true
		}
	}

	return false
}

func (td testDay) PartOne(inputData string) (string, error) {
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

func (td testDay) PartTwo(inputData string) (string, error) {
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

func (td testDay) Day() int {
	return 201505
}

func (td testDay) Heading() string {
	return "--- (2015) Day 5: Doesn't He Have Intern-Elves For This? ---"
}

func (td testDay) GetTestData() ([]string, []string) {
	return []string{
			"ugknbfddgicrmopn",
			"aaa",
			"jchzalrnumimnmhp",
			"haegwjzuvuyypxyu",
			"dvszwmarrgswjxmb",
		},
		[]string{
			"qjhvhtzxzqqjkmpb",
			"xxyxx",
			"uurcxstgmygtbstg",
			"ieodomkazucvgmuy",
			"aaa",
			"aaaa",
			"xilodxfuxphuiiii",
		}
}
