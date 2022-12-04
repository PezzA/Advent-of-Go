package Day202204

import (
	"fmt"
	"strconv"
	"strings"
)

type elfRange struct {
	start int
	end   int
}

type elfPair struct {
	first  elfRange
	second elfRange
}

func getData(input string) []elfPair {
	lines := strings.Split(input, "\n")

	data := make([]elfPair, 0)

	for _, line := range lines {
		elves := strings.Split(line, ",")

		elf1bits := strings.Split(elves[0], "-")
		elf2bits := strings.Split(elves[1], "-")

		elf1Start, _ := strconv.Atoi(elf1bits[0])
		elf1End, _ := strconv.Atoi(elf1bits[1])
		elf2Start, _ := strconv.Atoi(elf2bits[0])
		elf2End, _ := strconv.Atoi(elf2bits[1])

		data = append(data, elfPair{elfRange{elf1Start, elf1End}, elfRange{elf2Start, elf2End}})
	}
	return data
}

func (ep elfPair) hasWhollyContainedElf() bool {
	if ep.first.start >= ep.second.start && ep.first.end <= ep.second.end {
		return true
	}
	if ep.second.start >= ep.first.start && ep.second.end <= ep.first.end {
		return true
	}
	return false
}

func (ep elfPair) hasAnyOverlap() bool {

	if ep.hasWhollyContainedElf() {
		return true
	}

	if ep.first.start <= ep.second.end && ep.first.end >= ep.second.start {
		return true
	}

	if ep.first.start >= ep.second.start && ep.first.start <= ep.second.end {
		return true
	}

	return false
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	pairs := getData(inputData)
	count := 0
	for _, pair := range pairs {
		if pair.hasWhollyContainedElf() {
			count++
		}
	}
	return fmt.Sprint(count)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	pairs := getData(inputData)
	count := 0
	for _, pair := range pairs {
		if pair.hasAnyOverlap() {
			count++
		}
	}
	return fmt.Sprint(count)
}
