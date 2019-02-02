package Day201915

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pezza/advent-of-code/puzzles/Common"
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
#######`)
	})

	It("Should be able to parse the puzzle input", func() {
		By("Actors in the right position")
		Expect(c[0][0].PosType).Should(Equal(1))
		Expect(c[1][1].PosType).Should(Equal(2))
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
#######`)

		mob := cave.getAllMobs()[0]
		_, acquired := cave.getTarget(mob)

		Expect(acquired).Should(Equal(false))
	})

	It("Should be able to get through narrow gaps", func() {
		cave := getData(`#######
#.....#
#E###G#
#..#..#
#######`)

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
#######`)

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
#######`)

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
#######`)

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
})
