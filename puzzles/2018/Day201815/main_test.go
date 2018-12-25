package Day201815

import (
	"testing"

	"fmt"

	. "github.com/onsi/gomega"
	"github.com/pezza/advent-of-code/puzzles/Common"
)

var moveTest = `#########
#G..G..G#
#.......#
#.......#
#G..E..G#
#.......#
#.......#
#G..G..G#
#########`

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)
	cave := getData(moveTest)

	mob := cave.getAllMobs()[0]

	nextTarget, found := cave.getTarget(mob)
	Expect(found).Should(Equal(true))
	Expect(nextTarget).Should(Equal(common.Point{4, 3}))

	moveLoc := cave.getDistanceMap(nextTarget).filterAdjacent(mob.Point).getTop()

	Expect(moveLoc).Should(Equal(common.Point{2, 1}))
}

func Test_Move(t *testing.T) {
	RegisterTestingT(t)
	cave := getData(moveTest)
	cave.draw(nil)
	cave, _, _ = cave.runRound()
	cave.draw(nil)
	cave, _, _ = cave.runRound()
	cave.draw(nil)
	cave, _, _ = cave.runRound()
	cave.draw(nil)
}

var b0 = `#######
#.G...#
#...EG#
#.#.#G#
#..G#E#
#.....#
#######`

var b1 = `#######
#G..#E#
#E#E.E#
#G.##.#
#...#E#
#...E.#
#######`

var b2 = `#######   
#E..EG#
#.#G.E#
#E.##E#
#G..#.#
#..E#.#   
#######`

var b3 = `#######   
#E.G#.#
#.#G..#
#G.#.G#   
#G..#.#
#...E.#
#######`

var b4 = `#######   
#.E...#   
#.#..G#
#.###.#   
#E#G#G#   
#...#G#
#######`

var b5 = `#########   
#G......#
#.E.#...#
#..##..G#
#...##..#   
#...#...#
#.G...G.#   
#.....G.#   
#########`

var real = Entry.PuzzleInput()

func Test_Battlefield1(t *testing.T) {
	RegisterTestingT(t)
	Expect(runBattle(b0)).Should(Equal(27730))
	Expect(runBattle(b1)).Should(Equal(36334))
	Expect(runBattle(b2)).Should(Equal(39514))
	Expect(runBattle(b3)).Should(Equal(27755))
	Expect(runBattle(b4)).Should(Equal(28944))
	Expect(runBattle(b5)).Should(Equal(18740))
	Expect(runBattle(real)).Should(Not(Equal(261036)))
	Expect(runBattle(real)).Should(Not(Equal(263780)))

	fmt.Println("not 261036 (too low). not 263780")
}

func Test_m1(t *testing.T) {
	RegisterTestingT(t)

	cave := getData(`#######
#######
#.E..G#
#.#####
#G#####
#######
#######`)

	cave.runRound()
	cave.draw(nil)

}

func Test_m2(t *testing.T) {
	RegisterTestingT(t)

	cave := getData(`####
#GG#
#.E#
####`)

	cave.runRound()
	cave.draw(nil)

}

func Test_m3(t *testing.T) {
	RegisterTestingT(t)

	cave := getData(`########
#..E..G#
#G######
########`)

	cave.runRound()
	cave.draw(nil)
}
func Test_m4(t *testing.T) {
	RegisterTestingT(t)

	cave := getData(`#######
#.E..G#
#.....#
#G....#
#######`)

	cave.runRound()
	cave.draw(nil)
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
