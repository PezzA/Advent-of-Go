package assembunny

type RegisterSet = map[string]int

type Program = []Instruction

type Instruction struct {
	instruction string
	arg1        Argument
	arg2        Argument
}

type Argument struct {
	isRegister bool
	register   string
	value      int
}
