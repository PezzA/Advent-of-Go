package Day202023

import (
	"container/list"
	"fmt"
	"strconv"
	"testing"

	"github.com/pezza/advent-of-code/common"

	. "github.com/onsi/gomega"
)

func Test_ReadData(t *testing.T) {
	RegisterTestingT(t)
	Expect(getData(Entry.PuzzleInput())).Should(Equal([]int{6, 8, 5, 9, 7, 4, 2, 1, 3}))
}

func Test_RunRound(t *testing.T) {
	RegisterTestingT(t)

	Expect(getData("389125467").runRound(0)).Should(Equal(circle{3, 2, 8, 9, 1, 5, 4, 6, 7}))

	Expect(getData("325467891").runRound(2)).Should(Equal(circle{7, 2, 5, 8, 9, 1, 3, 4, 6}))
}
func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	circ := getData(Entry.PuzzleInput())

	targetCup := circ[0]

	for count := 0; count <= 99; count++ {
		circ, targetCup = circ.runRound(targetCup)
	}

	fmt.Println(circ.print(targetCup))
	_, index := common.ContainsInt(circ, 1)

	fmt.Println(index)
	val := ""

	for i := 1; i <= len(circ)-1; i++ {
		index = incIndex(index, len(circ))

		val += strconv.Itoa(circ[index])
	}

	fmt.Println(val)
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)
	circData := getData(Entry.PuzzleInput())

	circle := list.New()

	for i := range circData {
		circle.PushBack(circData[i])
	}

	for i := 10; i <= 1000000; i++ {
		circle.PushBack(i)
	}

	locs := make(map[int]*list.Element, 0)

	for e := circle.Front(); e != nil; e = e.Next() {
		locs[e.Value.(int)] = e
	}

	next := circle.Front()

	for i := 0; i < 10000000; i++ {
		next = listRunRound(circle, next, locs)
	}

	cupOne := locs[1]

	nextP1 := cupOne.Next().Value.(int)
	nextP2 := cupOne.Next().Next().Value.(int)

	fmt.Println(nextP1 * nextP2)
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
