package Day201901

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string) {
	return 2019, 1, "The Tyranny of the Rocket Equation"
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
	basicFuel := int(math.Floor(float64(module)/3) - 2)

	if basicFuel < 0 {
		basicFuel = 0
	}
	return basicFuel
}

func fuelFuelRequirement(module int) int {
	req := 0

	for {
		fuel := fuelRequirement(module)

		if fuel == 0 {
			break
		}
		req += fuel
		module = fuel
	}
	return req
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
		fuelNeeded += fuelFuelRequirement(mod)
	}

	return fmt.Sprintf("%v", fuelNeeded)
}
