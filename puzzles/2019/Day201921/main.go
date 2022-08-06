package Day201921

import (
	"fmt"
	"sync"

	"github.com/pezza/advent-of-code/puzzles/2019/intcode"
)

const newLine = 10

func inputString(input string, in chan int64) {
	for _, char := range input {
		in <- int64(char)
	}
	in <- int64(newLine)
}

func runProgram(input string, springCode string) int64 {
	in, out := make(chan int64, 0), make(chan int64, 0)
	output := int64(0)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		vm := intcode.New(input)
		vm.RunProgram(nil, nil, in, out, nil)
		close(out)
	}()

	go func() {
		defer wg.Done()
		for val := range out {
			if val > 128 {
				output = val
				continue
			}
			fmt.Print(string(val))
		}
	}()

	inputString(springCode, in)
	wg.Wait()

	return output
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	output := runProgram(inputData, `NOT A J
NOT B T
AND D T
OR T J
NOT C T
AND D T
OR T J
WALK`)

	return fmt.Sprintf("%v", output)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	output := runProgram(inputData, `NOT A J
NOT B T
AND D T
OR T J
NOT C T
AND D T
AND H T
OR T J
RUN`)
	return fmt.Sprintf("%v", output)
}
