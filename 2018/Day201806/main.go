package Day201806

import (
	"fmt"
	"strings"

	"github.com/pezza/advent-of-code/Common"
)

var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string) {
	return 2018, 6, "Chronal Coordinates"
}

// end of scaffolding

type numberedPoint struct {
	id string
	common.Point
}

type infinitePlane map[common.Point]string

func getData(inputData string) []numberedPoint {
	points := make([]numberedPoint, 0)

	startLabel := 65
	for _, line := range strings.Split(inputData, "\n") {
		point := numberedPoint{string(startLabel), common.Point{}}
		fmt.Sscanf(line, "%d, %d", &point.X, &point.Y)
		startLabel++
		points = append(points, point)
	}
	return points
}

func getBounds(points []numberedPoint) (common.Point, common.Point) {
	minX, minY, maxX, maxY := -1, -1, -1, -1

	for _, point := range points {
		if minX == -1 || point.X < minX {
			minX = point.X
		}

		if maxX == -1 || point.X > maxX {
			maxX = point.X
		}

		if minY == -1 || point.Y < minY {
			minY = point.Y
		}

		if maxY == -1 || point.Y > maxX {
			maxY = point.Y
		}
	}

	return common.Point{minX, minY}, common.Point{maxX, maxY}
}

func getMDistance(source common.Point, target common.Point) int {
	return common.Abs(source.X-target.X) + common.Abs(source.Y-target.Y)

}

func getNearestPoint(source common.Point, targets []numberedPoint) string {
	distances := make(map[string]int)

	for _, targetPoint := range targets {
		distance := getMDistance(source, targetPoint.Point)
		if distance == 0 {
			return targetPoint.id
		}
		distances[targetPoint.id] = distance
	}

	min := -1

	for _, v := range distances {
		if min == -1 || v < min {
			min = v
		}
	}

	id := ""
	for k, v := range distances {
		if v == min {
			if id != "" {
				return "."
			}
			id = k
		}
	}

	return strings.ToLower(id)
}

func getPlane(points []numberedPoint) infinitePlane {
	min, max := getBounds(points)

	plane := make(infinitePlane)

	for x := min.X; x <= max.X; x++ {
		for y := min.Y; y <= max.Y; y++ {
			plane[common.Point{x, y}] = getNearestPoint(common.Point{x, y}, points)
		}
	}

	return plane
}

func countBorders(min common.Point, max common.Point, points []numberedPoint) map[string]int {
	counts := make(map[string]int)

	for i := min.X; i <= max.X; i++ {
		counts[strings.ToUpper(getNearestPoint(common.Point{i, min.Y}, points))]++
		counts[strings.ToUpper(getNearestPoint(common.Point{i, max.Y}, points))]++
	}

	for i := min.Y + 1; i < max.Y; i++ {
		counts[strings.ToUpper(getNearestPoint(common.Point{min.X, i}, points))]++
		counts[strings.ToUpper(getNearestPoint(common.Point{max.X, i}, points))]++
	}

	return counts
}

func printPlane(min common.Point, max common.Point, plane infinitePlane) {
	for x := min.X; x <= max.X; x++ {
		for y := min.Y; y <= max.Y; y++ {
			fmt.Print(plane[common.Point{y, x}])
		}
		fmt.Print("\n")
	}
}

// method count the id's on the edges. if the number is going down for any, it's not infinite.  when we have no more reductions.  we have hit the upper bounds,
// and can count the areas
func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	return fmt.Sprintf(" -- Not Yet Implemented --")
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	return fmt.Sprintf(" -- Not Yet Implemented --")
}
