package Day201922

import (
	"fmt"
	"testing"

	"github.com/pezza/advent-of-code/puzzles/common"

	. "github.com/onsi/gomega"
)

func Test_ReadData(t *testing.T) {
	RegisterTestingT(t)

}

var testData = `deal into new stack
cut -2
deal with increment 7
cut 8
cut -4
deal with increment 7
cut 3
deal with increment 9
deal with increment 3
cut -1`

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	newDeck := getCards(10)

	Expect(newDeck.newStack()).Should(Equal(deck{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}))

	newDeck2 := getCards(10)
	Expect(newDeck2.cut(3)).Should(Equal(deck{3, 4, 5, 6, 7, 8, 9, 0, 1, 2}))

	newDeck3 := getCards(10)
	Expect(newDeck3.cut(-4)).Should(Equal(deck{6, 7, 8, 9, 0, 1, 2, 3, 4, 5}))

	newDeck4 := getCards(10)
	Expect(newDeck4.increment(3)).Should(Equal(deck{0, 7, 4, 1, 8, 5, 2, 9, 6, 3}))

	testDeck := getCards(10)

	actions := getData(testData)

	for _, a := range actions {
		switch a.action {
		case "cut":
			testDeck = testDeck.cut(a.mod)
		case "increment":
			testDeck = testDeck.increment(a.mod)
		case "new":
			testDeck = testDeck.newStack()
		}
	}

	Expect(testDeck).Should(Equal(deck{9, 2, 5, 8, 1, 4, 7, 0, 3, 6}))

	partOneDeck := getCards(10007)
	partOneActions := getData(Entry.PuzzleInput())

	for _, a := range partOneActions {
		switch a.action {
		case "cut":
			partOneDeck = partOneDeck.cut(a.mod)
		case "increment":
			partOneDeck = partOneDeck.increment(a.mod)
		case "new":
			partOneDeck = partOneDeck.newStack()
		}
	}

	for i := range partOneDeck {
		if partOneDeck[i] == 2019 {
			fmt.Println(i)
		}
	}

	fmt.Println(partOneDeck)
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

	fmt.Println(119315717514047 - 101741582076661)

	fmt.Println(119315717514047 / 10007)
	partOneDeck := getCards(10007)
	partOneDeckCmp := getCards(10007)
	partOneActions := getData(Entry.PuzzleInput())

	inc := 1
	for {
		for i := range partOneDeck {
			if partOneDeck[i] == 2020 {
				fmt.Println(i)
			}
		}
		for _, a := range partOneActions {
			switch a.action {
			case "cut":
				partOneDeck = partOneDeck.cut(a.mod)
			case "increment":
				partOneDeck = partOneDeck.increment(a.mod)
			case "new":
				partOneDeck = partOneDeck.newStack()
			}
		}

		if common.Same(partOneDeck, partOneDeckCmp) {
			fmt.Println(inc)
		}
		inc++
	}

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
