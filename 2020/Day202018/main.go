package Day202018

import (
	"fmt"
	"strconv"
	"strings"
)

func getData(input string) []string {
	return strings.Split(input, "\n")
}

func getDepth(input string) int {
	currDepth, maxDepth := 0, 0

	for i := range input {
		if input[i] == 40 {
			currDepth++
			if currDepth > maxDepth {
				maxDepth = currDepth
			}
		}

		if input[i] == 41 {
			currDepth--
		}
	}
	return maxDepth
}

func getExpressionAtDepth(input string, depth int) (string, int, int) {
	currDepth := 0
	track := false
	expression := ""
	start, end := 0, 0

	for i := range input {
		if input[i] == 40 {
			currDepth++
			if currDepth == depth {
				track = true
				start = i + 1
				continue
			}
		}
		// not worried about inner expressions as
		// always getting the deepest.
		if input[i] == 41 {
			currDepth--
			if track {
				end = i - 1
				break
			}
		}

		if track {
			expression += string(input[i])
		}
	}
	return expression, start, end
}

func findPlus(input string) int {
	bits := strings.Split(input, " ")

	for i := range bits {
		if bits[i] == "+" {
			return i
		}
	}
	return 0
}

func resolvePlus(input string) string {
	for index := findPlus(input); index != 0; index = findPlus(input) {
		bits := strings.Split(input, " ")

		left, right := bits[index-1], bits[index+1]

		leftNum, _ := strconv.Atoi(left)
		rightNum, _ := strconv.Atoi(right)

		sum := leftNum + rightNum

		newBits := bits[:index-1]
		newBits = append(newBits, strconv.Itoa(sum))
		newBits = append(newBits, bits[index+2:]...)

		input = strings.Join(newBits, " ")
	}

	return input
}

func evaluateExpression(input string, advanced bool) int {
	if advanced {
		input = resolvePlus(input)
	}

	bits := strings.Split(input, " ")

	acc := 0
	op := ""
	for _, bit := range bits {
		if bit == "+" || bit == "*" {
			op = bit
			continue
		}

		newNum, _ := strconv.Atoi(bit)

		switch op {
		case "":
			acc = newNum
		case "+":
			acc += newNum
		case "*":
			acc *= newNum
		}
	}

	return acc
}

func evaluateLine(input string, advanced bool) int {
	for depth := getDepth(input); depth != 0; depth = getDepth(input) {
		exp, start, end := getExpressionAtDepth(input, depth)
		replacement := evaluateExpression(exp, advanced)
		input = input[:start-1] + strconv.Itoa(replacement) + input[end+2:]
	}

	return evaluateExpression(input, advanced)
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	data := getData(inputData)

	tot := 0
	for i := range data {
		fmt.Println(data[i])
		tot += evaluateLine(data[i], false)
	}

	return fmt.Sprintf("%v", tot)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	data := getData(inputData)

	tot := 0
	for i := range data {
		fmt.Println(data[i])
		tot += evaluateLine(data[i], true)
	}

	return fmt.Sprintf("%v", tot)
}
