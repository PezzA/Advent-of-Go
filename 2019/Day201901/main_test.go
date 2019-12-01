package Day201901

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	Expect(getData(Entry.PuzzleInput())[0]).Should(Equal(141589))
	Expect(getData(Entry.PuzzleInput())[1]).Should(Equal(93261))

	Expect(fuelRequirement(100756)).Should(Equal(33583))
	Expect(fuelRequirement(1969)).Should(Equal(654))
	Expect(fuelRequirement(12)).Should(Equal(2))
	Expect(fuelRequirement(14)).Should(Equal(2))
	Expect(fuelRequirement(2)).Should(Equal(0))

	fmt.Println("Not 46996")
	fmt.Println(Entry.PartOne(Entry.PuzzleInput(), nil))

}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

	Expect(fuelFuelRequirement(100756)).Should(Equal(50346))

	fmt.Println(Entry.PartTwo(Entry.PuzzleInput(), nil))
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
