package Day201718

import (
	"bufio"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type dayEntry bool

var Entry dayEntry

func (td dayEntry) Describe() (int, int, string, int) {
	return 2017, 18, "Duet", 0
}

type instruction struct {
	command  string
	operand1 string
	operand2 string
}

type registers map[string]int

type chip struct {
	registers        registers
	lastSound        int
	recoveredSound   int
	index            int
	sendChannel      chan int
	receiveChannel   chan int
	messagesSent     int
	messagesReceived int
	deadlocked       bool
	instructions     int
}

var instructionRegex = regexp.MustCompile(`(set|mul|jgz|add|mod|snd|rcv) (\S+)( )*(\S*)`)

func loadInstructions(input string) []instruction {

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

func (c chip) runInstruction(i instruction) chip {

	switch i.command {
	case "set":
		targetVal, err := strconv.Atoi(i.operand2)

		if err != nil {
			c.registers[i.operand1] = c.getRegister(i.operand2)
		} else {
			c.registers[i.operand1] = targetVal
		}
		c.index++
	case "snd":
		if c.sendChannel != nil {
			targetVal, err := strconv.Atoi(i.operand1)

			if err != nil {
				targetVal = c.getRegister(i.operand1)
			}

			c.sendChannel <- targetVal
			c.messagesSent++
			c.index++
			break
		}

		c.lastSound = c.getRegister(i.operand1)
		c.index++
	case "add":
		targetVal, err := strconv.Atoi(i.operand2)

		if err != nil {
			c.registers[i.operand1] += c.getRegister(i.operand2)
		} else {
			c.registers[i.operand1] += targetVal
		}
		c.index++
	case "mul":
		targetVal, err := strconv.Atoi(i.operand2)

		if err != nil {
			c.registers[i.operand1] *= c.getRegister(i.operand2)
		} else {
			c.registers[i.operand1] *= targetVal
		}
		c.index++
	case "mod":
		targetVal, err := strconv.Atoi(i.operand2)

		if err != nil {
			c.registers[i.operand1] = c.registers[i.operand1] % c.getRegister(i.operand2)
		} else {
			c.registers[i.operand1] = c.registers[i.operand1] % targetVal
		}
		c.index++
	case "rcv":
		if c.receiveChannel != nil {

			select {
			case intVal := <-c.receiveChannel:
				c.registers[i.operand1] = intVal
				c.index++
				c.messagesReceived++
			case <-time.After(time.Millisecond * 10):
				c.deadlocked = true
			}

			break
		}

		if c.getRegister(i.operand1) == 0 {
			c.index++
			break
		}

		c.recoveredSound = c.lastSound
		c.index++
	case "jgz":
		val, err := strconv.Atoi(i.operand1)

		if err != nil {
			val = c.getRegister(i.operand1)
		}

		if val > 0 {
			shift, err := strconv.Atoi(i.operand2)

			if err != nil {
				c.index += c.getRegister(i.operand2)
			}

			c.index += shift
			break
		}
		c.index++
	}

	c.instructions++
	return c
}

func runProgram(ins []instruction, c chip, progressChannel chan int, updateChan chan []string) int {
	for {
		c = c.runInstruction(ins[c.index])

		if c.recoveredSound > 0 || c.index < 0 || c.index > len(ins) || c.deadlocked {
			break
		}

	}

	if progressChannel != nil {
		progressChannel <- c.messagesSent
	}
	return c.recoveredSound
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	c := chip{
		registers:      make(registers, 0),
		lastSound:      0,
		recoveredSound: 0,
		index:          0,
	}

	val := runProgram(loadInstructions(inputData), c, nil, nil)

	return strconv.Itoa(val)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {

	atob := make(chan int, 10000)
	btoa := make(chan int, 10000)

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

	go runProgram(loadInstructions(inputData), c, progress1, nil)
	go runProgram(loadInstructions(inputData), c2, progress2, updateChan)

	<-progress1
	prog1messages := <-progress2

	return strconv.Itoa(prog1messages)
}
