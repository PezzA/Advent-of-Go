package Day202025

import (
	"fmt"
	"strconv"
	"strings"
)

func getData(input string) (int, int) {
	lines := strings.Split(input, "\n")

	cardPublicKey, _ := strconv.Atoi(lines[0])
	doorPublicKey, _ := strconv.Atoi(lines[1])

	return cardPublicKey, doorPublicKey
}

func transForm(in int, subject int) int {
	val := in * subject
	return val % 20201227
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	cardPublicKey, doorPublicKey := getData(inputData)
	val := 1

	cardLoopSize := 0
	for i := 0; true; i++ {
		val = transForm(val, 7)
		if cardPublicKey == val {
			cardLoopSize = i
			break
		}
	}

	encVal := 1
	for i := 0; i <= cardLoopSize; i++ {
		encVal = transForm(encVal, doorPublicKey)
	}

	return fmt.Sprintf("%v", encVal)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	return fmt.Sprintf("%v", " -- Not Yet Implemented --")
}
