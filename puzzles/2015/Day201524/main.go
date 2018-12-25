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

	for index, val := range source {
		if val != compare[index] {
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
	totPres := len(presents)
	loop := int32(math.Pow(float64(2), float64(len(presents)+1)))

	combos := make([][]int, 0)
	min := 0
	//	currentEntanglement := 0

	// loop through each possible combination
	for i := int32(0); i < loop; i++ {
		currentCobmination, currentTotal, currentBits := make([]int, 0), 0, 0

		for powers := 1; powers <= totPres; powers++ {
			digit := powers - 1

			currPower := int32(math.Pow(2, float64(powers)))

			if i&currPower != 0 {
				currentCobmination = append(currentCobmination, presents[digit])

				currentBits++
				if currentBits > min && min != 0 {
					break
				}
				currentTotal += presents[digit]
				if currentTotal > targetWeight {
					continue
				}

				if currentTotal > targetWeight {
					break
				}

				// gone over abort
				if currentTotal == targetWeight {
					if min == 0 || currentBits <= min {
						min = currentBits
						combos = addNew(combos, currentCobmination)
					}

					break
				}
			}
		}

		if updateChan != nil {
			if i%10000000 == 0 {
				updateChan <- []string{fmt.Sprintf("%v of %v combinations checked, %v matches found.", i, loop, len(combos))}
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

func (td dayEntry) Describe() (int, int, string) {
	return 2015, 24, "It Hangs in the Balance"
}

//31112183811
func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	presents := getPresents(inputData)

	weight := totalweight(presents)

	combos := combinations(weight/3, presents, updateChan)

	sort.Slice(combos, func(i, j int) bool {

		if len(combos[i]) == len(combos[j]) {
			return getQuantumEntanglement(combos[i]) < getQuantumEntanglement(combos[j])
		}
		return len(combos[i]) < len(combos[j])
	})

	return fmt.Sprintf("%v", getQuantumEntanglement(combos[0]))
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	presents := getPresents(inputData)

	weight := totalweight(presents)

	combos := combinations(weight/4, presents, updateChan)

	sort.Slice(combos, func(i, j int) bool {

		if len(combos[i]) == len(combos[j]) {
			return getQuantumEntanglement(combos[i]) < getQuantumEntanglement(combos[j])
		}
		return len(combos[i]) < len(combos[j])
	})

	return fmt.Sprintf("%v", getQuantumEntanglement(combos[0]))
}
