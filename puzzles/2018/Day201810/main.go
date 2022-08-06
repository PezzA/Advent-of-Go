package Day201810

import (
	"fmt"
	"strings"

	"github.com/pezza/advent-of-code/puzzles/common"
)

var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string, int) {
	return 2018, 10, "The Stars Align"
}

type star struct {
	px int
	py int
	vx int
	vy int
}

func getBounds(input []star) (common.Point, common.Point) {
	minX, minY, maxX, maxY := -1, -1, -1, -1

	for _, star := range input {
		if minX == -1 || star.px < minX {
			minX = star.px
		}

		if maxX == -1 || star.px > maxX {
			maxX = star.px
		}

		if minY == -1 || star.py < minY {
			minY = star.py
		}

		if maxY == -1 || star.py > maxX {
			maxY = star.py
		}
	}

	return common.Point{minX, minY}, common.Point{maxX, maxY}
}

func getData(input string) []star {
	stars := make([]star, 0)

	for _, line := range strings.Split(input, "\n") {
		px, py, vx, vy := 0, 0, 0, 0
		fmt.Sscanf(line, "position=<%d,  %d> velocity=<%d, %d>", &px, &py, &vx, &vy)

		stars = append(stars, star{px, py, vx, vy})
	}

	return stars
}
func getBoundSize(min common.Point, max common.Point) int {
	return common.Abs(max.Y-min.Y) + common.Abs(max.X-min.X)
}

func updateStars(input []star) []star {
	returnValue := make([]star, 0)

	for _, oldStar := range input {
		returnValue = append(returnValue, star{oldStar.px + oldStar.vx, oldStar.py + oldStar.vy, oldStar.vx, oldStar.vy})
	}

	return returnValue
}

// I would imagine that when the starts converge, the bounds will be at thier smallest point
func getFocusFrame(input []star) ([]star, int) {
	boundSize, prevBoundSize := 0, getBoundSize(getBounds(input))

	var newStars []star
	prevStars := input
	count := 0
	for {
		newStars = updateStars(prevStars)
		boundSize = getBoundSize(getBounds(newStars))

		if boundSize > prevBoundSize {
			break
		}

		prevBoundSize = boundSize
		prevStars = newStars
		count++
	}

	return prevStars, count
}

// YOU HAVE TO OBSERVE THE OUTPUT!!
func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	stars := getData(Entry.PuzzleInput())

	focusStars, _ := getFocusFrame(stars)
	min, max := getBounds(focusStars)

	x := max.X - min.X
	y := max.Y - min.Y

	retVal := ""

	for ly := 0; ly <= y; ly++ {
		for lx := 0; lx <= x; lx++ {

			found := false

			for _, drawStar := range focusStars {
				if lx+min.X == drawStar.px && ly+min.Y == drawStar.py {
					retVal += "#"
					found = true
					break
				}
			}
			if !found {
				retVal += "."
			}
		}

		retVal += "\n"
	}
	updateChan <- []string{"See output below"}
	return retVal
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	stars := getData(Entry.PuzzleInput())
	_, seconds := getFocusFrame(stars)
	return fmt.Sprintf("%v", seconds)
}
