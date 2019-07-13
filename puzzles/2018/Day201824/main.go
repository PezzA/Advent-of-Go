package Day201824

import (
	"fmt"
)

var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string) {
	return 2018, 24, "Immune System Simulator 20XX"
}

type group struct {
	id         int
	units      int
	hp         int
	attack     int
	initative  int
	attackType string
	weaknesses []string
	immunites  []string
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	return fmt.Sprintf(" -- Not Yet Implemented --")
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	return fmt.Sprintf(" -- Not Yet Implemented --")
}
