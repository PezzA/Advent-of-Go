package Day201903

import (
	"testing"

	"github.com/pezza/advent-of-code/common"

	. "github.com/onsi/gomega"
)

func Test_ReadData(t *testing.T) {
	RegisterTestingT(t)
	green, red := getData(Entry.PuzzleInput())
	Expect(green[1]).Should(Equal("U336"))
	Expect(red[1]).Should(Equal("U227"))

	Expect(getVector(string(green[1][0]))).Should(Equal(common.Point{X: 0, Y: 1}))
}

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	green, red := getData(`R8,U5,L5,D3
U7,R6,D4,L4`)
	c, s := runWires(green, red)
	Expect(c).Should(Equal(6))
	Expect(s).Should(Equal(30))

	green, red = getData(`R75,D30,R83,U83,L12,D49,R71,U7,L72
U62,R66,U55,R34,D71,R55,D58,R83`)
	c, s = runWires(green, red)
	Expect(c).Should(Equal(159))
	Expect(s).Should(Equal(610))

	green, red = getData(`R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51
U98,R91,D20,R16,D67,R40,U7,R15,U6,R7`)
	c, s = runWires(green, red)
	Expect(c).Should(Equal(135))
	Expect(s).Should(Equal(410))

	green, red = getData(Entry.PuzzleInput())
	c, s = runWires(green, red)

	// answers!
	Expect(c).Should(Equal(1195))
	Expect(s).Should(Equal(91518))
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
