package Day202024

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

var example1 = `sesenwnenenewseeswwswswwnenewsewsw
neeenesenwnwwswnenewnwwsewnenwseswesw
seswneswswsenwwnwse
nwnwneseeswswnenewneswwnewseswneseene
swweswneswnenwsewnwneneseenw
eesenwseswswnenwswnwnwsewwnwsene
sewnenenenesenwsewnenwwwse
wenwwweseeeweswwwnwwe
wsweesenenewnwwnwsenewsenwwsesesenwne
neeswseenwwswnwswswnw
nenwswwsewswnenenewsenwsenwnesesenew
enewnwewneswsewnwswenweswnenwsenwsw
sweneswneswneneenwnewenewwneswswnese
swwesenesewenwneswnwwneseswwne
enesenwswwswneneswsenwnewswseenwsese
wnwnesenesenenwwnenwsewesewsesesew
nenewswnwewswnenesenwnesewesw
eneswnwswnwsenenwnwnwwseeswneewsenese
neswnwewnwnwseenwseesewsenwsweewe
wseweeenwnesenwwwswnew`

func Test_ReadData(t *testing.T) {
	RegisterTestingT(t)

	data := getData(Entry.PuzzleInput())

	Expect(len(data[0])).Should(Equal(20))
}

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	data := getData(Entry.PuzzleInput())

	f := make(floor, 0)

	for i := range data {
		f.flipTile(data[i])
	}

	fmt.Println(f.countBlackTiles())
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

	data := getData(Entry.PuzzleInput())
	f := make(floor, 0)

	f.flipAllTiles(data)

	for i := 0; i < 100; i++ {
		f = f.automata()
	}

	fmt.Println(f.countBlackTiles())

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
