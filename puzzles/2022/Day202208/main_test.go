package Day202208

import (
	"testing"

	. "github.com/onsi/gomega"
	"github.com/pezza/advent-of-code/puzzles/common"
)

const testData string = `30373
25512
65332
33549
35390`

func Test_ReadData(t *testing.T) {
	RegisterTestingT(t)

	grid, width, height := getData(Entry.PuzzleInput())

	Expect(grid[0][0]).Should(Equal(2))
	Expect(grid[0][width-1]).Should(Equal(3))
	Expect(grid[height-1][0]).Should(Equal(2))
	Expect(grid[height-1][width-1]).Should(Equal(3))

}

func Test_TestScenicScore(t *testing.T) {
	RegisterTestingT(t)
	grid, _, _ := getData(testData)

	Expect(getScenicScore(grid, common.Point{2, 1})).Should(Equal(4))
	Expect(getScenicScore(grid, common.Point{2, 3})).Should(Equal(8))
}

func Test_TestFromUp(t *testing.T) {
	RegisterTestingT(t)
	grid, _, _ := getData(testData)
	canSee := canViewFromUp(grid, common.Point{2, 1})

	Expect(canSee).Should(Equal(true))
}

func Test_CountTrees(t *testing.T) {
	RegisterTestingT(t)

	grid, _, _ := getData(testData)

	Expect(getTreeCount(grid)).Should(Equal(21))
}

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	Expect(Entry.PartOne(Entry.PuzzleInput(), nil)).Should(Equal("1681"))
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

	Expect(Entry.PartTwo(Entry.PuzzleInput(), nil)).Should(Equal("201684"))
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
