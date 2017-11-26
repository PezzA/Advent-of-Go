package Day201619

import (
	"strconv"
	"testing"
)

func Test_EndToEnd(t *testing.T) {
	day1data, _ := Entry.GetTestData()

	circleSize, _ := strconv.Atoi(day1data[0])

	elfCircle := makeElfCircle(circleSize)
	finished := false

	elf := elfCircle.Front()

	for !finished {

		if next := elf.Next(); next != nil {
			elfCircle.Remove(next)
		} else {
			elfCircle.Remove(elfCircle.Front())
		}

		elf = elf.Next()

		if elf == nil {
			elf = elfCircle.Front()
		}

		finished = elfCircle.Len() == 1
	}
}
