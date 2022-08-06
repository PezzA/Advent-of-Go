package Day202019

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type expression struct {
	char     string
	subRules [][]int
}

func getData(input string) (map[int]expression, []string, int, int) {
	lines := strings.Split(input, "\n")

	expressions := make(map[int]expression, 0)
	tests := make([]string, 0)

	insA, insB := 0, 0
	processingTests := false
	for _, line := range lines {
		if line == "" {
			processingTests = true
			continue
		}

		if !processingTests {
			bits := strings.Split(line, ":")
			ins, _ := strconv.Atoi(bits[0])

			if strings.TrimSpace(bits[1])[1] == 98 {
				expressions[ins] = expression{"b", [][]int{}}
				insB = ins
				continue
			}

			if strings.TrimSpace(bits[1])[1] == 97 {
				expressions[ins] = expression{"a", [][]int{}}
				insA = ins
				continue
			}

			subRules := strings.Split(bits[1], "|")

			subRuleSlice := make([][]int, len(subRules))
			for r, rule := range subRules {
				ruleBits := strings.Split(strings.TrimSpace(rule), " ")

				ruleSlice := make([]int, len(ruleBits))
				for i := range ruleBits {
					ruleSlice[i], _ = strconv.Atoi(ruleBits[i])
				}

				subRuleSlice[r] = ruleSlice
			}

			expressions[ins] = expression{"", subRuleSlice}

			continue
		}

		tests = append(tests, line)
	}

	return expressions, tests, insA, insB
}

func (e expression) getText(insA, insB int, exps map[int]expression) string {
	val := ""

	for _, sr := range e.subRules {

		if val != "" {
			val += "|"
		}

		for _, r := range sr {
			if r == insA || r == insB {
				val += exps[r].char
				continue
			}
			val += "(" + exps[r].getText(insA, insB, exps) + ")"
		}
	}

	return val
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	exp, tests, aIns, bIns := getData(inputData)

	re := regexp.MustCompile("^" + exp[0].getText(aIns, bIns, exp) + "$")
	tot := 0

	for _, t := range tests {
		if re.MatchString(t) {
			tot++
		}
	}

	return fmt.Sprintf("%v", tot)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	exp, tests, aIns, bIns := getData(inputData)

	tot := 0

	exp42 := exp[42].getText(aIns, bIns, exp)
	exp31 := exp[31].getText(aIns, bIns, exp)

	// 6 is arbitrary based on how long the input data is.
	for i := 1; i < 6; i++ {
		expression := fmt.Sprintf("^(%v)+(%v){%d}(%v){%d}$", exp42, exp42, i, exp31, i)
		re := regexp.MustCompile(expression)
		for _, t := range tests {
			if re.MatchString(t) {
				tot++
			}
		}
	}

	return fmt.Sprintf("%v", tot)
}
