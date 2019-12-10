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

	belt := getData(testB)

	printMap(belt)

	newBelt := getVisibleMap(common.Point{1, 8}, belt)

	printMap(newBelt)
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
