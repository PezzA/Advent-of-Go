package Day201619

import (
	"container/list"
	"strconv"
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

func (td dayEntry) PuzzleInput() string {
	return "3005290"
}
