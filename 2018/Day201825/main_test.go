package Day201825

import (
	"testing"

	"fmt"

	. "github.com/onsi/gomega"
)

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)
	points := getData(Entry.PuzzleInput())

	Expect(points[0]).Should(Equal(point{-1, -5, 2, -5}))

	Expect(Get4DMDistance(point{0, 0, 0, 0}, point{3, 0, 0, 0})).Should(Equal(3))
	Expect(Get4DMDistance(point{0, 0, 0, 6}, point{9, 0, 0, 0})).Should(Equal(15))

	count := 0
	for len(points) > 0 {
		cons := getConstellation(points)
		points = removeConstellationFromList(cons, points)
		count++
		fmt.Println(count, len(points))
	}

	fmt.Println(count)
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
