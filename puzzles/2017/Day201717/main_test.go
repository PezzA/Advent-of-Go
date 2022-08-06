package Day201717

import (
	"container/list"
	"fmt"
	"testing"

	"github.com/pezza/AoC2017/Common"

	. "github.com/onsi/gomega"
)

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	Expect(common.GetWrappedIndex(0, 1, 3)).Should(Equal(0))
	Expect(common.GetWrappedIndex(0, 2, 3)).Should(Equal(1))

	sl := list.New()
	sl.PushFront(0)
	steps := 3

	pos := sl.Front()
	pos = addValue(1, steps, pos, sl)

	test := sl.Front()
	Expect(test.Value).Should(Equal(0))
	test = test.Next()
	Expect(test.Value).Should(Equal(1))

	pos = addValue(2, steps, pos, sl)
	test = sl.Front()
	Expect(test.Value).Should(Equal(0))
	test = test.Next()
	Expect(test.Value).Should(Equal(2))
	test = test.Next()
	Expect(test.Value).Should(Equal(1))

	pos = addValue(3, steps, pos, sl)
	test = sl.Front()
	Expect(test.Value).Should(Equal(0))
	test = test.Next()
	Expect(test.Value).Should(Equal(2))
	test = test.Next()
	Expect(test.Value).Should(Equal(3))
	test = test.Next()
	Expect(test.Value).Should(Equal(1))

	pos = addValue(4, steps, pos, sl)
	test = sl.Front()
	Expect(test.Value).Should(Equal(0))
	test = test.Next()
	Expect(test.Value).Should(Equal(2))
	test = test.Next()
	Expect(test.Value).Should(Equal(4))
	test = test.Next()
	Expect(test.Value).Should(Equal(3))
	test = test.Next()
	Expect(test.Value).Should(Equal(1))

	pos = addValue(5, steps, pos, sl)
	test = sl.Front()
	Expect(test.Value).Should(Equal(0))
	test = test.Next()
	Expect(test.Value).Should(Equal(5))
	test = test.Next()
	Expect(test.Value).Should(Equal(2))
	test = test.Next()
	Expect(test.Value).Should(Equal(4))
	test = test.Next()
	Expect(test.Value).Should(Equal(3))
	test = test.Next()
	Expect(test.Value).Should(Equal(1))

	newsl := list.New()
	newpos := newsl.PushFront(0)

	for i := 1; i < 2018; i++ {
		newpos = addValue(i, steps, newpos, newsl)
	}

	fmt.Println(newpos.Next().Value)

}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

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
