package Day202012

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/pezza/advent-of-code/common"
)

type instruction struct {
	ins      string
	distance int
}

type point struct {
	x int
	y int
}

type ship struct {
	point
	direction int
}

func getData(input string) []instruction {
	lines := strings.Split(input, "\n")
	ins := make([]instruction, len(lines))

	for i := range lines {
		newIns := lines[i][:1]
		newDis, _ := strconv.Atoi(lines[i][1:])

		ins[i] = instruction{
			ins:      newIns,
			distance: newDis,
		}
	}
	return ins
}

func getStartShip() ship {
	return ship{
		point:     point{x: 0, y: 0},
		direction: 90,
	}
}

func (p *point) move(direction string, distance int) {
	switch direction {
	case "N":
		p.y += distance
	case "S":
		p.y -= distance
	case "E":
		p.x += distance
	case "W":
		p.x -= distance
	}
}

func degreeToDirection(degree int) string {
	switch degree {
	case 0:
		return "N"
	case 90:
		return "E"
	case 180:
		return "S"
	case 270:
		return "W"
	}

	log.Fatal("invalid degree")
	return ""
}

func (p *point) rotate(degree int) {
	radians := float64(degree) * (math.Pi / 180.0)
	cos, sin := math.Cos(radians), math.Sin(radians)
	origX, origY := float64(p.x), float64(p.y)

	p.x = int(math.Round((origX * cos) - (origY * sin)))
	p.y = int(math.Round((origY * cos) + (origX * sin)))
}

func (s *ship) processRevisedIns(ins instruction, wp point) point {
	switch ins.ins {
	case "N", "E", "S", "W":
		wp.move(ins.ins, ins.distance)
		return wp
	case "F":
		s.x += ins.distance * wp.x
		s.y += ins.distance * wp.y
	case "L":
		wp.rotate(ins.distance)
	case "R":
		wp.rotate(-ins.distance)
	}

	return wp
}

func (s *ship) processIns(ins instruction) {
	switch ins.ins {
	case "N", "E", "S", "W":
		s.move(ins.ins, ins.distance)
	case "F":
		s.move(degreeToDirection(s.direction), ins.distance)
	case "L":
		s.direction = rotateLeft(s.direction, ins.distance)
	case "R":
		s.direction = rotateRight(s.direction, ins.distance)
	}
}

func rotateRight(degree int, turnAmount int) int {
	degree += turnAmount

	if degree >= 360 {
		degree -= 360
	}
	return degree
}

func rotateLeft(degree int, turnAmount int) int {
	degree -= turnAmount

	if degree < 0 {
		degree += 360
	}
	return degree
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	ins := getData(inputData)

	ship := getStartShip()

	for i := range ins {
		ship.processIns(ins[i])
	}

	return fmt.Sprintf("%v", common.Abs(ship.x)+common.Abs(ship.y))
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	ins := getData(Entry.PuzzleInput())

	ship := getStartShip()
	waypoint := point{10, 1}

	for i := range ins {
		waypoint = ship.processRevisedIns(ins[i], waypoint)
	}

	return fmt.Sprintf("%v", common.Abs(ship.x)+common.Abs(ship.y))
}
