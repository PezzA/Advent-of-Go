package Day201712

import (
	"bufio"
	"strconv"
	"strings"
)

// Entry holds wraps the data and runner interfaces for this puzzle
var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string) {
	return 2017, 12, "Digital Plumber"
}

type programList map[int][]int

func getPrograms(input string) programList {
	progList := make(programList, 0)

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		bits := strings.Split(scanner.Text(), " ")

		index := 0
		childList := make([]int, 0)

		for bitIndex, val := range bits {
			if bitIndex == 0 {
				index, _ = strconv.Atoi(val)
			}

			if bitIndex > 1 {
				val = strings.Replace(val, ",", "", 1)
				child, _ := strconv.Atoi(val)
				childList = append(childList, child)
			}

		}

		progList[index] = childList
	}
	return progList
}

// from index, recursively travel each pipe until we hit a program with no new programs then return
func travelPipes(index int, seenPrograms map[int]bool, pl programList) map[int]bool {
	if ok, _ := seenPrograms[index]; !ok {
		seenPrograms[index] = true
	}

	program := pl[index]

	for _, childIndex := range program {
		if ok, _ := seenPrograms[childIndex]; !ok {
			travelPipes(childIndex, seenPrograms, pl)
		}
	}

	return seenPrograms
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) (string, error) {
	programs := getPrograms(inputData)
	seenList := travelPipes(0, make(map[int]bool, 0), programs)
	return strconv.Itoa(len(seenList)), nil
}

// Coarse implementation, but getting in early for points.  Once you have a group, remove them from the list of all
// programs, detect next group and repeat until length of programs is zero.
func (td dayEntry) PartTwo(inputData string, updateChan chan []string) (string, error) {
	programs := getPrograms(inputData)

	seenList := travelPipes(0, make(map[int]bool, 0), programs)

	for key := range seenList {
		delete(programs, key)
	}

	groups := 1

	for len(programs) > 0 {
		groups++
		index := 0
		for key := range programs {
			index = key
			break
		}
		seenList := travelPipes(index, make(map[int]bool, 0), programs)

		for key := range seenList {
			delete(programs, key)
		}
	}

	return strconv.Itoa(groups), nil
}
