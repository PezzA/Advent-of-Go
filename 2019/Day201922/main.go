package Day201922

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pezza/advent-of-code/common"
)

type action struct {
	action string
	mod    int
}

type deck []int

func getCards(size int) deck {
	var cards deck
	for i := 0; i < size; i++ {
		cards = append(cards, i)
	}
	return cards
}

func getData(input string) []action {
	lines := strings.Split(input, "\n")
	var actionList []action
	for _, line := range lines {
		bits := strings.Fields(line)
		if bits[0] == "cut" {
			mod, _ := strconv.Atoi(bits[1])
			actionList = append(actionList, action{"cut", mod})
		} else if bits[2] == "increment" {
			mod, _ := strconv.Atoi(bits[3])
			actionList = append(actionList, action{"increment", mod})
		} else if bits[2] == "new" {
			actionList = append(actionList, action{"new", 0})
		}
	}
	return actionList
}

func (d deck) newStack() deck {
	var out deck
	for index := range d {
		out = append(out, d[len(d)-index-1])
	}
	return out
}

func (d deck) cut(in int) deck {
	if in > 0 {
		return append(d[in:], d[:in]...)
	}

	cut := common.Abs(in)
	return append(d[len(d)-cut:], d[:len(d)-cut]...)
}

func (d deck) increment(offset int) deck {
	newDeck := make(deck, len(d))

	counter := 0
	for i := range d {
		newDeck[counter] = d[i]
		counter += offset
		if counter >= len(d) {
			counter -= len(d)
		}
	}

	return newDeck
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	return fmt.Sprintf("%v", " -- Not Yet Implemented --")
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	return fmt.Sprintf("%v", " -- Not Yet Implemented --")
}
