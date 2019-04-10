package Day201816

import (
	"fmt"

	"github.com/pezza/advent-of-code/puzzles/2018/ChronalCompiler"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("OpCodes", func() {
	It("Should be able to read the test data", func() {
		_, program := getData(Entry.PuzzleInput())

		Expect(program[0]).Should(Equal([]int{7, 3, 2, 0}))
		Expect(program[1]).Should(Equal([]int{7, 2, 1, 1}))
		Expect(program[len(program)-1]).Should(Equal([]int{1, 3, 0, 0}))
	})

	It("Should be able to determine the code map from the tests given", func() {
		tests, _ := getData(Entry.PuzzleInput())
		codeMap := determineCodeMap(ChronalCompiler.GetOpCodes(), tests)

		Expect(codeMap[11]).Should(Equal("eqrr"))
		Expect(codeMap[1]).Should(Equal("addr"))
	})

	It("Should be able to run a program", func() {
		tests, program := getData(Entry.PuzzleInput())
		codeMap := determineCodeMap(ChronalCompiler.GetOpCodes(), tests)
		startingSet := ChronalCompiler.RegisterSet{0: 0, 1: 0, 2: 0, 3: 0}

		opCodes := ChronalCompiler.GetOpCodes()

		for _, line := range program {
			startingSet = opCodes[codeMap[line[0]]].Process(startingSet, line[1], line[2], line[3])
			fmt.Println(line, codeMap[line[0]], startingSet)
		}

		fmt.Println(startingSet)
		fmt.Println("Not 617")
	})

	It("Should be able to process all the op-codes", func() {
		ops := ChronalCompiler.GetOpCodes()
		By("addr")
		reg := make(ChronalCompiler.RegisterSet)
		reg[0] = 1
		reg[1] = 1

		// Reg 0 + Reg 1 = 2
		Expect(ops["addr"].Process(reg.DeepCopy(), 0, 1, 3)).Should(
			Equal(ChronalCompiler.RegisterSet{0: 1, 1: 1, 3: 2}))

		Expect(ops["addr"].Process(reg.DeepCopy(), 0, 1, 1)).Should(
			Equal(ChronalCompiler.RegisterSet{0: 1, 1: 2}))

		By("addi")
		reg = make(ChronalCompiler.RegisterSet)
		reg[0] = 1
		reg[1] = 4

		// Reg 0 + Reg 1 = (4 + 4) = 8
		Expect(ops["addi"].Process(reg.DeepCopy(), 1, 4, 3)).Should(
			Equal(ChronalCompiler.RegisterSet{0: 1, 1: 4, 3: 8}))

		Expect(ops["addi"].Process(reg.DeepCopy(), 0, 1, 3)).Should(
			Equal(ChronalCompiler.RegisterSet{0: 1, 1: 4, 3: 2}))

		By("mulr")
		reg = make(ChronalCompiler.RegisterSet)
		reg[0] = 1
		reg[1] = 1

		// Reg 0 + Reg 1 = 2
		Expect(ops["mulr"].Process(reg.DeepCopy(), 0, 1, 3)).Should(
			Equal(ChronalCompiler.RegisterSet{0: 1, 1: 1, 3: 1}))

		Expect(ops["mulr"].Process(reg.DeepCopy(), 0, 1, 1)).Should(
			Equal(ChronalCompiler.RegisterSet{0: 1, 1: 1}))

		By("muli")
		reg = make(ChronalCompiler.RegisterSet)
		reg[0] = 1
		reg[1] = 4

		// Reg 0 + Reg 1 = (4 + 4) = 8
		Expect(ops["muli"].Process(reg.DeepCopy(), 1, 4, 3)).Should(
			Equal(ChronalCompiler.RegisterSet{0: 1, 1: 4, 3: 16}))

		Expect(ops["muli"].Process(reg.DeepCopy(), 0, 10, 3)).Should(
			Equal(ChronalCompiler.RegisterSet{0: 1, 1: 4, 3: 10}))

		By("setr")
		reg = make(ChronalCompiler.RegisterSet)
		reg[0] = 1
		reg[1] = 4

		Expect(ops["setr"].Process(reg.DeepCopy(), 1, -1, 2)).Should(
			Equal(ChronalCompiler.RegisterSet{0: 1, 1: 4, 2: 4}))

		Expect(ops["setr"].Process(reg.DeepCopy(), 0, -1, 1)).Should(
			Equal(ChronalCompiler.RegisterSet{0: 1, 1: 1}))

		By("seti")
		reg = make(ChronalCompiler.RegisterSet)
		reg[0] = 5
		reg[1] = 22

		Expect(ops["seti"].Process(reg.DeepCopy(), 1, -1, 0)).Should(
			Equal(ChronalCompiler.RegisterSet{0: 1, 1: 22}))

		Expect(ops["seti"].Process(reg.DeepCopy(), -1, -1, 2)).Should(
			Equal(ChronalCompiler.RegisterSet{0: 5, 1: 22, 2: -1}))

		By("banr")
		reg = make(ChronalCompiler.RegisterSet)
		reg[0] = 5
		reg[1] = 3

		Expect(ops["banr"].Process(reg.DeepCopy(), 0, 1, 3)).Should(
			Equal(ChronalCompiler.RegisterSet{0: 5, 1: 3, 3: 1}))

		Expect(ops["banr"].Process(reg.DeepCopy(), 0, 1, 0)).Should(
			Equal(ChronalCompiler.RegisterSet{0: 1, 1: 3}))

		By("bani")
		reg = make(ChronalCompiler.RegisterSet)
		reg[0] = 3
		reg[1] = 11

		Expect(ops["bani"].Process(reg.DeepCopy(), 0, 2, 2)).Should(
			Equal(ChronalCompiler.RegisterSet{0: 3, 1: 11, 2: 2}))

		Expect(ops["bani"].Process(reg.DeepCopy(), 0, 2, 0)).Should(
			Equal(ChronalCompiler.RegisterSet{0: 2, 1: 11}))

		By("borr")
		reg = make(ChronalCompiler.RegisterSet)
		reg[0] = 5
		reg[1] = 3

		Expect(ops["borr"].Process(reg.DeepCopy(), 0, 1, 3)).Should(
			Equal(ChronalCompiler.RegisterSet{0: 5, 1: 3, 3: 7}))

		Expect(ops["borr"].Process(reg.DeepCopy(), 0, 1, 0)).Should(
			Equal(ChronalCompiler.RegisterSet{0: 7, 1: 3}))

		By("bori")
		reg = make(ChronalCompiler.RegisterSet)
		reg[0] = 3
		reg[1] = 11

		Expect(ops["bori"].Process(reg.DeepCopy(), 0, 5, 2)).Should(
			Equal(ChronalCompiler.RegisterSet{0: 3, 1: 11, 2: 7}))

		Expect(ops["bori"].Process(reg.DeepCopy(), 0, 5, 0)).Should(
			Equal(ChronalCompiler.RegisterSet{0: 7, 1: 11}))

		By("gtir")
		reg = make(ChronalCompiler.RegisterSet)
		reg[0] = 4
		reg[1] = 5
		Expect(ops["gtir"].Process(reg.DeepCopy(), 5, 1, 2)).Should(
			Equal(ChronalCompiler.RegisterSet{0: 4, 1: 5, 2: 0}))
		Expect(ops["gtir"].Process(reg.DeepCopy(), 6, 1, 2)).Should(
			Equal(ChronalCompiler.RegisterSet{0: 4, 1: 5, 2: 1}))
		Expect(ops["gtir"].Process(reg.DeepCopy(), 4, 1, 2)).Should(
			Equal(ChronalCompiler.RegisterSet{0: 4, 1: 5, 2: 0}))

		By("gtri")
		reg = make(ChronalCompiler.RegisterSet)
		reg[0] = 4
		reg[1] = 5
		Expect(ops["gtri"].Process(reg.DeepCopy(), 0, 3, 2)).Should(
			Equal(ChronalCompiler.RegisterSet{0: 4, 1: 5, 2: 1}))
		Expect(ops["gtri"].Process(reg.DeepCopy(), 0, 4, 2)).Should(
			Equal(ChronalCompiler.RegisterSet{0: 4, 1: 5, 2: 0}))
		Expect(ops["gtri"].Process(reg.DeepCopy(), 0, 5, 2)).Should(
			Equal(ChronalCompiler.RegisterSet{0: 4, 1: 5, 2: 0}))

		By("gtrr")
		reg = make(ChronalCompiler.RegisterSet)
		reg[0] = 4
		reg[1] = 5
		reg[2] = 5
		Expect(ops["gtrr"].Process(reg.DeepCopy(), 0, 1, 3)).Should(
			Equal(ChronalCompiler.RegisterSet{0: 4, 1: 5, 2: 5, 3: 0}))
		Expect(ops["gtrr"].Process(reg.DeepCopy(), 1, 0, 3)).Should(
			Equal(ChronalCompiler.RegisterSet{0: 4, 1: 5, 2: 5, 3: 1}))
		Expect(ops["gtrr"].Process(reg.DeepCopy(), 1, 2, 3)).Should(
			Equal(ChronalCompiler.RegisterSet{0: 4, 1: 5, 2: 5, 3: 0}))

		By("eqir")
		reg = make(ChronalCompiler.RegisterSet)
		reg[0] = 4
		reg[1] = 5
		reg[2] = 5
		Expect(ops["eqir"].Process(reg.DeepCopy(), 5, 1, 3)).Should(
			Equal(ChronalCompiler.RegisterSet{0: 4, 1: 5, 2: 5, 3: 1}))
		Expect(ops["eqir"].Process(reg.DeepCopy(), 4, 1, 3)).Should(
			Equal(ChronalCompiler.RegisterSet{0: 4, 1: 5, 2: 5, 3: 0}))

		By("eqri")
		reg = make(ChronalCompiler.RegisterSet)
		reg[0] = 4
		reg[1] = 5
		reg[2] = 5
		Expect(ops["eqri"].Process(reg.DeepCopy(), 0, 4, 3)).Should(
			Equal(ChronalCompiler.RegisterSet{0: 4, 1: 5, 2: 5, 3: 1}))
		Expect(ops["eqri"].Process(reg.DeepCopy(), 1, 4, 3)).Should(
			Equal(ChronalCompiler.RegisterSet{0: 4, 1: 5, 2: 5, 3: 0}))

		By("eqrr")
		reg = make(ChronalCompiler.RegisterSet)
		reg[0] = 4
		reg[1] = 5
		reg[2] = 5
		Expect(ops["eqrr"].Process(reg.DeepCopy(), 0, 1, 3)).Should(
			Equal(ChronalCompiler.RegisterSet{0: 4, 1: 5, 2: 5, 3: 0}))
		Expect(ops["eqrr"].Process(reg.DeepCopy(), 1, 2, 3)).Should(
			Equal(ChronalCompiler.RegisterSet{0: 4, 1: 5, 2: 5, 3: 1}))

		Expect(len(getTestMatches(ops, createQuick(`Before: [3, 2, 1, 1]
9 2 1 2
After:  [3, 2, 2, 1]`)))).Should(Equal(3))

		Expect(len(getTestMatches(ops, createQuick(`Before: [2, 1, 2, 3]
1 3 3 2
After:  [2, 1, 3, 3]
`)))).Should(Equal(6))

		Expect(len(getTestMatches(ops, createQuick(`Before: [2, 3, 2, 3]
14 3 2 3
After:  [2, 3, 2, 2]`)))).Should(Equal(2))

		Expect(len(getTestMatches(ops, createQuick(`Before: [3, 3, 1, 1]
8 2 3 1
After:  [3, 1, 1, 1]`)))).Should(Equal(7))

		Expect(len(getTestMatches(ops, createQuick(`Before: [0, 1, 0, 0]
0 3 1 3
After:  [0, 1, 0, 1]`)))).Should(Equal(5))

		Expect(len(getTestMatches(ops, createQuick(`Before: [2, 1, 1, 2]
15 0 1 1
After:  [2, 3, 1, 2]`)))).Should(Equal(4))

		Expect(len(getTestMatches(ops, createQuick(`Before: [0, 2, 2, 1]
15 0 2 1
After:  [0, 2, 2, 1]`)))).Should(Equal(4))

		Expect(len(getTestMatches(ops, createQuick(`Before: [0, 2, 1, 0]
3 0 0 1
After:  [0, 0, 1, 0]`)))).Should(Equal(13))

		Expect(len(getTestMatches(ops, createQuick(`Before: [3, 1, 2, 0]
0 3 1 2
After:  [3, 1, 1, 0]`)))).Should(Equal(5))

		Expect(len(getTestMatches(ops, createQuick(`Before: [0, 1, 2, 1]
5 1 3 1
After:  [0, 1, 2, 1]`)))).Should(Equal(8))

		Expect(len(getTestMatches(ops, createQuick(`Before: [2, 2, 2, 3]
11 3 3 1
After:  [2, 9, 2, 3]`)))).Should(Equal(2))

		Expect(len(getTestMatches(ops, createQuick(`Before: [2, 1, 0, 3]
14 2 1 2
After:  [2, 1, 1, 3]`)))).Should(Equal(5))

	})
})

func createQuick(input string) test {
	a1, b1, c1, d1, a2, b2, c2, d2, a3, b3, c3, d3 := 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0

	fmt.Sscanf(input, `Before: [%d, %d, %d, %d]
%d %d %d %d
After:  [%d, %d, %d, %d]`, &a1, &b1, &c1, &d1, &a2, &b2, &c2, &d2, &a3, &b3, &c3, &d3)

	return test{
		ChronalCompiler.RegisterSet{0: a1, 1: b1, 2: c1, 3: d1},
		[]int{a2, b2, c2, d2},
		ChronalCompiler.RegisterSet{0: a3, 1: b3, 2: c3, 3: d3},
	}
}
