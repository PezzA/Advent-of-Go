package Day201522

import (
	"fmt"
	"math"
	"strings"
)

type spell struct {
	name           string
	cost           int
	damage         int
	heal           int
	duration       int
	damageOverTime int
	shieldOverTime int
	manaOverTime   int
}

type spellStack map[string]spell

var spellBook = spellStack{
	"missile": {
		name:   "missile",
		cost:   53,
		damage: 4,
	},
	"drain": {
		name:   "drain",
		cost:   73,
		damage: 2,
		heal:   2,
	},
	"shield": {
		name:           "shield",
		cost:           113,
		duration:       6,
		shieldOverTime: 7,
	},
	"poison": {
		name:           "poison",
		cost:           173,
		duration:       6,
		damageOverTime: 3,
	},
	"recharge": {
		name:         "recharge",
		cost:         229,
		duration:     5,
		manaOverTime: 101,
	},
}

type character struct {
	health int
	damage int
	mana   int
	armour int
}

type gameState struct {
	player    character
	boss      character
	manaSpent int
	stack     spellStack
	casts     []*gameState
}

type logger func(string, int)

func getStartingCharacters(bossHealth, bossDamage int) (character, character) {
	return character{
			health: 50,
			mana:   500,
		}, character{
			health: bossHealth,
			damage: bossDamage,
		}
}

func updateStack(
	player character,
	boss character,
	stack spellStack,
	depth int,
	log logger,
) (character, character, spellStack) {
	player.armour = 0

	newSpellStack := make(spellStack, 0)

	for k, v := range stack {

		prefix := ""
		if v.shieldOverTime > 0 {
			prefix = fmt.Sprintf("Player gains %d armour from shield", v.shieldOverTime)
			player.armour = v.shieldOverTime
		}

		if v.damageOverTime > 0 {
			prefix = fmt.Sprintf("Boss takes %d damage from poison", v.damageOverTime)
			boss.health -= v.damageOverTime
		}

		if v.manaOverTime > 0 {
			prefix = fmt.Sprintf("Player gains %d mana from recharge", v.manaOverTime)
			player.mana += v.manaOverTime
		}

		v.duration -= 1

		if v.duration == 0 {
			log(fmt.Sprintf("%s, %s expires", prefix, k), depth)
			delete(stack, k)
		} else {
			log(fmt.Sprintf("%s, %s has %d turns remaining", prefix, k, v.duration), depth)
			newSpellStack[k] = v
		}
	}

	return player, boss, newSpellStack
}

func battleRound(
	player character,
	boss character,
	stack spellStack,
	targetSpell spell,
	hardMode bool,
	depth int,
	log logger,
) (character, character, spellStack, int) {

	// update stack
	log("-- Player Turn --", depth)
	if hardMode {
		log("! Player loses 1 health in hard mode", depth)
		player.health -= 1

		if player.health <= 0 {
			log("This kills the player and the boss wins", depth)
			return player, boss, stack, 0
		}
	}

	player, boss, stack = updateStack(player, boss, stack, depth, log)
	log(fmt.Sprintf("* Player: %d hp %d mana", player.health, player.mana), depth)
	log(fmt.Sprintf("* Boss: %d hp", boss.health), depth)
	if boss.health <= 0 {
		log("This kills the boss and the player wins", depth)
		return player, boss, stack, 0
	}

	// apply target spell
	if targetSpell.damage > 0 {
		log(fmt.Sprintf("Boss takes %d damage from %s", targetSpell.damage, targetSpell.name), depth)
		boss.health -= targetSpell.damage
	}

	if targetSpell.heal > 0 {
		log(fmt.Sprintf("Player heals for %d  from %s", targetSpell.heal, targetSpell.name), depth)
		player.health += targetSpell.heal
	}

	if targetSpell.duration > 0 {
		log(fmt.Sprintf("Player casts %s for %d mana", targetSpell.name, targetSpell.cost), depth)
		stack[targetSpell.name] = targetSpell
	}

	player.mana -= targetSpell.cost

	if boss.health <= 0 {
		log("This kills the boss and the player wins", depth)
		return player, boss, stack, targetSpell.cost
	}

	// update stack
	log("-- Boss Turn --", depth)

	player, boss, stack = updateStack(player, boss, stack, depth, log)
	log(fmt.Sprintf("* Player: %d hp %d mana", player.health, player.mana), depth)
	log(fmt.Sprintf("* Boss: %d hp", boss.health), depth)
	if boss.health <= 0 {
		log("This kills the boss and the player wins", depth)
		return player, boss, stack, targetSpell.cost
	}

	dmg := boss.damage - player.armour

	if dmg <= 0 {
		dmg = 1
	}

	log(fmt.Sprintf("Player takes %d damage", dmg), depth)
	player.health -= dmg

	if player.health <= 0 {
		log("This kills the player and the boss wins", depth)
		return player, boss, stack, targetSpell.cost
	}

	return player, boss, stack, targetSpell.cost
}

func runBattle(player character, boss character, hardMode bool, log logger) *gameState {
	stack := make(spellStack, 0)

	gs := &gameState{
		player:    player,
		boss:      boss,
		manaSpent: 0,
		stack:     stack,
		casts:     make([]*gameState, 0),
	}

	for _, spell := range spellBook {
		gs.casts = append(gs.casts, resolveBattle(gs, spell, hardMode, 0, log))
	}

	return gs
}

func log(input string, depth int) {
	fmt.Print(fmt.Sprintf("%s%d. %s\n", strings.Repeat(" ", depth*2), depth, input))
}

func nullLog(input string, depth int) {
}

var minWin = math.MaxInt32

func resolveBattle(gs *gameState, s spell, hardmode bool, depth int, log logger) *gameState {
	log(s.name, depth)

	player, boss, stack, spent := battleRound(gs.player, gs.boss, gs.stack, s, hardmode, depth, log)

	newGs := &gameState{
		player:    player,
		boss:      boss,
		stack:     stack,
		manaSpent: gs.manaSpent + spent,
		casts:     make([]*gameState, 0),
	}

	if newGs.manaSpent > minWin {
		return newGs
	}

	if boss.health <= 0 && newGs.manaSpent < minWin {
		minWin = newGs.manaSpent
	}

	if player.health <= 0 || boss.health <= 0 {
		return newGs
	}

	for _, spell := range spellBook {
		if canCast(stack, player.mana, spell) {
			newGs.casts = append(newGs.casts, resolveBattle(newGs, spell, hardmode, depth+1, log))
		}
	}

	return newGs
}

func canCast(stack spellStack, mana int, s spell) bool {
	if s.cost > mana {
		return false
	}

	if v, ok := stack[s.name]; ok {
		if v.duration > 1 {
			return false
		}
	}

	return true
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	player, boss := getStartingCharacters(58, 9)

	minWin = math.MaxInt32
	runBattle(player, boss, false, nullLog)
	return fmt.Sprintf("%v", minWin)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	player, boss := getStartingCharacters(58, 9)
	minWin = math.MaxInt32
	runBattle(player, boss, true, nullLog)
	return fmt.Sprintf("%v", minWin)
}
