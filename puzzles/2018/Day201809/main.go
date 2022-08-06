package Day201809

import (
	"container/list"
	"fmt"
)

var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string, int) {
	return 2018, 9, "Marble Mania", 0
}

func getData(input string) (int, int) {
	players, lastMarbleValue := 0, 0
	fmt.Sscanf(input, "%d players; last marble is worth %d points", &players, &lastMarbleValue)
	return players, lastMarbleValue
}

func playGame(players int, marbles int) int {

	playerScores := make(map[int]int)

	l := list.New()
	l.PushBack(0)
	currMarble := l.PushBack(1)

	currPlayer := 2
	for marble := 2; marble <= marbles; marble++ {

		if marble%23 == 0 {
			playerScores[currPlayer] += marble

			for i := 0; i < 7; i++ {
				if currMarble.Prev() == nil {
					currMarble = l.Back()
				} else {
					currMarble = currMarble.Prev()
				}
			}
			var nextItem *list.Element
			if currMarble.Next() == nil {
				nextItem = l.Front()
			} else {
				nextItem = currMarble.Next()
			}
			playerScores[currPlayer] += currMarble.Value.(int)
			l.Remove(currMarble)
			currMarble = nextItem

		} else {
			if currMarble.Next() == nil {
				currMarble = l.Front()
			} else {
				currMarble = currMarble.Next()
			}

			currMarble = l.InsertAfter(marble, currMarble)
		}

		currPlayer++

		if currPlayer > players {
			currPlayer = 1
		}

	}

	maxScore := -1

	for _, v := range playerScores {
		if maxScore == -1 || v > maxScore {
			maxScore = v
		}
	}

	return maxScore
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	players, marble := getData(Entry.PuzzleInput())
	return fmt.Sprintf("%v", playGame(players, marble))
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	players, marble := getData(Entry.PuzzleInput())
	return fmt.Sprintf("%v", playGame(players, marble*100))
}
