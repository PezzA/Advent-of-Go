package Day201513

import (
	"fmt"
	"testing"

	"github.com/pezza/AoC2017/Common"

	. "github.com/onsi/gomega"
)

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	seatingList := getData(Entry.PuzzleInput())

	seatingList.addNeutralGuest("John")
	guestList := seatingList.getGuestList()

	fmt.Println(seatingList.scoreList(guestList))

	resultChan := make(chan []string)

	go permutation(guestList, make([]string, 0), resultChan)

	i := 0
	modifier := 0
	for gl := range resultChan {
		i++
		//	fmt.Println(perm, i)
		currMod := seatingList.scoreList(gl)
		if currMod > modifier {
			modifier = currMod
			fmt.Println(i, modifier)
		}
		if i == common.Factorial(len(guestList)) {
			break
		}
	}

	fmt.Println(modifier)

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
