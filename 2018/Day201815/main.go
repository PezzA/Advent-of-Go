package Day201815

import (
	"fmt"
	"strings"

	"github.com/pezza/advent-of-code/Common"
)

var Entry dayEntry

type dayEntry bool

type distanceMap map[common.Point]int

func (td dayEntry) Describe() (int, int, string) {
	return 2018, 15, "Getting the boilerplate in place"
}

type position struct {
	posType int
	mob
}

type cave [][]position

type mob struct {
	id      int
	faction string
	common.Point
	hp  int
	atk int
}

const startingHp = 200
const startingAtk = 3

func getData(input string) cave {
	cave := make(cave, 0)

	for y, line := range strings.Split(input, "\n") {
		cave = append(cave, make([]position, len(line)))
		for x, char := range line {
			switch char {
			case 35: //#
				cave[y][x] = position{1, mob{}}
			case 69, 71: //E, G
				cave[y][x] = position{2, mob{len(cave.getAllMobs()), string(char), common.Point{X: x, Y: y}, startingHp, startingAtk}}
			}
		}
	}

	return cave
}

func (c cave) runRound() (cave, bool, bool) {
	killList := make([]mob, 0)
	victorEmerged, fullRound := false, false

	all := c.getAllMobs()
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

			if c[tp.Y][tp.X].posType == 2 && c[tp.Y][tp.X].faction != m.faction {
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

				if c[next.Y][next.X].posType == 0 {
					c[m.Y][m.X] = position{0, mob{}}
					c[next.Y][next.X] = position{2, mob{m.id, m.faction, common.Point{X: next.X, Y: next.Y}, m.hp, m.atk}}
					m.Point = next
				}
			}
		}

		enemyMap = make(distanceMap)
		// check again? do we have an adjacent enemy?
		for _, dir := range common.Cardinal4 {
			tp := m.Point.Add(dir)

			if c[tp.Y][tp.X].posType == 2 && c[tp.Y][tp.X].faction != m.faction {
				enemyMap[tp] = 0
			}
		}

		if len(enemyMap) > 0 {
			// if we do, attack!
			minHps := c.getMinHp(enemyMap)
			target := minHps.getTop()

			c[target.Y][target.X].hp -= m.atk

			if c[target.Y][target.X].hp <= 0 {
				// mob killed
				faction := c[target.Y][target.X].faction
				killList = append(killList, c[target.Y][target.X].mob)
				c[target.Y][target.X].posType = 0
				c[target.Y][target.X].mob = mob{}

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

func (c cave) draw(d distanceMap) {
	for y := range c {
		lineMobs := make([]mob, 0)
		for x := range c[y] {

			switch c[y][x].posType {
			case 0:
				if d != nil && d[common.Point{X: x, Y: y}] != 0 {
					fmt.Printf("%v", d[common.Point{X: x, Y: y}])
				} else {
					fmt.Print(".")
				}

			case 1:
				fmt.Print("#")
			case 2:
				lineMobs = append(lineMobs, c[y][x].mob)
				fmt.Print(c[y][x].id)
			}
		}

		sep := "  "
		for _, m := range lineMobs {
			fmt.Printf("%v%v-%v[%v,%v](%v)", sep, m.faction, m.id, m.X, m.Y, m.hp)
			sep = ", "
		}
		fmt.Print("\n")
	}
}

func (c cave) getAllMobs() []mob {
	fl := make([]mob, 0)
	for y := range c {
		for x := range c[y] {
			if c[y][x].posType == 2 {
				fl = append(fl, c[y][x].mob)
			}
		}
	}
	return fl
}

func (c cave) getFaction(i string) []mob {
	fl := make([]mob, 0)

	for y := range c {
		for x := range c[y] {
			if c[y][x].posType == 2 && c[y][x].faction == i {
				fl = append(fl, c[y][x].mob)
			}
		}
	}
	return fl
}

func (c cave) getDistanceMap(start common.Point) distanceMap {

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

func (c cave) minimumSteps(start common.Point, step int, dMap distanceMap) distanceMap {
	if dMap == nil {
		dMap = make(distanceMap)
	}
	for _, dir := range common.Cardinal4 {
		tp := start.Add(dir)
		if c[tp.Y][tp.X].posType == 0 {
			if dMap[tp] == 0 || step < dMap[tp] {
				dMap[tp] = step
				c.minimumSteps(tp, step+1, dMap)
			}
		}
	}
	return dMap
}

func (c cave) move(m mob, p common.Point) cave {
	c[m.Y][m.X].posType = 0
	c[m.Y][m.X].mob = mob{}
	c[p.Y][p.X].posType = 2
	c[p.Y][p.X].mob = m
	c[p.Y][p.X].mob.Point.X = p.X
	c[p.Y][p.X].mob.Point.Y = p.Y
	return c
}

func (c cave) getMinHp(dm distanceMap) distanceMap {
	minHp := -1
	for k := range dm {
		if minHp == -1 || c[k.Y][k.X].hp < minHp {
			minHp = c[k.Y][k.X].hp
		}
	}

	minimums := make(distanceMap)

	for k := range dm {
		if c[k.Y][k.X].hp == minHp {
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
			remainingHp += m.hp
		}
	} else {
		finals := cave.getFaction("G")

		for _, m := range finals {
			remainingHp += m.hp
		}
	}

	return round * remainingHp
}

func (c cave) getTarget(m mob) (common.Point, bool) {
	targetList := make(distanceMap)

	dMap := c.minimumSteps(m.Point, 1, nil)

	faction := "G"
	if m.faction == "G" {
		faction = "E"
	}

	for _, m := range c.getFaction(faction) {
		for _, dir := range common.Cardinal4 {
			tp := m.Point.Add(dir)
			if c[tp.Y][tp.X].posType == 0 && dMap[tp] != 0 {
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
	return fmt.Sprintf(" -- Not Yet Implemented --")
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	return fmt.Sprintf(" -- Not Yet Implemented --")
}
