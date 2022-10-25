package Day201708

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

func getTestData() string {
	return `b inc 5 if a > 1
a inc 1 if b < 5
c dec -10 if a >= 1
c inc -20 if c == 10`
}
func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	Expect(parseInstruction("b inc 5 if a > 1")).Should(Equal(instruction{"b", "inc", 5, "a", ">", 1}))
	Expect(parseInstruction("b inc -5 if a > -1")).Should(Equal(instruction{"b", "inc", -5, "a", ">", -1}))

	r, _ := runProgram(make(registers, 0), getInstructions(getTestData()))

	fmt.Println(r)

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
