package Day202211

import (
	"fmt"
	"math/big"
	"sort"
	"strconv"
	"strings"
)

type operation struct {
	operator string
	amount   *big.Int
}

type monkey struct {
	items       []*big.Int
	op          operation
	test        *big.Int
	trueTarget  int
	falseTarget int
	handleCount int
}

func getData(input string) []monkey {
	lines := strings.Split(input, "\n")

	monkeys := make([]monkey, 0)

	for i := 0; i < len(lines); {

		//monkey line
		i++
		//items
		itemLine := strings.Replace(lines[i], " ", "", -1)
		itemBits := strings.Split(itemLine, ":")
		rawItems := strings.Split(itemBits[1], ",")
		items := make([]*big.Int, 0)
		for _, rawItem := range rawItems {
			val, _ := strconv.Atoi(rawItem)
			items = append(items, big.NewInt(int64(val)))
		}
		i++
		//operation
		operator, amount := "", ""
		fmt.Sscanf(lines[i], "  Operation: new = old %s %s", &operator, &amount)
		if amount == "old" {
			amount = "2"
			operator = "^"
		}

		operationVal, _ := strconv.Atoi(amount)
		i++
		//test
		test := 0
		fmt.Sscanf(lines[i], "  Test: divisible by %d", &test)
		i++
		//trueTarget
		trueTarget := 0
		fmt.Sscanf(lines[i], "    If true: throw to monkey %d", &trueTarget)
		i++
		//falseTarget
		falseTarget := 0
		fmt.Sscanf(lines[i], "    If false: throw to monkey %d", &falseTarget)
		i++
		i++

		monkey := monkey{
			items,
			operation{operator, big.NewInt(int64(operationVal))},
			big.NewInt(int64(test)),
			trueTarget,
			falseTarget,
			0,
		}

		monkeys = append(monkeys, monkey)

	}

	return monkeys
}

func runRound(monkeys []monkey, dodiv bool) []monkey {
	zero := big.NewInt(0)

	for i := 0; i < len(monkeys); i++ {
		if len(monkeys[i].items) == 0 {
			continue
		}

		m := monkeys[i]

		for itemI := 0; itemI < len(monkeys[i].items); itemI++ {
			item := m.items[itemI]

			worry := calcWorry(item, m.op)

			if dodiv {
				worry = worry.Div(worry, big.NewInt(int64(3)))
			}

			checked := big.NewInt(0)
			checked = checked.Mod(worry, m.test)

			if checked.Cmp(zero) == 0 {
				monkeys[m.trueTarget].items = append(monkeys[m.trueTarget].items, worry)
			} else {
				monkeys[m.falseTarget].items = append(monkeys[m.falseTarget].items, worry)
			}

			monkeys[i].handleCount++
		}

		monkeys[i].items = []*big.Int{}
	}
	return monkeys
}

func calcWorry(start *big.Int, op operation) *big.Int {
	switch op.operator {
	case "+":
		return start.Add(start, op.amount)
	case "*":
		return start.Mul(start, op.amount)
	case "^":
		return start.Mul(start, start)
	}
	return big.NewInt(int64(-1))
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	monkeys := getData(inputData)

	for i := 0; i < 20; i++ {
		fmt.Println(i)
		monkeys = runRound(monkeys, true)
	}

	handleLevels := make([]int, 0)

	for _, m := range monkeys {
		handleLevels = append(handleLevels, m.handleCount)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(handleLevels)))

	return fmt.Sprint(handleLevels[0] * handleLevels[1])
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	monkeys := getData(inputData)

	for i := 0; i < 10000; i++ {
		fmt.Println(i)
		fmt.Println(i)
		fmt.Println(i)
		fmt.Println(i)
		monkeys = runRound(monkeys, false)
	}

	handleLevels := make([]int, 0)

	for _, m := range monkeys {
		handleLevels = append(handleLevels, m.handleCount)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(handleLevels)))

	return fmt.Sprint(handleLevels[0] * handleLevels[1])

}
