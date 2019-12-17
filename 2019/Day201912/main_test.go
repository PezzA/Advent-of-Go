package Day201912

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

var testData = `<x=-1, y=0, z=2>
<x=2, y=-10, z=-7>
<x=4, y=-8, z=8>
<x=3, y=5, z=-1>`

func Test_ReadData(t *testing.T) {
	RegisterTestingT(t)

	bodies := getData(Entry.PuzzleInput())
	Expect(bodies[0]).Should(Equal(planet{vector3d{-16, 15, -9}, vector3d{0, 0, 0}}))
	Expect(bodies[3]).Should(Equal(planet{vector3d{-3, 18, 9}, vector3d{0, 0, 0}}))
}

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	Expect(Entry.PartOne(Entry.PuzzleInput(), nil)).Should(Equal("10664"))
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

	pairs := getPairs(4)
	bodies := getData(Entry.PuzzleInput())

	for i := 0; i < 1000000; i++ {
		bodies.applyGravity(pairs)
		bodies.applyVelocity()

		fmt.Println(bodies[0].position.x)
	}

	//not 188395899048
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
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
