package Day202019

import (
	"fmt"
	"regexp"
	"testing"

	. "github.com/onsi/gomega"
)

func Test_ReadData(t *testing.T) {
	RegisterTestingT(t)
	exp, tests, _, _ := getData(Entry.PuzzleInput())
	Expect(exp[110].char).Should(Equal("b"))
	Expect(tests[0]).Should(Equal("aaababaaaaaabbbaaaabbaaa"))

	fmt.Println(exp[0])
	fmt.Println(exp[8])
	fmt.Println(exp[11])
}

func Test_GetText(t *testing.T) {
	RegisterTestingT(t)

	exp, tests, aIns, bIns := getData(Entry.PuzzleInput())

	fmt.Println(exp[0])
	fmt.Println(exp[8])
	fmt.Println(exp[11])
	exp42 := exp[42].getText(aIns, bIns, exp)
	exp31 := exp[31].getText(aIns, bIns, exp)

	tot := 0

	for i := 1; i < 20; i++ {
		expression := fmt.Sprintf("^(%v)+(%v){%d}(%v){%d}$", exp42, exp42, i, exp31, i)
		re := regexp.MustCompile(expression)
		for _, t := range tests {
			if re.MatchString(t) {
				tot++
			}
		}
	}

	Expect(tot).Should(Equal(2))

}

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

}

func Test_PartTwo(t *testing.T) {

	// 480 too high
	// 433 too high

	// 424 ~

	// 299 too low
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
