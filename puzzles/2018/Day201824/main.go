package Day201824

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string, int) {
	return 2018, 24, "Immune System Simulator 20XX", 0
}

const immuneArmy = "Immune"
const infectionArmy = "Infection"

type group struct {
	army       string
	id         int
	units      int
	hp         int
	attack     int
	initative  int
	attackType string
	weaknesses []string
	immunites  []string
}

type battle struct {
	attackInitiative int
	attackID         int
	attackArmy       string
	defendID         int
}

var groupParser = regexp.MustCompile(`(?P<units>[0-9]*) units each with (?P<HP>[0-9]*) hit points( \((?P<mods>.*)*\))* with an attack that does (?P<atk>[0-9]*) (?P<atkType>\S*) damage at initiative (?P<init>[0-9]*)`)
var groupParserNames = groupParser.SubexpNames()

func (g group) String() string {
	return fmt.Sprintf("%v group %v. U:%v HP:%v A:%v AT:%v W:%v I:%v", g.army, g.id, g.units, g.hp, g.attack, g.attackType, g.weaknesses, g.immunites)
}

func (b battle) String() string {
	return fmt.Sprintf("[%v] %v group %v attacks defending group %v.", b.attackInitiative, b.attackArmy, b.attackID, b.defendID)
}

func parseAttributes(input string) ([]string, []string) {
	input = strings.TrimSpace(input)
	weaknesses, immunities := []string{}, []string{}

	if input == "" {
		return weaknesses, immunities
	}

	for _, line := range strings.Split(input, ";") {
		words := strings.Fields(line)

		if words[0] == "weak" {
			weaknesses = words[2:]

			for index := range weaknesses {
				weaknesses[index] = strings.TrimSpace(strings.Replace(weaknesses[index], ",", "", -1))
			}
		}

		if words[0] == "immune" {
			immunities = words[2:]
			for index := range immunities {
				immunities[index] = strings.TrimSpace(strings.Replace(immunities[index], ",", "", -1))
			}
		}
	}

	return weaknesses, immunities
}

func getData(input string) ([]group, []group) {
	immuneGroup := make([]group, 0)
	infectionGroup := make([]group, 0)

	lines := strings.Split(input, "\n")

	isInfectionList := false
	for _, line := range lines {
		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}
		if line == "Immune System:" {
			isInfectionList = false
			continue
		}
		if line == "Infection:" {
			isInfectionList = true
			continue
		}

		match := groupParser.FindStringSubmatch(line)
		matchData := make(map[string]string, 0)

		for index, matchName := range groupParserNames {
			if index != 0 && matchName != "" {
				matchData[matchName] = match[index]
			}
		}

		weaknesses, immunities := parseAttributes(matchData["mods"])

		units, _ := strconv.Atoi(matchData["units"])
		hp, _ := strconv.Atoi(matchData["HP"])
		atk, _ := strconv.Atoi(matchData["atk"])
		init, _ := strconv.Atoi(matchData["init"])

		if isInfectionList {
			infectionGroup = append(infectionGroup, group{
				id:         len(infectionGroup) + 1,
				army:       infectionArmy,
				units:      units,
				hp:         hp,
				attack:     atk,
				initative:  init,
				attackType: matchData["atkType"],
				weaknesses: weaknesses,
				immunites:  immunities,
			})
		} else {
			immuneGroup = append(immuneGroup, group{
				id:         len(immuneGroup) + 1,
				army:       immuneArmy,
				units:      units,
				hp:         hp,
				attack:     atk,
				initative:  init,
				attackType: matchData["atkType"],
				weaknesses: weaknesses,
				immunites:  immunities,
			})
		}
	}

	return immuneGroup, infectionGroup
}

func (g group) effectivePower() int {
	return g.units * g.attack
}

func (g group) deepCopy() group {
	return group{
		id:         g.id,
		army:       g.army,
		units:      g.units,
		hp:         g.hp,
		attack:     g.attack,
		initative:  g.initative,
		attackType: g.attackType,
		weaknesses: g.weaknesses,
		immunites:  g.immunites,
	}
}

func (g group) receiveAttack(attacker group) (group, int) {
	attack := attacker.attackingPower(g)

	if attack == 0 {
		return g, 0
	}

	unitsLost := attack / g.hp

	if unitsLost > g.units {
		unitsLost = g.units
	}

	g.units = g.units - unitsLost

	return g, unitsLost
}

func getInitativeOrder(a1 []group, a2 []group) []group {
	fullList := make([]group, 0)

	for _, g := range a1 {
		fullList = append(fullList, g.deepCopy())
	}

	for _, g := range a2 {
		fullList = append(fullList, g.deepCopy())
	}

	sort.Slice(fullList, func(i, j int) bool {
		return fullList[i].initative >= fullList[j].initative
	})

	return fullList
}

func getGroup(i int, army []group) group {
	for _, g := range army {
		if g.id == i {
			return g
		}
	}

	return group{}
}

func getTargetOrder(a1 []group, a2 []group) []group {

	fullList := make([]group, 0)

	for _, g := range a1 {
		fullList = append(fullList, g.deepCopy())
	}

	for _, g := range a2 {
		fullList = append(fullList, g.deepCopy())
	}

	sort.Slice(fullList, func(i, j int) bool {
		if fullList[i].effectivePower() == fullList[j].effectivePower() {
			return fullList[i].initative >= fullList[j].initative
		}

		return fullList[i].effectivePower() > fullList[j].effectivePower()
	})

	return fullList
}

func (g group) attackingPower(target group) int {
	for _, immunity := range target.immunites {
		if g.attackType == immunity {
			return 0
		}
	}

	for _, weakness := range target.weaknesses {
		if g.attackType == weakness {
			return g.effectivePower() * 2
		}
	}

	return g.effectivePower()
}

func isAlreadyAssailed(g group, battleList []battle) bool {
	for _, b := range battleList {
		if b.defendID == g.id && g.army != b.attackArmy {
			return true
		}
	}

	return false
}

func resolveAttackOrder(immunes []group, infections []group) []battle {
	battles, priorityList := make([]battle, 0), getTargetOrder(immunes, infections)

	for _, contender := range priorityList {
		checkList := immunes

		if contender.army == immuneArmy {
			checkList = infections
		}

		bestAttack, currentTarget := 0, group{}
		for _, candidate := range checkList {
			if isAlreadyAssailed(candidate, battles) {
				continue
			}

			attack := contender.attackingPower(candidate)

			if attack == 0 || attack < bestAttack {
				continue
			}

			if attack > bestAttack {
				bestAttack = attack
				currentTarget = candidate
				continue
			}

			if attack == bestAttack {
				if candidate.effectivePower() > currentTarget.effectivePower() {
					bestAttack = attack
					currentTarget = candidate
					continue
				}

				if candidate.effectivePower() == currentTarget.effectivePower() {
					if candidate.initative > currentTarget.initative {
						bestAttack = attack
						currentTarget = candidate
					}
				}
			}
		}

		if bestAttack > 0 {
			battles = append(battles, battle{contender.initative, contender.id, contender.army, currentTarget.id})
		}
	}

	// sort battles into descending order of initiative
	sort.Slice(battles, func(i, j int) bool {
		return battles[i].attackInitiative >= battles[j].attackInitiative
	})

	return battles
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	return fmt.Sprintf(" -- Not Yet Implemented --")
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	return fmt.Sprintf(" -- Not Yet Implemented --")
}
