package Day201802

import (
	"fmt"
	"strings"
)

// Entry holds wraps the data and runner interfaces for this puzzle
var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string, int) {
	return 2018, 2, "Inventory Management System", 0
}

func getData(input string) []string {
	return strings.Split(input, "\n")
}

func getCounts(input string) map[rune]int {
	counts := make(map[rune]int)

	for _, char := range input {
		counts[char]++
	}

	return counts
}

func hasSameLetterCount(min int, counts map[rune]int) bool {
	for _, v := range counts {
		if v == min {
			return true
		}
	}

	return false
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	ids := getData(inputData)

	twos, threes := 0, 0
	for _, id := range ids {
		counts := getCounts(id)

		if hasSameLetterCount(2, counts) {
			twos++
		}

		if hasSameLetterCount(3, counts) {
			threes++
		}
	}
	return fmt.Sprintf("%v", twos*threes)
}

func findCommonLetters(ids []string) string {
	for _, val := range ids {
		for _, testVal := range ids {
			diffs := 0
			foundIndex := 0
			for index := range val {
				if val[index] != testVal[index] {
					diffs++
					foundIndex = index
				}
			}

			if diffs == 1 {
				return val[:foundIndex] + val[foundIndex+1:]
			}
		}
	}

	return ""
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	return fmt.Sprint(findCommonLetters(getData(inputData)))
}
