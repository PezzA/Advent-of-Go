package Day201910

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/pezza/advent-of-code/common"
)

type asteriodAngle struct {
	angle float64
	dist  int
	rot   int
	pos   common.Point
}

// vector2D are taken from my own package https://github.com/PezzA/maths
type vector2D struct {
	X float64
	Y float64
}

type point2D common.Point

func (p point2D) getVectorToPoint(p2 point2D) vector2D {
	return vector2D{
		X: float64(p2.X - p.X),
		Y: float64(p2.Y - p.Y),
	}
}

//Normalised returns the normalised unit versin of a Vector2D
func (v vector2D) normalised() vector2D {
	return v.divideBy(v.getLength())
}

func (v vector2D) divideBy(s float64) vector2D {
	return vector2D{
		X: v.X / s,
		Y: v.Y / s,
	}
}

func (v vector2D) getLength() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v vector2D) dotProduct(v2 vector2D) float64 {
	return v.X*v2.X + v.Y*v2.Y
}

func (v vector2D) crossProduct(v2 vector2D) float64 {
	return v.X*v2.X + v.Y*v2.Y
}

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

func copyBelt(input [][]bool) [][]bool {
	o := make([][]bool, len(input))

	for i := range o {
		o[i] = make([]bool, len(input[i]))
	}

	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			o[y][x] = input[y][x]
		}
	}
	return o
}

func countAsteriods(i [][]bool) int {
	count := 0
	for y := 0; y < len(i); y++ {
		for x := 0; x < len(i[y]); x++ {
			if i[y][x] {
				count++
			}
		}
	}
	return count
}

func getBestAsteriod(belt [][]bool) (int, common.Point) {
	count, point := 0, common.Point{X: 0, Y: 0}
	for y := 0; y < len(belt); y++ {
		for x := 0; x < len(belt[y]); x++ {
			if !belt[y][x] {
				continue
			}

			astPos := common.Point{X: x, Y: y}

			newMap := getVisibleMap(astPos, belt)
			astVis := countAsteriods(newMap)
			//printMap(newMap, astPos)
			//fmt.Println(astVis)

			if astVis > count {
				count = astVis
				point = astPos
			}
		}
	}

	return count, point
}

func getVisibleMap(s common.Point, belt [][]bool) [][]bool {
	newBelt := copyBelt(belt)

	newBelt[s.Y][s.X] = false
	for y := 0; y < len(newBelt); y++ {
		for x := 0; x < len(newBelt); x++ {
			if !newBelt[y][x] || (x == s.X && y == s.Y) {
				continue
			}

			v := common.Point{
				X: x - s.X,
				Y: y - s.Y,
			}

			vp := getStep(v)

			start := common.Point{X: x, Y: y}
			for {
				start = start.Add(vp)

				if start.X < 0 || start.Y < 0 || start.X >= len(belt) || start.Y >= len(belt) {
					break
				}

				newBelt[start.Y][start.X] = false
			}
		}
	}

	return newBelt
}

func printMap(belt [][]bool, h common.Point) {
	for y := 0; y < len(belt); y++ {
		for x := 0; x < len(belt[y]); x++ {
			if x == h.X && y == h.Y {
				fmt.Print("@")
				continue
			}
			if !belt[y][x] {
				fmt.Print(".")
				continue
			}

			fmt.Print("#")

		}
		fmt.Print("\n")
	}
}

func getStep(v common.Point) common.Point {

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
	field := getData(inputData)
	count, _ := getBestAsteriod(field)
	return fmt.Sprintf("%v", count)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	field := getData(inputData)
	_, pos := getBestAsteriod(field)

	up := vector2D{0, 1} // swapped for array plotting
	right := vector2D{1, 0}

	pos2d := point2D{pos.X, pos.Y}

	astList := []asteriodAngle{}

	for y := 0; y < len(field); y++ {
		for x := 0; x < len(field[y]); x++ {
			if !field[y][x] || (y == pos2d.Y && x == pos2d.X) {
				continue
			}

			vec := pos2d.getVectorToPoint(point2D{x, y}).
				normalised()

			upP := vec.dotProduct(up)
			rightP := vec.dotProduct(right)

			angle := math.Atan2(upP, rightP)*(180/math.Pi) + 90

			if angle < 0 {
				angle = 360 - math.Abs(angle)
			}

			astList = append(astList, asteriodAngle{
				angle,
				common.GetMDistance(pos, common.Point{X: x, Y: y}),
				0,
				common.Point{X: x, Y: y},
			})
		}
	}

	// sort by angle and distance
	sort.Slice(astList, func(i, j int) bool {
		if astList[i].angle == astList[j].angle {
			return astList[i].dist < astList[j].dist
		}
		return astList[i].angle < astList[j].angle
	})

	// go thought the sorted list and where the angle is the same increment rotation
	prev := float64(-1)
	rot := 0
	for index := range astList {
		if astList[index].angle == prev {
			rot++
		} else {
			rot = 0
		}

		astList[index].rot = rot
		prev = astList[index].angle
	}

	// finally sort by rotation and angle
	sort.Slice(astList, func(i, j int) bool {
		if astList[i].rot == astList[j].rot {
			return astList[i].angle < astList[j].angle
		}
		return astList[i].rot < astList[j].rot
	})

	return fmt.Sprintf("%v", astList[199].pos.X*100+astList[199].pos.Y)
}
