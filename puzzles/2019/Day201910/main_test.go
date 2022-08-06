package Day201910

import (
	"testing"

	. "github.com/onsi/gomega"
	"github.com/pezza/advent-of-code/common"
)

var testA = `.#..#
.....
#####
....#
...##`

var testB = `......#.#.
#..#.#....
..#######.
.#.#.###..
.#..#.....
..#....#.#
#..#....#.
.##.#..###
##...#..#.
.#....####`

var testC = `#.#...#.#.
.###....#.
.#....#...
##.#.#.#.#
....#.#.#.
.##..###.#
..#...##..
..##....##
......#...
.####.###.`

var testD = `.#..#..###
####.###.#
....###.#.
..###.##.#
##.##.#.#.
....###..#
..#.#..#.#
#..#.#.###
.##...##.#
.....#.#..`

var testE = `.#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##`

func Test_ReadData(t *testing.T) {
	RegisterTestingT(t)

	//belt := getData(testA)

	//fmt.Println(belt)

	//copyBelt := copyBelt(belt)
	//fmt.Println(copyBelt)
}

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	count, pos := getBestAsteriod(getData(testA))
	Expect(count).Should(Equal(8))
	Expect(pos).Should(Equal(common.Point{X: 3, Y: 4}))

	count, pos = getBestAsteriod(getData(testB))
	Expect(count).Should(Equal(33))
	Expect(pos).Should(Equal(common.Point{X: 5, Y: 8}))

	count, pos = getBestAsteriod(getData(testC))
	Expect(count).Should(Equal(35))
	Expect(pos).Should(Equal(common.Point{X: 1, Y: 2}))

	count, pos = getBestAsteriod(getData(testD))
	Expect(count).Should(Equal(41))
	Expect(pos).Should(Equal(common.Point{X: 6, Y: 3}))

	count, pos = getBestAsteriod(getData(testE))
	Expect(count).Should(Equal(210))
	Expect(pos).Should(Equal(common.Point{X: 11, Y: 13}))

	Expect(Entry.PartOne(Entry.PuzzleInput(), nil)).Should(Equal("260"))

}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)
	Expect(Entry.PartTwo(Entry.PuzzleInput(), nil)).Should(Equal("608"))
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
