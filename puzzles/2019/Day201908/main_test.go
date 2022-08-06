package Day201908

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

func Test_ReadData(t *testing.T) {
	RegisterTestingT(t)

	fmt.Println(getData(Entry.PuzzleInput()))
}

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	pix := getData(Entry.PuzzleInput())

	layers := getLayers(pix)

	leastZeros := -1
	leastZeroIndex := 0
	for lIndex, l := range layers {
		counts := getPixelCounts(l)
		if leastZeros == -1 || counts[0] < leastZeros {
			leastZeros = counts[0]
			leastZeroIndex = lIndex
		}
	}

	tlCounts := getPixelCounts(layers[leastZeroIndex])

	Expect(tlCounts[1] * tlCounts[2]).Should(Equal(1862))

}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

	pix := getData(Entry.PuzzleInput())

	layers := getLayers(pix)

	for y := 0; y < displayHeight; y++ {
		for x := 0; x < displayWidth; x++ {

			switch getTopVisiblePixel(x+(y*displayWidth), layers) {
			case 0:
				fmt.Print(".")
			case 1:
				fmt.Print("#")

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
