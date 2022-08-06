package Day201812

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	pl, transforms := getData(Entry.PuzzleInput())

	Expect(pl[0]).Should(Equal(true))
	Expect(pl[1]).Should(Equal(true))
	Expect(pl[2]).Should(Equal(true))
	Expect(pl[3]).Should(Equal(false))

	Expect(transforms[0]).Should(Equal(transform{[]bool{false, false, true, false, true}, false}))
	Expect(transforms[1]).Should(Equal(transform{[]bool{true, false, false, true, true}, true}))

	npl, ntransforms := getData(`initial state: #..#.#..##......###...###

...## => #
..#.. => #
.#... => #
.#.#. => #
.#.## => #
.##.. => #
.#### => #
#.#.# => #
#.### => #
##.#. => #
##.## => #
###.. => #
###.# => #
####. => #`)

	npl2 := regenerate(npl, ntransforms)

	Expect(npl2[0]).Should(Equal(true))
	Expect(npl2[1]).Should(Equal(false))
	Expect(npl2[2]).Should(Equal(false))
	Expect(npl2[3]).Should(Equal(false))
	Expect(npl2[4]).Should(Equal(true))

	npl2 = regenerate(npl2, ntransforms)
	npl2 = regenerate(npl2, ntransforms)

	Expect(npl2[-1]).Should(Equal(true))

	npl3, transforms3 := getData(`initial state: #..#.#..##......###...###

...## => #
..#.. => #
.#... => #
.#.#. => #
.#.## => #
.##.. => #
.#### => #
#.#.# => #
#.### => #
##.#. => #
##.## => #
###.. => #
###.# => #
####. => #`)

	for i := 0; i < 20; i++ {
		npl3 = regenerate(npl3, transforms3)

	}

	count := 0

	for k, v := range npl3 {
		if v {
			count += k
		}
	}

	Expect(count).Should(Equal(325))

}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

	opl, tl := getData(Entry.PuzzleInput())

	count := 50000000000

	inLoop := false
	for !inLoop {
		pl := deepCopy(opl)
		pl = regenerate(pl, tl)

		inLoop = isShifted(opl, pl)

		opl = pl
		count--
	}

	finalcount := 0
	for k, _ := range opl {
		if opl[k] {
			finalcount += k + count
		}
	}

	fmt.Println(finalcount)
	// high 2100000000470
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
