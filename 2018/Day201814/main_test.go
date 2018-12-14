package Day201814

import (
	"testing"

	"fmt"

	. "github.com/onsi/gomega"
)

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	elf1, elf2, recipies := getRecipes(0, 1, []int{3, 7})
	Expect(elf1).Should(Equal(0))
	Expect(elf2).Should(Equal(1))
	Expect(recipies).Should(Equal([]int{3, 7, 1, 0}))

	elf1, elf2, recipies = getRecipes(elf1, elf2, recipies)

	Expect(elf1).Should(Equal(4))
	Expect(elf2).Should(Equal(3))
	Expect(recipies).Should(Equal([]int{3, 7, 1, 0, 1, 0}))

	elf1, elf2, recipies = 0, 1, []int{3, 7}

	for len(recipies) < 320900 {
		elf1, elf2, recipies = getRecipes(elf1, elf2, recipies)
	}

	fmt.Println(recipies[9:19])
	fmt.Println(recipies[5:15])
	fmt.Println(recipies[18:28])
	fmt.Println(recipies[2018:2028])
	fmt.Println(recipies[320851:320861])
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

	elf1, elf2, recipies := 0, 1, []int{3, 7}
	check := []int{3, 2, 0, 8, 5, 1}
	checkLen := len(check)

	foundIndex := -1

	prevSize := len(recipies)

	for foundIndex == -1 {
		elf1, elf2, recipies = getRecipes(elf1, elf2, recipies)

		if len(recipies) > checkLen {

			for i := len(recipies) - checkLen; i > len(recipies)-(prevSize-checkLen); i-- {
				isSame := true
				for index := range check {
					if recipies[i+index] != check[index] {
						isSame = false
						break
					}
				}

				if isSame {
					foundIndex = i
					break
				}
			}

		}
		prevSize = len(recipies)
		if len(recipies)%10000 == 0 {
			fmt.Println(len(recipies))
		}
	}

	fmt.Println(foundIndex)

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
