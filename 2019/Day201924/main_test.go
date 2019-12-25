package Day201924

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

func Test_ReadData(t *testing.T) {
	RegisterTestingT(t)

}

var testData = `....#
#..#.
#..##
..#..
#....`

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	bugs := getData(Entry.PuzzleInput())

	states := [][][]bool{}

	states = append(states, bugs)

	var found = false

	for !found {

		bugs = tickAutomata(bugs)

		for _, s := range states {
			if same(bugs, s) {
				found = true
			}
		}

		if !found {
			states = append(states, bugs)
		}
	}

	printBugs(bugs)
	fmt.Println(int64(getBioDiversity(bugs)))
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

	hive := make(hive, 0)

	hive[0] = getData(`....#
#..#.
#.?##
..#..
#....
`)

	for i := 1; i <= 10; i++ {
		hive[0-i] = getData(`.....'
.....
..?..
.....
.....
`)
	}

	hive[0+i] = getData(`.....'
.....
..?..
.....
.....
`)

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
