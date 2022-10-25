package Day201917

import (
	"fmt"
	"sync"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/pezza/advent-of-code/puzzles/2019/intcode"
)

func Test_ReadData(t *testing.T) {
	RegisterTestingT(t)

}

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)
	Expect(Entry.PartOne(Entry.PuzzleInput(), nil)).Should(Equal("2660"))
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)
	p := "A,B,A,C,A,B,C,B,C,B"
	a := "L,10,R,8,L,6,R,6"
	b := "L,8,L,8,R,8"
	c := "R,8,L,6,L,10,L,10"

	out := make(chan int64, 0)
	in := make(chan int64, 0)
	var dust []int64
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		vm := intcode.New(Entry.PuzzleInput())
		init := make(map[int64]int64, 0)
		init[0] = 2
		vm.RunProgram(init, nil, in, out, nil)
		close(out)
		wg.Done()
	}()

	go func() {
		for val := range out {
			dust = append(dust, val)
		}
		fmt.Println(dust)
		wg.Done()
	}()

	for _, char := range p {
		in <- int64(char)
	}
	in <- int64(10)

	for _, char := range a {
		in <- int64(char)
	}
	in <- int64(10)

	for _, char := range b {
		in <- int64(char)
	}
	in <- int64(10)

	for _, char := range c {
		in <- int64(char)
	}
	in <- int64(10)

	fmt.Println("opt")
	in <- int64(110)
	fmt.Println("end")

	in <- int64(10)

	wg.Wait()

	if len(dust) != 0 {
		for _, val := range dust {
			fmt.Print(string(byte(val)))
		}
	}
	fmt.Println("Not 1114999 (too high), not 190561")
}

func Benchmark_BenchPartOne(b *testing.B) {
	data := Entry.PuzzleInput()
	for n := 0; n < b.N; n++ {
		Entry.PartOne(data, nil)
	}
}

func Benchmark_BenchPartTwo(b *testing.B) {
	data := Entry.PuzzleInput()
	for n := 0; n < b.N; n++ {
		Entry.PartTwo(data, nil)
	}
}
