package Day201911

import (
	"fmt"
	"testing"

	"github.com/pezza/advent-of-code/common"

	"github.com/pezza/advent-of-code/2019/intcode"

	. "github.com/onsi/gomega"
)

func Test_ReadData(t *testing.T) {
	RegisterTestingT(t)

}

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	input := make(chan int64, 1)
	output := make(chan int64, 0)

	vm := intcode.New(Entry.PuzzleInput())

	go func() {
		vm.RunProgram(nil, nil, input, output)
		close(output)
	}()

	hull := make(map[common.Point]bool, 0)
	paintIns := true

	rob := robot{
		common.Point{0, 0},
		0,
	}

	input <- 0

	for val := range output {

		if paintIns {

			hull[rob.pos] = val == 1
		} else {

			rob.moveIns(int(val))

			if val, ok := hull[rob.pos]; ok {
				sig := int64(0)
				if val {
					sig = 1
				}

				input <- sig
			} else {
				input <- 0
			}

		}
		paintIns = !paintIns
	}

	fmt.Println(len(hull))

}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

	hull := runRobot(Entry.PuzzleInput(), 1)

	minX, minY, maxX, maxY := -1, -1, 0, 0
	for k := range hull {
		if minX == -1 || k.X < minX {
			minX = k.X
		}

		if minY == -1 || k.Y < minY {
			minY = k.Y
		}

		if k.X > maxX {
			maxX = k.X
		}

		if k.Y > maxY {
			maxY = k.Y
		}
	}

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if val, ok := hull[common.Point{x, y}]; ok {
				if val {
					fmt.Print("#")
				} else {
					fmt.Print(" ")
				}
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
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
