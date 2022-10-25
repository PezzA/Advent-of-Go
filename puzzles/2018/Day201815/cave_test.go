package Day201815

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pezza/advent-of-code/puzzles/common"
)

var _ = Describe("Cave", func() {

	var (
		c cave
	)

	BeforeEach(func() {
		c = getData(`#######
#E..G.#
#...#.#
#.G.#G#
#######`, baseAttack)
	})

	It("Should be able to parse the puzzle input", func() {
		By("Actors in the right position")
		Expect(c[0][0].PosType).Should(Equal(wall))
		Expect(c[1][1].PosType).Should(Equal(occupied))
		Expect(c[1][1].mob).Should(Equal(mob{0, "E", common.Point{X: 1, Y: 1}, 200, 3}))

		By("Should contain 4 mobs")
		Expect(len(c.getAllMobs())).Should(Equal(4))

		By("One of which is an elf")
		Expect(len(c.getFaction("E"))).Should(Equal(1))

		By("And the other three are goblins")
		Expect(len(c.getFaction("G"))).Should(Equal(3))
	})

	It("Should be able to find a target for a mob", func() {
		mob := c.getAllMobs()[0]
		target, acquired := c.getTarget(mob)
		Expect(acquired).Should(Equal(true))
		Expect(target).Should(Equal(common.Point{X: 3, Y: 1}))
	})

	It("Should be able to detect that is has no valid targets", func() {
		cave := getData(`#######
#..#..#
#E.#.G#
#..#..#
#######`, baseAttack)

		mob := cave.getAllMobs()[0]
		_, acquired := cave.getTarget(mob)

		Expect(acquired).Should(Equal(false))
	})

	It("Should be able to get through narrow gaps", func() {
		cave := getData(`#######
#.....#
#E###G#
#..#..#
#######`, baseAttack)

		mob := cave.getAllMobs()[0]
		target, acquired := cave.getTarget(mob)

		Expect(acquired).Should(Equal(true))
		Expect(target).Should(Equal(common.Point{X: 5, Y: 1}))
	})

	It("Should be able to create a blank distance map with the same dimensions", func() {
		cave := getData(`#######
#..G..#
#.....#
#G.E.G#
#.....#
#..G..#
#######`, baseAttack)

		Expect(len(cave)).Should(Equal(7))

		for index := range cave {
			Expect(len(cave[index])).Should(Equal(7))
		}

		blankMap := cave.getBlankDistanceMap()

		Expect(len(blankMap)).Should(Equal(7))

		for index := range blankMap {
			Expect(len(blankMap[index])).Should(Equal(7))
		}
	})

	It("Should pick the target highest in reading order", func() {
		cave := getData(`#######
#..G..#
#.....#
#G.E.G#
#.....#
#..G..#
#######`, baseAttack)

		mob := cave.getFaction("E")[0]
		target, acquired := cave.getTarget(mob)

		Expect(acquired).Should(Equal(true))
		Expect(target).Should(Equal(common.Point{X: 3, Y: 2}))
	})

	It("Should take the right first starting step", func() {
		cave := getData(`#######
#E....#
#.....#
#.....#
#.....#
#....G#
#######`, baseAttack)

		By("Acquiring the right target")

		mob := cave.getFaction("E")[0]
		target, acquired := cave.getTarget(mob)

		Expect(acquired).Should(Equal(true))
		Expect(target).Should(Equal(common.Point{X: 5, Y: 4}))

		By("It should be able to get the distance map for the target")

		targetDMap := cave.getDistanceMap(target)

		Expect(targetDMap[1][2]).Should(Equal(6))
		Expect(targetDMap[2][1]).Should(Equal(6))

		By("It should be able to take the first right step")

		nextStep := cave.getNextStep(mob, target)
		Expect(nextStep).Should(Equal(common.Point{X: 2, Y: 1}))
	})

	It("Should be able to detect an adjacent enemy", func() {
		cave := getData(`#######
#.....#
#.....#
#.E.G.#
#.....#
#.....#
#######`, baseAttack)

		Expect(cave.getAdjacentEnemy(cave.getAllMobs()[0])).Should(Equal(false))

		cave = getData(`#######
#.....#
#.E...#
#.E.G.#
#.....#
#.....#
#######`, baseAttack)

		Expect(cave.getAdjacentEnemy(cave.getAllMobs()[0])).Should(Equal(false))

		cave = getData(`#######
#GGGGG#
#G...G#
#G.E.G#
#G...G#
#GGGGG#
#######`, baseAttack)

		hasEnemy, _ := cave.getAdjacentEnemy(cave.getAllMobs()[8])
		Expect(hasEnemy).Should(Equal(false))

		cave = getData(`#######
#.....#
#.....#
#.GE..#
#.....#
#.....#
#######`, baseAttack)

		hasEnemy, enemy := cave.getAdjacentEnemy(cave.getAllMobs()[1])
		Expect(hasEnemy).Should(Equal(true))
		Expect(enemy.Point).Should(Equal(common.Point{X: 2, Y: 3}))

		cave = getData(`#######
#.....#
#..G..#
#.GEG.#
#..G..#
#.....#
#######`, baseAttack)

		hasEnemy, enemy = cave.getAdjacentEnemy(cave.getAllMobs()[2])
		Expect(hasEnemy).Should(Equal(true))
		Expect(enemy.Point).Should(Equal(common.Point{X: 3, Y: 2}))

		hasEnemy, enemy = cave.getAdjacentEnemy(cave.getAllMobs()[0])
		Expect(hasEnemy).Should(Equal(true))
		Expect(enemy.Point).Should(Equal(common.Point{X: 3, Y: 3}))
	})

	It("Should be able to run battles!", func() {
		var battleTests = []struct {
			c              cave
			expectedRounds int
			expectedHp     int
			debug          bool
		}{
			{getData(`#######   
#.G...#
#...EG#
#.#.#G#
#..G#E#
#.....#
#######`, baseAttack), 47, 590, false},

			{getData(`#######
#G..#E#
#E#E.E#
#G.##.#
#...#E#
#...E.#
#######`, baseAttack), 37, 982, false},

			{getData(`#######
#E..EG#
#.#G.E#
#E.##E#
#G..#.#
#..E#.#
#######`, baseAttack), 46, 859, false},

			{getData(`#######
#E.G#.#
#.#G..#
#G.#.G#
#G..#.#
#...E.#
#######`, baseAttack), 35, 793, false},

			{getData(`#######
#.E...#
#.#..G#
#.###.#
#E#G#G#
#...#G#
#######`, baseAttack), 54, 536, false},

			{getData(`#########
#G......#
#.E.#...#
#..##..G#
#...##..#
#...#...#
#.G...G.#
#.....G.#
#########`, baseAttack), 20, 937, false},

			// Reddit based Tests cases, taken from https://github.com/ShaneMcC/aoc-2018/blob/master/15/tests
			{getData(`#####
#GG##
#.###
#..E#
#.#G#
#.E##
#####`, baseAttack), 71, 197, false},
		}

		for _, caveTest := range battleTests {
			rounds, c := caveTest.c.resolveBattle(caveTest.debug)
			Expect(rounds).Should(Equal(caveTest.expectedRounds))
			Expect(c.getRemainingHp()).Should(Equal(caveTest.expectedHp))
		}
	})

	It("Should be able to find the winning point for elves", func() {
		var battleTests = []struct {
			c              string
			expectedRounds int
			expectedHp     int
			expectedAttack int
		}{{`#######
#.G...#
#...EG#
#.#.#G#
#..G#E#
#.....#
#######`, 29, 172, 15},

			{`#######
#E..EG#
#.#G.E#
#E.##E#
#G..#.#
#..E#.#
#######`, 33, 948, 4},

			{`#######
#E.G#.#
#.#G..#
#G.#.G#
#G..#.#
#...E.#
#######`, 37, 94, 15},

			{`#######
#.E...#
#.#..G#
#.###.#
#E#G#G#
#...#G#
#######`, 39, 166, 12},

			{`#########
#G......#
#.E.#...#
#..##..G#
#...##..#
#...#...#
#.G...G.#
#.....G.#
#########`, 30, 38, 34},

			// Reddit based Tests cases, taken from https://github.com/ShaneMcC/aoc-2018/blob/master/15/tests

			{`####
##E#
#GG#
####`, 58, 29, 7},

			{`#################
##..............#
##........G.....#
####.....G....###
#....##......####
#...............#
##........GG....#
##.........E..#.#
#####.###...#####
#################`, 32, 11, 25},

			{`#####
#GG##
#.###
#..E#
#.#G#
#.E##
#####`, 68, 103, 6},

			{`#####
##########
#.E....G.#
#......###
#.G......#
##########`, 46, 8, 9},
		}

		for _, caveTest := range battleTests {
			attack, rounds, c := escalateBattle(caveTest.c)
			Expect(rounds).Should(Equal(caveTest.expectedRounds))
			Expect(attack).Should(Equal(caveTest.expectedAttack))
			Expect(c.getRemainingHp()).Should(Equal(caveTest.expectedHp))
		}
	})
})
