package Day201919

import (
	"fmt"
	"testing"

	"github.com/pezza/advent-of-code/common"

	. "github.com/onsi/gomega"
)

func Test_ReadData(t *testing.T) {
	RegisterTestingT(t)

}

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	count := 0

	for y := 0; y < 500; y++ {
		for x := 0; x < 500; x++ {
			val := scanPoint(Entry.PuzzleInput(), int64(x), int64(y))
			if val == 1 {
				count++
				fmt.Print("#")
				continue
			}
			fmt.Print(".")
		}
		fmt.Print("\n")
	}
	Expect(count).Should(Equal(186))
}

type lineRange struct {
	start int
	end   int
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

	beamArea := make(map[common.Point]bool, 0)

	mainY := 0

	limits := make(map[int]lineRange, 0)

	for {
		start, end := 0, 0
		for x := 0; x < mainY; x++ {
			val := scanPoint(Entry.PuzzleInput(), int64(x), int64(mainY))

			if val == 1 {
				if start == 0 {
					start = x
				} else {
					beamArea[common.Point{X: x, Y: mainY}] = true
					continue
				}
			} else {
				if start != 0 && end == 0 {
					end = x
				}
			}
		}

		limits[mainY] = lineRange{start, end}

		if end-start >= 100 {
			break
		}

		fmt.Println("seedscan", mainY, end-start)

		mainY++
	}

	offsetY := mainY - 110

	var pos common.Point
	for {
		// check to see if we have a 100x100
		hasFound := false

		fmt.Println("checking range for line", offsetY, ":", limits[offsetY])

		for x := limits[offsetY].start; x < limits[offsetY].end; x++ {
			if beamArea[common.Point{X: x, Y: offsetY}] {
				if checkDepth(x, offsetY, beamArea) {
					pos = common.Point{X: x, Y: offsetY}
					hasFound = true
					break
				}

			}
		}

		if hasFound {
			break
		}

		// add new scans

		start, end := 0, 0
		for x := 0; x < mainY; x++ {
			val := scanPoint(Entry.PuzzleInput(), int64(x), int64(mainY))
			if val == 1 {
				beamArea[common.Point{X: x, Y: mainY}] = true

				if start == 0 {
					start = x
				}
			} else {
				if start != 0 && end == 0 {
					end = x
				}
			}
		}

		limits[mainY] = lineRange{start: start, end: end}

		mainY++
		offsetY++
	}

	fmt.Println(pos)
	fmt.Println("not 8761079, or 8791082(too low) or 9291149")
	fmt.Println(pos.X*10000 + pos.Y)
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
