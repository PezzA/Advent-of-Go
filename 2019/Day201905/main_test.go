package Day201905

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

func Test_ReadData(t *testing.T) {
	RegisterTestingT(t)

}

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	Expect(parseOpCode(1002)).Should(Equal(opCode{2, 0, 1, 0}))
	Expect(parseOpCode(10001)).Should(Equal(opCode{1, 0, 0, 1}))
	Expect(parseOpCode(1)).Should(Equal(opCode{1, 0, 0, 0}))
	Expect(parseOpCode(99)).Should(Equal(opCode{99, 0, 0, 0}))

	codes := getListIntData(Entry.PuzzleInput())

	fmt.Println("not 14386578")
	fmt.Println(runProgram(codes, 1, false))
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)
	codes := getListIntData(`3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,
1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,
999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99`)
	fmt.Println(runProgram(codes, 7, true))

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
