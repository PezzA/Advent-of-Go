package Day202020

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

func Test_ReadData(t *testing.T) {
	RegisterTestingT(t)

	tiles := getData(Entry.PuzzleInput())

	Expect(tiles[0].id).Should(Equal(1579))
	Expect(tiles[1].id).Should(Equal(3413))

	fmt.Print(len(tiles))
}

func Test_Flip(t *testing.T) {
	RegisterTestingT(t)

	tiles := getData(Entry.PuzzleInput())

	tiles[0].printTile()
	fmt.Println()
	tiles[0].grid.flipLeftRight().printData()
	fmt.Println()
	tiles[0].grid.flipTopBottom().printData()
	fmt.Println()
	tiles[0].grid.rotate().printData()
}

func Test_matchingSides(t *testing.T) {
	RegisterTestingT(t)

	tiles := getData(Entry.PuzzleInput())

	tot := 1
	for _, t := range tiles {
		if len(t.matchingSides(tiles, [][]tile{})) == 2 {
			tot *= t.id
		}
	}

	fmt.Println(tot)
}

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)
	tiles := getData(Entry.PuzzleInput())
	tot := 1
	// don't have to put it all together, just need the corners
	for _, t := range tiles {
		if len(t.matchingSides(tiles, [][]tile{})) == 2 {
			fmt.Println(t.id)
			tot *= t.id
		}
	}
	fmt.Println(tot)
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

	tiles := getData(Entry.PuzzleInput())

	completedJig := solveJigSaw(tiles)

	for _, row := range completedJig {
		for _, tile := range row {
			fmt.Print(tile.id, " ")
		}
		fmt.Println()
	}
}

func Test_StripBorder(t *testing.T) {
	tiles := getData(Entry.PuzzleInput())

	completedJig := solveJigSaw(tiles)

	completedJig[0][0].grid.stripBorder().printData()

	grid := stitchTogether(completedJig)
	mask := getMonsterMask()
	cmp := grid.copyData()
	for i := 0; i < 4; i++ {
		cmp = cmp.rotate()

		fmt.Println(searchGrid(mask, cmp))
		cmpLR := cmp.flipLeftRight()
		fmt.Println(searchGrid(mask, cmpLR))
		cmpTB := cmp.flipTopBottom()
		fmt.Println(searchGrid(mask, cmpTB))

	}
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
