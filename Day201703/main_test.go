package Day201703

import (
	"fmt"
	"testing"
)

func Test_Stuff(t *testing.T) {
	fmt.Println(getSpiral(23, false))
}

func cheapAssert(expected int, actual int, t *testing.T) {
	if expected != actual {
		t.Error("expected", expected, "got", actual)
	}
}

func Benchmark_Day1(b *testing.B) {
	data := Entry.GetData()
	for n := 0; n < b.N; n++ {
		Entry.PartOne(data)
	}
}

func Benchmark_Day2(b *testing.B) {
	data := Entry.GetData()
	for n := 0; n < b.N; n++ {
		Entry.PartTwo(data)
	}
}
