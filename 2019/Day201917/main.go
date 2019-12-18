package Day201917

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/pezza/advent-of-code/common"

	"github.com/pezza/advent-of-code/2019/intcode"
)

type robot struct {
	direction int
	pos       common.Point
}

type view [][]byte

const platform = byte(35)

func getCameraView(program string) view {
	out := make(chan int64, 0)

	go func() {
		vm := intcode.New(program)
		vm.RunProgram(nil, nil, nil, out, nil)
		close(out)
	}()

	var output view
	var line []byte

	for val := range out {
		bVal := byte(val)

		switch bVal {
		case 10:
			output = append(output, line)
			line = []byte{}
		default:
			line = append(line, bVal)
		}
	}
	return output[:len(output)-1]
}

func (v view) getRobot() (robot, error) {
	for y := 0; y < len(v); y++ {
		for x := 0; x < len(v[y]); x++ {
			curr := v[y][x]

			switch curr {
			case 94:
				return robot{0, common.Point{X: x, Y: y}}, nil
			case 118:
				return robot{180, common.Point{X: x, Y: y}}, nil
			case 60:
				return robot{270, common.Point{X: x, Y: y}}, nil
			case 62:
				return robot{90, common.Point{X: x, Y: y}}, nil
			}
		}
	}

	return robot{}, errors.New("could not view robot in viewport")
}

func (v view) testDirection(start common.Point, direction int) (bool, error) {
	switch direction {
	case 0:
		if start.Y == 0 {
			return false, nil
		}

		return v[start.Y-1][start.X] == platform, nil
	case 180:
		if start.Y == len(v)-1 {
			return false, nil
		}
		return v[start.Y+1][start.X] == platform, nil
	case 90:
		if start.X == len(v[start.Y])-1 {
			return false, nil
		}
		return v[start.Y][start.X+1] == platform, nil
	case 270:
		if start.X == 0 {
			return false, nil
		}
		return v[start.Y][start.X-1] == platform, nil
	}

	return false, errors.New("unknown direction")
}

func (r *robot) move() {
	switch r.direction {
	case 0:
		r.pos = common.Point{X: r.pos.X, Y: r.pos.Y - 1}
	case 180:
		r.pos = common.Point{X: r.pos.X, Y: r.pos.Y + 1}
	case 90:
		r.pos = common.Point{X: r.pos.X + 1, Y: r.pos.Y}
	case 270:
		r.pos = common.Point{X: r.pos.X - 1, Y: r.pos.Y}
	}
}

func getDir(old, new int) string {
	if (old == 0 && new == 90) || (old == 90 && new == 180) || (old == 180 && new == 270) || (old == 270 && new == 0) {
		return "R"
	}
	return "L"
}

func (v view) getPath(rob robot) (map[common.Point]int, []string) {
	path := make(map[common.Point]int, 0)
	ins := []string{}

	mov := 0

	for {

		// log current position
		if _, ok := path[rob.pos]; ok {
			path[rob.pos]++
		} else {
			path[rob.pos] = 1
		}

		// test direction
		if ok, _ := v.testDirection(rob.pos, rob.direction); ok {
			rob.move()
			mov++
			continue
		}

		var tests []int
		switch rob.direction {
		case 0:
			tests = []int{270, 90}
		case 180:
			tests = []int{90, 270}
		case 90:
			tests = []int{0, 180}
		case 270:
			tests = []int{180, 0}
		}

		dirFound := false

		for _, testDir := range tests {
			if ok, _ := v.testDirection(rob.pos, testDir); ok {
				if mov != 0 {
					ins = append(ins, strconv.Itoa(mov))
				}

				ins = append(ins, getDir(rob.direction, testDir))
				rob.direction = testDir
				rob.move()
				mov = 1
				dirFound = true
				break
			}
		}

		if !dirFound {
			ins = append(ins, strconv.Itoa(mov))
			break
		}
	}

	return path, ins
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	data := getCameraView(Entry.PuzzleInput())
	robot, _ := data.getRobot()
	path, ins := data.getPath(robot)

	totAP := 0
	for k, v := range path {
		if v > 1 {
			totAP += k.X * k.Y
		}
	}

	fmt.Println(ins)
	return fmt.Sprintf("%v", totAP)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	return fmt.Sprintf("%v", " -- Not Yet Implemented --")
}
