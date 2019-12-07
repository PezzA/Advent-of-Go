package Day201907

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/pezza/advent-of-code/common"
)

func Test_ReadData(t *testing.T) {
	RegisterTestingT(t)

	//fmt.Println(getListIntData(Entry.PuzzleInput()))

}

var testa = "3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0"
var testb = "3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0"
var testc = "3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0"

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	//Expect(amplifySignal([]int{4, 3, 2, 1, 0}, testa)).Should(Equal(43210))
	//Expect(amplifySignal([]int{0, 1, 2, 3, 4}, testb)).Should(Equal(54321))
	//Expect(amplifySignal([]int{1, 0, 4, 3, 2}, testc)).Should(Equal(65210))

	maxSignal := 0

	for _, perm := range common.GetPermuations([]int{0, 1, 2, 3, 4}) {
		result := straightAmplifySignal(perm, Entry.PuzzleInput())

		if result > maxSignal {
			maxSignal = result
		}
	}

	//answer!
	Expect(maxSignal).Should(Equal(440880))
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

	maxSignal := 0

	for _, perm := range common.GetPermuations([]int{5, 6, 7, 8, 9}) {
		result := feedbackAmplifysignal(perm, Entry.PuzzleInput())

		if result > maxSignal {
			maxSignal = result
		}
	}

	fmt.Println(maxSignal)
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
