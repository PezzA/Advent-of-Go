package Day202203

import (
	"fmt"
	"strings"
)

type rucksack string

func getData(input string) []rucksack {
	lines := strings.Split(input, "\n")

	data := []rucksack{}

	for _, line := range lines {
		data = append(data, rucksack(line))
	}

	return data
}

func (r rucksack) getCommonItem() rune {
	cmp1, cmp2 := splitInHalf(string(r))

	for _, char1 := range cmp1 {
		for _, char2 := range cmp2 {
			if char1 == char2 {
				return char1
			}
		}
	}
	return rune(0)
}

func getItemPriority(input rune) int {
	val := int(input)
	if val >= 97 {
		return val - 96
	}

	return val - 38
}

func getCommonItemType(r1, r2, r3 rucksack) rune {
	for _, char1 := range string(r1) {
		for _, char2 := range string(r2) {
			if char1 == char2 {
				for _, char3 := range string(r3) {
					if char1 == char3 {
						return char1
					}
				}
			}
		}
	}

	return rune(0)
}

// assumes always even lengthed input
func splitInHalf(input string) (string, string) {
	pos := len(input) / 2
	return input[:pos], input[pos:]
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	data := getData(inputData)
	total := 0

	for _, rs := range data {
		total += getItemPriority(rs.getCommonItem())
	}

	return fmt.Sprint(total)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	data := getData(inputData)
	total := 0

	for i := 0; i < len(data); i += 3 {
		total += getItemPriority(getCommonItemType(data[i], data[i+1], data[i+2]))
	}
	return fmt.Sprint(total)
}
