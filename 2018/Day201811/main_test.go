package Day201811

import (
	"testing"

	"fmt"

	. "github.com/onsi/gomega"
	"github.com/pezza/advent-of-code/Common"
)

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	Expect(getPowerLevel(common.Point{3, 5}, 8)).Should(Equal(4))

	Expect(getPowerLevel(common.Point{122, 79}, 57)).Should(Equal(-5))
	Expect(getPowerLevel(common.Point{101, 153}, 71)).Should(Equal(4))
	Expect(getPowerLevel(common.Point{217, 196}, 39)).Should(Equal(0))

	gridSerial := getData(Entry.PuzzleInput())
	maxPower := -1
	var powerPoint common.Point
	for x := 1; x <= 298; x++ {
		for y := 1; y <= 298; y++ {
			testPoint := common.Point{x, y}
			squarePower := getPowerLevel(common.Point{x, y}, gridSerial)
			squarePower += getPowerLevel(common.Point{x + 1, y}, gridSerial)
			squarePower += getPowerLevel(common.Point{x + 2, y}, gridSerial)
			squarePower += getPowerLevel(common.Point{x, y + 1}, gridSerial)
			squarePower += getPowerLevel(common.Point{x + 1, y + 1}, gridSerial)
			squarePower += getPowerLevel(common.Point{x + 2, y + 1}, gridSerial)
			squarePower += getPowerLevel(common.Point{x, y + 2}, gridSerial)
			squarePower += getPowerLevel(common.Point{x + 1, y + 2}, gridSerial)
			squarePower += getPowerLevel(common.Point{x + 2, y + 2}, gridSerial)

			if maxPower == -1 || squarePower > maxPower {
				maxPower = squarePower
				powerPoint = testPoint
			}
			fmt.Println("computing", x, y)
		}
	}

	fmt.Println(powerPoint)
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

	gridSerial := getData(Entry.PuzzleInput())
	maxPower := -1
	var powerPoint common.Point
	squareSize := -1

	computedGrid := make([][]int, 300)

	for index := range computedGrid {
		computedGrid[index] = make([]int, 300)
	}

	for index := range computedGrid {
		for subIndex := range computedGrid[index] {
			computedGrid[index][subIndex] = getPowerLevel(common.Point{subIndex + 1, index + 1}, gridSerial)
		}
	}

	fmt.Println(computedGrid)
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
				for zx := x; zx < z; zx++ {
					for zy := y; zy < z; zy++ {

						squareCount += computedGrid[zx-1][zy-1]
					}
				}

				if maxPower == -1 || squareCount > maxPower {
					fmt.Println("setting max", maxPower)
					maxPower = squareCount
					powerPoint = common.Point{x, y}
					squareSize = z
				}
			}
		}
		fmt.Println("size ", z, offSet, count)
	}
	// not 20,21,31 // not 21,20,31
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
