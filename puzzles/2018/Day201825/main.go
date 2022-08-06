package Day201825

import (
	"fmt"
	"strings"

	"github.com/pezza/advent-of-code/puzzles/common"
)

var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string, int) {
	return 2018, 25, "Four-Dimensional Adventure"
}

type point struct {
	x int
	y int
	z int
	j int
}

type constellation []point

func getData(input string) []point {
	points := make([]point, 0)
	for _, line := range strings.Split(input, "\n") {
		a, b, c, d := 0, 0, 0, 0
		fmt.Sscanf(line, "%d,%d,%d,%d", &a, &b, &c, &d)
		points = append(points, point{a, b, c, d})
	}
	return points
}

func notExists(p point, pl []point) bool {
	for _, tp := range pl {
		if p == tp {
			return false
		}
	}
	return true
}

func Get4DMDistance(source point, target point) int {
	return common.Abs(source.x-target.x) + common.Abs(source.y-target.y) + common.Abs(source.z-target.z) + common.Abs(source.j-target.j)
}

func getConstellation(stars []point) constellation {
	startingStar := stars[0]
	constel := constellation{startingStar}

	starsAdded := true
	// until we can't find anything else to add to the constellation
	for starsAdded {
		starsAdded = false
		// loop through each star
		for _, star := range stars {
			// check it against each star in the current constellation
			for _, constelStar := range constel {
				// ignore if the same
				if star == constelStar {
					continue
				}

				if Get4DMDistance(star, constelStar) <= 3 && notExists(star, constel) {
					constel = append(constel, star)
					starsAdded = true
				}
			}
		}
	}

	return constel
}

func removeConstellationFromList(cons constellation, stars []point) []point {
	newStars := make([]point, 0)

	for _, testStar := range stars {
		found := false

		for _, consStar := range cons {
			if testStar == consStar {
				found = true
			}
		}

		if !found {
			newStars = append(newStars, testStar)
		}
	}

	return newStars
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	points := getData(inputData)
	count := 0
	for len(points) > 0 {
		cons := getConstellation(points)
		points = removeConstellationFromList(cons, points)
		count++
	}

	return fmt.Sprintf("%v", count)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	return fmt.Sprintf(" -- Not Yet Implemented --")
}
