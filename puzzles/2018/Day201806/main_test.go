package Day201806

import (
	"testing"

	"fmt"

	. "github.com/onsi/gomega"
	"github.com/pezza/advent-of-code/puzzles/common"
)

var testInput = `1, 1
1, 6
8, 3
3, 4
5, 5
8, 9`

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	pointList := getData(testInput)

	Expect(pointList[0]).Should(Equal(numberedPoint{"A", "A", common.Point{X: 1, Y: 1}}))
	Expect(pointList[1]).Should(Equal(numberedPoint{"B", "A", common.Point{X: 1, Y: 6}}))

	min, max := getBounds(pointList)
	Expect(min).Should(Equal(common.Point{X: 1, Y: 1}))
	Expect(max).Should(Equal(common.Point{X: 8, Y: 9}))

	Expect(common.Point{X: -3, Y: -3}.GetMDistance(common.Point{X: 3, Y: 3})).Should(Equal(12))
	Expect(common.Point{X: 3, Y: 3}.GetMDistance(common.Point{X: -3, Y: -3})).Should(Equal(12))

	testPoint := common.Point{X: 0, Y: 0}
	Expect(getNearestPoint(testPoint, pointList)).Should(Equal("a"))

	testPoint = common.Point{X: 1, Y: 1}
	Expect(getNearestPoint(testPoint, pointList)).Should(Equal("A"))
	testPoint = common.Point{X: 10, Y: 10}
	Expect(getNearestPoint(testPoint, pointList)).Should(Equal("f"))

	testPoint = common.Point{X: 5, Y: 0}
	Expect(getNearestPoint(testPoint, pointList)).Should(Equal("."))

}

func Test_PrintGrid(t *testing.T) {
	pointList := getData(Entry.PuzzleInput())

	min, max := getBounds(pointList)

	//printPlane(min, max, getPlane(pointList, 0))

	hitExtremety := false

	count := 0

	// keep drawing increasing borders until none of the items are getting smaller

	var lastCounts countMap
	for !hitExtremety {
		incMin := common.Point{X: min.X - count, Y: min.Y - count}
		incMax := common.Point{X: max.X + count, Y: max.Y + count}

		counts := getCounts(getBorders(incMin, incMax, pointList))
		preCounts := counts

		hitExtremety = !hasDecreasing(preCounts, counts)

		preCounts = counts
		lastCounts = counts
		count++
	}

	nonInfinites := getNonIncludedPoints(lastCounts, pointList)

	plane := getPlane(pointList, count)

	printPlane(min, max, plane)
	finalCounts := getCounts(plane)

	fmt.Println(finalCounts)
	areaMax := -1
	for key, val := range finalCounts {

		if _, ok := nonInfinites[key]; !ok {
			continue
		}
		if areaMax == -1 || val > areaMax {
			areaMax = val
		}
	}

	fmt.Println(areaMax)

}
func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)
}

func Benchmark_BenchPartOne(b *testing.B) {
	data := Entry.PuzzleInput()
	for n := 0; n < b.N; n++ {
		Entry.PartOne(data, nil)
	}
}

func Benchmark_BenchPartTwo(b *testing.B) {
	data := Entry.PuzzleInput()
	for n := 0; n < b.N; n++ {
		Entry.PartTwo(data, nil)
	}
}
