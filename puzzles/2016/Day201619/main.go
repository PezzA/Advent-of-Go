package Day201619

import (
	"container/list"
	"fmt"
	"strconv"
	"time"
)

// Entry holds wraps the data and runner interfaces for this puzzle
var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string, int) {
	return 2016, 19, "An Elephant Named Joseph", 1
}

func getAdjacentElfHops(circleSize int, position int) int {
	offset := circleSize / 2

	newPos := position + offset

	if newPos <= circleSize {
		return newPos
	}
	return newPos - circleSize
}

func makeElfCircle(circleSize int) *list.List {
	elves := list.New()
	for i := 0; i < circleSize; i++ {
		elves.PushBack(i + 1)
	}
	return elves
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	circleSize, _ := strconv.Atoi(inputData)

	elfCircle := makeElfCircle(circleSize)
	elf := elfCircle.Front()

	for {

		if next := elf.Next(); next != nil {
			elfCircle.Remove(next)
		} else {
			elfCircle.Remove(elfCircle.Front())
		}

		elf = elf.Next()

		if elf == nil {
			elf = elfCircle.Front()
		}

		if elfCircle.Len() == 1 {
			break
		}
	}

	return strconv.Itoa(elfCircle.Front().Value.(int))
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	circleSize, _ := strconv.Atoi(inputData)

	elfCircle := makeElfCircle(circleSize)
	elf := elfCircle.Front()

	fmt.Println("Starting with Elf", elf.Value)

	start := time.Now()
	for {

		// move forward by half the number of current elves (rounding down)
		increment := circleSize / 2
		//	fmt.Println("Moving forward by ", increment, elfCircle.Len())
		startElf := elf

		for i := 0; i < increment; i++ {
			elf = elf.Next()

			if elf == nil {
				elf = elfCircle.Front()
			}
		}

		//fmt.Println("removing elf ", elf.Value)

		elfCircle.Remove(elf)
		circleSize--
		elf = startElf

		elf = elf.Next()

		if elf == nil {
			elf = elfCircle.Front()
		}

		if circleSize%100 == 0 {
			updateChan <- []string{fmt.Sprint(circleSize, "-", time.Since(start), "-")}
			start = time.Now()
		}
		if circleSize == 1 {
			break
		}
	}

	return strconv.Itoa(elfCircle.Front().Value.(int))
}

func (td dayEntry) PuzzleInput() string {
	return "3005290"
}
