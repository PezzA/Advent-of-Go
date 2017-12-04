package Day201704

import (
	"testing"
)

func Test_Stuff(t *testing.T) {

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
