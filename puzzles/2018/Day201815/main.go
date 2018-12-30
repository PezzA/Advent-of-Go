package Day201815

import (
	"fmt"
	"strings"

	"github.com/pezza/advent-of-code/puzzles/Common"
)

// Puzzle BoilerPlate //
var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string) {
	return 2018, 15, "Getting the boilerplate in place"
}

type distanceMap map[common.Point]int

// Exported Types (used in github.com/pezza/advent-of-wasm)
type Position struct {
	PosType int
	Mob
}

type Cave [][]Position

type Mob struct {
	Id      int
	Faction string
	common.Point
	Hp  int
	Atk int
}

const startingHp = 200
const startingAtk = 3

func getData(input string) Cave {
	cave := make(Cave, 0)

	for y, line := range strings.Split(input, "\n") {
		cave = append(cave, make([]Position, len(line)))
		for x, char := range line {
			switch char {
			case 35: //#
				cave[y][x] = Position{1, Mob{}}
			case 69, 71: //E, G
				cave[y][x] = Position{2, Mob{len(cave.GetAllMobs()), string(char), common.Point{X: x, Y: y}, startingHp, startingAtk}}
			}
		}
	}

	return cave
}

func (c Cave) runRound() (Cave, bool, bool) {
	killList := make([]Mob, 0)
	victorEmerged, fullRound := false, false

	all := c.GetAllMobs()
	for roundIndex, m := range all {

		alreadyDiedThisRound := false
		for _, kl := range killList {
			if m.X == kl.X && m.Y == kl.Y {
				alreadyDiedThisRound = true
				continue
			}
		}

		if alreadyDiedThisRound {
			continue
		}

		enemyMap := make(distanceMap)
		// do we have an adjacent enemy?
		for _, dir := range common.Cardinal4 {
			tp := m.Point.Add(dir)

			if c[tp.Y][tp.X].PosType == 2 && c[tp.Y][tp.X].Faction != m.Faction {
				enemyMap[tp] = 0
			}
		}

		// no? move towards the nearest if possible
		if len(enemyMap) == 0 {
			target, acquired := c.getTarget(m)

			if acquired {
				dm := c.getDistanceMap(target)
				adj := dm.filterAdjacent(m.Point)
				next := adj.getTop()

				if c[next.Y][next.X].PosType == 0 {
					c[m.Y][m.X] = Position{0, Mob{}}
					c[next.Y][next.X] = Position{2, Mob{m.Id, m.Faction, common.Point{X: next.X, Y: next.Y}, m.Hp, m.Atk}}
					m.Point = next
				}
			}
		}

		enemyMap = make(distanceMap)
		// check again? do we have an adjacent enemy?
		for _, dir := range common.Cardinal4 {
			tp := m.Point.Add(dir)

			if c[tp.Y][tp.X].PosType == 2 && c[tp.Y][tp.X].Faction != m.Faction {
				enemyMap[tp] = 0
			}
		}

		if len(enemyMap) > 0 {
			// if we do, attack!
			minHps := c.getMinHp(enemyMap)
			target := minHps.getTop()

			c[target.Y][target.X].Hp -= m.Atk

			if c[target.Y][target.X].Hp <= 0 {
				// mob killed
				faction := c[target.Y][target.X].Faction
				killList = append(killList, c[target.Y][target.X].Mob)
				c[target.Y][target.X].PosType = 0
				c[target.Y][target.X].Mob = Mob{}

				// check to see if we have an exit condition
				if len(c.getFaction(faction)) == 0 {
					victorEmerged = true
					fullRound = roundIndex == len(all)-1
					break
				}
			}
		}
	}

	return c, victorEmerged, fullRound
}

func (c Cave) draw(d distanceMap) {
	for y := range c {
		lineMobs := make([]Mob, 0)
		for x := range c[y] {

			switch c[y][x].PosType {
			case 0:
				if d != nil && d[common.Point{X: x, Y: y}] != 0 {
					fmt.Printf("%v", d[common.Point{X: x, Y: y}])
				} else {
					fmt.Print(".")
				}

			case 1:
				fmt.Print("#")
			case 2:
				lineMobs = append(lineMobs, c[y][x].Mob)
				fmt.Print(c[y][x].Id)
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

func (c Cave) GetAllMobs() []Mob {
	fl := make([]Mob, 0)
	for y := range c {
		for x := range c[y] {
			if c[y][x].PosType == 2 {
				fl = append(fl, c[y][x].Mob)
			}
		}
	}
	return fl
}

func (c Cave) getFaction(i string) []Mob {
	fl := make([]Mob, 0)

	for y := range c {
		for x := range c[y] {
			if c[y][x].PosType == 2 && c[y][x].Faction == i {
				fl = append(fl, c[y][x].Mob)
			}
		}
	}
	return fl
}

func (c Cave) getDistanceMap(start common.Point) distanceMap {

	dm := make(distanceMap)

	dm[start] = 0
	return c.minimumSteps(start, 1, dm)
}

func (dm distanceMap) filterAdjacent(point common.Point) distanceMap {
	retVal := make(distanceMap)

	for _, dir := range common.Cardinal4 {
		tp := point.Add(dir)

		if _, ok := dm[tp]; ok {
			retVal[tp] = dm[tp]
		}
	}
	return retVal
}

func (c Cave) minimumSteps(start common.Point, step int, dMap distanceMap) distanceMap {
	if dMap == nil {
		dMap = make(distanceMap)
	}
	for _, dir := range common.Cardinal4 {
		tp := start.Add(dir)
		if c[tp.Y][tp.X].PosType == 0 {
			if dMap[tp] == 0 || step < dMap[tp] {
				dMap[tp] = step
				c.minimumSteps(tp, step+1, dMap)
			}
		}
	}
	return dMap
}

func (c Cave) move(m Mob, p common.Point) Cave {
	c[m.Y][m.X].PosType = 0
	c[m.Y][m.X].Mob = Mob{}
	c[p.Y][p.X].PosType = 2
	c[p.Y][p.X].Mob = m
	c[p.Y][p.X].Mob.Point.X = p.X
	c[p.Y][p.X].Mob.Point.Y = p.Y
	return c
}

func (c Cave) getMinHp(dm distanceMap) distanceMap {
	minHp := -1
	for k := range dm {
		if minHp == -1 || c[k.Y][k.X].Hp < minHp {
			minHp = c[k.Y][k.X].Hp
		}
	}

	minimums := make(distanceMap)

	for k := range dm {
		if c[k.Y][k.X].Hp == minHp {
			minimums[k] = 1
		}
	}
	return minimums

}

// given a distancemap, will return the
func (dm distanceMap) getTop() common.Point {
	minPoint := -1
	var targetLocation common.Point

	for k, v := range dm {

		// if its the same value, use it if it is higher up in reading order
		if minPoint == v {
			if k.Y < targetLocation.Y {
				minPoint = v
				targetLocation = k
			} else if k.Y == targetLocation.Y {
				if k.X < targetLocation.X {
					minPoint = v
					targetLocation = k
				}
			}
		}

		// if it's closer, user it
		if minPoint == -1 || v < minPoint {
			minPoint = v
			targetLocation = k
		}
	}

	return targetLocation
}

func runBattle(input string) int {
	cave := getData(input)

	factionRisesVictorious, fullRound := false, false

	round := 1
	for {
		cave, factionRisesVictorious, fullRound = cave.runRound()

		if factionRisesVictorious {
			if !fullRound {
				round--
			}
			break
		}
		round++
	}

	remainingHp := 0
	if len(cave.getFaction("E")) != 0 {
		finals := cave.getFaction("E")

		for _, m := range finals {
			remainingHp += m.Hp
		}
	} else {
		finals := cave.getFaction("G")

		for _, m := range finals {
			remainingHp += m.Hp
		}
	}

	return round * remainingHp
}

func (c Cave) getTarget(m Mob) (common.Point, bool) {
	targetList := make(distanceMap)

	dMap := c.minimumSteps(m.Point, 1, nil)

	faction := "G"
	if m.Faction == "G" {
		faction = "E"
	}

	for _, m := range c.getFaction(faction) {
		for _, dir := range common.Cardinal4 {
			tp := m.Point.Add(dir)
			if c[tp.Y][tp.X].PosType == 0 && dMap[tp] != 0 {
				targetList[tp] = dMap[tp]
			}
		}
	}

	if len(targetList) == 0 {
		return common.Point{}, false
	}

	return targetList.getTop(), true
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {

	val := runBattle(inputData)
	return fmt.Sprintf("%v", val)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	return fmt.Sprintf(" -- Not Yet Implemented --")
}

var cav Cave

func (td dayEntry) Page() int { return 45 }
func (td dayEntry) Start(inputData string) interface{} {

	if inputData == "" {
		inputData = Entry.PuzzleInput()
	}

	cav = getData(inputData)

	return cav

}

func (td dayEntry) Prev() interface{} {
	cav, _, _ := cav.runRound()
	return cav
}

func (td dayEntry) Next() interface{} {
	cav, _, _ := cav.runRound()
	return cav
}
func (td dayEntry) End() interface{} { return "" }
