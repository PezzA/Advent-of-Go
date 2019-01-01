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

// Exported Types (used in github.com/pezza/advent-of-wasm)
type Position struct {
	PosType int
	Mob
}

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
