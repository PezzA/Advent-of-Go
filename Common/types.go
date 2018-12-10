package common

type Point struct {
	X int
	Y int
}

func GetMDistance(source Point, target Point) int {
	return Abs(source.X-target.X) + Abs(source.Y-target.Y)

}
