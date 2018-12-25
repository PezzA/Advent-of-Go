package Day201811

import (
	"testing"

	"fmt"

	. "github.com/onsi/gomega"
	"github.com/pezza/advent-of-code/puzzles/Common"
)

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	Expect(getPowerLevel(common.Point{3, 5}, 8)).Should(Equal(4))

	Expect(getPowerLevel(common.Point{122, 79}, 57)).Should(Equal(-5))
	Expect(getPowerLevel(common.Point{101, 153}, 71)).Should(Equal(4))
	Expect(getPowerLevel(common.Point{217, 196}, 39)).Should(Equal(0))

	gridSerial := getData(Entry.PuzzleInput())

	computedGrid := make([][]int, 300)

	for index := range computedGrid {
		computedGrid[index] = make([]int, 300)
	}

	for index := range computedGrid {
		for subIndex := range computedGrid[index] {
			computedGrid[index][subIndex] = getPowerLevel(common.Point{index + 1, subIndex + 1}, gridSerial)
		}
	}

	maxPower := -1
	var powerPoint common.Point
	for x := 1; x <= 298; x++ {
		for y := 1; y <= 298; y++ {
			testPoint := common.Point{x, y}
			squarePower := computedGrid[x-1][y-1]
			squarePower += computedGrid[x][y-1]
			squarePower += computedGrid[x+1][y-1]
			squarePower += computedGrid[x-1][y]
			squarePower += computedGrid[x][y]
			squarePower += computedGrid[x+1][y]
			squarePower += computedGrid[x-1][y+1]
			squarePower += computedGrid[x][y+1]
			squarePower += computedGrid[x+1][y+1]

			if maxPower == -1 || squarePower > maxPower {
				maxPower = squarePower
				powerPoint = testPoint
			}
		}
	}

	fmt.Println(powerPoint)
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

	gridSerial := getData(Entry.PuzzleInput())
	//gridSerial := 18
	maxPower := -1
	var powerPoint common.Point
	squareSize := -1

	computedGrid := make([][]int, 300)

	for index := range computedGrid {
		computedGrid[index] = make([]int, 300)
	}

	for index := range computedGrid {
		for subIndex := range computedGrid[index] {
			computedGrid[index][subIndex] = getPowerLevel(common.Point{index + 1, subIndex + 1}, gridSerial)
		}
	}
	//for each size
	for z := 1; z <= 300; z++ {
		count := 0
		offSet := 301 - z

		// for each instance of that size
		for x := 1; x <= offSet; x++ {
			for y := 1; y <= offSet; y++ {
				count++
				squareCount := 0

				// for each cell in that size
				for zx := x; zx < x+z; zx++ {
					for zy := y; zy < y+z; zy++ {

						squareCount += computedGrid[zx-1][zy-1]

					}
				}

				if maxPower == -1 || squareCount > maxPower {
					maxPower = squareCount
					powerPoint = common.Point{x, y}
					squareSize = z
				}
			}
		}
		fmt.Println(z, maxPower)

	}
	// not 20,21,31 // not 21,20,31 // not 22,21,32
	fmt.Println(powerPoint, squareSize)
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
