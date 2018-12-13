package Day201813

import (
	"testing"

	"fmt"

	. "github.com/onsi/gomega"
	"github.com/pezza/advent-of-code/Common"
)

var testData = `/---\
|   |  /----\
| /-+--+-\  |
| | |  | ^  |
\-+-/  \-+--/
  \------/   `

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	grid, carts := getData(Entry.PuzzleInput())

	crashed := false
	var point common.Point
	for {
		for index := range carts {
			carts[index], crashed = moveCart(carts[index], grid, carts)
			if crashed {
				point = common.Point{carts[index].x, carts[index].y}
			}
		}

		if crashed {
			break
		}
	}

	fmt.Println(point)
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
