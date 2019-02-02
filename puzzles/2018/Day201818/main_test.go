package Day201818

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	land := getData(`.#.#...|#.
.....#|##|
.|..|...#.
..|#.....#
#.#|||#|#|
...#.||...
.|....|...
||...#|.#|
|.||||..|.
...#.|..|.`)

	/*open, tree, lumber := land.getCounts(common.Point{X: 0, Y: 0})
	Expect(open).Should(Equal(2))
	Expect(tree).Should(Equal(1))
	Expect(lumber).Should(Equal(1))

	open, tree, lumber = land.getCounts(common.Point{X: 1, Y: 1})
	Expect(open).Should(Equal(5))
	Expect(tree).Should(Equal(3))
	Expect(lumber).Should(Equal(1))
	*/
	for i := 0; i < 10; i++ {
		land = land.transform()
	}

	Expect(land.getTypeCount(treeArea) * land.getTypeCount(lumberYard)).Should(Equal(1147))

	l1 := getData(Entry.PuzzleInput())
	l2 := getData(Entry.PuzzleInput())
	Expect(l1.sameAs(l2)).Should(Equal(true))

	l1[3][3] = 4
	Expect(l1.sameAs(l2)).Should(Equal(false))
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

	land := getData(Entry.PuzzleInput())
	lands := []acreage{land}

	for i := 1; i <= 1000000000; i++ {
		land = land.transform()

		if ok, foundIndex := contains(lands, land); ok {
			loopSpan := 1000000000 - i
			loopLength := i - foundIndex

			fmt.Println(foundIndex, foundIndex+(loopSpan%loopLength))

			fmt.Println(lands[foundIndex+(loopSpan%loopLength)].getTypeCount(treeArea) * lands[foundIndex+(loopSpan%loopLength)].getTypeCount(lumberYard))

			break
		} else {
			lands = append(lands, land)
		}
	}

	fmt.Println("not 197938,205737")

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
