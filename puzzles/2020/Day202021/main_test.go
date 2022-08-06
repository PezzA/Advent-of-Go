package Day202021

import (
	"fmt"
	"sort"
	"testing"

	"github.com/pezza/advent-of-code/common"

	. "github.com/onsi/gomega"
)

func Test_func(t *testing.T) {
	RegisterTestingT(t)
}

func Test_ReadData(t *testing.T) {
	RegisterTestingT(t)

	data, ingList, allList := getData(Entry.PuzzleInput())
	Expect(data[0].allergens).Should(Equal([]string{"shellfish", "fish"}))

	Expect(common.Contains(allList, "shellfish")).Should(Equal(true))
	Expect(common.Contains(ingList, "lflxjnh")).Should(Equal(true))
}

func Test_getPotentials(t *testing.T) {
	RegisterTestingT(t)

	foodList, _, allergenList := getData(Entry.PuzzleInput())

	potentialAllergens := make([]string, 0)

	allergenMap := make(map[string][]string, 0)

	for _, allergen := range allergenList {
		candidates := getPotentialAllergenCandidates(allergen, foodList)

		allergenMap[allergen] = candidates

		for _, candidate := range candidates {
			if !common.Contains(potentialAllergens, candidate) {
				potentialAllergens = append(potentialAllergens, candidate)
			}
		}
	}

	foundList := make([]allergen, 0)

	for {
		if len(allergenMap) == 0 {
			break
		}

		for k, v := range allergenMap {
			if len(v) == 1 {
				foundList = append(foundList, allergen{
					k, v[0],
				})

				delete(allergenMap, k)
				break
			}
		}

		for k, v := range allergenMap {
			for index, ing := range v {
				for _, found := range foundList {
					if ing == found.ingredient {
						allergenMap[k] = append(v[:index], v[index+1:]...)
					}
				}
			}
		}
	}

	sort.Slice(foundList, func(i, j int) bool {
		return foundList[i].name < foundList[j].name
	})

	val := ""
	for i := range foundList {

		if val != "" {
			val += ","
		}
		val += foundList[i].ingredient
	}

	fmt.Println(val)
}



func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

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
