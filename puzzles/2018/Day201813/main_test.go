package Day201813

import (
	"fmt"
	"sort"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/pezza/advent-of-code/puzzles/common"
)

var testData = `/->-\        
|   |  /----\
| /-+--+-\  |
| | |  | v  |
\-+-/  \-+--/
  \------/`

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	grid, carts := getData(Entry.PuzzleInput())

	crashed := false
	var point common.Point
	for {
		for index := range carts {

			carts[index], _, crashed = moveCart(carts[index], grid, carts)
			if crashed {
				point = common.Point{X: carts[index].x, Y: carts[index].y}
				break
			}
		}

		// sort the carts into x y order
		sort.Slice(carts, func(i, j int) bool {
			if carts[i].y == carts[j].y {
				return carts[i].x < carts[j].x
			}
			return carts[i].y < carts[j].y
		})

		if crashed {
			break
		}
	}

	fmt.Println(point)
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)
	grid, carts := getData(Entry.PuzzleInput())

	for {
		crashIds := make([]int, 0)
		for index := range carts {
			crashed := false
			var otherCrashedCart cart

			carts[index], otherCrashedCart, crashed = moveCart(carts[index], grid, carts)

			if crashed {
				crashIds = append(crashIds, carts[index].id)
				crashIds = append(crashIds, otherCrashedCart.id)
			}
		}

		if len(crashIds) > 0 {
			fmt.Println("before", carts)
			fmt.Println(crashIds)

			for _, val := range crashIds {
				foundIndex := 0
				for searchIndex := range carts {
					if carts[searchIndex].id == val {
						foundIndex = searchIndex
						break
					}
				}
				fmt.Println("removing car at index number", foundIndex, "should have id", val)
				carts = append(carts[:foundIndex], carts[foundIndex+1:]...)
			}
			fmt.Println("after", carts)
		}

		if len(carts) == 1 {
			fmt.Println(carts)
			break
		}

		// sort the carts into x y order
		sort.Slice(carts, func(i, j int) bool {
			if carts[i].y == carts[j].y {
				return carts[i].x < carts[j].x
			}
			return carts[i].y < carts[j].y
		})
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
