package Day201513

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pezza/advent-of-code/common"
)

// Entry holds wraps the data and runner interfaces for this puzzle
var Entry dayEntry

type dayEntry bool

type seating struct {
	guest   string
	partner string
}

type modifierList map[seating]int

func getData(input string) modifierList {
	modList := make(modifierList)
	for _, line := range strings.Split(input, "\n") {
		bits := strings.Fields(line)

		guest, sign, modifier, partner := bits[0], bits[2], bits[3], strings.Replace(bits[10], ".", "", -1)
		mod, _ := strconv.Atoi(modifier)

		if sign == "lose" {
			mod = -mod
		}

		modList[seating{guest, partner}] = mod
	}

	return modList
}

func (ml modifierList) getGuestList() []string {
	guestList := make([]string, 0)

	for k := range ml {
		if !common.Contains(guestList, k.guest) {
			guestList = append(guestList, k.guest)
		}
		if !common.Contains(guestList, k.partner) {
			guestList = append(guestList, k.partner)
		}
	}

	return guestList
}

// should work ok for small modifiers
func getZeroWrappedIndex(currIndex int, maxIndex int, modifier int) int {
	if modifier == 0 {
		return currIndex
	}

	if modifier > 0 {
		for i := 0; i < modifier; i++ {
			if currIndex+1 > maxIndex {
				currIndex = 0
			} else {
				currIndex++
			}
		}
	}

	if modifier < 0 {
		for i := modifier; i < 0; i++ {
			if currIndex == 0 {
				currIndex = maxIndex
			} else {
				currIndex--
			}
		}
	}

	return currIndex
}

func (ml modifierList) addNeutralGuest(name string) {

	guestList := ml.getGuestList()

	for _, val := range guestList {
		ml[seating{name, val}] = 0
		ml[seating{val, name}] = 0
	}
}
func (ml modifierList) scoreList(input []string) int {
	score := 0
	for index, val := range input {
		leftPartner := input[getZeroWrappedIndex(index, len(input)-1, 1)]
		rightPartner := input[getZeroWrappedIndex(index, len(input)-1, -1)]

		score += ml[seating{val, leftPartner}]
		score += ml[seating{val, rightPartner}]
	}

	return score
}
func permutation(valueList []string, acc []string, resultChan chan []string) {
	if len(valueList) == 1 {
		resultChan <- append(acc, valueList[0])
		return
	}
	for _, value := range valueList {
		permutation(common.Remove(valueList, value), append(acc, value), resultChan)
	}
}

func (td dayEntry) Describe() (int, int, string) {
	return 2015, 13, "Knights of the Dinner Table"
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	seatingList := getData(Entry.PuzzleInput())
	guestList := seatingList.getGuestList()
	resultChan := make(chan []string)

	go permutation(guestList, make([]string, 0), resultChan)

	i := 0
	modifier := 0
	for gl := range resultChan {
		i++
		//	fmt.Println(perm, i)
		currMod := seatingList.scoreList(gl)
		if currMod > modifier {
			modifier = currMod
		}
		if i == common.Factorial(len(guestList)) {
			break
		}
	}
	return fmt.Sprintf("%v", modifier)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	seatingList := getData(Entry.PuzzleInput())
	seatingList.addNeutralGuest("John")
	guestList := seatingList.getGuestList()
	resultChan := make(chan []string)

	go permutation(guestList, make([]string, 0), resultChan)

	i := 0
	modifier := 0
	for gl := range resultChan {
		i++
		//	fmt.Println(perm, i)
		currMod := seatingList.scoreList(gl)
		if currMod > modifier {
			modifier = currMod

		}
		if i == common.Factorial(len(guestList)) {
			break
		}
	}

	return fmt.Sprintf("%v", modifier)
}
