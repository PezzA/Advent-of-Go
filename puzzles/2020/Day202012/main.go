package Day202012

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pezza/advent-of-code/puzzles/common"
)

type instruction struct {
	ins      string
	distance int
}

type ship struct {
	common.Point
	direction common.Degree
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
		Point:     common.Point{X: 0, Y: 0},
		direction: common.EastDegree,
	}
}

func (s *ship) processRevisedIns(ins instruction, wp common.Point) common.Point {
	switch ins.ins {
	case "N", "E", "S", "W":
		wp.MoveOrdinal(common.GetDegreeOrdinal(ins.ins), ins.distance)
		return wp
	case "F":
		s.X += ins.distance * wp.X
		s.Y += ins.distance * wp.Y
	case "L":
		wp.RotateOrigin(common.Degree(ins.distance))
	case "R":
		wp.RotateOrigin(common.Degree(-ins.distance))
	}

	return wp
}

func (s *ship) processIns(ins instruction) {
	switch ins.ins {
	case "N", "E", "S", "W":
		s.MoveOrdinal(common.GetDegreeOrdinal(ins.ins), ins.distance)
	case "F":
		s.MoveOrdinal(s.direction, ins.distance)
	case "L":
		s.direction = s.direction.OrientateLeft(ins.distance)
	case "R":
		s.direction = s.direction.OrientateRight(ins.distance)
	}
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	ins := getData(inputData)

	ship := getStartShip()

	for i := range ins {
		ship.processIns(ins[i])
	}

	return fmt.Sprintf("%v", ship.GetMDistanceOrigin())
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	ins := getData(Entry.PuzzleInput())

	ship := getStartShip()
	waypoint := common.Point{X: 10, Y: 1}

	for i := range ins {
		waypoint = ship.processRevisedIns(ins[i], waypoint)
	}

	return fmt.Sprintf("%v", ship.GetMDistanceOrigin())
}
