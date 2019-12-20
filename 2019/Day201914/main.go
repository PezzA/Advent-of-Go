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

type recipe struct {
	inputs []ingredient
	output ingredient
}

type recipeBook []recipe

func parseIngredient(input string) ingredient {
	input = strings.TrimSpace(input)
	amount, name := 0, ""
	fmt.Sscanf(input, "%d %v", &amount, &name)

	return ingredient{
		amount,
		name,
	}
}

func getData(input string) recipeBook {
	lines := strings.Split(input, "\n")
	recipes := recipeBook{}

	for _, line := range lines {
		bits := strings.Split(line, "=>")

		var inputs []ingredient
		inputData := strings.Split(bits[0], ",")

		for _, data := range inputData {
			inputs = append(inputs, parseIngredient(data))
		}

		output := parseIngredient(bits[1])
		recipes = append(recipes, recipe{inputs, output})
	}
	return recipes
}

func (rb recipeBook) getByOutput(class string) (recipe, error) {
	for _, recipe := range rb {
		if recipe.output.class == class {
			return recipe, nil
		}
	}
	return recipe{}, fmt.Errorf(fmt.Sprintf("Could not find recipe with output %v", class))
}

type oreRequirement struct {
	amount int
	class  string
}

func (rb recipeBook) startCrawl(amountRequired int, name string, depth int) []oreRequirement {

	fmt.Printf("%v%-10v%-10v\n",
		strings.Repeat(" ", depth*4),
		name,
		amountRequired)

	recipe, _ := rb.getByOutput(name)
	amountToCraft := int(math.Ceil(float64(amountRequired) / float64(recipe.output.amount)))

	var amounts []oreRequirement

	for _, rec := range recipe.inputs {
		if rec.class == "ORE" {
			amounts = append(amounts, oreRequirement{amountRequired, recipe.output.class})
		} else {
			amounts = append(amounts, rb.startCrawl(rec.amount*amountToCraft, rec.class, depth+1)...)
		}
	}

	return amounts
}

func getTotalOreNeeded(input string) int {
	data := getData(input)

	amounts := data.startCrawl(1, "FUEL", 0)

	for _, am := range amounts {
		fmt.Println(am)
	}
	totals := make(map[string]int, 0)

	for _, amt := range amounts {
		if _, ok := totals[amt.class]; ok {
			totals[amt.class] += amt.amount
		} else {
			totals[amt.class] = amt.amount
		}
	}

	totalOre := 0

	fmt.Printf("%-10v\t\t%-10v\t\t\t%-10v\t\t%-10v\t\t%-10v\t\t%-10v\n",
		"Recipe",
		"Needed",
		"Per Craft",
		"Base Ore",
		"crafts",
		"totalore")
	for k, v := range totals {
		recipe, _ := data.getByOutput(k)

		crafts := int(math.Ceil(float64(v) / float64(recipe.output.amount)))

		fmt.Printf("%-10v\t\t%-10v\t\t\t%-10v\t\t%-10v\t\t%-10v\t\t%-10v\n",
			recipe.output.class,
			v,
			recipe.output.amount,
			recipe.inputs[0].amount,
			crafts,
			crafts*recipe.inputs[0].amount)

		totalOre += crafts * recipe.inputs[0].amount
	}

	fmt.Println()
	return totalOre
}
func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	return fmt.Sprintf("%v", " -- Not Yet Implemented --")
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	return fmt.Sprintf("%v", " -- Not Yet Implemented --")
}
