package Day201816

import (
	"fmt"
	"strings"

	"github.com/pezza/advent-of-code/puzzle-support/chronalcompiler"
)

type dayEntry bool

// Boiler Plate
var Entry dayEntry

func (td dayEntry) Describe() (int, int, string, int) {
	return 2018, 16, "Chronal Classification", 2
}

// Start
type test struct {
	before chronalcompiler.RegisterSet
	codes  []int
	after  chronalcompiler.RegisterSet
}

func (t test) String() string {
	return fmt.Sprintf("%v -> %v -> %v", t.before.String(), t.codes, t.after.String())
}

func getData(input string) ([]test, [][]int) {
	tests := make([]test, 0)
	codes := make([][]int, 0)
	lines := strings.Split(input, "\n")
	for i := 0; i < len(lines); i++ {
		if strings.HasPrefix(lines[i], "Before") {
			a1, b1, c1, d1 := 0, 0, 0, 0
			fmt.Sscanf(lines[i], "Before: [%d, %d, %d, %d]", &a1, &b1, &c1, &d1)
			i++
			a2, b2, c2, d2 := 0, 0, 0, 0
			fmt.Sscanf(lines[i], "%d %d %d %d", &a2, &b2, &c2, &d2)
			i++
			a3, b3, c3, d3 := 0, 0, 0, 0
			fmt.Sscanf(lines[i], "After:  [%d, %d, %d, %d]", &a3, &b3, &c3, &d3)

			tests = append(tests, test{
				chronalcompiler.RegisterSet{0: a1, 1: b1, 2: c1, 3: d1},
				[]int{a2, b2, c2, d2},
				chronalcompiler.RegisterSet{0: a3, 1: b3, 2: c3, 3: d3}})
			continue
		}

		if lines[i] != "" {
			a, b, c, d := 0, 0, 0, 0
			fmt.Sscanf(lines[i], "%d %d %d %d", &a, &b, &c, &d)
			codes = append(codes, []int{a, b, c, d})
		}
	}

	return tests, codes
}

func getTestMatches(ocs chronalcompiler.OpCodes, t test) []string {
	codes := make([]string, 0)

	for k, v := range ocs {

		regSet := t.before.DeepCopy()
		regSet[t.codes[3]] = v.Process(regSet, t.codes[1], t.codes[2])

		if t.after.Same(regSet) {
			codes = append(codes, k)
		}
	}

	return codes
}

func getUniques(codes chronalcompiler.OpCodes, tests []test, uniques map[int]string) {
	for _, test := range tests {
		results := getTestMatches(codes, test)

		if len(results) == 1 {
			uniques[test.codes[0]] = results[0]
			return
		}
	}
}

func determineCodeMap(codes chronalcompiler.OpCodes, tests []test) map[int]string {
	codeMap := make(map[int]string, 0)

	for len(codes) > 0 {
		getUniques(codes, tests, codeMap)

		for _, v := range codeMap {
			delete(codes, v)
		}
	}

	return codeMap
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	tests, _ := getData(inputData)

	count := 0
	for index := range tests {
		results := getTestMatches(chronalcompiler.GetOpCodes(), tests[index])

		if len(results) >= 3 {
			count++
		}
	}

	return fmt.Sprintf("%v", count)
}

func convertCodeMapToProgram(codes map[int]string, rawProgram [][]int) []chronalcompiler.Instruction {
	prog := make([]chronalcompiler.Instruction, len(rawProgram))

	for i := range rawProgram {
		prog[i] = chronalcompiler.Instruction{OpCode: codes[rawProgram[i][0]], A: rawProgram[i][1], B: rawProgram[i][2], C: rawProgram[i][3]}
	}

	return prog
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	tests, rawProgram := getData(inputData)

	// get the instruction lookup list
	insLookup := determineCodeMap(chronalcompiler.GetOpCodes(), tests)

	// convert the raw codes into actual instructions
	program := convertCodeMapToProgram(insLookup, rawProgram)

	// run
	endRegSet := chronalcompiler.RunProgram(program, 4, nil, nil)

	return fmt.Sprint(endRegSet[0])
}
