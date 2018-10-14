package Day201509

import (
	"fmt"
	"strconv"
	"strings"
)

// Entry holds wraps the data and runner interfaces for this puzzle
var Entry dayEntry

type dayEntry bool

type route struct {
	start  string
	finish string
}

type routeList map[route]int

func contains(list []string, test string) bool {
	for _, val := range list {
		if val == test {
			return true
		}
	}
	return false
}

func factorial(n int) (result int) {
	if n > 0 {
		result = n * factorial(n-1)
		return result
	}
	return 1
}

func getData(input string) routeList {
	routes := make(routeList, 0)
	for _, line := range strings.Split(input, "\n") {
		bits := strings.Fields(line)
		distance, _ := strconv.Atoi(bits[4])
		routes[route{bits[0], bits[2]}] = distance
	}
	return routes
}

func (rl routeList) extractDestinations() []string {
	destinations := make([]string, 0)
	for k := range rl {
		if !contains(destinations, k.start) {
			destinations = append(destinations, k.start)
		}
		if !contains(destinations, k.finish) {
			destinations = append(destinations, k.finish)
		}
	}
	return destinations
}

func removeDestination(input []string, target string) []string {
	retVal := make([]string, 0)

	for _, val := range input {
		if val != target {
			retVal = append(retVal, val)
		}
	}

	return retVal
}
func permutation(destinations []string, acc []string, resultChan chan []string) {
	if len(destinations) == 1 {
		resultChan <- append(acc, destinations[0])
		return
	}
	for _, destination := range destinations {
		permutation(removeDestination(destinations, destination), append(acc, destination), resultChan)
	}
}

func (td dayEntry) Describe() (int, int, string) {
	return 2015, 9, "All in a Single Night"
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	destinations := getData(inputData).extractDestinations()

	resultChan := make(chan []string)

	go permutation(destinations, make([]string, 0), resultChan)

	i := 0
	for range resultChan {
		i++
		fmt.Println(i)

		if i == factorial(len(destinations)) {
			close(resultChan)
		}
	}
	return fmt.Sprintf(" -- Not Yet Implemented --")
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	return fmt.Sprintf(" -- Not Yet Implemented --")
}
