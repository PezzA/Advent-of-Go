package Day201806

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pezza/advent-of-code/puzzles/common"
)

var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string, int) {
	return 2018, 6, "Chronal Coordinates", 0
}

// end of scaffolding

type numberedPoint struct {
	id     string
	charid string
	common.Point
}

type infinitePlane map[common.Point]string
type distancePlane map[common.Point]int

type countMap map[string]int

func getData(inputData string) []numberedPoint {
	points := make([]numberedPoint, 0)

	startLabel := 65
	for index, line := range strings.Split(inputData, "\n") {
		point := numberedPoint{strconv.Itoa(index), string(index + 33), common.Point{}}
		fmt.Sscanf(line, "%d, %d", &point.X, &point.Y)
		startLabel++
		points = append(points, point)
	}
	return points
}

// given the point list, this will return the top left and bottom right bounds of all the points
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

func getCombinedPoint(source common.Point, targets []numberedPoint) int {
	combinedDistance := 0

	for _, targetPoint := range targets {
		combinedDistance += source.GetMDistance(targetPoint.Point)
	}
	return combinedDistance
}

func getNearestPoint(source common.Point, targets []numberedPoint) string {
	distances := make(countMap)

	for _, targetPoint := range targets {
		distance := source.GetMDistance(targetPoint.Point)
		if distance == 0 {
			return targetPoint.charid
		}
		distances[targetPoint.charid] = distance
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

	return id
}

func getPlane(points []numberedPoint, offSet int) infinitePlane {
	min, max := getBounds(points)

	newmin := common.Point{min.X - offSet, min.Y - offSet}
	newmax := common.Point{max.X + offSet, max.Y + offSet}

	plane := make(infinitePlane)

	for x := newmin.X; x <= newmax.X; x++ {
		for y := newmin.Y; y <= newmax.Y; y++ {
			plane[common.Point{x, y}] = getNearestPoint(common.Point{x, y}, points)
		}
	}

	return plane
}

func getCombinedPlane(points []numberedPoint, offSet int) distancePlane {
	min, max := getBounds(points)

	newmin := common.Point{min.X - offSet, min.Y - offSet}
	newmax := common.Point{max.X + offSet, max.Y + offSet}

	plane := make(distancePlane)

	for x := newmin.X; x <= newmax.X; x++ {
		for y := newmin.Y; y <= newmax.Y; y++ {
			plane[common.Point{x, y}] = getCombinedPoint(common.Point{x, y}, points)
		}
	}

	return plane
}

func getBorders(min common.Point, max common.Point, points []numberedPoint) infinitePlane {
	border := make(infinitePlane)

	for i := min.X; i <= max.X; i++ {
		border[common.Point{i, min.Y}] = getNearestPoint(common.Point{i, min.Y}, points)
		border[common.Point{i, max.Y}] = getNearestPoint(common.Point{i, max.Y}, points)
	}

	for i := min.Y + 1; i < max.Y; i++ {
		border[common.Point{min.X, i}] = getNearestPoint(common.Point{min.X, i}, points)
		border[common.Point{max.X, i}] = getNearestPoint(common.Point{max.X, i}, points)
	}

	return border
}

func getCombinedBorders(min common.Point, max common.Point, points []numberedPoint) distancePlane {
	border := make(distancePlane)

	for i := min.X; i <= max.X; i++ {
		border[common.Point{i, min.Y}] = getCombinedPoint(common.Point{i, min.Y}, points)
		border[common.Point{i, max.Y}] = getCombinedPoint(common.Point{i, max.Y}, points)
	}

	for i := min.Y + 1; i < max.Y; i++ {
		border[common.Point{min.X, i}] = getCombinedPoint(common.Point{min.X, i}, points)
		border[common.Point{max.X, i}] = getCombinedPoint(common.Point{max.X, i}, points)
	}

	return border
}

func getCounts(plane infinitePlane) countMap {
	counts := make(countMap)
	for _, v := range plane {
		counts[v]++
	}
	return counts
}

func getNonIncludedPoints(counts countMap, points []numberedPoint) map[string]bool {

	returnPoints := make(map[string]bool)

	for _, val := range points {

		if _, ok := counts[val.charid]; !ok {
			returnPoints[val.charid] = true
		}
	}

	return returnPoints
}

func hasDecreasing(prev countMap, new countMap) bool {
	// for each item int the prev map, it should be smaller or non existant in new map

	for k, prevValue := range prev {
		if newValue, ok := new[k]; ok {
			if prevValue > newValue {
				return true
			}
		} else {
			return true
		}
	}
	return false
}

func printPlane(min common.Point, max common.Point, plane infinitePlane) {
	for x := min.X; x <= max.X; x++ {
		for y := min.Y; y <= max.Y; y++ {

			if val, ok := plane[common.Point{x, y}]; ok {
				fmt.Print(val)
			} else {
				fmt.Print("#")
			}
		}
		fmt.Print("\n")
	}
}

// method count the id's on the edges. if the number is going down for any, it's not infinite.  when we have no more reductions.  we have hit the upper bounds,
// and can count the areas
func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	pointList := getData(Entry.PuzzleInput())

	min, max := getBounds(pointList)

	hitExtremety := false

	count := 0

	// keep drawing increasing borders until none of the items are getting smaller

	var lastCounts countMap
	for !hitExtremety {
		incMin := common.Point{min.X - count, min.Y - count}
		incMax := common.Point{max.X + count, max.Y + count}

		counts := getCounts(getBorders(incMin, incMax, pointList))
		preCounts := counts

		hitExtremety = !hasDecreasing(preCounts, counts)

		preCounts = counts
		lastCounts = counts
		count++
	}

	nonInfinites := getNonIncludedPoints(lastCounts, pointList)
	plane := getPlane(pointList, count)
	finalCounts := getCounts(plane)

	areaMax := -1
	for key, val := range finalCounts {

		if _, ok := nonInfinites[key]; !ok {
			continue
		}
		if areaMax == -1 || val > areaMax {
			areaMax = val
		}
	}

	return fmt.Sprintf("%v", areaMax)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	pointList := getData(Entry.PuzzleInput())

	min, max := getBounds(pointList)

	hitExtremety := false

	count := 0

	// keep drawing increasing borders all of them are greater or equal to 10,000

	for !hitExtremety {
		incMin := common.Point{min.X - count, min.Y - count}
		incMax := common.Point{max.X + count, max.Y + count}

		counts := getCombinedBorders(incMin, incMax, pointList)
		hitExtremety = true
		for _, v := range counts {
			if v < 10000 {
				hitExtremety = false
				break
			}
		}

		count++
	}

	finalPlane := getCombinedPlane(pointList, count)

	finalCount := 0
	for _, val := range finalPlane {
		if val < 10000 {
			finalCount++
		}
	}
	return fmt.Sprintf("%v", finalCount)
}
