package Day201724

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

func getTestBridges() []component {
	return getComponents(`0/2
2/2
2/3
3/4
3/5
0/1
10/1
9/10`)
}

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	components := getTestBridges()
	//components := getComponents(Entry.PuzzleInput())

	bridges := make([][]component, 0)

	for index := range components {

		part1Zero := components[index].port1 == 0
		part2Zero := components[index].port2 == 0

		if part1Zero || part2Zero {
			newComp := component{components[index].port1, part1Zero, components[index].port2, part2Zero}

			newList := make([]component, len(components))
			copy(newList, components)

			getBridges(newComp, make([]component, 0), append(newList[:index], newList[index+1:]...), &bridges)
		}
	}

	fmt.Println("what")
	for index := range bridges {
		fmt.Println(bridges[index])
	}
	fmt.Println("the")

}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

}

func Benchmark_PartOne(b *testing.B) {
	data := Entry.PuzzleInput()
	for n := 0; n < b.N; n++ {
		Entry.PartOne(data, nil)
	}
}

func Benchmark_PartTwo(b *testing.B) {
	data := Entry.PuzzleInput()
	for n := 0; n < b.N; n++ {
		Entry.PartTwo(data, nil)
	}
}
