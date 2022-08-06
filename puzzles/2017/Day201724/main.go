package Day201724

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type dayEntry bool

var Entry dayEntry

func (td dayEntry) Describe() (int, int, string, int) {
	return 2017, 24, "Electromagnetic Moat", 4
}

type component struct {
	port1     int
	port1used bool
	port2     int
	port2used bool
}

func getComponents(input string) []component {
	components := make([]component, 0)
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()

		bits := strings.Split(line, "/")

		port1, _ := strconv.Atoi(bits[0])
		port2, _ := strconv.Atoi(bits[1])

		components = append(components, component{port1, false, port2, false})
	}

	return components
}

func (c component) String() string {

	bit1 := ""
	if c.port1used {
		//	bit1 = tm.Color(fmt.Sprintf("%v", c.port1), tm.RED)
	} else {
		//	bit1 = tm.Color(fmt.Sprintf("%v", c.port1), tm.GREEN)
	}

	bit2 := ""
	if c.port2used {
		//	bit2 = tm.Color(fmt.Sprintf("%v", c.port2), tm.RED)
	} else {
		//	bit2 = tm.Color(fmt.Sprintf("%v", c.port2), tm.GREEN)
	}
	return fmt.Sprintf("%v/%v", bit1, bit2)
}

func getNewComponentList(cl []component) []component {
	newList := make([]component, 0)
	for index := range cl {
		newList = append(newList, cl[index])
	}
	return newList
}

func isMatch(source []component, compare []component) bool {
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

func addNew(combos [][]component, test []component) [][]component {
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

func dedupeList(cl [][]component) [][]component {
	newList := make([][]component, 0)

	for index := range cl {

		newList = addNew(newList, cl[index])

	}

	return newList
}

func getMatchingPin(src component, dst component) (bool, int, int) {

	if src.port1 == dst.port1 && !src.port1used && !dst.port1used {
		return true, 1, 1
	}

	if src.port1 == dst.port2 && !src.port1used && !dst.port2used {
		return true, 1, 2
	}

	if src.port2 == dst.port1 && !src.port2used && !dst.port1used {
		return true, 2, 1
	}

	if src.port2 == dst.port2 && !src.port2used && !dst.port2used {
		return true, 2, 2
	}

	return false, 0, 0
}

var longest = 0
var longestLength = 0

func getBridges(cmp component, currentBridge []component, componentsleft []component, bridges *[][]component) {
	currentBridge = append(currentBridge, cmp)

	for index := range componentsleft {
		match, _, cmp2 := getMatchingPin(cmp, componentsleft[index])

		if match {

			newComp := component{
				componentsleft[index].port1,
				componentsleft[index].port1used,
				componentsleft[index].port2,
				componentsleft[index].port2used,
			}

			if cmp2 == 1 {
				newComp.port1used = true
			} else {
				newComp.port2used = true
			}

			newList := make([]component, len(componentsleft))
			copy(newList, componentsleft)

			newComponentList := append(newList[:index], newList[index+1:]...)

			getBridges(newComp, currentBridge, newComponentList, bridges)
		}
	}
	fmt.Println(currentBridge)

	newbridges := *bridges
	newbridges = append(newbridges, currentBridge)
}

func getBridgeCombos(input []component) [][]component {
	bridges := make([][]component, 0)

	for index := range input {

		part1Zero := input[index].port1 == 0
		part2Zero := input[index].port2 == 0

		if part1Zero || part2Zero {
			newComp := component{input[index].port1, part1Zero, input[index].port2, part2Zero}

			newList := make([]component, len(input))
			copy(newList, input)

			getBridges(newComp, make([]component, 0), append(newList[:index], newList[index+1:]...), &bridges)
		}
	}
	return bridges
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {

	//components := getComponents(Entry.PuzzleInput())
	//combos := getBridgeCombos(components)

	//for
	return fmt.Sprintf("%v", "something someday")
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {

	components := getComponents(Entry.PuzzleInput())

	combos := getBridgeCombos(components)
	return fmt.Sprintf("%v", len(combos))
}
