package Day201522

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	player := character{
		class:  clPlayer,
		health: 50,
		dmg:    0,
		mana:   500,
	}

	boss := getBoss(Entry.PuzzleInput())

	outcomes := runBattle(player, boss)

	lowMan := -1
	for _, o := range outcomes {
		if o.winner == clPlayer {
			if lowMan == -1 || o.mana < lowMan {
				lowMan = o.mana
			}
		}
	}

	fmt.Println(len(outcomes))
	fmt.Println("not 1415, 1362")
	fmt.Println(lowMan)
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
