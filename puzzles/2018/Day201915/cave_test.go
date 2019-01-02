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
})
