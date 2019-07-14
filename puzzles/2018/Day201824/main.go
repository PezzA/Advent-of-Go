package Day201824

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string) {
	return 2018, 24, "Immune System Simulator 20XX"
}

type group struct {
	id         int
	units      int
	hp         int
	attack     int
	initative  int
	attackType string
	weaknesses []string
	immunites  []string
}

var groupParser = regexp.MustCompile(`(?P<units>[0-9]*) units each with (?P<HP>[0-9]*) hit points( \((?P<mods>.*)*\))* with an attack that does (?P<atk>[0-9]*) (?P<atkType>\S*) damage at initiative (?P<init>[0-9]*)`)
var groupParserNames = groupParser.SubexpNames()

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
	for index, line := range lines {
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

		newGroup := group{
			id:         index,
			units:      units,
			hp:         hp,
			attack:     atk,
			initative:  init,
			attackType: matchData["atkType"],
			weaknesses: weaknesses,
			immunites:  immunities,
		}

		if isInfectionList {
			infectionGroup = append(infectionGroup, newGroup)
		} else {
			immuneGroup = append(immuneGroup, newGroup)
		}
	}

	return immuneGroup, infectionGroup
}
func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	return fmt.Sprintf(" -- Not Yet Implemented --")
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	return fmt.Sprintf(" -- Not Yet Implemented --")
}
