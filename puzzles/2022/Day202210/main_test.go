package Day202210

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

func Test_ReadData(t *testing.T) {
	RegisterTestingT(t)

	data := getData(Entry.PuzzleInput())

	Expect(data[0]).Should(Equal(instruction{"addx", 1}))
	Expect(data[len(data)-1]).Should(Equal(instruction{"noop", 0}))
}

const sampleProg1 string = `noop
addx 3
addx -5`

const sampleProg2 string = `addx 15
addx -11
addx 6
addx -3
addx 5
addx -1
addx -8
addx 13
addx 4
noop
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx -35
addx 1
addx 24
addx -19
addx 1
addx 16
addx -11
noop
noop
addx 21
addx -15
noop
noop
addx -3
addx 9
addx 1
addx -3
addx 8
addx 1
addx 5
noop
noop
noop
noop
noop
addx -36
noop
addx 1
addx 7
noop
noop
noop
addx 2
addx 6
noop
noop
noop
noop
noop
addx 1
noop
noop
addx 7
addx 1
noop
addx -13
addx 13
addx 7
noop
addx 1
addx -33
noop
noop
noop
addx 2
noop
noop
noop
addx 8
noop
addx -1
addx 2
addx 1
noop
addx 17
addx -9
addx 1
addx 1
addx -3
addx 11
noop
noop
addx 1
noop
addx 1
noop
noop
addx -13
addx -19
addx 1
addx 3
addx 26
addx -30
addx 12
addx -1
addx 3
addx 1
noop
noop
noop
addx -9
addx 18
addx 1
addx 2
noop
noop
addx 9
noop
noop
noop
addx -1
addx 2
addx -37
addx 1
addx 3
noop
addx 15
addx -21
addx 22
addx -6
addx 1
noop
addx 2
addx 1
noop
addx -10
noop
noop
addx 20
addx 1
addx 2
addx 2
addx -6
addx -11
noop
noop
noop`

func Test_RunProg(t *testing.T) {
	RegisterTestingT(t)

	data := getData(sampleProg1)

	history := runProg(data)

	for i := 1; i <= len(history); i++ {
		fmt.Println(i, history[i])
	}
}

func Test_RunProg2(t *testing.T) {
	RegisterTestingT(t)

	data := getData(sampleProg2)

	history := runProg(data)

	//for i := 1; i <= len(history); i++ {
	//	fmt.Println(i, history[i])
	//}
	Expect(history[20-1]).Should(Equal(21))
	Expect(history[60-1]).Should(Equal(19))
	Expect(history[100-1]).Should(Equal(18))
	Expect(history[140-1]).Should(Equal(21))
	Expect(history[180-1]).Should(Equal(16))
	Expect(history[220-1]).Should(Equal(18))
	ss := (history[20-1] * 20) + (history[60-1] * 60) + (history[100-1] * 100) + (history[140-1] * 140) + (history[180-1] * 180) + (history[220-1] * 220)

	Expect(ss).Should(Equal(13140))
}

func Test_GetScan(t *testing.T) {
	RegisterTestingT(t)

	data := getData(Entry.PuzzleInput())
	history := runProg(data)

	output := runScan(history)
	fmt.Println(len(history))
	fmt.Println(output[0:39])
	fmt.Println(output[40:79])
	fmt.Println(output[80:119])
	fmt.Println(output[120:159])
	fmt.Println(output[160:199])
	fmt.Println(output[200:239])

	fmt.Println(output)
}

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	Expect(Entry.PartOne(Entry.PuzzleInput(), nil)).Should(Equal("14780"))
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

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
