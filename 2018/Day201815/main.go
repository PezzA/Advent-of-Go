package Day201815

import (
	"fmt"
	"strings"

	"github.com/pezza/advent-of-code/common"
)

// Puzzle BoilerPlate //
var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string) {
	return 2018, 15, "Beverage Bandits"
}

type posType int

// Exported Types (used in github.com/pezza/advent-of-wasm)
type position struct {
	PosType posType
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

const (
	space posType = iota
	wall
	occupied
)

const (
	untraversed   = -1
	startingPoint = 0
	startingHp    = 200
	baseAttack    = 3
)

func getData(input string, elfAttack int) cave {
	cave := make(cave, 0)

	mobCount := 0
	for y, line := range strings.Split(input, "\n") {
		cave = append(cave, make([]position, len(line)))
		for x, char := range line {
			switch char {
			// #
			case 35:
				cave[y][x] = position{wall, mob{}}
			// E, G
			case 69, 71:
				atk := baseAttack
				if string(char) == "E" {
					atk = elfAttack
				}
				cave[y][x] = position{occupied, mob{
					mobCount,
					string(char),
					common.Point{X: x, Y: y}, startingHp, atk}}
				mobCount++
			}
		}
	}

	return cave
}

func (c cave) getBlankDistanceMap() distanceMap {
	dm := make(distanceMap, 0)
	for index := range c {
		row := make([]int, len(c[index]))
		for index := range row {
			row[index] = untraversed
		}
		dm = append(dm, row)
	}

	return dm
}

func (c cave) getDistanceMap(start common.Point) distanceMap {
	dm := c.getBlankDistanceMap()
	dm[start.Y][start.X] = startingPoint
	return c.minimumSteps(start, 1, dm)
}

func (c cave) minimumSteps(start common.Point, step int, dMap distanceMap) distanceMap {
	if dMap == nil {
		panic("No distance map passed")
	}
	for _, dir := range common.Cardinal4 {
		tp := start.Add(dir)
		if c[tp.Y][tp.X].PosType == space {
			if dMap[tp.Y][tp.X] == untraversed || step < dMap[tp.Y][tp.X] {
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

func (c cave) draw(dm distanceMap, round int) {
	fmt.Println("Round", round)
	for y := range c {
		lineMobs := make([]mob, 0)
		for x := range c[y] {
			switch c[y][x].PosType {
			case space:
				if dm != nil {
					if dm[y][x] == -1 {
						fmt.Print("0")
					} else {
						fmt.Print(dm[y][x])
					}
				} else {
					fmt.Print(".")
				}
			case wall:
				fmt.Print("#")
			case occupied:
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
	fmt.Println()
}

func (c cave) getAllMobs() []mob {
	fl := make([]mob, 0)
	for y := range c {
		for x := range c[y] {
			if c[y][x].PosType == occupied {
				fl = append(fl, c[y][x].mob)
			}
		}
	}
	return fl
}

func contains(m mob, l []mob) bool {
	for _, listMob := range l {
		if m.Id == listMob.Id {
			return true
		}
	}
	return false
}

func (c cave) getAdjacentEnemy(m mob) (bool, mob) {
	// get all the adjacent enemies
	enemies := make([]mob, 0)

	minHp := -1
	for _, dir := range common.Cardinal4ReadingOrder {
		tp := m.Point.Add(dir)

		if c[tp.Y][tp.X].PosType == occupied && c[tp.Y][tp.X].mob.Faction != m.Faction {
			enemies = append(enemies, c[tp.Y][tp.X].mob)
			if minHp == -1 || c[tp.Y][tp.X].Hp < minHp {
				minHp = c[tp.Y][tp.X].Hp
			}
		}
	}

	// no adacent enemies detected
	if len(enemies) == 0 {
		return false, mob{}
	}

	for index := range enemies {
		if enemies[index].Hp == minHp {
			return true, enemies[index]
		}
	}

	return false, mob{}
}

func (c cave) runRound(round int) (bool, cave) {
	killList := make([]mob, 0)

	// loop over each mob
	for _, currMob := range c.getAllMobs() {
		// is the round over?
		testFaction := "E"

		if currMob.Faction == "E" {
			testFaction = "G"
		}

		// end the partial round
		if len(c.getFaction(testFaction)) == 0 {
			return false, c
		}

		// is the mob we are trying to process in the round kill list?
		if contains(currMob, killList) {
			// no need to process, skip to next mob
			continue
		}

		if hasTarget, _ := c.getAdjacentEnemy(currMob); !hasTarget {
			// Move
			if target, canMove := c.getTarget(currMob); canMove {

				nextPosition := c.getNextStep(currMob, target)

				targetMob := c[currMob.Y][currMob.X]
				c[currMob.Y][currMob.X] = position{space, mob{}}

				currMob = mob{
					targetMob.Id,
					targetMob.Faction,
					common.Point{X: nextPosition.X, Y: nextPosition.Y},
					targetMob.Hp,
					targetMob.Atk}

				c[nextPosition.Y][nextPosition.X] = position{occupied, currMob}
			}
		}

		if hasTarget, target := c.getAdjacentEnemy(currMob); hasTarget {
			// Attack
			c[target.Y][target.X].Hp -= currMob.Atk

			// check to see if the mob has been killed
			if c[target.Y][target.X].Hp <= 0 {
				killList = append(killList, c[target.Y][target.X].mob)
				c[target.Y][target.X] = position{space, mob{}}

			}
		}
	}
	// end the full round
	return true, c
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

		// if it's closer, use it
		if minPoint == -1 || distanceVal < minPoint {
			minPoint = distanceVal
			targetLocation = pl[index]
		}
	}

	return targetLocation
}

// getAdjacentPoints will return each adjacent position to the mob that registers
// on the distance map.
func (m mob) getAdjacentPoints(dmap distanceMap) []common.Point {
	points := make([]common.Point, 0)

	for _, p := range common.Cardinal4 {
		tp := m.Point.Add(p)
		if dmap[tp.Y][tp.X] > -1 {
			points = append(points, tp)
		}
	}
	return points
}

func (c cave) getNextStep(m mob, t common.Point) common.Point {
	dMap := c.getDistanceMap(t)

	starts := m.getAdjacentPoints(dMap)

	next := getFirstInReadingOrder(starts, dMap)

	return next
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
			if c[tp.Y][tp.X].PosType == space && dMap[tp.Y][tp.X] >= 0 {
				targetList = append(targetList, tp)
			}
		}
	}

	if len(targetList) == 0 {
		return common.Point{}, false
	}

	return getFirstInReadingOrder(targetList, dMap), true
}

func (c cave) getFaction(faction string) []mob {
	fl := make([]mob, 0)

	for y := range c {
		for x := range c[y] {
			if c[y][x].PosType == occupied && c[y][x].mob.Faction == faction {
				fl = append(fl, c[y][x].mob)
			}
		}
	}
	return fl
}

func (c cave) resolveBattle(debug bool) (int, cave) {
	rounds := 0

	if debug {
		c.draw(nil, rounds)
	}

	for {

		round := -1

		if debug {
			round = rounds
		}

		fullRound, c := c.runRound(round)

		if fullRound {
			rounds++
		}

		if debug {
			c.draw(nil, rounds)
		}

		if len(c.getFaction("E")) == 0 || len(c.getFaction("G")) == 0 {
			break
		}
	}

	return rounds, c
}

func escalateBattle(inputData string) (int, int, cave) {
	elfAttack := baseAttack + 1

	for {
		cave := getData(inputData, elfAttack)
		startingElves := len(cave.getFaction("E"))

		rounds, caveEndState := cave.resolveBattle(false)

		if len(caveEndState.getFaction("E")) == startingElves {
			return elfAttack, rounds, caveEndState
		}

		elfAttack++
	}
}

func (c cave) getRemainingHp() int {
	hp := 0
	for _, mob := range c.getAllMobs() {
		hp += mob.Hp
	}

	return hp
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	cave := getData(inputData, baseAttack)
	rounds, caveEndState := cave.resolveBattle(false)
	return fmt.Sprintf("%v", caveEndState.getRemainingHp()*rounds)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	_, rounds, c := escalateBattle(inputData)
	return fmt.Sprintf("%v", rounds*c.getRemainingHp())
}
