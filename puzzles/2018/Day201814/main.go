package Day201814

import (
	"container/list"
	"fmt"
	"strconv"
)

var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string, int) {
	return 2018, 14, "Chocolate Charts", 0
}

func getData(input string) int {
	ret, _ := strconv.Atoi(input)
	return ret
}

func isResult(test string, scores *list.List) bool {
	back := scores.Back()
	check := ""

	for i := 0; i < len(test); i++ {
		check += strconv.Itoa(back.Value.(int))
		back = back.Prev()
	}

	if check == test {
		return true
	}
	return false
}

func getRecipes(elf1 *list.Element, elf2 *list.Element, scores *list.List, test string) (*list.Element, *list.Element, int) {
	sum := elf1.Value.(int) + elf2.Value.(int)

	for _, char := range strconv.Itoa(sum) {
		newScore, _ := strconv.Atoi(string(char))
		scores.PushBack(newScore)

		if test != "" && scores.Len() > 20 {
			if isResult(test, scores) {
				return nil, nil, scores.Len() - len(test)
			}
		}
	}

	loopVal := elf1.Value.(int) + 1
	for i := 0; i < loopVal; i++ {
		if elf1.Next() == nil {
			elf1 = scores.Front()
		} else {
			elf1 = elf1.Next()
		}
	}
	loopVal = elf2.Value.(int) + 1
	for i := 0; i < loopVal; i++ {
		if elf2.Next() == nil {
			elf2 = scores.Front()
		} else {
			elf2 = elf2.Next()
		}
	}

	return elf1, elf2, 0
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {

	maxVal := getData(inputData)

	scores := list.New()
	elf1 := scores.PushBack(3)
	elf2 := scores.PushBack(7)
	output := ""
	for {
		elf1, elf2, _ = getRecipes(elf1, elf2, scores, "")

		if scores.Len() > maxVal+10 {
			pos := scores.Front()
			for i := 0; i < maxVal; i++ {
				pos = pos.Next()
			}

			for i := 0; i < 10; i++ {
				output += strconv.Itoa(pos.Value.(int))
				pos = pos.Next()
			}
			break
		}
	}

	return output
}

func Reverse(s string) string {
	var reverse string
	for i := len(s) - 1; i >= 0; i-- {
		reverse += string(s[i])
	}
	return reverse
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	scores := list.New()
	elf1 := scores.PushBack(3)
	elf2 := scores.PushBack(7)

	reverseString := Reverse(inputData)
	val := 0
	for {
		elf1, elf2, val = getRecipes(elf1, elf2, scores, reverseString)

		if val != 0 {
			break
		}
	}

	return fmt.Sprintf("%v", val)
}
