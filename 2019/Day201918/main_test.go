package Day201918

import (
	"strconv"
	"testing"

	. "github.com/onsi/gomega"
)

var tests = []struct {
	i string
	o string
}{
	{
		i: `#########
#b.A.@.a#
#########`,
		o: "8"},

	{
		i: `########################
#f.D.E.e.C.b.A.@.a.B.c.#
######################.#
#d.....................#
########################`,
		o: `86`,
	},

	{
		i: `########################
#...............b.C.D.f#
#.######################
#.....@.a.B.c.d.A.e.F.g#
########################`,
		o: `132`,
	},

	{
		i: `#################
#i.G..c...e..H.p#
########.########
#j.A..b...f..D.o#
########@########
#k.E..a...g..B.n#
########.########
#l.F..d...h..C.m#
#################`,
		o: `136`,
	},

	{
		i: `########################
#@..............ac.GI.b#
###d#e#f################
###A#B#C################
###g#h#i################
########################`,
		o: `81`,
	},
}

func Test_ReadData(t *testing.T) {
	RegisterTestingT(t)

	for index, test := range tests {

		if index != 3 {
			continue
		}
		tunnels, objects, player := getData(test.i)

		routes := solve(tunnels, objects, player, 1, 0)

		min := -1
		for _, val := range routes {
			if min == -1 || val < min {
				min = val
			}
		}

		Expect(strconv.Itoa(min)).Should(Equal(test.o))
	}
}

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

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
