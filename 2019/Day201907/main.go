package Day201907

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/pezza/advent-of-code/2019/intcode"
)

func feedbackAmplifysignal(phaseSettings []int, program string) int {

	fmt.Println("making channels")
	aInput := make(chan int, 0)
	bInput := make(chan int, 0)
	cInput := make(chan int, 0)
	dInput := make(chan int, 0)
	eInput := make(chan int, 0)

	var wg sync.WaitGroup

	wg.Add(4)

	go intcode.RunProgram(getListIntData(program), nil, false, aInput, bInput, "A", &wg)
	go intcode.RunProgram(getListIntData(program), nil, false, bInput, cInput, "B", &wg)
	go intcode.RunProgram(getListIntData(program), nil, false, cInput, dInput, "C", &wg)
	go intcode.RunProgram(getListIntData(program), nil, false, dInput, eInput, "D", &wg)
	go intcode.RunProgram(getListIntData(program), nil, false, eInput, aInput, "E", nil)

	aInput <- phaseSettings[0]
	bInput <- phaseSettings[1]
	cInput <- phaseSettings[2]
	dInput <- phaseSettings[3]
	eInput <- phaseSettings[4]

	aInput <- 0

	wg.Wait()

	fmt.Println("we done here")
	return <-aInput
}

func straightAmplifySignal(phaseSettings []int, program string) int {
	aInput := make(chan int, 0)
	bInput := make(chan int, 0)
	cInput := make(chan int, 0)
	dInput := make(chan int, 0)
	eInput := make(chan int, 0)
	signal := make(chan int, 0)

	go intcode.RunProgram(getListIntData(program), nil, false, aInput, bInput, "A", nil)
	go intcode.RunProgram(getListIntData(program), nil, false, bInput, cInput, "B", nil)
	go intcode.RunProgram(getListIntData(program), nil, false, cInput, dInput, "C", nil)
	go intcode.RunProgram(getListIntData(program), nil, false, dInput, eInput, "D", nil)
	go intcode.RunProgram(getListIntData(program), nil, false, eInput, signal, "E", nil)

	aInput <- phaseSettings[0]
	bInput <- phaseSettings[1]
	cInput <- phaseSettings[2]
	dInput <- phaseSettings[3]
	eInput <- phaseSettings[4]

	aInput <- 0

	return <-signal
}

func getListIntData(input string) []int {
	retval := []int{}

	for _, i := range strings.Split(input, ",") {
		newVal, _ := strconv.Atoi(i)

		retval = append(retval, newVal)
	}

	return retval
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	return fmt.Sprintf("%v", " -- Not Yet Implemented --")
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	return fmt.Sprintf("%v", " -- Not Yet Implemented --")
}
