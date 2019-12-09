package Day201909

import (
	"fmt"
	"testing"

	"github.com/pezza/advent-of-code/2019/intcode"

	. "github.com/onsi/gomega"
)

func Test_ReadData(t *testing.T) {
	RegisterTestingT(t)

}

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	codes := getListIntData(`109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99`)
	Expect(intcode.RunProgram(codes, []int64{}, false, nil, nil, "", nil)).Should(Equal([]int64{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}))

	codes = getListIntData(`104,1125899906842624,99`)
	Expect(intcode.RunProgram(codes, []int64{}, false, nil, nil, "", nil)).Should(Equal([]int64{1125899906842624}))

	codes = getListIntData(`1102,34915192,34915192,7,4,7,99,0`)
	Expect(intcode.RunProgram(codes, []int64{}, false, nil, nil, "", nil)).Should(Equal([]int64{1219070632396864}))

	codes = getListIntData(Entry.PuzzleInput())
	fmt.Println(intcode.RunProgram(codes, []int64{1}, false, nil, nil, "", nil))
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

	codes := getListIntData(Entry.PuzzleInput())
	fmt.Println(intcode.RunProgram(codes, []int64{2}, false, nil, nil, "", nil))

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
