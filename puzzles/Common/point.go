package common

import "fmt"

type Point struct {
	X int
	Y int
}

func GetMDistance(s Point, t Point) int {
	return Abs(s.X-t.X) + Abs(s.Y-t.Y)
}

var Cardinal4 = []Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

func (p Point) Add(t Point) Point {
	return Point{p.X + t.X, p.Y + t.Y}
}

func (p Point) String() string {
	return fmt.Sprintf("<%v,%v>", p.X, p.Y)
}
