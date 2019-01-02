package Day201915

import (
	"fmt"
	"strings"

	"github.com/pezza/advent-of-code/puzzles/Common"
)

// Puzzle BoilerPlate //
var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string) {
	return 2019, 15, "Getting the boilerplate in place"
}

// Exported Types (used in github.com/pezza/advent-of-wasm)
type position struct {
	PosType int
	mob
}

type mob struct {
	Id      int
	Faction string
	common.Point
	Hp  int
	Atk int
}

type cave [][]position
type distanceMap [][]int

const startingHp = 200
const startingAtk = 3

func getData(input string) cave {
	cave := make(cave, 0)

	mobCount := 0
	for y, line := range strings.Split(input, "\n") {
		cave = append(cave, make([]position, len(line)))
		for x, char := range line {
			switch char {
			case 35: //#
				cave[y][x] = position{1, mob{}}
			case 69, 71: //E, G
				cave[y][x] = position{2, mob{
					mobCount,
					string(char),
					common.Point{X: x, Y: y}, startingHp, startingAtk}}
				mobCount++
			}
		}
	}

	return cave
}

func (c cave) getBlankDistanceMap() distanceMap {
	dm := make(distanceMap, len(c))
	for index := range dm {
		dm[index] = make([]int, len(c[index]))
	}
	return dm
}

func (c cave) getDistanceMap(start common.Point) distanceMap {
	dm := c.getBlankDistanceMap()
	dm[start.Y][start.X] = -1
	return c.minimumSteps(start, 1, dm)
}

func (c cave) minimumSteps(start common.Point, step int, dMap distanceMap) distanceMap {
	if dMap == nil {
		panic("No distance map passed")
	}
	for _, dir := range common.Cardinal4 {
		tp := start.Add(dir)
		if c[tp.Y][tp.X].PosType == 0 {
			if dMap[tp.Y][tp.X] == 0 || step < dMap[tp.Y][tp.X] {
				dMap[tp.Y][tp.X] = step
				c.minimumSteps(tp, step+1, dMap)
			}
		}
	}
	return dMap
}

func (dm distanceMap) draw() {
	fmt.Println()
	for y := range dm {
		for x := range dm[y] {
			if dm[y][x] == -1 {
				fmt.Print(".")
			} else {
				fmt.Print(dm[y][x])
			}
		}
		fmt.Print("\n")
	}
}

func (c cave) draw(dm distanceMap) {
	fmt.Println()
	for y := range c {
		lineMobs := make([]mob, 0)
		for x := range c[y] {
			switch c[y][x].PosType {
			case 0:
				if dm != nil {
					if dm[y][x] == -1 {
						fmt.Print("0")
					} else {
						fmt.Print(dm[y][x])
					}
				} else {
					fmt.Print(".")
				}
			case 1:
				fmt.Print("#")
			case 2:
				lineMobs = append(lineMobs, c[y][x].mob)
				fmt.Print(c[y][x].Faction)
			}
		}

		sep := "  "
		for _, m := range lineMobs {
			fmt.Printf("%v%v-%v[%v,%v](%v)", sep, m.Faction, m.Id, m.X, m.Y, m.Hp)
			sep = ", "
		}
		fmt.Print("\n")
	}
}

func (c cave) getAllMobs() []mob {
	fl := make([]mob, 0)
	for y := range c {
		for x := range c[y] {
			if c[y][x].PosType == 2 {
				fl = append(fl, c[y][x].mob)
			}
		}
	}
	return fl
}

func getFirstInReadingOrder(pl []common.Point, dm distanceMap) common.Point {
	minPoint := -1
	var targetLocation common.Point

	for index := range pl {
		distanceVal := dm[pl[index].Y][pl[index].X]

		// if its the same value, use it if it is higher up in reading order
		if minPoint == distanceVal {
			if pl[index].Y < targetLocation.Y {
				minPoint = distanceVal
				targetLocation = pl[index]
			} else if pl[index].Y == targetLocation.Y {
				if pl[index].X < targetLocation.X {
					minPoint = distanceVal
					targetLocation = pl[index]
				}
			}
		}

		// if it's closer, user it
		if minPoint == -1 || distanceVal < minPoint {
			minPoint = distanceVal
			targetLocation = pl[index]
		}
	}

	return targetLocation
}

func (c cave) getTarget(m mob) (common.Point, bool) {
	targetList := make([]common.Point, 0)

	dMap := c.getDistanceMap(m.Point)

	faction := "G"
	if m.Faction == "G" {
		faction = "E"
	}

	for _, fm := range c.getFaction(faction) {
		for _, dir := range common.Cardinal4 {
			tp := fm.Point.Add(dir)
			if c[tp.Y][tp.X].PosType == 0 && dMap[tp.Y][tp.X] > 0 {
				targetList = append(targetList, tp)
			}
		}
	}

	if len(targetList) == 0 {
		return common.Point{}, false
	}

	return getFirstInReadingOrder(targetList, dMap), true
}

func (c cave) getFaction(i string) []mob {
	fl := make([]mob, 0)

	for y := range c {
		for x := range c[y] {
			if c[y][x].PosType == 2 && c[y][x].Faction == i {
				fl = append(fl, c[y][x].mob)
			}
		}
	}
	return fl
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	return fmt.Sprintf(" -- Not Yet Implemented --")
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	return fmt.Sprintf(" -- Not Yet Implemented --")
}
