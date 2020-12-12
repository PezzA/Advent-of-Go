package common

import (
	"fmt"
	"math"
)

type Degree int

const (
	NorthDegree Degree = 0
	EastDegree  Degree = 90
	SouthDegree Degree = 180
	WestDegree  Degree = 270
)

type Point struct {
	X int
	Y int
}

func (p Point) GetMDistance(t Point) int {
	return Abs(p.X-t.X) + Abs(p.Y-t.Y)
}

func (p Point) GetMDistanceOrigin() int {
	return Abs(p.X) + Abs(p.Y)
}

var North = Point{X: 0, Y: -1}
var East = Point{X: 1, Y: 0}
var South = Point{X: 0, Y: 1}
var West = Point{X: -1, Y: 0}

var Cardinal4ReadingOrder = []Point{
	{X: 0, Y: -1}, // North
	{X: -1, Y: 0}, // West
	{X: 1, Y: 0},  // East
	{X: 0, Y: 1},  // South
}

var Cardinal4 = []Point{
	{X: 0, Y: -1}, // North
	{X: 1, Y: 0},  // East
	{X: 0, Y: 1},  // South
	{X: -1, Y: 0}, // West
}

var Cardinal8 = []Point{
	{X: 0, Y: -1},
	{X: 1, Y: -1},
	{X: 1, Y: 0},
	{X: 1, Y: 1},
	{X: 0, Y: 1},
	{X: -1, Y: 1},
	{X: -1, Y: 0},
	{X: -1, Y: -1},
}

func (p Point) Add(t Point) Point {
	return Point{p.X + t.X, p.Y + t.Y}
}

func (p Point) String() string {
	return fmt.Sprintf("<%v,%v>", p.X, p.Y)
}

// Rotate will rotate a Point around the origin
func (p *Point) RotateOrigin(degree Degree) {
	radians := float64(degree) * (math.Pi / 180.0)
	cos, sin := math.Cos(radians), math.Sin(radians)
	origX, origY := float64(p.X), float64(p.Y)

	p.X = int(math.Round((origX * cos) - (origY * sin)))
	p.Y = int(math.Round((origY * cos) + (origX * sin)))
}

func (d Degree) OrientateRight(turnAmount int) Degree {
	newD := d + Degree(turnAmount)

	if newD >= 360 {
		newD -= 360
	}

	return newD
}

func (d Degree) OrientateLeft(turnAmount int) Degree {
	newD := d - Degree(turnAmount)

	if newD < 0 {
		newD += 360
	}
	return newD
}

func GetDegreeOrdinal(o string) Degree {
	switch o {
	case "N":
		return NorthDegree
	case "S":
		return SouthDegree
	case "E":
		return EastDegree
	case "W":
		return WestDegree

	}

	return 0
}

func (p *Point) MoveOrdinal(direction Degree, distance int) {
	switch direction {
	case NorthDegree:
		p.Y += distance
	case SouthDegree:
		p.Y -= distance
	case EastDegree:
		p.X += distance
	case WestDegree:
		p.X -= distance
	}
}
