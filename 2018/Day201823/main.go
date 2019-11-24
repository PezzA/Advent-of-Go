package Day201823

import (
	"fmt"
	"strings"

	"github.com/pezza/advent-of-code/common"
)

var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string) {
	return 2018, 23, "Experimental Emergency Teleportation"
}

type point3D struct {
	x int
	y int
	z int
}

type nanobot struct {
	point3D
	r int
}

func (p point3D) getMDistance(t point3D) int {
	return common.Abs(p.x-t.x) + common.Abs(p.y-t.y) + common.Abs(p.z-t.z)
}

func getData(input string) []nanobot {
	lines := strings.Split(input, "\n")
	bots := make([]nanobot, 0)

	for _, line := range lines {
		x, y, z, r := 0, 0, 0, 0
		fmt.Sscanf(line, "pos=<%d,%d,%d>, r=%d", &x, &y, &z, &r)

		bots = append(bots, nanobot{point3D: point3D{x: x, y: y, z: z}, r: r})
	}

	return bots
}

func getHighestRadius(bots []nanobot) nanobot {
	found, foundBot := false, nanobot{point3D{0, 0, 0}, 0}

	for _, bot := range bots {
		if found == false || bot.r > foundBot.r {
			found = true
			foundBot = bot
		}
	}

	return foundBot
}

func (n nanobot) getBotsInRange(bots []nanobot) []nanobot {
	rangeBots := make([]nanobot, 0)

	for _, bot := range bots {
		if n.getMDistance(bot.point3D) <= n.r {
			rangeBots = append(rangeBots, bot)
		}
	}

	return rangeBots
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	bots := getData(inputData)
	rBot := getHighestRadius(bots)
	botsInRange := rBot.getBotsInRange(bots)

	return fmt.Sprintf("%v", len(botsInRange))
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	return fmt.Sprintf(" -- Not Yet Implemented --")
}
