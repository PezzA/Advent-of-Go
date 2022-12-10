package Day202209

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_ReadData(t *testing.T) {
	RegisterTestingT(t)

	instructions := getData(Entry.PuzzleInput())

	Expect(instructions[0]).Should(Equal(instruction{"L", 1}))
	Expect(instructions[len(instructions)-1]).Should(Equal(instruction{"U", 19}))
}

const sampleData1 string = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`

const sampleData2 string = `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 2`

func Test_ProcessMoves(t *testing.T) {
	RegisterTestingT(t)
	moves := getData(sampleData1)
	tailLocations := processMoves(moves, false)
	Expect(len(tailLocations)).Should(Equal(13))
}

func Test_Print(t *testing.T) {
	RegisterTestingT(t)
	moves := getData(sampleData1)
	processMoves(moves, true)
}

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)
	// not 6084
	Expect(Entry.PartOne(Entry.PuzzleInput(), nil)).Should(Equal("6181"))
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)
	// not 26
	Expect(Entry.PartTwo(Entry.PuzzleInput(), nil)).Should(Equal("2386"))
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
