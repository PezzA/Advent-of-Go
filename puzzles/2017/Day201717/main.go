package Day201717

import (
	"container/list"
	"fmt"
)

type dayEntry bool

var Entry dayEntry

func (td dayEntry) Describe() (int, int, string, int) {
	return 2017, 17, "Spinlock", 0
}

func addValue(val int, steps int, pos *list.Element, sl *list.List) *list.Element {
	for i := 0; i < steps; i++ {
		pos = pos.Next()

		if pos == nil {
			pos = sl.Front()
		}
	}
	return sl.InsertAfter(val, pos)
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	newsl, steps := list.New(), 355
	newpos := newsl.PushFront(0)

	for i := 1; i < 2018; i++ {
		newpos = addValue(i, steps, newpos, newsl)
	}

	return fmt.Sprint(newpos.Next().Value)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	newsl, steps := list.New(), 355
	newpos := newsl.PushFront(0)

	for i := 1; i < 50000000; i++ {
		newpos = addValue(i, steps, newpos, newsl)

		if i%10000 == 0 {
			updateChan <- []string{fmt.Sprintf("Running %v of %v", i, 50000000)}
		}
	}

	return fmt.Sprint(newsl.Front().Next().Value)
}

func (td dayEntry) PuzzleInput() string {
	return "355"
}
