package Day201814

import (
	"fmt"
	"strconv"
)

var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string) {
	return 2018, 14, "Chocolate Charts"
}

func getData(input string) int {
	ret, _ := strconv.Atoi(input)
	return ret
}

func getRecipes(elf1 int, elf2 int, scores []int) (int, int, []int) {
	sum := scores[elf1] + scores[elf2]

	for _, char := range strconv.Itoa(sum) {
		newScore, _ := strconv.Atoi(string(char))
		scores = append(scores, newScore)
	}

	elf1Pos := (elf1 + scores[elf1] + 1) % len(scores)
	elf2Pos := (elf2 + scores[elf2] + 1) % len(scores)

	return elf1Pos, elf2Pos, scores
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	return fmt.Sprintf(" -- Not Yet Implemented --")
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	return fmt.Sprintf(" -- Not Yet Implemented --")
}
