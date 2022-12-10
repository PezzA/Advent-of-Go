package Day202210

import (
	"fmt"
	"strconv"
	"strings"
)

type instruction struct {
	ins string
	val int
}

func getData(input string) []instruction {
	lines := strings.Split(input, "\n")
	insList := make([]instruction, 0)

	for _, line := range lines {
		bits := strings.Split(line, " ")

		if bits[0] == "addx" {
			val, _ := strconv.Atoi(bits[1])
			insList = append(insList, instruction{"addx", val})
			continue
		}

		insList = append(insList, instruction{"noop", 0})
	}
	return insList
}

func runProg(prog []instruction) map[int]int {
	cycle, reg := 1, 1
	history := make(map[int]int, 0)
	for _, ins := range prog {
		history[cycle] = reg
		cycle++

		if ins.ins == "addx" {
			reg += ins.val
			history[cycle] = reg
			cycle++
		}
	}
	return history
}

func runScan(history map[int]int) string {
	output := ""
	for i := 1; i < len(history); i++ {

		scan := i % 40

		if scan == history[i] || scan == history[i]-1 || scan == history[i]+1 {
			output += "#"
			continue
		}
		output += "."
	}
	return output
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	program := getData(inputData)
	history := runProg(program)

	ss := (history[20-1] * 20) + (history[60-1] * 60) + (history[100-1] * 100) + (history[140-1] * 140) + (history[180-1] * 180) + (history[220-1] * 220)
	return fmt.Sprint(ss)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	return fmt.Sprintf("%v", " -- Not Yet Implemented --")
}
