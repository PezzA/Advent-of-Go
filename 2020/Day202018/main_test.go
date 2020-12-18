package Day202018

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

func Test_ReadData(t *testing.T) {
	RegisterTestingT(t)
}

func Test_GetDepth(t *testing.T) {
	RegisterTestingT(t)

	Expect(getDepth("5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))")).Should(Equal(2))
	Expect(getDepth("1 + 2 * 3 + 4 * 5 + 6")).Should(Equal(0))

	input := "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))"
	exp, start, end := getExpressionAtDepth(input, 2)
	Expect(exp).Should(Equal("8 + 6 * 4"))
	Expect(input[start:end]).Should(Equal(exp))
}

func Test_EvaluateExpression(t *testing.T) {
	RegisterTestingT(t)
	Expect(evaluateExpression("1 + 2 * 3 + 4 * 5 + 6", false)).Should(Equal(71))
}

func Test_EvaluateLine(t *testing.T) {
	RegisterTestingT(t)

	Expect(evaluateLine("1 + 2 * 3 + 4 * 5 + 6", false)).Should(Equal(71))
	Expect(evaluateLine("2 * 3 + (4 * 5)", false)).Should(Equal(26))
	Expect(evaluateLine("5 + (8 * 3 + 9 + 3 * 4 * 3)", false)).Should(Equal(437))
	Expect(evaluateLine("5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", false)).Should(Equal(12240))
	Expect(evaluateLine("((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", false)).Should(Equal(13632))
	Expect(evaluateLine("8 + ((9 * 7) + 2) + (4 * (9 * 3 * 9 + 3 + 8) + 6 + 5 + 8)", false)).Should(Equal(1108))
}

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	data := getData(Entry.PuzzleInput())

	tot := 0
	for i := range data {
		fmt.Println(data[i])
		tot += evaluateLine(data[i], false)
	}

	fmt.Println(tot)
}

func Test_ResolvePlus(t *testing.T) {
	RegisterTestingT(t)

	Expect(resolvePlus("2 + 2")).Should(Equal(4))
}

func Test_AdvancedResolve(t *testing.T) {
	RegisterTestingT(t)

	Expect(evaluateLine("1 + 2 * 3 + 4 * 5 + 6", true)).Should(Equal(231))
	Expect(evaluateLine("1 + (2 * 3) + (4 * (5 + 6))", true)).Should(Equal(51))
	Expect(evaluateLine("2 * 3 + (4 * 5)", true)).Should(Equal(46))
	Expect(evaluateLine("5 + (8 * 3 + 9 + 3 * 4 * 3)", true)).Should(Equal(1445))
	Expect(evaluateLine("5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", true)).Should(Equal(669060))
	Expect(evaluateLine("((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", true)).Should(Equal(23340))
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

	data := getData(Entry.PuzzleInput())

	tot := 0
	for i := range data {
		fmt.Println(data[i])
		tot += evaluateLine(data[i], true)
	}

	fmt.Println(tot)
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
