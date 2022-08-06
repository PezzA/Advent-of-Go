package Day201515

import (
	"fmt"
	"strconv"
	"strings"
)

// Entry holds wraps the data and runner interfaces for this puzzle
var Entry dayEntry

type dayEntry bool

type ingredient struct {
	name       string
	capacity   int
	durability int
	flavour    int
	texture    int
	calories   int
}

type recipeComponent struct {
	name   string
	amount int
}

type recipe []recipeComponent

func getIngredients(input string) map[string]ingredient {
	ingredients := make(map[string]ingredient, 0)

	for _, line := range strings.Split(input, "\n") {
		bits := strings.Fields(line)
		name := strings.Replace(bits[0], ":", "", -1)

		capacity, _ := strconv.Atoi(strings.Replace(bits[2], ",", "", -1))
		durability, _ := strconv.Atoi(strings.Replace(bits[4], ",", "", -1))
		flavour, _ := strconv.Atoi(strings.Replace(bits[6], ",", "", -1))
		texture, _ := strconv.Atoi(strings.Replace(bits[8], ",", "", -1))
		calories, _ := strconv.Atoi(strings.Replace(bits[10], ",", "", -1))

		ingredients[name] = ingredient{name, capacity, durability, flavour, texture, calories}
	}

	return ingredients
}

func getIngredientList(ingredientList map[string]ingredient) []string {
	ingredients := make([]string, 0)
	for k := range ingredientList {
		ingredients = append(ingredients, k)
	}
	return ingredients
}

func scoreFormula(ingredientBook map[string]ingredient, formula []recipeComponent) (int, int) {
	capacity, durability, flavour, texture, calories := 0, 0, 0, 0, 0

	for _, val := range formula {

		capacity += val.amount * ingredientBook[val.name].capacity
		durability += val.amount * ingredientBook[val.name].durability
		flavour += val.amount * ingredientBook[val.name].flavour
		texture += val.amount * ingredientBook[val.name].texture
		calories += val.amount * ingredientBook[val.name].calories
	}

	if capacity < 0 {
		capacity = 0
	}

	if durability < 0 {
		durability = 0
	}
	if flavour < 0 {
		flavour = 0
	}
	if texture < 0 {
		texture = 0
	}

	return capacity * durability * flavour * texture, calories
}

// recursive
func kenwoodChef(ingredientList []string, composition recipe, formulaChan chan recipe) {
	if len(composition) == 0 {
		defer close(formulaChan)
	}

	compositionTotal := 0

	for _, comp := range composition {
		compositionTotal += comp.amount
	}

	numIngredients := len(ingredientList)

	for i := 1; i < 100; i++ {
		newComposition := append(composition, recipeComponent{ingredientList[len(composition)], i})

		if compositionTotal+i == 100 && len(newComposition) == numIngredients {
			formulaChan <- newComposition
			return
		}

		if compositionTotal+1 > 100 {
			return
		}

		if len(newComposition) < numIngredients {
			kenwoodChef(ingredientList, newComposition, formulaChan)
		}
	}
}

func (td dayEntry) Describe() (int, int, string, int) {
	return 2015, 15, "Science for Hungry People", 2
}
func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	ingredientBook := getIngredients(Entry.PuzzleInput())
	recipeChan := make(chan recipe)

	go kenwoodChef(getIngredientList(ingredientBook), make([]recipeComponent, 0), recipeChan)

	bestscore := 0
	for formula := range recipeChan {

		formulaScore, _ := scoreFormula(ingredientBook, formula)

		if formulaScore > bestscore {
			bestscore = formulaScore
		}
	}
	return fmt.Sprintf("%v", bestscore)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	ingredientBook := getIngredients(Entry.PuzzleInput())
	recipeChan := make(chan recipe)

	go kenwoodChef(getIngredientList(ingredientBook), make([]recipeComponent, 0), recipeChan)

	bestscore := 0
	for formula := range recipeChan {

		formulaScore, calories := scoreFormula(ingredientBook, formula)

		if formulaScore > bestscore && calories == 500 {
			bestscore = formulaScore
		}
	}
	return fmt.Sprintf("%v", bestscore)
}
