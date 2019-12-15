package Day201914

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

func Test_ReadData(t *testing.T) {
	RegisterTestingT(t)

	recipes := getData(Entry.PuzzleInput())

	Expect(recipes[0]).Should(Equal(recipe{
		[]ingredient{
			ingredient{3, "CFGBR"},
			ingredient{9, "PFMFC"},
			ingredient{2, "FQFPN"},
		},
		ingredient{2, "PKPWN"},
	}))

	for _, r := range recipes {
		fmt.Println(r)
	}
}

var testData = `10 ORE => 10 A
1 ORE => 1 B
7 A, 1 B => 1 C
7 A, 1 C => 1 D
7 A, 1 D => 1 E
7 A, 1 E => 1 FUEL`

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	data := getData(Entry.PuzzleInput())

	data.startCrawl()
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
