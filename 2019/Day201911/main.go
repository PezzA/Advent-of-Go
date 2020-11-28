package Day201911

import (
	"fmt"

	"github.com/pezza/advent-of-code/2019/intcode"
	"github.com/pezza/advent-of-code/common"
)

type robot struct {
	pos       common.Point
	direction int
}

func (r *robot) moveIns(ins int) {

	if ins == 0 {
		r.direction -= 90
	} else {
		r.direction += 90
	}

	if r.direction == -90 {
		r.direction = 270
	}

	if r.direction == 360 {
		r.direction = 0
	}

	switch r.direction {
	case 0:
		r.pos = r.pos.Add(common.Cardinal4[0])
	case 90:
		r.pos = r.pos.Add(common.Cardinal4[1])
	case 180:
		r.pos = r.pos.Add(common.Cardinal4[2])
	case 270:
		r.pos = r.pos.Add(common.Cardinal4[3])
	}
}

func runRobot(inputProg string, startSignal int64) map[common.Point]bool {
	input := make(chan int64, 1)
	output := make(chan int64, 0)

	vm := intcode.New(inputProg)

	go func() {
		vm.RunProgram(nil, nil, input, output, nil)
		close(output)
	}()

	hull := make(map[common.Point]bool, 0)
	paintIns := true

	rob := robot{
		common.Point{X: 0, Y: 0},
		0,
	}

	input <- startSignal

	for val := range output {

		if paintIns {

			hull[rob.pos] = val == 1
		} else {

			rob.moveIns(int(val))

			if val, ok := hull[rob.pos]; ok {
				sig := int64(0)
				if val {
					sig = 1
				}

				input <- sig
			} else {
				input <- 0
			}

		}
		paintIns = !paintIns
	}

	return hull
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	hull := runRobot(inputData, 0)
	return fmt.Sprintf("%v", len(hull))
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	hull := runRobot(inputData, 1)

	paintJob := ""
	minX, minY, maxX, maxY := -1, -1, 0, 0
	for k := range hull {
		if minX == -1 || k.X < minX {
			minX = k.X
		}

		if minY == -1 || k.Y < minY {
			minY = k.Y
		}

		if k.X > maxX {
			maxX = k.X
		}

		if k.Y > maxY {
			maxY = k.Y
		}
	}

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if val, ok := hull[common.Point{X: x, Y: y}]; ok {
				if val {
					paintJob += "#"
				} else {
					paintJob += " "
				}
			} else {
				paintJob += " "
			}
		}
		paintJob += "\n"
	}

	return fmt.Sprintf("%v", paintJob)
}
