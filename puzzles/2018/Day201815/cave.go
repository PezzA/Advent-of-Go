package Day201815

import (
	"fmt"

	common "github.com/pezza/advent-of-code/puzzles/Common"
)

type Cave [][]Position

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
