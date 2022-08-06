package Day201723

import (
	"bufio"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type dayEntry bool

var Entry dayEntry

type instruction struct {
	command  string
	operand1 string
	operand2 string
}

type registers map[string]int

type chip struct {
	registers      registers
	mulInvocations int
	index          int
	instructions   int
}

var instructionRegex = regexp.MustCompile(`(set|mul|jnz|sub|mod|snd|rcv) (\S+)( )*(\S*)`)

func getInstructions(input string) []instruction {

	ins := make([]instruction, 0)
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()

		matches := instructionRegex.FindStringSubmatch(line)

		ins = append(ins, instruction{matches[1], matches[2], matches[4]})
	}

	return ins
}

func (c chip) getRegister(reg string) int {
	if val, ok := c.registers[reg]; ok {
		return val
	}

	c.registers[reg] = 0
	return 0
}

func (c chip) getActualVal(reg string) int {
	targetVal, err := strconv.Atoi(reg)

	if err != nil {
		return c.getRegister(reg)
	}

	return targetVal
}

func (c chip) runInstruction(i instruction) chip {

	switch i.command {
	case "set":
		c.registers[i.operand1] = c.getActualVal(i.operand2)
		c.index++
	case "sub":
		c.registers[i.operand1] -= c.getActualVal(i.operand2)
		c.index++
	case "mul":
		c.registers[i.operand1] *= c.getActualVal(i.operand2)
		c.index++
		c.mulInvocations++
	case "jnz":

		//if i.operand1 == "1" && i.operand2 == "10" {
		//	fmt.Println(i, "\t a:", c.getRegister("a"), "\t b:", c.getRegister("b"), "\t c:", c.getRegister("c"), "\t d:", c.getRegister("d"), "\t e:", c.getRegister("e"), "\t f:", c.getRegister("f"), "\t g:", c.getRegister("g"), "\t h:", c.getRegister("h"))
		//	fmt.Println(c.getRegister("b") / c.getRegister("e"))
		//}

		test := c.getActualVal(i.operand1)
		if test != 0 {
			shift := c.getActualVal(i.operand2)
			c.index += shift
		} else {
			c.index++
		}
	}

	c.instructions++

	return c
}

func runProgram(ins []instruction, c chip, updateChan chan []string) int {
	for {
		if c.index < 0 || c.index >= len(ins) {
			break
		}
		c = c.runInstruction(ins[c.index])
	}

	return c.mulInvocations
}

func (td dayEntry) Describe() (int, int, string, int) {
	return 2017, 23, "Getting the boilerplate in place"
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	ins := getInstructions(Entry.PuzzleInput())

	c := chip{
		registers:      make(registers, 0),
		index:          0,
		mulInvocations: 0,
	}

	return fmt.Sprintf("%v", runProgram(ins, c, updateChan))
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	/*
			set b 93
		set c b				// final value of c
		jnz a 2   // IF DEBUG
		jnz 1 5	  // QUICK
		mul b 100			//DEBUG
		sub b -100000		//DEBUG
		set c b				//DEBUG
		sub c -17000		//DEBUG + final value of c

		set f 1				// init f
		set d 2				// init d
			set e 2			// init e
				set g d
				mul g e
				sub g b
				jnz g 2		// if g is zero
				set f 0		// 		f is set to found
				sub e -1	// else e++
				set g e
				sub g b
				jnz g -8
			sub d -1		// d++
			set g d
			sub g b
			jnz g -13
		jnz f 2				// if f == 0 then h++, but what is f doing?
		sub h -1
		set g b
		sub g c
		jnz g 2		// if loop == control
		jnz 1 3			EXIT
		sub b -17  // loop counter UPDATE, should loop 1000 times
		jnz 1 -23
	*/

	nonprimes := 0
	for i := 109300; i <= 126300; i += 17 {
		if !isPrime(i) {
			nonprimes++
		}
	}

	return fmt.Sprintf("%v", nonprimes)
}

func isPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}
