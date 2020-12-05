package Day202005

import (
	"fmt"
	"strings"
)

type boardingPass struct {
	rowSpec  string
	seatSpec string
}

func getdata(input string) []boardingPass {
	data := make([]boardingPass, 0)

	for _, line := range strings.Split(input, "\n") {

		data = append(data, boardingPass{
			rowSpec:  line[:7],
			seatSpec: line[7:],
		})
	}

	return data
}

func getHalf(getFront bool, min int, max int) (int, int) {
	half := ((max - min) + 1) / 2

	if getFront {
		return min, (min + half) - 1
	}

	return min + half, max
}

func (bp boardingPass) getRow() int {

	min, max := 0, 127

	for _, c := range bp.rowSpec {
		if string(c) == "F" {
			min, max = getHalf(true, min, max)
		} else {
			min, max = getHalf(false, min, max)
		}
	}

	return min
}

func (bp boardingPass) getSeat() int {

	min, max := 0, 7

	for _, c := range bp.seatSpec {
		if string(c) == "L" {
			min, max = getHalf(true, min, max)
		} else {
			min, max = getHalf(false, min, max)
		}
	}

	return min
}

func (bp boardingPass) getSeatId() int {
	return (bp.getRow() * 8) + bp.getSeat()
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	highest := 0

	for _, bp := range getdata(Entry.PuzzleInput()) {
		if bp.getSeatId() > highest {
			highest = bp.getSeatId()
		}
	}
	return fmt.Sprintf("%v", highest)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	plan := make(map[int]bool, 0)
	for _, bp := range getdata(Entry.PuzzleInput()) {
		plan[bp.getSeatId()] = true
	}

	missing := 0
	for start := 8; start < (127*8)+8; start++ {
		if _, ok := plan[start]; !ok {
			missing = start
			break
		}
	}
	return fmt.Sprintf("%v", missing)
}
