package Day201919

import (
	"fmt"

	"github.com/pezza/advent-of-code/2019/intcode"
	"github.com/pezza/advent-of-code/common"
)

func scanPoint(input string, x, y int64) int64 {
	out := make(chan int64, 0)
	in := make(chan int64, 0)
	go func() {
		vm := intcode.New(input)
		vm.RunProgram(nil, nil, in, out, nil)
		close(out)
	}()

	in <- x
	in <- y
	close(in)
	return <-out
}

func checkDepth(ox, oy int, area map[common.Point]bool) bool {
	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			if _, ok := area[common.Point{X: x + ox, Y: y + oy}]; !ok {
				return false
			}

		}
	}
	return true
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	return fmt.Sprintf("%v", " -- Not Yet Implemented --")
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	return fmt.Sprintf("%v", " -- Not Yet Implemented --")
}
