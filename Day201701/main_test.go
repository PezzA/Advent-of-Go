package Day201701

import (
	"testing"
)

func Test_Stuff(t *testing.T) {
	index := getWrappedIndex(0, 10, 0)

	test := 0
	if index != 0 {
		t.Error("expecting", test, "got", index)
	}

	index = getWrappedIndex(0, 10, 10)

	if index != 0 {
		t.Error("expecting", test, "got", index)
	}

	index = getWrappedIndex(0, 10, 1)

	if index != 1 {
		t.Error("expecting", test, "got", index)
	}

	index = getWrappedIndex(0, 10, 15)

	if index != 5 {
		t.Error("expecting", test, "got", index)
	}
}

func Benchmark_PartOne(b *testing.B) {
	data := Entry.GetData()
	for n := 0; n < b.N; n++ {
		Entry.PartOne(data)
	}
}

func Benchmark_PartTwo(b *testing.B) {
	data := Entry.GetData()
	for n := 0; n < b.N; n++ {
		Entry.PartTwo(data)
	}
}
