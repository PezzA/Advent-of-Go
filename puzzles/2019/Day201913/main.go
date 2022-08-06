package Day201913

import (
	"fmt"

	"github.com/pezza/advent-of-code/2019/intcode"
	"github.com/pezza/advent-of-code/common"
)

func runGame(input string, freePlay bool) int {

	in := make(chan int64, 0)
	out := make(chan int64, 0)
	request := make(chan bool, 0)

	// run the program
	go func() {
		vm := intcode.New(input)
		init := make(map[int64]int64, 0)
		if freePlay {
			init[0] = 2
		}
		vm.RunProgram(init, nil, in, out, request)
		close(out)
		close(request)
	}()

	// handle output
	screen := make(map[common.Point]int64, 0)
	x, y := 0, 0
	step := 1
	score := 0
	ball, paddle := common.Point{X: 0, Y: 0}, common.Point{X: 0, Y: 0}
	inputDone, outputDone := false, false

	for {
		select {
		case val, ok := <-out:
			if !ok {
				outputDone = true
				break
			}
			switch step {
			case 1:
				x = int(val)
				step = 2
			case 2:
				y = int(val)
				step = 3
			case 3:
				if x == -1 && y == 0 {
					score = int(val)
				} else {
					pos := common.Point{X: int(x), Y: int(y)}
					screen[pos] = val

					if val == 3 {
						paddle = pos
					} else if val == 4 {
						ball = pos
					}
				}

				step = 1
			}
		case _, ok := <-request:
			if !ok {
				inputDone = true
				break
			}
			if ball.X < paddle.X {
				in <- -1
			} else if ball.X > paddle.X {
				in <- 1
			} else {
				in <- 0
			}
		}

		if inputDone && outputDone {
			break
		}
	}

	return score
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
