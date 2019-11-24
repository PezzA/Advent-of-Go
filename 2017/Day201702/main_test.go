package Day201702

import (
	"testing"
)

func Test_Stuff(t *testing.T) {
	cheapAssert(8, diff("5 1 9 5"), t)
	cheapAssert(4, diff("7 5 3"), t)
	cheapAssert(6, diff("2 4 6 8"), t)

	res, err := evenDivide("5 9 2 8")

	if err != nil {
		t.Error(err)
	}

	cheapAssert(4, res, t)

	res, err = evenDivide("9 4 7 3")

	if err != nil {
		t.Error(err)
	}

	cheapAssert(3, res, t)

	res, err = evenDivide("3 8 6 5")

	if err != nil {
		t.Error(err)
	}

	cheapAssert(2, res, t)

}

func cheapAssert(expected int, actual int, t *testing.T) {
	if expected != actual {
		t.Error("expected", expected, "got", actual)
	}
}
func Benchmark_PartOne(b *testing.B) {
	data := Entry.PuzzleInput()
	for n := 0; n < b.N; n++ {
		Entry.PartOne(data)
	}
}

func Benchmark_PartTwo(b *testing.B) {
	data := Entry.PuzzleInput()
	for n := 0; n < b.N; n++ {
		Entry.PartTwo(data)
	}
}
