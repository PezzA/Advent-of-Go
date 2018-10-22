package Day201518

import (
	"testing"

	"fmt"

	. "github.com/onsi/gomega"
)

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	startFrame := getData(`
.#.#.#
...##.
#....#
..#...
#.#..#
####..`)

	for i := 0; i < 4; i++ {
		startFrame = startFrame.getNewFrame()
	}
	startFrame.drawFrame()

	fmt.Println(startFrame.pixelOnCount())

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
