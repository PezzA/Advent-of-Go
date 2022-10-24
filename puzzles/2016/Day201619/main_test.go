package Day201619

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

	elfCircle := makeElfCircle(3005290)
	elf := elfCircle.Front()

	fmt.Println("Starting with Elf", elf.Value)
	for {

		// move forward by half the number of current elves (rounding down)
		increment := elfCircle.Len() / 2
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

		elf = startElf

		elf = elf.Next()

		if elf == nil {
			elf = elfCircle.Front()
		}

		if elfCircle.Len()%100 == 0 {
			fmt.Println((elfCircle.Len()))
		}
		if elfCircle.Len() == 1 {
			break
		}
	}

	fmt.Print(elf.Value)
}

func Benchmark_PartOne(b *testing.B) {
	data := Entry.PuzzleInput()
	for n := 0; n < b.N; n++ {
		Entry.PartOne(data, nil)
	}
}

func Benchmark_PartTwo(b *testing.B) {
	data := Entry.PuzzleInput()
	for n := 0; n < b.N; n++ {
		Entry.PartTwo(data, nil)
	}
}
