package common

type Point struct {
	X int
	Y int
}

func GetMDistance(source Point, target Point) int {
	return Abs(source.X-target.X) + Abs(source.Y-target.Y)

}

var Cardinal4 = []Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

func (s Point) Add(t Point) Point {
	return Point{s.X + t.X, s.Y + t.Y}
}
