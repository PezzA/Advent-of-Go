package Day202011

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

func Test_ReadData(t *testing.T) {
	RegisterTestingT(t)

	data := getData(`L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`)

	fmt.Println(data)
}

func Test_AdjacentOccupiedSeats(t *testing.T) {
	RegisterTestingT(t)
	plan := getData(`#.LL.L#.##
#LLLLLL.L#
L.L.L..L..
#LLL.LL.L#
#.LL.LL.LL
#.LLLL#.##
..L.L.....
#LLLLLLLL#
#.LLLLLL.L
#.#LLLL.##`)

	Expect(plan.adjacentOccupiedSeats(0, 0)).Should(Equal(1))
	Expect(plan.adjacentOccupiedSeats(3, 0)).Should(Equal(0))
	Expect(plan.adjacentOccupiedSeats(9, 9)).Should(Equal(1))
	Expect(plan.adjacentOccupiedSeats(0, 9)).Should(Equal(1))
	Expect(plan.adjacentOccupiedSeats(9, 0)).Should(Equal(2))
}

func Test_NextState(t *testing.T) {
	RegisterTestingT(t)

	data := getData(`L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`)

	plan := data.nextState(4, false)

	Expect(plan.equals(getData(`#.##.##.##
#######.##
#.#.#..#..
####.##.##
#.##.##.##
#.#####.##
..#.#.....
##########
#.######.#
#.#####.##`))).Should(Equal(true))

	plan = plan.nextState(4, false)

	for i := range plan {
		fmt.Println(string(plan[i]))
	}

	Expect(plan.equals(getData(`#.LL.L#.##
#LLLLLL.L#
L.L.L..L..
#LLL.LL.L#
#.LL.LL.LL
#.LLLL#.##
..L.L.....
#LLLLLLLL#
#.LLLLLL.L
#.#LLLL.##`))).Should(Equal(true))

}

func Test_OccupiedSeats(t *testing.T) {
	RegisterTestingT(t)

	data := getData(`#.#L.L#.##
#LLL#LL.L#
L.#.L..#..
#L##.##.L#
#.#L.LL.LL
#.#L#L#.##
..L.L.....
#L#L##L#L#
#.LLLLLL.L
#.#L#L#.##`)

	Expect(data.countOccupiedSeats()).Should(Equal(37))
}

func Test_Equals(t *testing.T) {
	RegisterTestingT(t)
	data := getData(`L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`)

	data1 := getData(`L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`)

	data2 := getData(`L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.#L.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`)

	Expect(data.equals(data1)).Should(Equal(true))
	Expect(data.equals(data2)).Should(Equal(false))
}

func Test_ExtendedCheck(t *testing.T) {
	RegisterTestingT(t)

	plan := getData(`L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`)

	plan2 := plan.nextState(5, true)

	Expect(plan2.equals(getData(`#.##.##.##
#######.##
#.#.#..#..
####.##.##
#.##.##.##
#.#####.##
..#.#.....
##########
#.######.#
#.#####.##`))).Should(Equal(true))

	plan3 := plan2.nextState(5, true)

	for i := range plan3 {
		fmt.Println(string(plan3[i]))
	}
	Expect(plan3.equals(getData(`#.LL.LL.L#
#LLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLL#
#.LLLLLL.L
#.LLLLL.L#`))).Should(Equal(true))

	plan4 := plan3.nextState(5, true)

	Expect(plan4.equals(getData(`#.L#.##.L#
#L#####.LL
L.#.#..#..
##L#.##.##
#.##.#L.##
#.#####.#L
..#.#.....
LLL####LL#
#.L#####.L
#.L####.L#`))).Should(Equal(true))

	plan5 := plan4.nextState(5, true)

	Expect(plan5.equals(getData(`#.L#.L#.L#
#LLLLLL.LL
L.L.L..#..
##LL.LL.L#
L.LL.LL.L#
#.LLLLL.LL
..L.L.....
LLLLLLLLL#
#.LLLLL#.L
#.L#LL#.L#`))).Should(Equal(true))

	plan6 := plan5.nextState(5, true)

	Expect(plan6.equals(getData(`#.L#.L#.L#
#LLLLLL.LL
L.L.L..#..
##L#.#L.L#
L.L#.#L.L#
#.L####.LL
..#.#.....
LLL###LLL#
#.LLLLL#.L
#.L#LL#.L#`))).Should(Equal(true))

	plan7 := plan6.nextState(5, true)

	Expect(plan7.equals(getData(`#.L#.L#.L#
#LLLLLL.LL
L.L.L..#..
##L#.#L.L#
L.L#.LL.L#
#.LLLL#.LL
..#.L.....
LLL###LLL#
#.LLLLL#.L
#.L#LL#.L#`))).Should(Equal(true))

	plan8 := plan7.nextState(5, true)

	Expect(plan8.equals(getData(`#.L#.L#.L#
#LLLLLL.LL
L.L.L..#..
##L#.#L.L#
L.L#.LL.L#
#.LLLL#.LL
..#.L.....
LLL###LLL#
#.LLLLL#.L
#.L#LL#.L#`))).Should(Equal(true))

	plan9 := plan8.nextState(5, true)

	Expect(plan9.equals(getData(`#.L#.L#.L#
#LLLLLL.LL
L.L.L..#..
##L#.#L.L#
L.L#.LL.L#
#.LLLL#.LL
..#.L.....
LLL###LLL#
#.LLLLL#.L
#.L#LL#.L#`))).Should(Equal(true))
}

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	data := getData(Entry.PuzzleInput())

	occ := 0

	newState := data.nextState(4, false)

	for {
		newState2 := newState.nextState(4, false)

		if newState2.equals(newState) {
			occ = newState.countOccupiedSeats()
			break
		}

		newState = newState2
	}

	Expect(occ).Should(Equal(2354))
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

	data := getData(Entry.PuzzleInput())

	occ := 0

	newState := data.nextState(5, true)

	for {
		newState2 := newState.nextState(5, true)

		if newState2.equals(newState) {
			occ = newState.countOccupiedSeats()
			break
		}

		newState = newState2
	}

	Expect(occ).Should(Equal(2354))

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
