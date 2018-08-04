package Day201510

import (
	"fmt"
	"strconv"
)

// Entry holds wraps the data and runner interfaces for this puzzle
var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string) {
	return 2015, 10, "Elves Look, Elves Say"
}

func lookSay(input string) string {
	outString := ""
	curVal, count := rune(0), 0
	for _, val := range input {
		if curVal != val {

			if count != 0 {
				outString += strconv.Itoa(count) + string(curVal)
			}
			count = 1
			curVal = val
		} else {
			count += 1
		}
	}
	outString += strconv.Itoa(count) + string(curVal)

	return outString
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {

	data := Entry.PuzzleInput()
	count := 0
	for count < 40 {
		updateChan <- []string{"On iteration : " + strconv.Itoa(count)}
		data = lookSay(data)
		count++
	}
	return fmt.Sprintf("%v", len(data))
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	data := Entry.PuzzleInput()
	count := 0
	for count < 50 {
		updateChan <- []string{"On iteration : " + strconv.Itoa(count)}
		data = lookSay(data)
		count++
	}
	return fmt.Sprintf("%v", len(data))
}
