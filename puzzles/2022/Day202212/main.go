package Day202212

import (
	"fmt"
	"strings"

	"github.com/pezza/advent-of-code/puzzles/common"
)

type vista map[common.Point]rune

func getData(input string) (vista, common.Point, common.Point) {
	lines := strings.Split(input, "\n")

	v := make(vista, 0)
	start := common.Point{0, 0}
	end := common.Point{0, 0}

	for y, line := range lines {
		for x, char := range line {
			if char == 83 {
				start = common.Point{x, y}
				v[start] = 83
			}
			if char == 69 {
				end = common.Point{x, y}
			}
			v[common.Point{x, y}] = char
		}
	}

	return v, start, end
}

type command struct {
	value int
	loc   common.Point
}

func NewCommand(v int, loc common.Point) command {
	return command{v, loc}
}

func startWalking(v vista, start, end common.Point) {
	distances := make(map[common.Point]int, 0)

	queue := []command{NewCommand(0, start)}

	for len(queue) > 0 {

	}

}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	return fmt.Sprintf("%v", " -- Not Yet Implemented --")
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	return fmt.Sprintf("%v", " -- Not Yet Implemented --")
}
