package Day201704

import (
	"bufio"
	"sort"
	"strconv"
	"strings"
)

var Entry testDay

type testDay bool

func stringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}

func sortStringByCharacter(s string) string {
	r := stringToRuneSlice(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

func noAnagram(input string) bool {
	bits := strings.Fields(input)
	sortedBits := make([]string, 0)
	counts := make(map[string]int, 0)

	for _, val := range bits {
		sortedBits = append(sortedBits, sortStringByCharacter(val))
	}

	for _, val := range sortedBits {
		if _, ok := counts[val]; ok {
			return false
		}

		counts[val] = 1
	}

	return true
}

func noDupedWord(input string) bool {
	bits := strings.Fields(input)

	counts := make(map[string]int, 0)

	for _, val := range bits {
		if _, ok := counts[val]; ok {
			return false
		}

		counts[val] = 1
	}
	return true
}

func (td testDay) PartOne(inputData string) (string, error) {

	total := 0
	scanner := bufio.NewScanner(strings.NewReader(inputData))
	for scanner.Scan() {
		if noDupedWord(scanner.Text()) {
			total++
		}
	}

	return strconv.Itoa(total), nil
}

func (td testDay) PartTwo(inputData string) (string, error) {
	total := 0
	scanner := bufio.NewScanner(strings.NewReader(inputData))
	for scanner.Scan() {
		if noDupedWord(scanner.Text()) && noAnagram(scanner.Text()) {
			total++
		}
	}

	return strconv.Itoa(total), nil
}

func (td testDay) Day() int {
	return 201704
}

func (td testDay) Heading() string {
	return "--- (2017) Day 4: High-Entropy Passphrases ---"
}
