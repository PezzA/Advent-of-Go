package Day201901

import (
	"fmt"
	"strconv"
	"strings"
)

var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string, int) {
	return 2019, 1, "The Tyranny of the Rocket Equation", 0
}

func getData(input string) []int {
	lines := strings.Split(input, "\n")

	modules := []int{}

	for index := range lines {
		f, _ := strconv.Atoi(lines[index])
		modules = append(modules, f)
	}

	return modules
}

func fuelRequirement(module int) int {
	req := module/3 - 2
	if req < 0 {
		return 0
	}
	return req

}

func fuel(module int) int {
	req := fuelRequirement(module)
	if req == 0 {
		return 0
	}
	return req + fuel(req)
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	fuelNeeded := 0

	for _, mod := range getData(inputData) {
		fuelNeeded += fuelRequirement(mod)
	}

	return fmt.Sprintf("%v", fuelNeeded)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	fuelNeeded := 0

	for _, mod := range getData(inputData) {
		fuelNeeded += fuel(mod)
	}

	return fmt.Sprintf("%v", fuelNeeded)
}
