package Day201915

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pezza/advent-of-code/puzzles/Common"
)

var _ = Describe("Cave", func() {
	cave := getData(`#######
#E..G.#
#...#.#
#.G.#G#
#######`)

	It("Should be able to parse the puzzle input", func() {
		By("Actors in the right position")
		Expect(cave[0][0].PosType).Should(Equal(1))
		Expect(cave[1][1].PosType).Should(Equal(2))
		Expect(cave[1][1].mob).Should(Equal(mob{0, "E", common.Point{1, 1}, 200, 3}))

		By("Should contain 4 mobs")
		Expect(len(cave.getAllMobs())).Should(Equal(4))

		By("One of which is an elf")
		Expect(len(cave.getFaction("E"))).Should(Equal(1))

		By("And the other three are goblins")
		Expect(len(cave.getFaction("G"))).Should(Equal(3))
	})
})
