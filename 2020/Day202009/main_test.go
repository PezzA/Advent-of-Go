package Day202009

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

func Test_ReadData(t *testing.T) {
	RegisterTestingT(t)

	data := getData(Entry.PuzzleInput())
	Expect(data[0]).Should(Equal(17))
	Expect(data[len(data)-1]).Should(Equal(57233154052801))
}

var testData = getData(`35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`)

func Test_GetPreAmble(t *testing.T) {
	RegisterTestingT(t)

	Expect(getPreamble(5, testData, 5)).Should(Equal([]int{35, 20, 15, 25, 47}))
	Expect(getPreamble(10, testData, 5)).Should(Equal([]int{40, 62, 55, 65, 95}))
}

func Test_CheckValidIndex(t *testing.T) {
	RegisterTestingT(t)

	Expect(checkValidIndex(102, []int{40, 62, 55, 65, 95})).Should(Equal(true))
	Expect(checkValidIndex(103, []int{40, 62, 55, 65, 95})).Should(Equal(false))
}

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	fmt.Println(checkSequence(getData(Entry.PuzzleInput()), 25))
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

	wl := getWeaknessList(10884537, getData(Entry.PuzzleInput()))

	min, max := -1, 0

	for i := range wl {
		if wl[i] < min || min == -1 {
			min = wl[i]
		}

		if wl[i] > max {
			max = wl[i]
		}
	}

	Expect(min + max).Should(Equal(62))
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
