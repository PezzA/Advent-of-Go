package Day201502

import "strings"
import "strconv"
import "bufio"
import "sort"

// Entry holds wraps the data and runner interfaces for this puzzle
var Entry testDay

type testDay bool

type present struct {
	l int
	h int
	w int
}

func (p present) surfaceArea() int {
	return (2 * p.l * p.w) + (2 * p.w * p.h) + (2 * p.h * p.l)
}

func (p present) sortedInts() []int {
	intList := []int{p.l, p.h, p.w}
	sort.Ints(intList)
	return intList
}

func (p present) littleBitExtra() int {
	intList := p.sortedInts()
	return intList[0] * intList[1]
}

func (p present) paperRequired() int {
	return p.surfaceArea() + p.littleBitExtra()
}

func (p present) ribbonLength() int {
	intList := p.sortedInts()
	return (intList[0] * 2) + (intList[1] * 2)
}

func (p present) bowSize() int {
	return p.h * p.l * p.w
}

func (p present) ribbonRequired() int {
	return p.ribbonLength() + p.bowSize()
}

func parsePresent(input string) present {
	fields := strings.Split(strings.TrimSpace(input), "x")

	l, _ := strconv.Atoi(fields[0])
	h, _ := strconv.Atoi(fields[1])
	w, _ := strconv.Atoi(fields[2])

	return present{
		l: l,
		h: h,
		w: w,
	}
}

func (td testDay) PartOne(inputData string) (string, error) {
	total := 0
	scanner := bufio.NewScanner(strings.NewReader(inputData))
	for scanner.Scan() {
		present := parsePresent(scanner.Text())
		total += present.paperRequired()
	}

	return strconv.Itoa(total), nil
}

func (td testDay) PartTwo(inputData string) (string, error) {
	total := 0
	scanner := bufio.NewScanner(strings.NewReader(inputData))
	for scanner.Scan() {
		present := parsePresent(scanner.Text())
		total += present.ribbonRequired()
	}

	return strconv.Itoa(total), nil
}

func (td testDay) Day() int {
	return 201502
}

func (td testDay) Heading() string {
	return "--- (2015) Day 2: I Was Told There Would Be No Math ---"
}

func (td testDay) GetTestData() ([]string, []string) {
	return []string{
			"2x3x4",
			"1x1x10",
		}, []string{
			"2x3x4",
			"1x1x10",
		}
}
