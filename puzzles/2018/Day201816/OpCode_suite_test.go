package Day201816

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("OpCodes", func() {

	It("Should be able to process all the op-codes", func() {

		ops := getOpCodes()

		By("addr")
		reg := make(registerSet)
		reg[0] = 1
		reg[1] = 1

		// Reg 0 + Reg 1 = 2
		Expect(ops["addr"].process(reg.deepCopy(), 0, 1, 3)).Should(
			Equal(registerSet{0: 1, 1: 1, 3: 2}))

		Expect(ops["addr"].process(reg.deepCopy(), 0, 1, 1)).Should(
			Equal(registerSet{0: 1, 1: 2}))

		By("addi")
		reg = make(registerSet)
		reg[0] = 1
		reg[1] = 4

		// Reg 0 + Reg 1 = (4 + 4) = 8
		Expect(ops["addi"].process(reg.deepCopy(), 1, 4, 3)).Should(
			Equal(registerSet{0: 1, 1: 4, 3: 8}))

		Expect(ops["addi"].process(reg.deepCopy(), 0, 1, 3)).Should(
			Equal(registerSet{0: 1, 1: 4, 3: 2}))

		By("mulr")
		reg = make(registerSet)
		reg[0] = 1
		reg[1] = 1

		// Reg 0 + Reg 1 = 2
		Expect(ops["mulr"].process(reg.deepCopy(), 0, 1, 3)).Should(
			Equal(registerSet{0: 1, 1: 1, 3: 1}))

		Expect(ops["mulr"].process(reg.deepCopy(), 0, 1, 1)).Should(
			Equal(registerSet{0: 1, 1: 1}))

		By("muli")
		reg = make(registerSet)
		reg[0] = 1
		reg[1] = 4

		// Reg 0 + Reg 1 = (4 + 4) = 8
		Expect(ops["muli"].process(reg.deepCopy(), 1, 4, 3)).Should(
			Equal(registerSet{0: 1, 1: 4, 3: 16}))

		Expect(ops["muli"].process(reg.deepCopy(), 0, 10, 3)).Should(
			Equal(registerSet{0: 1, 1: 4, 3: 10}))

		By("setr")
		reg = make(registerSet)
		reg[0] = 1
		reg[1] = 4

		Expect(ops["setr"].process(reg.deepCopy(), 1, -1, 2)).Should(
			Equal(registerSet{0: 1, 1: 4, 2: 4}))

		Expect(ops["setr"].process(reg.deepCopy(), 0, -1, 1)).Should(
			Equal(registerSet{0: 1, 1: 1}))

		By("seti")
		reg = make(registerSet)
		reg[0] = 5
		reg[1] = 22

		Expect(ops["seti"].process(reg.deepCopy(), 1, -1, 0)).Should(
			Equal(registerSet{0: 1, 1: 22}))

		Expect(ops["seti"].process(reg.deepCopy(), -1, -1, 2)).Should(
			Equal(registerSet{0: 5, 1: 22, 2: -1}))

		By("banr")
		reg = make(registerSet)
		reg[0] = 5
		reg[1] = 3

		Expect(ops["banr"].process(reg.deepCopy(), 0, 1, 3)).Should(
			Equal(registerSet{0: 5, 1: 3, 3: 1}))

		Expect(ops["banr"].process(reg.deepCopy(), 0, 1, 0)).Should(
			Equal(registerSet{0: 1, 1: 3}))

		By("bani")
		reg = make(registerSet)
		reg[0] = 3
		reg[1] = 11

		Expect(ops["bani"].process(reg.deepCopy(), 0, 2, 2)).Should(
			Equal(registerSet{0: 3, 1: 11, 2: 2}))

		Expect(ops["bani"].process(reg.deepCopy(), 0, 2, 0)).Should(
			Equal(registerSet{0: 2, 1: 11}))

		By("borr")
		reg = make(registerSet)
		reg[0] = 5
		reg[1] = 3

		Expect(ops["borr"].process(reg.deepCopy(), 0, 1, 3)).Should(
			Equal(registerSet{0: 5, 1: 3, 3: 7}))

		Expect(ops["borr"].process(reg.deepCopy(), 0, 1, 0)).Should(
			Equal(registerSet{0: 7, 1: 3}))

		By("bori")
		reg = make(registerSet)
		reg[0] = 3
		reg[1] = 11

		Expect(ops["bori"].process(reg.deepCopy(), 0, 5, 2)).Should(
			Equal(registerSet{0: 3, 1: 11, 2: 7}))

		Expect(ops["bori"].process(reg.deepCopy(), 0, 5, 0)).Should(
			Equal(registerSet{0: 7, 1: 11}))

		By("gtir")
		reg = make(registerSet)
		reg[0] = 4
		reg[1] = 5
		Expect(ops["gtir"].process(reg.deepCopy(), 5, 1, 2)).Should(
			Equal(registerSet{0: 4, 1: 5, 2: 0}))
		Expect(ops["gtir"].process(reg.deepCopy(), 6, 1, 2)).Should(
			Equal(registerSet{0: 4, 1: 5, 2: 1}))
		Expect(ops["gtir"].process(reg.deepCopy(), 4, 1, 2)).Should(
			Equal(registerSet{0: 4, 1: 5, 2: 0}))

		By("gtri")
		reg = make(registerSet)
		reg[0] = 4
		reg[1] = 5
		Expect(ops["gtri"].process(reg.deepCopy(), 0, 3, 2)).Should(
			Equal(registerSet{0: 4, 1: 5, 2: 1}))
		Expect(ops["gtri"].process(reg.deepCopy(), 0, 4, 2)).Should(
			Equal(registerSet{0: 4, 1: 5, 2: 0}))
		Expect(ops["gtri"].process(reg.deepCopy(), 0, 5, 2)).Should(
			Equal(registerSet{0: 4, 1: 5, 2: 0}))

		By("gtrr")
		reg = make(registerSet)
		reg[0] = 4
		reg[1] = 5
		reg[2] = 5
		Expect(ops["gtrr"].process(reg.deepCopy(), 0, 1, 3)).Should(
			Equal(registerSet{0: 4, 1: 5, 2: 5, 3: 0}))
		Expect(ops["gtrr"].process(reg.deepCopy(), 1, 0, 3)).Should(
			Equal(registerSet{0: 4, 1: 5, 2: 5, 3: 1}))
		Expect(ops["gtrr"].process(reg.deepCopy(), 1, 2, 3)).Should(
			Equal(registerSet{0: 4, 1: 5, 2: 5, 3: 0}))

		By("eqir")
		reg = make(registerSet)
		reg[0] = 4
		reg[1] = 5
		reg[2] = 5
		Expect(ops["eqir"].process(reg.deepCopy(), 5, 1, 3)).Should(
			Equal(registerSet{0: 4, 1: 5, 2: 5, 3: 1}))
		Expect(ops["eqir"].process(reg.deepCopy(), 4, 1, 3)).Should(
			Equal(registerSet{0: 4, 1: 5, 2: 5, 3: 0}))

		By("eqri")
		reg = make(registerSet)
		reg[0] = 4
		reg[1] = 5
		reg[2] = 5
		Expect(ops["eqri"].process(reg.deepCopy(), 0, 4, 3)).Should(
			Equal(registerSet{0: 4, 1: 5, 2: 5, 3: 1}))
		Expect(ops["eqri"].process(reg.deepCopy(), 1, 4, 3)).Should(
			Equal(registerSet{0: 4, 1: 5, 2: 5, 3: 0}))

		By("eqrr")
		reg = make(registerSet)
		reg[0] = 4
		reg[1] = 5
		reg[2] = 5
		Expect(ops["eqrr"].process(reg.deepCopy(), 0, 1, 3)).Should(
			Equal(registerSet{0: 4, 1: 5, 2: 5, 3: 0}))
		Expect(ops["eqrr"].process(reg.deepCopy(), 1, 2, 3)).Should(
			Equal(registerSet{0: 4, 1: 5, 2: 5, 3: 1}))

		tests, _ := getData(Entry.PuzzleInput())

		testList(ops, tests)

		fmt.Println("NOT 527")
	})
})
