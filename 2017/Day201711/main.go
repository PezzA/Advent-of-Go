package Day201711

import (
	"sort"
	"strconv"
	"strings"
)

// Entry holds wraps the data and runner interfaces for this puzzle
var Entry testDay

type testDay bool

func travel(input string) (int, int) {
	directions := strings.Split(input, ",")

	x, y, max := 0, 0, 0

	for _, val := range directions {
		switch val {
		case "n":
			y++
		case "ne":
			x++
		case "se":
			x++
			y--
		case "s":
			y--
		case "sw":
			x--
		case "nw":
			x--
			y++
		}

		vals := []int{x, y, -(x + y)}

		sort.Ints(vals)

		if vals[2] > max {
			max = vals[2]
		}
	}

	vals := []int{x, y, -(x + y)}

	sort.Ints(vals)

	return vals[2], max
}

func (td testDay) PartOne(inputData string) (string, error) {
	distance, _ := travel(inputData)
	return strconv.Itoa(distance), nil
}

func (td testDay) PartTwo(inputData string) (string, error) {
	_, maxDistance := travel(inputData)
	return strconv.Itoa(maxDistance), nil
}

func (td testDay) Day() int {
	return 11
}

func (td testDay) Year() int {
	return 2017
}

func (td testDay) Title() string {
	return "Hex Ed"
}
