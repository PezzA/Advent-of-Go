package Day201702

import (
	"bufio"
	"errors"
	"math"
	"strconv"
	"strings"
)

// Entry holds wraps the data and runner interfaces for this puzzle
var Entry testDay

type testDay bool

// assumes whitespace delimited string of values
func diff(input string) int {
	bits := strings.Fields(input)

	initVal, _ := strconv.Atoi(bits[0])
	min := initVal
	max := initVal

	for _, val := range bits {
		intVal, _ := strconv.Atoi(val)

		if intVal > max {
			max = intVal
		}

		if intVal < min {
			min = intVal
		}
	}

	return max - min
}

func evenDivide(input string) (int, error) {
	bits := strings.Fields(input)

	for _, first := range bits {
		for _, second := range bits {

			if first == second {
				continue
			}

			fFloat, _ := strconv.ParseFloat(first, 64)
			sFloat, _ := strconv.ParseFloat(second, 64)

			dividedResult := fFloat / sFloat

			intpart, div := math.Modf(dividedResult)

			if div == 0 {
				return int(intpart), nil
			}
		}
	}

	return 0, errors.New("did not find dividing numbers")
}

func (td testDay) PartOne(inputData string) (string, error) {
	total := 0
	scanner := bufio.NewScanner(strings.NewReader(inputData))
	for scanner.Scan() {
		total += diff(scanner.Text())
	}

	return strconv.Itoa(total), nil

}

func (td testDay) PartTwo(inputData string) (string, error) {
	total := 0
	scanner := bufio.NewScanner(strings.NewReader(inputData))
	for scanner.Scan() {
		subtotal, _ := evenDivide(scanner.Text())
		total += subtotal
	}

	return strconv.Itoa(total), nil
}

func (td testDay) Day() int {
	return 201702
}

func (td testDay) Heading() string {
	return "--- (2017) Day 2: Corruption Checksum ---"
}
