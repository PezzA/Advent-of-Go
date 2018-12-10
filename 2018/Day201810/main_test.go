package Day201810

import (
	"testing"

	"fmt"

	. "github.com/onsi/gomega"
)

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	stars := getData(Entry.PuzzleInput())

	focusStars, _ := getFocusFrame(stars)
	min, max := getBounds(focusStars)

	x := max.X - min.X
	y := max.Y - min.Y
	for ly := 0; ly <= y; ly++ {
		for lx := 0; lx <= x; lx++ {

			found := false

			for _, drawStar := range focusStars {
				if lx+min.X == drawStar.px && ly+min.Y == drawStar.py {
					fmt.Print("#")
					found = true
					break
				}
			}
			if !found {
				fmt.Print(".")
			}
		}

		fmt.Print("\n")
	}

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
