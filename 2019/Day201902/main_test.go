package Day201902

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	Expect(runProgram(getListIntData("1,9,10,3,2,3,11,0,99,30,40,50"))[0]).Should(Equal(3500))
	Expect(runProgram(getListIntData("1,0,0,0,99"))[0]).Should(Equal(2))
	Expect(runProgram(getListIntData("2,4,4,5,99,0"))[5]).Should(Equal(9801))
	Expect(runProgram(getListIntData("1,1,1,4,99,5,6,0,99"))[0]).Should(Equal(30))

	codes := getListIntData(Entry.PuzzleInput())

	codes[1] = 12
	codes[2] = 2

	//fmt.Println("not 1028272, 1028270")
	Expect(runProgram(codes)[0]).Should(Equal(3562672))
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

	done := false
	actNoun, actVerb := 0, 0
	for noun := 0; noun <= 100; noun++ {
		for verb := 0; verb <= 100; verb++ {
			codes := getListIntData(Entry.PuzzleInput())
			codes[1] = noun
			codes[2] = verb

			if runProgram(codes)[0] == 19690720 {
				actNoun = noun
				actVerb = verb
				done = true
			}

			if done {
				break
			}
		}
		if done {
			break
		}
	}

	Expect(100*actNoun + actVerb).Should(Equal(8250))
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
