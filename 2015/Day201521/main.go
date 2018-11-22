package Day201521

import (
	"fmt"
	"strconv"
	"strings"
)

// Entry holds wraps the data and runner interfaces for this puzzle
var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string) {
	return 2015, 21, "RPG Simulator 20XX"
}

type character struct {
	hp    int
	def   int
	atk   int
	items []item
}

type item struct {
	name string
	def  int
	atk  int
	cost int
}

func getWeapons() []item {
	return []item{
		item{"Dagger", 0, 4, 8},
		item{"Shortsword", 0, 5, 10},
		item{"Warhammer", 0, 6, 25},
		item{"Longsword", 0, 7, 40},
		item{"Greataxe", 0, 8, 74},
	}
}

func getArmour() []item {
	return []item{
		item{"none", 0, 0, 0},
		item{"Leather", 1, 0, 13},
		item{"Chainmail", 2, 0, 31},
		item{"Splintmail", 3, 0, 53},
		item{"Bandemail", 4, 0, 75},
		item{"Platemail", 5, 0, 102},
	}
}

func getRings() []item {
	return []item{
		item{"none", 0, 0, 0},
		item{"Dmg +1", 0, 1, 25},
		item{"Dmg +2", 0, 2, 50},
		item{"Dmg +3", 0, 3, 100},
		item{"Defense +1", 1, 0, 20},
		item{"Defense +2", 2, 0, 40},
		item{"Defense +3", 3, 0, 80},
	}
}

func getPlayer() character {
	return character{100, 0, 0, make([]item, 0)}
}

func getData(input string) character {
	lines := strings.Split(input, "\n")

	hp, _ := strconv.Atoi(strings.Fields(lines[0])[2])
	atk, _ := strconv.Atoi(strings.Fields(lines[1])[1])
	def, _ := strconv.Atoi(strings.Fields(lines[2])[1])

	return character{hp, def, atk, make([]item, 0)}
}

func resolveCombat(player character, boss character) bool {

	for _, item := range player.items {
		player.atk += item.atk
		player.def += item.def
	}

	playerAttack := player.atk - boss.def
	if playerAttack < 1 {
		playerAttack = 1
	}

	bossAttack := boss.atk - player.def
	if bossAttack < 1 {
		bossAttack = 1
	}

	for {
		boss.hp -= playerAttack
		if boss.hp <= 0 {
			return true
		}

		player.hp -= bossAttack
		if player.hp <= 0 {
			return false
		}
	}
}

func cycleGear(player character, boss character) int {
	itemSets := make([][]item, 0)

	for _, weapon := range getWeapons() {
		for _, armour := range getArmour() {
			for _, ring1 := range getRings() {
				for _, ring2 := range getRings() {
					// can have 2 of the same type of ring (except from none)
					if ring1.name == ring2.name && ring1.name != "none" {
						continue
					}

					player.items = []item{weapon, armour, ring1, ring2}

					if resolveCombat(player, boss) {
						itemSets = append(itemSets, player.items)
					}
				}
			}
		}
	}

	minCost := -1

	for _, iset := range itemSets {
		cost := iset[0].cost + iset[1].cost + iset[2].cost + iset[3].cost

		if cost < minCost || minCost == -1 {
			minCost = cost
		}
	}

	return minCost
}

func mindControl(player character, boss character) int {
	itemSets := make([][]item, 0)

	for _, weapon := range getWeapons() {
		for _, armour := range getArmour() {
			for _, ring1 := range getRings() {
				for _, ring2 := range getRings() {
					// can have 2 of the same type of ring (except from none)
					if ring1.name == ring2.name && ring1.name != "none" {
						continue
					}

					player.items = []item{weapon, armour, ring1, ring2}

					if !resolveCombat(player, boss) {
						itemSets = append(itemSets, player.items)
					}
				}
			}
		}
	}

	maxCost := 0

	for _, iset := range itemSets {
		cost := iset[0].cost + iset[1].cost + iset[2].cost + iset[3].cost

		if cost > maxCost {
			maxCost = cost
		}
	}

	return maxCost
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	return fmt.Sprintf(" %v", cycleGear(getPlayer(), getData(inputData)))
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	return fmt.Sprintf(" %v", mindControl(getPlayer(), getData(inputData)))
}
