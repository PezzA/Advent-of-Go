package Day201905

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/pezza/advent-of-code/2019/intcode"
)

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	codes := getListIntData(Entry.PuzzleInput())
	//codes = correctTests(codes)
	fmt.Println(intcode.RunProgram(codes, 1, true))
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

	t1 := getListIntData(`3,3,1108,-1,8,3,4,3,99`)
	Expect(intcode.RunProgram(t1, 8, true)[0]).Should(Equal(1))

	t1 = getListIntData(`3,9,8,9,10,9,4,9,99,-1,8`)
	Expect(intcode.RunProgram(t1, 8, true)[0]).Should(Equal(1))

	t1 = getListIntData(`3,9,8,9,10,9,4,9,99,-1,8`)
	Expect(intcode.RunProgram(t1, 7, true)[0]).Should(Equal(0))

	t1 = getListIntData(`3,9,7,9,10,9,4,9,99,-1,8`)
	Expect(intcode.RunProgram(t1, 7, true)[0]).Should(Equal(1))

	t1 = getListIntData(`3,9,7,9,10,9,4,9,99,-1,8`)
	Expect(intcode.RunProgram(t1, 8, true)[0]).Should(Equal(0))

	t1 = getListIntData(`3,3,1108,-1,8,3,4,3,99`)
	Expect(intcode.RunProgram(t1, 7, true)[0]).Should(Equal(0))

	//t1 = getListIntData(`3,3,1107,-1,8,3,4,3,99`)
	//Expect(runProgram(t1, 8, true)[0]).Should(Equal(1))
	//t1 = getListIntData(`3,3,1107,-1,8,3,4,3,99`)
	//Expect(runProgram(t1, 7, true)[0]).Should(Equal(0))

	codes := getListIntData(`3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99`)
	Expect(intcode.RunProgram(codes, 7, true)[0]).Should(Equal(999))

	codes = getListIntData(`3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99`)
	Expect(intcode.RunProgram(codes, 8, true)[0]).Should(Equal(1000))

	codes = getListIntData(`3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99`)
	Expect(intcode.RunProgram(codes, 9, true)[0]).Should(Equal(1001))

	codes = getListIntData(Entry.PuzzleInput())
	fmt.Println(intcode.RunProgram(codes, 5, false))

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
