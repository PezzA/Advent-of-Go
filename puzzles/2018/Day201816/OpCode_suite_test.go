package Day201816

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("OpCodes", func() {

	It("Should be able to work out all opcodes", func() {
		ops := getOpCodes()
		testCodeDebug(ops, createQuick(`Before: [0, 1, 2, 1]
5 1 3 1
After:  [0, 1, 2, 1]`))
	})

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

		Expect(len(testCode(ops, createQuick(`Before: [3, 2, 1, 1]
9 2 1 2
After:  [3, 2, 2, 1]`)))).Should(Equal(3))

		Expect(len(testCode(ops, createQuick(`Before: [2, 1, 2, 3]
1 3 3 2
After:  [2, 1, 3, 3]
`)))).Should(Equal(6))

		Expect(len(testCode(ops, createQuick(`Before: [2, 3, 2, 3]
14 3 2 3
After:  [2, 3, 2, 2]`)))).Should(Equal(2))

		Expect(len(testCode(ops, createQuick(`Before: [3, 3, 1, 1]
8 2 3 1
After:  [3, 1, 1, 1]`)))).Should(Equal(7))

		Expect(len(testCode(ops, createQuick(`Before: [0, 1, 0, 0]
0 3 1 3
After:  [0, 1, 0, 1]`)))).Should(Equal(5))

		Expect(len(testCode(ops, createQuick(`Before: [2, 1, 1, 2]
15 0 1 1
After:  [2, 3, 1, 2]`)))).Should(Equal(4))

		Expect(len(testCode(ops, createQuick(`Before: [0, 2, 2, 1]
15 0 2 1
After:  [0, 2, 2, 1]`)))).Should(Equal(4))

		Expect(len(testCode(ops, createQuick(`Before: [0, 2, 1, 0]
3 0 0 1
After:  [0, 0, 1, 0]`)))).Should(Equal(13))

		Expect(len(testCode(ops, createQuick(`Before: [3, 1, 2, 0]
0 3 1 2
After:  [3, 1, 1, 0]`)))).Should(Equal(5))

		Expect(len(testCode(ops, createQuick(`Before: [0, 1, 2, 1]
5 1 3 1
After:  [0, 1, 2, 1]`)))).Should(Equal(8))

		Expect(len(testCode(ops, createQuick(`Before: [2, 2, 2, 3]
11 3 3 1
After:  [2, 9, 2, 3]`)))).Should(Equal(2))

		Expect(len(testCode(ops, createQuick(`Before: [2, 1, 0, 3]
14 2 1 2
After:  [2, 1, 1, 3]`)))).Should(Equal(5))

		fmt.Println("NOT 527")
	})
})

/*






Before: [0, 3, 2, 2]
4 0 2 0
After:  [0, 3, 2, 2]
// 12


Before: [1, 0, 3, 3]
12 3 1 1
After:  [1, 3, 3, 3]
// 6


Before: [0, 2, 3, 2]
7 2 3 3
After:  [0, 2, 3, 2]
// 2


Before: [1, 1, 2, 3]
10 1 3 3
After:  [1, 1, 2, 3]
// 4


Before: [2, 2, 2, 1]
6 3 2 1
After:  [2, 3, 2, 1]
// 5


Before: [2, 1, 2, 0]
2 0 3 1
After:  [2, 3, 2, 0]
// 1


Before: [1, 2, 3, 3]
11 2 3 1
After:  [1, 9, 3, 3]
// 2


Before: [2, 1, 2, 2]
13 1 3 2
After:  [2, 1, 3, 2]

*/
func createQuick(input string) test {
	a1, b1, c1, d1, a2, b2, c2, d2, a3, b3, c3, d3 := 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0

	fmt.Sscanf(input, `Before: [%d, %d, %d, %d]
%d %d %d %d
After:  [%d, %d, %d, %d]`, &a1, &b1, &c1, &d1, &a2, &b2, &c2, &d2, &a3, &b3, &c3, &d3)

	return test{
		registerSet{0: a1, 1: b1, 2: c1, 3: d1},
		[]int{a2, b2, c2, d2},
		registerSet{0: a3, 1: b3, 2: c3, 3: d3},
	}
}
