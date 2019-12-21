package Day201921

import (
	"fmt"
	"sync"

	"github.com/pezza/advent-of-code/2019/intcode"
)

const newLine = 10

func inputString(input string, in chan int64) {
	for _, char := range input {
		in <- int64(char)
	}
	in <- int64(newLine)
}

func runProgram(input string, springCode string) {
	out := make(chan int64, 0)
	in := make(chan int64, 0)
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

	fmt.Println(output)
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	return fmt.Sprintf("%v", " -- Not Yet Implemented --")
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	return fmt.Sprintf("%v", " -- Not Yet Implemented --")
}
