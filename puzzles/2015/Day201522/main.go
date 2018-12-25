package Day201522

import (
	"fmt"
	"strconv"
	"strings"
)

// Entry holds wraps the data and runner interfaces for this puzzle
var Entry dayEntry

type dayEntry bool

type effect struct {
	duration int
	armour   int
	dmg      int
	mana     int
}

type character struct {
	health int
	dmg    int
	mana   int
	armour int
}

type spell struct {
	name string
	cost int
	dmg  int
	hp   int
	effect
}

func getSpells() []spell {
	spells := make([]spell, 0)
	spells = append(spells, spell{"Magic Missile", 53, 4, 0, effect{}})
	spells = append(spells, spell{"Drain", 73, 2, 2, effect{}})
	spells = append(spells, spell{"Shield", 113, 0, 0, effect{6, 7, 0, 0}})
	spells = append(spells, spell{"Poison", 173, 0, 0, effect{6, 0, 3, 0}})
	spells = append(spells, spell{"Recharge", 229, 0, 0, effect{5, 0, 0, 101}})
	return spells
}

func getBoss(input string) character {
	lines := strings.Split(input, "\n")
	hp, _ := strconv.Atoi(strings.Fields(lines[0])[2])
	dmg, _ := strconv.Atoi(strings.Fields(lines[1])[1])
	return character{hp, dmg, 0, 0}
}

func combat(player character, boss character, stack []effect, spells []spell, depth int, result chan<- int, manaSpent int, playerAttack bool) {
	if depth == 0 {
		defer close(result)
	}

	armourModifier := 0
	// resolve effects
	for index, ef := range stack {
		if ef.mana != 0 {
			player.mana += ef.mana
		}
		if ef.armour != 0 {
			armourModifier += ef.armour
		}
		if ef.dmg != 0 {
			boss.health -= ef.dmg
		}

		// decrement effect duration
		stack[index].duration -= 1
		// if effect duration is zero, remove from stack.
		if stack[index].duration == 0 {
			stack = append(stack[:index], stack[index+1:]...)
		}
	}

	// check to see if boss has died
	if boss.health <= 0 {
		result <- manaSpent
		return
	}

	if !playerAttack {
		bossAttack := boss.dmg - armourModifier
		if bossAttack <= 0 {
			bossAttack = 1
		}

		player.health -= bossAttack

		if player.health <= 0 {
			return
		}
	}

	if playerAttack {
		spellCast := false
		for _, sp := range spells {

			sp.armour += 1
			combat(player, boss, stack, spells, depth+1, result, manaSpent, !playerAttack)
		}

		if !spellCast {
			return
		}
	}
}

func (td dayEntry) Describe() (int, int, string) {
	return 2015, 22, "Wizard Simulator 20XX"
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	return fmt.Sprintf(" -- Not Yet Implemented --")
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	return fmt.Sprintf(" -- Not Yet Implemented --")
}
