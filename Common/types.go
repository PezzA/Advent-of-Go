package common

import "github.com/pezza/aoc2017/Common"

type Point struct {
	X int
	Y int
}

func GetMDistance(source Point, target Point) int {
	return common.Abs(source.X-target.X) + common.Abs(source.Y-target.Y)

}
