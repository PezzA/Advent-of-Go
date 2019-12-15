package common

import (
	"fmt"
)

type Point struct {
	X int
	Y int
}

func GetMDistance(s Point, t Point) int {
	return Abs(s.X-t.X) + Abs(s.Y-t.Y)
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
