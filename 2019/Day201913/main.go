package Day201913

import (
	"fmt"
	"time"

	"github.com/pezza/advent-of-code/2019/intcode"
	"github.com/pezza/advent-of-code/common"
)

func runGame(input string, freePlay bool) map[common.Point]int64 {

	in := make(chan int64, 0)
	out := make(chan int64, 0)

	go func() {
		vm := intcode.New(input)
		if freePlay {
			vm.SetMemory(0, 2)
		}
		vm.RunProgram(nil, nil, in, out)
		close(out)
	}()

	screen := make(map[common.Point]int64, 0)
	x, y := 0, 0
	step := 1
	score := 0

	for val := range out {
		switch step {
		case 1:
			x = int(val)
			step = 2
		case 2:
			y = int(val)
			step = 3
		case 3:
			fmt.Println(x, y)
			if x == -1 && y == 0 {
				score = int(val)
				fmt.Println(score)
				printScreen(screen)
				time.Sleep(time.Second)
				in <- int64(-1)
			}

			pos := common.Point{int(x), int(y)}
			screen[pos] = val
			step = 1
		}
	}
	return screen
}

func printScreen(screen map[common.Point]int64) {
	maxX, maxY := 0, 0

	for k, _ := range screen {
		if k.X > maxX {
			maxX = k.X
		}
		if k.Y > maxY {
			maxY = k.Y
		}
	}

	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			switch screen[common.Point{x, y}] {
			case 0:
				fmt.Print(" ")
			case 1:
				fmt.Print("#")
			case 2:
				fmt.Print("-")
			case 4:
				fmt.Print("*")
			case 3:
				fmt.Print("=")
			}
		}
		fmt.Print("\n")
	}
}
func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	return fmt.Sprintf("%v", " -- Not Yet Implemented --")
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	return fmt.Sprintf("%v", " -- Not Yet Implemented --")
}
