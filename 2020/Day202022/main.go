package Day202022

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pezza/advent-of-code/common"
)

type deck []int

func getData(input string) (deck, deck) {
	capPlayer1, capPlayer2 := false, false
	deck1, deck2 := make(deck, 0), make(deck, 0)

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		if line == "Player 1:" {
			capPlayer1 = true
			continue
		}

		if line == "Player 2:" {
			capPlayer2 = true
			capPlayer1 = false
			continue
		}

		num, _ := strconv.Atoi(line)

		if capPlayer1 {
			deck1 = append(deck1, num)
		}

		if capPlayer2 {
			deck2 = append(deck2, num)
		}
	}

	return deck1, deck2
}

func (d deck) draw() (int, deck) {
	return d[0], d[1:]
}

func (d deck) score() int {
	score := 0
	for i := range d {
		score += d[i] * (len(d) - i)
	}

	return score
}
func runGame(p1Deck deck, p2Deck deck) (deck, deck) {
	var p1Card, p2Card int

	for {
		p1Card, p1Deck = p1Deck.draw()
		p2Card, p2Deck = p2Deck.draw()

		if p1Card > p2Card {
			p1Deck = append(p1Deck, p1Card)
			p1Deck = append(p1Deck, p2Card)
		} else {
			p2Deck = append(p2Deck, p2Card)
			p2Deck = append(p2Deck, p1Card)
		}
		if len(p1Deck) == 0 || len(p2Deck) == 0 {
			break
		}
	}

	return p1Deck, p2Deck
}

type state struct {
	p1State deck
	p2State deck
}

func isPrevState(p1 deck, p2 deck, states []state) bool {
	for i := range states {
		if common.Same(states[i].p1State, p1) && common.Same(states[i].p2State, p2) {
			return true
		}
	}
	return false
}

func runRecursiveGame(p1Deck deck, p2Deck deck, states []state) (deck, deck) {
	var p1Card, p2Card int

	for {
		// infinite game prevention rule
		if isPrevState(p1Deck, p2Deck, states) {
			return p1Deck, []int{}
		} else {
			states = append(states, state{common.DeepCopy(p1Deck), common.DeepCopy(p2Deck)})
		}

		// draw cards
		p1Card, p1Deck = p1Deck.draw()
		p2Card, p2Deck = p2Deck.draw()

		p1CardWinner := p1Card > p2Card

		// do recurse
		if p1Card <= len(p1Deck) && p2Card <= len(p2Deck) {
			subDeck1 := common.DeepCopy(p1Deck[:p1Card])
			subDeck2 := common.DeepCopy(p2Deck[:p2Card])
			rDeck1, _ := runRecursiveGame(subDeck1, subDeck2, make([]state, 0))
			p1CardWinner = len(rDeck1) > 0
		}

		// card check
		if p1CardWinner {
			p1Deck = append(p1Deck, p1Card)
			p1Deck = append(p1Deck, p2Card)
		} else {
			p2Deck = append(p2Deck, p2Card)
			p2Deck = append(p2Deck, p1Card)
		}

		if len(p1Deck) == 0 || len(p2Deck) == 0 {
			break
		}
	}

	return p1Deck, p2Deck
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	d1, d2 := getData(inputData)
	d1, d2 = runGame(d1, d2)

	score := 0
	if len(d1) > 0 {
		score = d1.score()
	} else {
		score = d2.score()
	}

	return fmt.Sprintf("%v", score)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	d1, d2 := getData(Entry.PuzzleInput())
	d1, d2 = runRecursiveGame(d1, d2, []state{})

	score := 0
	if len(d1) > 0 {
		score = d1.score()
	} else {
		score = d2.score()
	}
	return fmt.Sprintf("%v", score)
}
