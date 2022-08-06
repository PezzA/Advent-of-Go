package Day201925

import (
	"fmt"
	"sync"

	"github.com/pezza/advent-of-code/puzzles/2019/intcode"
)

func inputString(input string, in chan int64) {
	for _, char := range input {
		in <- int64(char)
	}
	in <- int64(10)
}

func runProgram(input string) int64 {
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
			fmt.Print(string(val))
		}
	}()

	inputString("south\ntake cake\nnorth", in)
	inputString("north\ntake mutex\nsouth", in)
	inputString("north\neast\neast\neast\ntake whirled peas\nwest\nwest\nwest\nsouth", in)
	inputString("west\ntake space law space brochure\nnorth\ntake loom\nsouth\nsouth\ntake hologram\nwest\ntake manifold", in)
	inputString("east\nnorth\neast\nsouth\nwest\nsouth\ntake easter egg\nsouth\nsouth", in)

	items := []string{"whirled peas", "hologram", "cake", "space law space brochure", "loom", "mutex", "manifold", "easter egg"}

	potentials := []string{"hologram", "space law space brochure", "loom", "mutex", "easter egg", "manifold"}
	for i := range items {
		inputString(fmt.Sprintf("drop %v", items[i]), in)
	}

	for i := 0; i < 64; i++ {
		fmt.Println("######################################", i)
		if i&1 == 1 {
			inputString(fmt.Sprintf("take %v", potentials[0]), in)
		}

		if i&2 == 2 {
			inputString(fmt.Sprintf("take %v", potentials[1]), in)
		}

		if i&4 == 4 {
			inputString(fmt.Sprintf("take %v", potentials[2]), in)
		}

		if i&8 == 8 {
			inputString(fmt.Sprintf("take %v", potentials[3]), in)
		}

		if i&16 == 16 {
			inputString(fmt.Sprintf("take %v", potentials[4]), in)
		}

		if i&32 == 32 {
			inputString(fmt.Sprintf("take %v", potentials[5]), in)
		}

		inputString("south", in)

		if i&1 == 1 {
			inputString(fmt.Sprintf("drop %v", potentials[0]), in)
		}

		if i&2 == 2 {
			inputString(fmt.Sprintf("drop %v", potentials[1]), in)
		}

		if i&4 == 4 {
			inputString(fmt.Sprintf("drop %v", potentials[2]), in)
		}

		if i&8 == 8 {
			inputString(fmt.Sprintf("drop %v", potentials[3]), in)
		}

		if i&16 == 16 {
			inputString(fmt.Sprintf("drop %v", potentials[4]), in)
		}

		if i&32 == 32 {
			inputString(fmt.Sprintf("drop %v", potentials[5]), in)
		}

	}

	// check point
	/*
			inputString(`south
		west
		south
		south
		south`, in)
	*/
	wg.Wait()

	return output
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	return fmt.Sprintf("%v", " -- Not Yet Implemented --")
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	return fmt.Sprintf("%v", " -- Not Yet Implemented --")
}
