package Day201601

import (
	"errors"
	"strconv"
	"strings"
)

// Entry holds wraps the data and runner interfaces for this puzzle
var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string, int) {
	return 2016, 1, "No Time for a Taxicab", 2
}

// Point represents a basic 2d co-ordinate
type Point struct {
	x int
	y int
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

type instruction struct {
	direction string
	steps     int
}

type position struct {
	p      Point
	facing int
}

func getlength(p Point) int {
	return abs(p.x) + abs(p.y)
}

func solveRawLocation(inputData string) (string, error) {
	var currentPosition = position{Point{0, 0}, 0}

	for _, ins := range parseInstructions(inputData) {
		currentPosition, _ = move(currentPosition, ins)
	}

	return strconv.Itoa(getlength(currentPosition.p)), nil
}

func containsPoint(points []Point, point Point) bool {
	for _, testPoint := range points {
		if (testPoint.x == point.x) && (testPoint.y == point.y) {
			return true
		}
	}
	return false
}

func solveActualLocation(inputData string) (string, error) {
	var currentPosition = position{Point{0, 0}, 0}

	points := make([]Point, 0)

	for _, ins := range parseInstructions(inputData) {
		newPoints := make([]Point, 0)

		currentPosition, newPoints = move(currentPosition, ins)

		for _, testPoint := range newPoints {
			if containsPoint(points, testPoint) {
				return strconv.Itoa(getlength(testPoint)), nil
			}
		}

		points = append(points, newPoints...)
	}

	return "-", errors.New("no intersections found")
}

func move(pos position, ins instruction) (position, []Point) {
	var newOrientation = orientate(ins.direction, pos.facing)

	points := make([]Point, 0)

	newPoint := pos.p

	for i := 0; i < ins.steps; i++ {
		newPoint = translate(newOrientation, 1, newPoint)
		points = append(points, newPoint)
	}

	return position{newPoint, newOrientation}, points
}

func parseInstructions(input string) []instruction {
	instructions := make([]instruction, 0)

	for _, c := range strings.Split(input, ",") {
		var input = strings.TrimSpace(c)

		step, _ := strconv.Atoi(input[1:len(input)])

		instructions = append(instructions, instruction{direction: input[0:1], steps: step})
	}

	return instructions
}

func orientate(direction string, currentOrientation int) int {
	switch direction {
	case "L":
		if currentOrientation == 0 {
			return 270
		}
		return currentOrientation - 90
	case "R":
		if currentOrientation == 270 {
			return 0
		}
		return currentOrientation + 90
	default:
		return 0
	}
}

func translate(orientation int, shift int, startPoint Point) Point {
	switch orientation {
	case 0:
		return Point{startPoint.x, startPoint.y + shift}
	case 90:
		return Point{startPoint.x + shift, startPoint.y}
	case 180:
		return Point{startPoint.x, startPoint.y - shift}
	case 270:
		return Point{startPoint.x - shift, startPoint.y}
	}

	return Point{0, 0}
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {

	result, _ := solveRawLocation(inputData)
	return result
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	result, _ := solveActualLocation(inputData)
	return result
}
