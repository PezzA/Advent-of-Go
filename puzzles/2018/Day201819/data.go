package Day201819

func (td dayEntry) PuzzleInput() string {
	return `#ip 5
addi 5 16 5
seti 1 0 4
seti 1 8 1
mulr 4 1 3
eqrr 3 2 3
addr 3 5 5
addi 5 1 5
addr 4 0 0
addi 1 1 1
gtrr 1 2 3
addr 5 3 5
seti 2 4 5
addi 4 1 4
gtrr 4 2 3
addr 3 5 5
seti 1 7 5
mulr 5 5 5
addi 2 2 2
mulr 2 2 2
mulr 5 2 2
muli 2 11 2
addi 3 6 3
mulr 3 5 3
addi 3 9 3
addr 2 3 2
addr 5 0 5
seti 0 5 5
setr 5 9 3
mulr 3 5 3
addr 5 3 3
mulr 5 3 3
muli 3 14 3
mulr 3 5 3
addr 2 3 2
seti 0 1 0
seti 0 0 5`
}

var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string, int) {
	return 2018, 19, "Go With The Flow", 2
}
