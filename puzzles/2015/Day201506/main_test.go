package Day201506

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	lightGrid := getGrid(1000, 1000)
	lightGrid = lightGrid.turnOn(0, 0, 999, 999)
	Expect(lightGrid.countLights()).Should(Equal(1000000))

	lightGrid = getGrid(1000, 1000)
	lightGrid = lightGrid.toggle(0, 0, 999, 0)
	Expect(lightGrid.countLights()).Should(Equal(1000))

	lightGrid = getGrid(1000, 1000)
	lightGrid = lightGrid.turnOn(0, 0, 999, 999)
	lightGrid = lightGrid.turnOff(499, 499, 500, 500)
	Expect(lightGrid.countLights()).Should(Equal(999996))
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
