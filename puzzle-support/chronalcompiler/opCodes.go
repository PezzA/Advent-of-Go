package chronalcompiler

type Processor interface {
	Process(input RegisterSet, a int, b int) int
}

type OpCodes map[string]Processor

// create the 16 types
type addr bool
type addi bool
type mulr bool
type muli bool
type setr bool
type seti bool
type banr bool
type bani bool
type borr bool
type bori bool
type gtir bool
type gtri bool
type gtrr bool
type eqir bool
type eqri bool
type eqrr bool

func (op addr) Process(input RegisterSet, a int, b int) int {
	return input[a] + input[b]
}

func (op addi) Process(input RegisterSet, a int, b int) int {
	return input[a] + b
}

func (op mulr) Process(input RegisterSet, a int, b int) int {
	return input[a] * input[b]
}

func (op muli) Process(input RegisterSet, a int, b int) int {
	return input[a] * b
}

func (op setr) Process(input RegisterSet, a int, b int) int {
	return input[a]
}

func (op seti) Process(input RegisterSet, a int, b int) int {
	return a
}

func (op banr) Process(input RegisterSet, a int, b int) int {
	return input[a] & input[b]
}

func (op bani) Process(input RegisterSet, a int, b int) int {
	return input[a] & b
}

func (op borr) Process(input RegisterSet, a int, b int) int {
	return input[a] | input[b]
}

func (op bori) Process(input RegisterSet, a int, b int) int {
	return input[a] | b
}

func (op gtir) Process(input RegisterSet, a int, b int) int {
	if a > input[b] {
		return 1
	} else {
		return 0
	}

}

func (op gtri) Process(input RegisterSet, a int, b int) int {
	if input[a] > b {
		return 1
	} else {
		return 0
	}
}

func (op gtrr) Process(input RegisterSet, a int, b int) int {
	if input[a] > input[b] {
		return 1
	} else {
		return 0
	}
}

func (op eqir) Process(input RegisterSet, a int, b int) int {
	if a == input[b] {
		return 1
	} else {
		return 0
	}
}

func (op eqri) Process(input RegisterSet, a int, b int) int {
	if input[a] == b {
		return 1
	} else {
		return 0
	}
}

func (op eqrr) Process(input RegisterSet, a int, b int) int {
	if input[a] == input[b] {
		return 1
	} else {
		return 0
	}
}

func GetOpCodes() OpCodes {
	ops := make(OpCodes)
	ops["addr"] = new(addr)
	ops["addi"] = new(addi)
	ops["mulr"] = new(mulr)
	ops["muli"] = new(muli)
	ops["setr"] = new(setr)
	ops["seti"] = new(seti)
	ops["banr"] = new(banr)
	ops["bani"] = new(bani)
	ops["borr"] = new(borr)
	ops["bori"] = new(bori)
	ops["gtir"] = new(gtir)
	ops["gtri"] = new(gtri)
	ops["gtrr"] = new(gtrr)
	ops["eqir"] = new(eqir)
	ops["eqri"] = new(eqri)
	ops["eqrr"] = new(eqrr)

	return ops
}

var OpCodeSortList = []string{
	"addr",
	"addi",
	"mulr",
	"muli",
	"setr",
	"seti",
	"banr",
	"bani",
	"borr",
	"bori",
	"gtir",
	"gtri",
	"gtrr",
	"eqir",
	"eqri",
	"eqrr",
}
