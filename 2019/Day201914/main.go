package Day201914

import (
	"fmt"
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

		inputs := []ingredient{}
		inputData := strings.Split(bits[0], ",")

		for _, data := range inputData {
			inputs = append(inputs, parseIngredient(data))
		}

		output := parseIngredient(bits[1])
		recipes = append(recipes, recipe{inputs, output})
	}
	return recipes
}

// checked the data, each reaction has a unique output
func (rb recipeBook) getByOutput(class string) (recipe, error) {
	for _, recipe := range rb {
		if recipe.output.class == class {
			return recipe, nil
		}
	}
	return recipe{}, fmt.Errorf(fmt.Sprintf("Could not find recipe with output %v", class))
}

func (rb recipeBook) startCrawl() {

	fuel, _ := rb.getByOutput("FUEL")

	rb.processCrawl(fuel, 0)
}

func (rb recipeBook) processCrawl(r recipe, depth int) int {
	fmt.Printf("%v%v\n", strings.Repeat(" ", depth*4), r.output.class)

	for _, ing := range r.inputs {

		if ing.class == "ORE" {
			return ing.amount
			fmt.Printf("%vORE (%v)\n", strings.Repeat(" ", depth*4), ing.amount)
		} else {
			rec, _ := rb.getByOutput(ing.class)
			rb.processCrawl(rec, depth+1)
		}
	}
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	return fmt.Sprintf("%v", " -- Not Yet Implemented --")
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	return fmt.Sprintf("%v", " -- Not Yet Implemented --")
}
