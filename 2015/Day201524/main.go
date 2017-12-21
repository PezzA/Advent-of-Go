package Day201524

import (
	"bufio"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

// Entry holds wraps the data and runner interfaces for this puzzle
var Entry dayEntry

type dayEntry bool

func getPresents(inputData string) []int {
	presents := make([]int, 0)
	scanner := bufio.NewScanner(strings.NewReader(inputData))
	for scanner.Scan() {
		line := scanner.Text()
		present, _ := strconv.Atoi(line)
		presents = append(presents, present)
	}

	return presents
}

func totalweight(presents []int) int {
	weight := 0
	for _, val := range presents {
		weight += val
	}
	return weight
}

func isMatch(source []int, compare []int) bool {
	if len(source) != len(compare) {
		return false
	}

	for index := range source {
		if source[index] != compare[index] {
			return false
		}
	}

	return true
}

func addNew(combos [][]int, test []int) [][]int {
	if len(combos) == 0 {
		return append(combos, test)
	}

	for _, val := range combos {
		if isMatch(val, test) {
			return combos
		}
	}

	return append(combos, test)
}

func combinations(targetWeight int, presents []int, updateChan chan []string) [][]int {
	totPres := int32(len(presents))
	loop := int32(math.Pow(float64(2), float64(len(presents))))

	loop100, update := loop/100, 0
	combos := make([][]int, 0)

	for i := int32(0); i < loop; i++ {
		test := make([]int, 0)

		currVal := 0

		for powers := int32(0); powers < totPres; powers++ {
			currPower := int32(1)

			if powers > 0 {
				currPower = int32(math.Pow(float64(2), float64(powers+1)))
			}

			if i&currPower != 0 {
				test = append(test, presents[powers])
				currVal += presents[powers]
				if currVal > targetWeight {
					break
				}

				if currVal == targetWeight {

					combos = addNew(combos, test)

					break
				}
			}
		}

		if updateChan != nil {
			if i%loop100 == 0 {
				update++
				updateChan <- []string{fmt.Sprintf("%v%% complete, %v matches found.", update, len(combos))}
			}
		}

	}
	return combos
}

func getQuantumEntanglement(combo []int) int {
	qe := combo[0]

	for index, val := range combo {
		if index != 0 {
			qe *= val
		}
	}
	return qe
}

func takeSpecificNumber(targetWeight int, numToTake int, presents []int) [][]int {
	candidates := make([][]int, 0)
	for _, p1 := range presents {
		if p1 == targetWeight {
			candidates = append(candidates, []int{p1})
		}
	}
	return candidates
}

func (td dayEntry) Describe() (int, int, string) {
	return 2015, 24, "It Hangs in the Balance"
}

//31112183811
func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {

	presents := getPresents(inputData)
	combos := combinations(totalweight(presents)/3, presents, updateChan)

	sort.Slice(combos, func(i, j int) bool {

		if len(combos[i]) == len(combos[j]) {
			return getQuantumEntanglement(combos[i]) < getQuantumEntanglement(combos[j])
		}
		return len(combos[i]) < len(combos[j])
	})

	return fmt.Sprintf("%v", getQuantumEntanglement(combos[0]))
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	return fmt.Sprintf(" -- Not Yet Implemented --")
}
