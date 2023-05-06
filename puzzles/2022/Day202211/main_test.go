package Day202211

import (
	"fmt"
	"math/big"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/pezza/advent-of-code/puzzles/common"
)

const sampleData1 string = `Monkey 0:
  Starting items: 79, 98
  Operation: new = old * 19
  Test: divisible by 23
    If true: throw to monkey 2
    If false: throw to monkey 3

Monkey 1:
  Starting items: 54, 65, 75, 74
  Operation: new = old + 6
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 0

Monkey 2:
  Starting items: 79, 60, 97
  Operation: new = old * old
  Test: divisible by 13
    If true: throw to monkey 1
    If false: throw to monkey 3

Monkey 3:
  Starting items: 74
  Operation: new = old + 3
  Test: divisible by 17
    If true: throw to monkey 0
    If false: throw to monkey 1`

func Test_ReadData(t *testing.T) {
	RegisterTestingT(t)

	monkeys := getData(sampleData1)

	Expect(len(monkeys)).Should(Equal(4))

	Expect(monkeys[0].items).Should(Equal([]int{79, 98}))
	Expect(monkeys[0].op).Should(Equal(operation{"*", big.NewInt(int64(19))}))
	Expect(monkeys[0].test).Should(Equal(23))
	Expect(monkeys[0].trueTarget).Should(Equal(2))
	Expect(monkeys[0].falseTarget).Should(Equal(3))

	Expect(monkeys[2].op).Should(Equal(operation{"^", big.NewInt(int64(2))}))

	Expect(monkeys[len(monkeys)-1].items).Should(Equal([]int{74}))
	Expect(monkeys[len(monkeys)-1].op).Should(Equal(operation{"+", big.NewInt(int64(3))}))
	Expect(monkeys[len(monkeys)-1].test).Should(Equal(17))
	Expect(monkeys[len(monkeys)-1].trueTarget).Should(Equal(0))
	Expect(monkeys[len(monkeys)-1].falseTarget).Should(Equal(1))
}

func Test_Round(t *testing.T) {
	RegisterTestingT(t)

	monkeys := getData(sampleData1)

	fmt.Println(monkeys)
	for i := 0; i < 20; i++ {
		monkeys = runRound(monkeys, true)
	}

	fmt.Println(monkeys)
	//	Expect(Entry.PartOne(Entry.PuzzleInput(), nil)).Should(Equal(common.NOT_IMPLEMENTED))
}

func Test_BigRound(t *testing.T) {
	RegisterTestingT(t)

	monkeys := getData(sampleData1)

	fmt.Println(monkeys)
	for i := 0; i < 10000; i++ {
		monkeys = runRound(monkeys, false)
	}

	fmt.Println(monkeys)
	//	Expect(Entry.PartOne(Entry.PuzzleInput(), nil)).Should(Equal(common.NOT_IMPLEMENTED))
}
func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)
	Expect(Entry.PartOne(Entry.PuzzleInput(), nil)).Should(Equal("90882"))
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)
	// 31176137622 too high
	Expect(Entry.PartTwo(Entry.PuzzleInput(), nil)).Should(Equal(common.NOT_IMPLEMENTED))
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
