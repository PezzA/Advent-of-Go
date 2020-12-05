package Day202005

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

func Test_ReadData(t *testing.T) {
	RegisterTestingT(t)

	data := getdata(Entry.PuzzleInput())

	Expect(data[0].rowSpec).Should(Equal("BFFFFFF"))
	Expect(data[0].seatSpec).Should(Equal("RLL"))
}

func Test_GetHalf(t *testing.T) {
	RegisterTestingT(t)

	x1, x2 := getHalf(true, 0, 127)

	Expect(x1).Should(Equal(0))
	Expect(x2).Should(Equal(63))

	x1, x2 = getHalf(true, 32, 63)

	Expect(x1).Should(Equal(32))
	Expect(x2).Should(Equal(47))

	x1, x2 = getHalf(true, 44, 47)

	Expect(x1).Should(Equal(44))
	Expect(x2).Should(Equal(45))

	x1, x2 = getHalf(true, 44, 45)

	Expect(x1).Should(Equal(44))
	Expect(x2).Should(Equal(44))

	b1, b2 := getHalf(false, 0, 127)

	Expect(b1).Should(Equal(64))
	Expect(b2).Should(Equal(127))

	b1, b2 = getHalf(false, 0, 63)

	Expect(b1).Should(Equal(32))
	Expect(b2).Should(Equal(63))

	b1, b2 = getHalf(false, 32, 47)

	Expect(b1).Should(Equal(40))
	Expect(b2).Should(Equal(47))

	b1, b2 = getHalf(false, 40, 47)

	Expect(b1).Should(Equal(44))
	Expect(b2).Should(Equal(47))

	b1, b2 = getHalf(false, 44, 47)

	Expect(b1).Should(Equal(46))
	Expect(b2).Should(Equal(47))

	b1, b2 = getHalf(false, 46, 47)

	Expect(b1).Should(Equal(47))
	Expect(b2).Should(Equal(47))
}

func Test_GetRow(t *testing.T) {
	RegisterTestingT(t)

	bp := boardingPass{
		rowSpec:  "FBFBBFF",
		seatSpec: "RLR",
	}

	Expect(bp.getRow()).Should(Equal(44))
	Expect(bp.getSeat()).Should(Equal(5))
	Expect(bp.getSeatId()).Should(Equal(357))

	bp2 := boardingPass{
		rowSpec:  "BFFFBBF",
		seatSpec: "RRR",
	}

	Expect(bp2.getRow()).Should(Equal(70))
	Expect(bp2.getSeat()).Should(Equal(7))
	Expect(bp2.getSeatId()).Should(Equal(567))

	bp3 := boardingPass{
		rowSpec:  "FFFBBBF",
		seatSpec: "RRR",
	}

	Expect(bp3.getRow()).Should(Equal(14))
	Expect(bp3.getSeat()).Should(Equal(7))
	Expect(bp3.getSeatId()).Should(Equal(119))

	bp4 := boardingPass{
		rowSpec:  "BBFFBBF",
		seatSpec: "RLL",
	}

	Expect(bp4.getRow()).Should(Equal(102))
	Expect(bp4.getSeat()).Should(Equal(4))
	Expect(bp4.getSeatId()).Should(Equal(820))
}

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	highest := 0

	for _, bp := range getdata(Entry.PuzzleInput()) {
		if bp.getSeatId() > highest {
			highest = bp.getSeatId()
		}
	}

	fmt.Println(highest)
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

	plan := make(map[int]bool, 0)
	for _, bp := range getdata(Entry.PuzzleInput()) {
		plan[bp.getSeatId()] = true
	}

	missing := 0
	for start := 8; start < (127*8)+8; start++ {

		if _, ok := plan[start]; !ok {
			missing = start
			break
		}

	}

	fmt.Println(missing)

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
