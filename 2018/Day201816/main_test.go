package Day201816

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	ops := getOpCodes()

	reg := make(registerSet)
	reg[0] = 1
	reg[1] = 1

	Expect(ops["addr"].process(reg, 0, 1, 3)).Should(
		Equal(registerSet{0: 1, 1: 1, 3: 2}))

	tests, _ := getData(Entry.PuzzleInput())

	Expect(tests[0].before).Should(Equal(registerSet{0: 1, 1: 1, 2: 0, 3: 1}))
	Expect(tests[0].codes).Should(Equal([]int{0, 1, 0, 1}))
	Expect(tests[0].after).Should(Equal(registerSet{0: 1, 1: 1, 2: 0, 3: 1}))

	masterCount := 0
	testCount := 0
	for _, test := range tests {
		count := 0

		for _, v := range ops {
			if v.process(test.before.deepCopy(), test.codes[1], test.codes[2], test.codes[3]).Same(test.after.deepCopy()) {
				count++
			}
		}

		if count >= 3 {
			masterCount++
		}
		testCount++
	}
	fmt.Println(masterCount, testCount)
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
