package Day201718

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	ins := loadInstructions(`set a 1
add a 2
mul a a
mod a 5
snd a
set a 0
rcv a
jgz a -1
set a 1
jgz a -2`)

	c := chip{
		registers:      make(registers, 0),
		lastSound:      0,
		recoveredSound: 0,
		index:          0,
	}

	val := runProgram(ins, c, nil, nil)

	Expect(val).Should(Equal(4))
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

	ins := loadInstructions(`snd 1
		snd 2
		snd p
		rcv a
		rcv b
		rcv c
		rcv d`)

	atob := make(chan int, 10)
	btoa := make(chan int, 10)

	c := chip{
		registers:      make(registers, 0),
		lastSound:      0,
		recoveredSound: 0,
		index:          0,
		sendChannel:    atob,
		receiveChannel: btoa,
	}

	c.registers["p"] = 0
	c2 := chip{
		registers:      make(registers, 0),
		lastSound:      0,
		recoveredSound: 0,
		index:          0,
		sendChannel:    btoa,
		receiveChannel: atob,
	}
	c2.registers["p"] = 1

	progress1 := make(chan int, 0)
	progress2 := make(chan int, 0)

	go runProgram(ins, c, progress1, nil)
	go runProgram(ins, c2, progress2, nil)

	<-progress1
	prog1messages := <-progress2

	Expect(prog1messages).Should(Equal(3))
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
