package Day201910

import (
	"fmt"
	"testing"

	"github.com/pezza/advent-of-code/common"

	. "github.com/onsi/gomega"
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

var testD = `.#..##.###...#######
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

	belt := getData(testA)

	Expect(belt[0]).Should(Equal([]bool{false, true, false, false, true}))
}

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	Expect(getStep(common.Point{X: 1, Y: 0}, common.Point{X: 4, Y: 2})).
		Should(Equal(common.Point{X: 3, Y: 2}))

	Expect(getCommonLowestFactor(5, 10)).Should(Equal(5))
	Expect(getCommonLowestFactor(2, 8)).Should(Equal(2))

	Expect(getStep(common.Point{X: 0, Y: 0}, common.Point{X: 2, Y: 8})).
		Should(Equal(common.Point{X: 1, Y: 4}))

	Expect(getStep(common.Point{X: 9, Y: 9}, common.Point{X: 13, Y: 13})).
		Should(Equal(common.Point{X: 1, Y: 1}))

	Expect(getStep(common.Point{X: 9, Y: 9}, common.Point{X: 13, Y: 13})).
		Should(Equal(common.Point{X: 1, Y: 1}))

	belt := getData(testC)
	/*
		Expect(isVisible(common.Point{X: 3, Y: 4}, common.Point{X: 1, Y: 0}, belt)).
			Should(Equal(false))

		Expect(isVisible(common.Point{X: 4, Y: 0}, common.Point{X: 4, Y: 2}, belt)).
			Should(Equal(true))

		Expect(isVisible(common.Point{X: 4, Y: 0}, common.Point{X: 4, Y: 3}, belt)).
			Should(Equal(false))
	*/
	best := 0
	for y := 0; y < len(belt); y++ {
		for x := 0; x < len(belt[y]); x++ {

			if !belt[y][x] {
				//		fmt.Print(",")
				continue
			}

			count := getVisibleCount(common.Point{X: x, Y: y}, belt)
			//	fmt.Print(count)
			if count > best {
				best = count
			}
		}
		//	fmt.Print("\n")
	}

	fmt.Println(best)
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
