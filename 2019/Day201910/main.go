package Day201910

import (
	"fmt"
	"strings"

	"github.com/pezza/advent-of-code/common"
)

func getData(input string) [][]bool {
	belt := [][]bool{}

	lines := strings.Split(input, "\n")

	for _, y := range lines {
		line := []bool{}
		for _, x := range y {
			line = append(line, x == 35)
		}
		belt = append(belt, line)
	}
	return belt
}

func isVisible(s common.Point, t common.Point, belt [][]bool) bool {
	v := common.Point{
		X: t.X - s.X,
		Y: t.Y - s.Y,
	}

	// if either x or y is 1, nothing can be in the way
	if common.Abs(v.X) == 1 || common.Abs(v.Y) == 1 {
		return true
	}

	step := getStep(s, t)

	if step == v {
		return true
	}

	check := s
	// there is a lot of trust in this loop
	for {
		check = check.Add(step)

		// got to the target, can be seen
		if check == t {
			return true
		}

		if belt[check.Y][check.X] {
			// we are blocked
			return false
		}
	}
}

func getVisibleCount(t common.Point, belt [][]bool) int {
	vis := 0
	for y := 0; y < len(belt); y++ {
		for x := 0; x < len(belt[y]); x++ {
			if x == t.X && y == t.Y || !belt[y][x] {
				continue
			}

			if isVisible(t, common.Point{X: x, Y: y}, belt) {
				vis++
			}
		}
	}

	return vis
}

func getStep(s common.Point, t common.Point) common.Point {
	// get the vector between the 2 stars
	v := common.Point{
		X: t.X - s.X,
		Y: t.Y - s.Y,
	}

	if v.X == v.Y {
		return common.Point{X: v.X / common.Abs(v.X), Y: v.Y / common.Abs(v.Y)}
	}

	if v.X == 0 {
		return common.Point{X: 0, Y: v.Y / common.Abs(v.Y)}
	}

	if v.Y == 0 {
		return common.Point{X: v.X / common.Abs(v.X), Y: 0}
	}

	if common.Abs(v.X) == 1 || common.Abs(v.Y) == 1 {
		return v
	}

	clf := getCommonLowestFactor(v.X, v.Y)
	if clf == 0 {
		return v
	}

	return common.Point{X: v.X / clf, Y: v.Y / clf}
}

func getCommonLowestFactor(a, b int) int {
	for i := 2; i <= common.Abs(a) && i <= common.Abs(b); i++ {
		if a%i == 0 && b%i == 0 {
			return i
		}
	}

	return 0
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	return fmt.Sprintf("%v", " -- Not Yet Implemented --")
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	return fmt.Sprintf("%v", " -- Not Yet Implemented --")
}
