package Day202012

import (
	"testing"

	"github.com/pezza/advent-of-code/common"

	. "github.com/onsi/gomega"
)

func Test_ReadData(t *testing.T) {
	RegisterTestingT(t)

	data := getData(`F10
N3
F7
R90
F11`)

	Expect(data[0].ins).Should(Equal("F"))
	Expect(data[0].distance).Should(Equal(10))
	Expect(data[4].ins).Should(Equal("F"))
	Expect(data[4].distance).Should(Equal(1))
}

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	ins := getData(Entry.PuzzleInput())

	ship := getStartShip()

	for i := range ins {
		ship.processIns(ins[i])
	}

	// not 3004
	Expect(common.Abs(ship.x) + common.Abs(ship.y)).Should(Equal(1032))
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

	ins := getData(Entry.PuzzleInput())

	ship := getStartShip()
	waypoint := point{10, 1}

	for i := range ins {
		waypoint = ship.processRevisedIns(ins[i], waypoint)
	}

	// not 704519 high
	// not 165309 high
	Expect(common.Abs(ship.x) + common.Abs(ship.y)).Should(Equal(1032))

}

func Test_RevisedIns(t *testing.T) {
	RegisterTestingT(t)

	ship := getStartShip()
	waypoint := point{10, 1}

	waypoint = ship.processRevisedIns(instruction{"F", 10}, waypoint)

	Expect(ship.x).Should(Equal(100))
	Expect(ship.y).Should(Equal(10))
	Expect(waypoint.x).Should(Equal(10))
	Expect(waypoint.y).Should(Equal(1))

	waypoint = ship.processRevisedIns(instruction{"N", 3}, waypoint)

	Expect(ship.x).Should(Equal(100))
	Expect(ship.y).Should(Equal(10))
	Expect(waypoint.x).Should(Equal(10))
	Expect(waypoint.y).Should(Equal(4))

	waypoint = ship.processRevisedIns(instruction{"F", 7}, waypoint)

	Expect(ship.x).Should(Equal(170))
	Expect(ship.y).Should(Equal(38))
	Expect(waypoint.x).Should(Equal(10))
	Expect(waypoint.y).Should(Equal(4))

	waypoint = ship.processRevisedIns(instruction{"R", 90}, waypoint)

	Expect(ship.x).Should(Equal(170))
	Expect(ship.y).Should(Equal(38))
	Expect(waypoint.x).Should(Equal(4))
	Expect(waypoint.y).Should(Equal(-10))

	waypoint = ship.processRevisedIns(instruction{"F", 11}, waypoint)

	Expect(ship.x).Should(Equal(214))
	Expect(ship.y).Should(Equal(-72))
	Expect(waypoint.x).Should(Equal(4))
	Expect(waypoint.y).Should(Equal(-10))

	Expect(common.Abs(ship.x) + common.Abs(ship.y)).Should(Equal(286))

}

func Test_Rotate(t *testing.T) {
	RegisterTestingT(t)

	p := point{10, 4}

	p.rotate(90)

	Expect(p.x).Should(Equal(-4))
	Expect(p.y).Should(Equal(10))

	p = point{10, 4}

	p.rotate(360)

	Expect(p.x).Should(Equal(10))
	Expect(p.y).Should(Equal(4))

	p = point{10, 4}

	p.rotate(720)

	Expect(p.x).Should(Equal(10))
	Expect(p.y).Should(Equal(4))

	p = point{10, 4}

	p.rotate(-720)

	Expect(p.x).Should(Equal(10))
	Expect(p.y).Should(Equal(4))

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
