package Day201806

import (
	"fmt"
	"testing"

	"github.com/pezza/advent-of-code/Common"

	. "github.com/onsi/gomega"
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

	Expect(pointList[0]).Should(Equal(numberedPoint{"A", common.Point{1, 1}}))
	Expect(pointList[1]).Should(Equal(numberedPoint{"B", common.Point{1, 6}}))

	min, max := getBounds(pointList)
	Expect(min).Should(Equal(common.Point{1, 1}))
	Expect(max).Should(Equal(common.Point{8, 9}))

	Expect(getMDistance(common.Point{-3, -3}, common.Point{3, 3})).Should(Equal(12))
	Expect(getMDistance(common.Point{3, 3}, common.Point{-3, -3})).Should(Equal(12))

	testPoint := common.Point{0, 0}
	Expect(getNearestPoint(testPoint, pointList)).Should(Equal("a"))

	testPoint = common.Point{1, 1}
	Expect(getNearestPoint(testPoint, pointList)).Should(Equal("A"))
	testPoint = common.Point{10, 10}
	Expect(getNearestPoint(testPoint, pointList)).Should(Equal("f"))

	testPoint = common.Point{5, 0}
	Expect(getNearestPoint(testPoint, pointList)).Should(Equal("."))

}

func Test_PrintGrid(t *testing.T) {
	pointList := getData(testInput)

	min, max := getBounds(pointList)

	//printPlane(min, max, getPlane(pointList))

	fmt.Println(countBorders(min, max, pointList))
	fmt.Println(countBorders(common.Point{min.X - 1, min.Y - 1}, common.Point{max.X + 1, max.Y + 1}, pointList))

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
