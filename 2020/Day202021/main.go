package Day202021

import (
	"fmt"
	"sort"
	"strings"

	"github.com/pezza/advent-of-code/common"
)

type food struct {
	ingredients []string
	allergens   []string
}

type allergen struct {
	name       string
	ingredient string
}

func getData(input string) ([]food, []string, []string) {
	lines := strings.Split(input, "\n")
	data := make([]food, len(lines))

	masterIngList := make([]string, 0)
	masterAllList := make([]string, 0)

	for i, line := range lines {
		bits := strings.Split(line, " (contains ")

		ingredients := strings.Split(bits[0], " ")

		for _, ing := range ingredients {
			if !common.Contains(masterIngList, ing) {
				masterIngList = append(masterIngList, ing)
			}
		}

		allergens := strings.Split(bits[1], " ")
		allergens[len(allergens)-1] = strings.Replace(allergens[len(allergens)-1], ")", "", 1)

		for a := range allergens {
			allergens[a] = strings.Replace(allergens[a], ",", "", 1)
			if !common.Contains(masterAllList, allergens[a]) {
				masterAllList = append(masterAllList, allergens[a])
			}
		}

		data[i] = food{
			ingredients: ingredients,
			allergens:   allergens,
		}
	}

	return data, masterIngList, masterAllList
}

func getPotentialAllergenCandidates(allergen string, foods []food) []string {
	candList := make(map[string]int, 0)

	occ := 0
	for _, food := range foods {
		if !common.Contains(food.allergens, allergen) {
			continue
		}
		occ++
		for _, ing := range food.ingredients {
			if _, ok := candList[ing]; ok {
				candList[ing]++
			} else {
				candList[ing] = 1
			}
		}
	}

	candidates := make([]string, 0)
	for k, v := range candList {
		if v == occ {
			candidates = append(candidates, k)
		}
	}

	return candidates
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	foodList, _, allergenList := getData(inputData)

	potentialAllergens := make([]string, 0)
	for _, allergen := range allergenList {
		candidates := getPotentialAllergenCandidates(allergen, foodList)

		for _, candidate := range candidates {
			if !common.Contains(potentialAllergens, candidate) {
				potentialAllergens = append(potentialAllergens, candidate)
			}
		}
	}

	nonAllergenOccurence := 0
	for _, food := range foodList {
		for _, ingredient := range food.ingredients {
			if !common.Contains(potentialAllergens, ingredient) {
				nonAllergenOccurence++
			}
		}
	}

	return fmt.Sprintf("%v", nonAllergenOccurence)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
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

	return fmt.Sprintf("%v", val)
}
