package Day201907

import (
	"fmt"
	"sync"

	"github.com/pezza/advent-of-code/2019/intcode"
	"github.com/pezza/advent-of-code/common"
)

func feedbackAmplifysignal(phaseSettings []int64, program string) int64 {
	aInput := make(chan int64, 0)
	bInput := make(chan int64, 0)
	cInput := make(chan int64, 0)
	dInput := make(chan int64, 0)
	eInput := make(chan int64, 0)

	ampA := intcode.New(program)
	ampB := intcode.New(program)
	ampC := intcode.New(program)
	ampD := intcode.New(program)
	ampE := intcode.New(program)

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		ampA.RunProgram(nil, nil, aInput, bInput, nil)
		wg.Done()
	}()

	go ampB.RunProgram(nil, nil, bInput, cInput, nil)
	go ampC.RunProgram(nil, nil, cInput, dInput, nil)
	go ampD.RunProgram(nil, nil, dInput, eInput, nil)
	go ampE.RunProgram(nil, nil, eInput, aInput, nil)

	aInput <- phaseSettings[0]
	bInput <- phaseSettings[1]
	cInput <- phaseSettings[2]
	dInput <- phaseSettings[3]
	eInput <- phaseSettings[4]

	aInput <- 0

	wg.Wait()
	return <-aInput
}

func straightAmplifySignal(phaseSettings []int64, program string) int64 {
	aInput := make(chan int64, 0)
	bInput := make(chan int64, 0)
	cInput := make(chan int64, 0)
	dInput := make(chan int64, 0)
	eInput := make(chan int64, 0)
	signal := make(chan int64, 0)

	ampA := intcode.New(program)
	ampB := intcode.New(program)
	ampC := intcode.New(program)
	ampD := intcode.New(program)
	ampE := intcode.New(program)

	go ampA.RunProgram(nil, nil, aInput, bInput, nil)
	go ampB.RunProgram(nil, nil, bInput, cInput, nil)
	go ampC.RunProgram(nil, nil, cInput, dInput, nil)
	go ampD.RunProgram(nil, nil, dInput, eInput, nil)
	go ampE.RunProgram(nil, nil, eInput, signal, nil)

	aInput <- phaseSettings[0]
	bInput <- phaseSettings[1]
	cInput <- phaseSettings[2]
	dInput <- phaseSettings[3]
	eInput <- phaseSettings[4]

	aInput <- 0

	return <-signal
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	maxSignal := int64(0)

	for _, perm := range common.GetPermuations([]int64{0, 1, 2, 3, 4}) {
		result := straightAmplifySignal(perm, inputData)

		if result > maxSignal {
			maxSignal = result
		}
	}

	return fmt.Sprintf("%v", maxSignal)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	maxSignal := int64(0)

	for _, perm := range common.GetPermuations([]int64{5, 6, 7, 8, 9}) {
		result := feedbackAmplifysignal(perm, inputData)

		if result > maxSignal {
			maxSignal = result
		}
	}
	return fmt.Sprintf("%v", maxSignal)
}
