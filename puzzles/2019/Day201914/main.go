package Day201914

import (
	"fmt"
	"math"
	"strings"
)

type ingredient struct {
	amount int
	class  string
}

type assembly struct {
	inputs    []ingredient
	output    ingredient
	stockHeld int
}

var totalOre = 0
var reserves = 1000000000000

type assemblyLine []*assembly

func (a *assembly) request(amount int, al assemblyLine) bool {
	if a.stockHeld >= amount {
		a.stockHeld -= amount
		return true
	}

	amount -= a.stockHeld
	a.stockHeld = 0

	crafts := int(math.Ceil(float64(amount) / float64(a.output.amount)))
	a.stockHeld = crafts*a.output.amount - amount

	for c := 0; c < crafts; c++ {
		for _, input := range a.inputs {
			if input.class == "ORE" {
				if totalOre+input.amount > reserves {
					return false
				}
				totalOre += input.amount
				continue
			}
			inga, _ := al.getByOutput(input.class)
			if !inga.request(input.amount, al) {
				return false
			}
		}
	}

	return true
}

func parseIngredient(input string) ingredient {
	input = strings.TrimSpace(input)
	amount, name := 0, ""
	fmt.Sscanf(input, "%d %v", &amount, &name)

	return ingredient{
		amount,
		name,
	}
}

func getData(input string) assemblyLine {
	lines := strings.Split(input, "\n")
	al := assemblyLine{}

	for _, line := range lines {
		bits := strings.Split(line, "=>")

		var inputs []ingredient
		inputData := strings.Split(bits[0], ",")

		for _, data := range inputData {
			inputs = append(inputs, parseIngredient(data))
		}

		output := parseIngredient(bits[1])
		al = append(al, &assembly{inputs, output, 0})
	}
	return al
}

func (al assemblyLine) getByOutput(class string) (*assembly, error) {
	for _, assem := range al {
		if assem.output.class == class {
			return assem, nil
		}
	}
	return &assembly{}, fmt.Errorf(fmt.Sprintf("Could not find recipe with output %v", class))
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	return fmt.Sprintf("%v", " -- Not Yet Implemented --")
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	return fmt.Sprintf("%v", " -- Not Yet Implemented --")
}
