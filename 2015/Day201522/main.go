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
	name     string
	duration int
	armour   int
	dmg      int
	regen    int
}

type charClass int

const (
	clPlayer charClass = 0
	clBoss   charClass = 1
)

type character struct {
	class  charClass
	health int
	dmg    int
	mana   int
}

type spell struct {
	name string
	cost int
	dmg  int
	hp   int
	effect
}

var spellbook = []spell{
	spell{"Magic Missile", 53, 4, 0, effect{}},
	spell{"Drain", 73, 2, 2, effect{}},
	spell{"Shield", 113, 0, 0, effect{"Shield", 6, 7, 0, 0}},
	spell{"Poison", 173, 0, 0, effect{"Poison", 6, 0, 3, 0}},
	spell{"Recharge", 229, 0, 0, effect{"Recharge", 5, 0, 0, 101}},
}

type outCome struct {
	winner charClass
	mana   int
}

func runBattle(player, boss character) []outCome {

	outcomes := []outCome{}

	for _, sp := range spellbook {
		outcomes = append(outcomes, resolveBattle(player, boss, []effect{}, sp, 0, 0)...)
	}

	return outcomes
}

func resolveBattle(player, boss character, fx []effect, s spell, depth int, spent int) []outCome {

	spaces := strings.Repeat(" ", depth*4)

	// Player Turn
	fmt.Printf("\n%v-- Player turn -- (%v)\n", spaces, depth)
	fmt.Printf("%v- Player has %v hit points, %v armour, %v mana\n", spaces, player.health, 0, player.mana)
	fmt.Printf("%v- Boss has %v hit points\n", spaces, boss.health)
	for _, e := range fx {
		if e.name == "Poison" {
			boss.health -= e.dmg
			fmt.Printf("%vPoison deals %v damage; its timer is now %v\n", spaces, e.dmg, e.duration-1)
		}
		if e.name == "Recharge" {
			player.mana += e.regen
			fmt.Printf("%vRecharge proveds %v mana; its timer is now %v\n", spaces, e.regen, e.duration-1)
		}
		if e.name == "Shield" {
			fmt.Printf("%vShield is active and proveds %v armour; its timer is now %v\n", spaces, e.regen, e.duration-1)
		}
	}
	fx = tickEffects(fx)

	// is boss dead?
	if boss.health <= 0 {
		fmt.Printf("%v* boss has died\n", spaces)
		fmt.Printf("%v--------------------------------\n", spaces)
		return []outCome{outCome{clPlayer, spent}}
	}

	// deduct mana cost
	player.mana -= s.cost
	spent += s.cost

	// cast target spell
	if s.name == "Magic Missile" {
		fmt.Printf("%vPlayer casts %v, dealing %v damage\n", spaces, s.name, s.dmg)
		boss.health -= s.dmg
	} else if s.name == "Drain" {
		fmt.Printf("%vPlayer casts %v, dealing %v damage and healing for %v health\n", spaces, s.name, s.dmg, s.hp)
		boss.health -= s.dmg
		player.health += s.hp
	} else {
		fmt.Printf("%vPlayer casts %v\n", spaces, s.name)
		fx = append(fx, s.effect)
	}
	fmt.Printf("%v#player has now spent a total of %v mana\n", spaces, spent)

	// is boss dead?
	if boss.health <= 0 {
		fmt.Printf("%v* boss has died\n", spaces)
		fmt.Printf("%v--------------------------------\n", spaces)
		return []outCome{outCome{clPlayer, spent}}
	}

	fmt.Printf("\n%v-- Boss turn -- (%v)\n", spaces, depth)
	fmt.Printf("%v- Player has %v hit points, %v armour, %v mana\n", spaces, player.health, 0, player.mana)
	fmt.Printf("%v- Boss has %v hit points\n", spaces, boss.health)
	armourBuff := 0
	// boss Turn
	for _, e := range fx {
		if e.name == "Poison" {
			boss.health -= e.dmg
			fmt.Printf("%vPoison deals %v damage; its timer is now %v\n", spaces, e.dmg, e.duration-1)
		}
		if e.name == "Recharge" {
			player.mana += e.regen
			fmt.Printf("%vRecharge proveds %v mana; its timer is now %v\n", spaces, e.regen, e.duration-1)
		}

		if e.name == "Shield" {
			armourBuff = e.armour
			fmt.Printf("%vShield is active and proveds %v armour; its timer is now %v\n", spaces, e.regen, e.duration-1)
		}
	}
	fx = tickEffects(fx)

	// is boss dead?
	if boss.health <= 0 {
		fmt.Printf("%v* boss has died\n", spaces)
		fmt.Printf("%v--------------------------------\n", spaces)
		return []outCome{outCome{clPlayer, spent}}
	}

	atk := boss.dmg - armourBuff

	if atk < 1 {
		atk = 1
	}
	fmt.Printf("%vBoss attacks for %v - %v = %v damage\n", spaces, boss.dmg, armourBuff, atk)
	player.health -= atk

	// is player dead or oom?
	if player.health <= 0 || player.mana < 53 {
		fmt.Printf("%v* Player has died\n", spaces)
		fmt.Printf("%v--------------------------------\n", spaces)
		return []outCome{outCome{clBoss, spent}}
	}

	results := []outCome{}
	for _, sp := range spellbook {
		doSpell := true

		// only cast valid spells
		if player.mana < sp.cost {
			doSpell = false
		}

		// cant cast an existing effect
		if sp.name == "Poison" || sp.name == "Shield" || sp.name == "Recharge" {
			for _, f := range fx {
				if f.name == sp.name {
					doSpell = false
				}
			}
		}

		if doSpell {
			results = append(results, resolveBattle(player, boss, fx, sp, depth+1, spent)...)
		}
	}

	return results
}

func tickEffects(fx []effect) []effect {
	updatedFx := []effect{}

	// update effect list
	for _, e := range fx {
		if e.duration > 1 {
			updatedFx = append(updatedFx, effect{
				e.name,
				e.duration - 1,
				e.armour,
				e.dmg,
				e.regen,
			})
		}
	}

	return updatedFx
}

func getBoss(input string) character {
	lines := strings.Split(input, "\n")
	hp, _ := strconv.Atoi(strings.Fields(lines[0])[2])
	dmg, _ := strconv.Atoi(strings.Fields(lines[1])[1])
	return character{clBoss, hp, dmg, 0}
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
