package Day201509

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pezza/advent-of-code/puzzles/common"
)

// Entry holds wraps the data and runner interfaces for this puzzle
var Entry dayEntry

type dayEntry bool

type route struct {
	start  string
	finish string
}

type routeList map[route]int

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
		if !common.Contains(destinations, k.start) {
			destinations = append(destinations, k.start)
		}
		if !common.Contains(destinations, k.finish) {
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

func (rl routeList) getDistance(path []string) int {
	distance := 0
	for i := 0; i < len(path)-1; i++ {
		pathSection := route{path[i], path[i+1]}

		if rl[pathSection] == 0 {
			pathSection = route{path[i+1], path[i]}
		}
		distance += rl[pathSection]
		if i == len(path)-1 {
			break
		}
	}
	return distance
}

func (td dayEntry) Describe() (int, int, string, int) {
	return 2015, 9, "All in a Single Night", 2
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	destinationMap := getData(inputData)
	destinationList := destinationMap.extractDestinations()
	resultChan := make(chan []string)

	go permutation(destinationList, make([]string, 0), resultChan)

	i := 0
	minDistance := -1

	for path := range resultChan {
		i++
		distance := destinationMap.getDistance(path)
		if minDistance == -1 || distance < minDistance {
			minDistance = distance
		}
		if i == common.Factorial(len(destinationList)) {
			close(resultChan)
		}
	}

	return fmt.Sprintf("%v", minDistance)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	destinationMap := getData(inputData)
	destinationList := destinationMap.extractDestinations()
	resultChan := make(chan []string)

	go permutation(destinationList, make([]string, 0), resultChan)

	i := 0
	maxDistance := 0

	for path := range resultChan {
		i++
		distance := destinationMap.getDistance(path)
		if distance > maxDistance {
			maxDistance = distance
		}
		if i == common.Factorial(len(destinationList)) {
			close(resultChan)
		}
	}

	return fmt.Sprintf("%v", maxDistance)
}
