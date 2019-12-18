package Day201918

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/pezza/advent-of-code/common"
)

var testA = `#########
#b.A.@.a#
#########`

func Test_ReadData(t *testing.T) {
	RegisterTestingT(t)

	tunnels, objects, player := getData(testA)

	// printTunnels(tunnels, objects, player)

	visits := make(map[common.Point]bool, 0)
	fl := getPossibleMoves(tunnels, objects, player.pos, player.inventory, 1, visits)

	for index := range fl {
		fmt.Println(fl[index])
	}
}

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

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
