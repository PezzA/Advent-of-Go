package Day201816

import (
	"fmt"
	"strings"
)

var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string) {
	return 2018, 16, "Chronal Classification"
}

type registerSet map[int]int

func (rs registerSet) String() string {
	return fmt.Sprintf("[%v %v %v %v]", rs[0], rs[1], rs[2], rs[3])
}

func (t test) String() string {
	return fmt.Sprintf("%v -> %v -> %v", t.before.String(), t.codes, t.after.String())
}

type processor interface {
	process(input registerSet, a int, b int, c int) registerSet
}

type opCodes map[string]processor

// create the 16 types
type addr bool
type addi bool
type mulr bool
type muli bool
type setr bool
type seti bool
type banr bool
type bani bool
type borr bool
type bori bool
type gtir bool
type gtri bool
type gtrr bool
type eqir bool
type eqri bool
type eqrr bool

func (op addr) process(input registerSet, a int, b int, c int) registerSet {
	input[c] = input[a] + input[b]
	return input
}

func (op addi) process(input registerSet, a int, b int, c int) registerSet {
	input[c] = input[a] + b
	return input
}

func (op mulr) process(input registerSet, a int, b int, c int) registerSet {
	input[c] = input[a] * input[b]
	return input
}

func (op muli) process(input registerSet, a int, b int, c int) registerSet {
	input[c] = input[a] * b
	return input
}

func (op setr) process(input registerSet, a int, b int, c int) registerSet {
	input[c] = input[a]
	return input
}

func (op seti) process(input registerSet, a int, b int, c int) registerSet {
	input[c] = a
	return input
}

func (op banr) process(input registerSet, a int, b int, c int) registerSet {
	input[c] = input[a] & input[b]
	return input
}

func (op bani) process(input registerSet, a int, b int, c int) registerSet {
	input[c] = input[a] & b
	return input
}

func (op borr) process(input registerSet, a int, b int, c int) registerSet {
	input[c] = input[a] | input[b]
	return input
}

func (op bori) process(input registerSet, a int, b int, c int) registerSet {
	input[c] = input[a] | b
	return input
}

func (op gtir) process(input registerSet, a int, b int, c int) registerSet {
	if a > input[b] {
		input[c] = 1
	} else {
		input[c] = 0
	}
	return input
}

func (op gtri) process(input registerSet, a int, b int, c int) registerSet {
	if input[a] > b {
		input[c] = 1
	} else {
		input[c] = 0
	}
	return input
}

func (op gtrr) process(input registerSet, a int, b int, c int) registerSet {
	if input[a] > input[b] {
		input[c] = 1
	} else {
		input[c] = 0
	}
	return input
}

func (op eqir) process(input registerSet, a int, b int, c int) registerSet {
	if a == input[b] {
		input[c] = 1
	} else {
		input[c] = 0
	}
	return input
}

func (op eqri) process(input registerSet, a int, b int, c int) registerSet {
	if input[a] == b {
		input[c] = 1
	} else {
		input[c] = 0
	}
	return input
}

func (op eqrr) process(input registerSet, a int, b int, c int) registerSet {
	if input[a] == input[b] {
		input[c] = 1
	} else {
		input[c] = 0
	}
	return input
}

func getOpCodes() opCodes {
	ops := make(opCodes)
	ops["addr"] = new(addr)
	ops["addi"] = new(addi)
	ops["mulr"] = new(mulr)
	ops["muli"] = new(muli)
	ops["setr"] = new(setr)
	ops["seti"] = new(seti)
	ops["banr"] = new(banr)
	ops["bani"] = new(bani)
	ops["borr"] = new(borr)
	ops["bori"] = new(bori)
	ops["gtir"] = new(gtir)
	ops["gtri"] = new(gtri)
	ops["gtrr"] = new(gtrr)
	ops["eqir"] = new(eqir)
	ops["eqri"] = new(eqri)
	ops["eqrr"] = new(eqrr)

	return ops
}

var opCodeSort = []string{
	"addr",
	"addi",
	"mulr",
	"muli",
	"setr",
	"seti",
	"banr",
	"bani",
	"borr",
	"bori",
	"gtir",
	"gtri",
	"gtrr",
	"eqir",
	"eqri",
	"eqrr",
}

type test struct {
	before registerSet
	codes  []int
	after  registerSet
}

func (rs registerSet) deepCopy() registerSet {
	retVal := make(registerSet)
	for k, v := range rs {
		retVal[k] = v
	}
	return retVal
}

func getData(input string) ([]test, [][]int) {
	tests := make([]test, 0)
	codes := make([][]int, 0)
	lines := strings.Split(input, "\n")
	for index, line := range lines {
		if strings.HasPrefix(line, "Before") {
			a1, b1, c1, d1 := 0, 0, 0, 0
			fmt.Sscanf(lines[index], "Before: [%d, %d, %d, %d]", &a1, &b1, &c1, &d1)

			a2, b2, c2, d2 := 0, 0, 0, 0
			fmt.Sscanf(lines[index+1], "%d %d %d %d", &a2, &b2, &c2, &d2)

			a3, b3, c3, d3 := 0, 0, 0, 0
			fmt.Sscanf(lines[index+2], "After:  [%d, %d, %d, %d]", &a3, &b3, &c3, &d3)

			tests = append(tests, test{
				registerSet{0: a1, 1: b1, 2: c1, 3: d1},
				[]int{a2, b2, c2, d2},
				registerSet{0: a3, 1: b3, 2: c3, 3: d3}})
		}
	}

	return tests, codes
}

func (rs registerSet) Same(cmp registerSet) bool {
	return rs[0] == cmp[0] && rs[1] == cmp[1] && rs[2] == cmp[2] && rs[3] == cmp[3]
}

func testCode(ocs opCodes, t test) []string {
	codes := make([]string, 0)

	for k, v := range ocs {
		if t.after.Same(v.process(t.before.deepCopy(), t.codes[1], t.codes[2], t.codes[3])) {
			codes = append(codes, k)
		}
	}

	return codes
}

func testCodeDebug(ocs opCodes, t test) {
	for _, opName := range opCodeSort {
		processed := ocs[opName].process(t.before.deepCopy(), t.codes[1], t.codes[2], t.codes[3])

		fmt.Println(opName, t.before, t.codes, processed, t.after, processed.Same(t.after))
	}
}

func testList(ocs opCodes, tests []test) int {
	count := 0
	for index := range tests {
		results := testCode(ocs, tests[index])

		if len(results) >= 3 {

			count++
		}
	}

	return count
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	ocs := getOpCodes()

	tests, _ := getData(inputData)

	return fmt.Sprintf("%v", testList(ocs, tests))
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	return fmt.Sprintf(" -- Not Yet Implemented --")
}
