package Day201816

import (
	"fmt"
	"strings"

	"github.com/pezza/advent-of-code/puzzles/2018/ChronalCompiler"
)

// Boiler Plate
var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string) {
	return 2018, 16, "Chronal Classification"
}

// Start
type test struct {
	before ChronalCompiler.RegisterSet
	codes  []int
	after  ChronalCompiler.RegisterSet
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
				ChronalCompiler.RegisterSet{0: a1, 1: b1, 2: c1, 3: d1},
				[]int{a2, b2, c2, d2},
				ChronalCompiler.RegisterSet{0: a3, 1: b3, 2: c3, 3: d3}})
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

func getTestMatches(ocs ChronalCompiler.OpCodes, t test) []string {
	codes := make([]string, 0)

	for k, v := range ocs {
		if t.after.Same(v.Process(t.before.DeepCopy(), t.codes[1], t.codes[2], t.codes[3])) {
			codes = append(codes, k)
		}
	}

	return codes
}

func getUniques(codes ChronalCompiler.OpCodes, tests []test, uniques map[int]string) map[int]string {
	for _, test := range tests {
		results := getTestMatches(codes, test)

		if len(results) == 1 {
			uniques[test.codes[0]] = results[0]
		}
	}

	return uniques
}

func determineCodeMap(codes ChronalCompiler.OpCodes, tests []test) map[int]string {
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
		results := getTestMatches(ChronalCompiler.GetOpCodes(), tests[index])

		if len(results) >= 3 {

			count++
		}
	}

	return fmt.Sprintf("%v", count)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {

	return fmt.Sprintf(" -- Not Yet Implemented --")
}
