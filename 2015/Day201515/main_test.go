package Day201515

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

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

	// != 13324953600
	fmt.Println(bestscore)
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)
}

func Benchmark_BenchPartOne(b *testing.B) {
	data := Entry.PuzzleInput()
	for n := 0; n < b.N; n++ {
		Entry.PartOne(data, nil)
	}
}

func Benchmark_BenchPartTwo(b *testing.B) {
	data := Entry.PuzzleInput()
	for n := 0; n < b.N; n++ {
		Entry.PartTwo(data, nil)
	}
}
