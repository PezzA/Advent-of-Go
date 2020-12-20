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
	printData(flipLeftRight(tiles[0].active))
	fmt.Println()
	printData(flipTopBottom(tiles[0].active))
	fmt.Println()
	printData(rotate(tiles[0].active))
}

func Test_matchingSides(t *testing.T) {
	RegisterTestingT(t)

	tiles := getData(Entry.PuzzleInput())

	tot := 1
	for _, tile := range tiles {
		fmt.Println(tile.matchingSides(tiles))
		if tile.matchingSides(tiles) == 2 {
			tot *= tile.id
		}

	}

	fmt.Println(tot)
}

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

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
