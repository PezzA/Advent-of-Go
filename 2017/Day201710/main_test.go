package Day201710

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	Expect(getElements(5)).Should(Equal([]int{0, 1, 2, 3, 4}))

	elements := getElements(5)
	elements, currentPosition, skipSize := swap(elements, 3, 0, 0)

	Expect(currentPosition).Should(Equal(3))
	Expect(skipSize).Should(Equal(1))
	Expect(elements).Should(Equal([]int{2, 1, 0, 3, 4}))

	elements, currentPosition, skipSize = swap(elements, 4, currentPosition, skipSize)

	Expect(currentPosition).Should(Equal(3))
	Expect(skipSize).Should(Equal(2))
	Expect(elements).Should(Equal([]int{4, 3, 0, 1, 2}))
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

	Expect(getDenseHashSection([]int{65, 27, 9, 1, 4, 3, 40, 50, 91, 7, 6, 0, 2, 5, 68, 22})).Should(Equal(64))

	Expect(DoFullHash("")).Should(Equal("a2582a3a0e66e6e86e3812dcb672a272"))
	Expect(DoFullHash("AoC 2017")).Should(Equal("33efeb34ea91902bb2f59c9920caa6cd"))
	Expect(DoFullHash("1,2,3")).Should(Equal("3efbe78a8d82f29979031a4aa0b16a9d"))
	Expect(DoFullHash("1,2,4")).Should(Equal("63960835bcdc130f0b66d7ff4f6a5a8e"))
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
