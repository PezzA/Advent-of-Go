package Day201514

import (
	"fmt"
	"strconv"
	"strings"
)

// Entry holds wraps the data and runner interfaces for this puzzle
var Entry dayEntry

type dayEntry bool

type reindeer struct {
	name string
	physicality
	state
}

type physicality struct {
	speed   int
	stamina int
	rest    int
}
type state struct {
	distance int
	running  bool
	ticks    int
	wins     int
}

func getData(input string) []*reindeer {
	beasts := make([]*reindeer, 0)
	for _, deer := range strings.Split(input, "\n") {
		bits := strings.Fields(deer)
		speed, _ := strconv.Atoi(bits[3])
		stamina, _ := strconv.Atoi(bits[6])
		rest, _ := strconv.Atoi(bits[13])
		beasts = append(beasts, &reindeer{bits[0], physicality{speed, stamina, rest}, state{0, true, stamina, 0}})
	}
	return beasts
}

func (r *reindeer) tick() {
	if r.running {
		r.distance += r.speed
	}

	r.ticks--

	if r.ticks == 0 {
		if r.running {
			r.running = false
			r.ticks = r.rest
		} else {
			r.running = true
			r.ticks = r.stamina
		}
	}

}

func updateLeadingReindeer(input []*reindeer) {
	maxDistance := 0

	for _, deer := range input {
		if deer.distance > maxDistance {
			maxDistance = deer.distance
		}
	}

	for _, deer := range input {
		if deer.distance == maxDistance {
			deer.wins++
		}
	}

}

func (td dayEntry) Describe() (int, int, string, int) {
	return 2015, 14, "Reindeer Olympics", 2
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	deers := getData(inputData)

	for i := 0; i < 2503; i++ {
		for _, deer := range deers {
			deer.tick()
		}
	}

	maxDistance := 0

	for _, deer := range deers {
		if deer.distance > maxDistance {
			maxDistance = deer.distance
		}
	}
	return fmt.Sprintf("%v", maxDistance)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	deers := getData(inputData)

	for i := 0; i < 2503; i++ {
		for _, deer := range deers {
			deer.tick()
		}
		updateLeadingReindeer(deers)
	}

	maxWins := 0

	for _, deer := range deers {
		if deer.wins > maxWins {
			maxWins = deer.wins
		}
	}
	return fmt.Sprintf("%v", maxWins)
}
