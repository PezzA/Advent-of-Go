package Day201908

import (
	"fmt"
	"strconv"
)

var displayWidth = 25
var displayHeight = 6

func getData(input string) []int {
	vals := []int{}
	for i := range input {
		digit, _ := strconv.Atoi(string(input[i]))

		vals = append(vals, digit)
	}

	return vals
}

func getLayers(pix []int) [][]int {
	layers := [][]int{}

	for i := 0; i < len(pix); i += displayWidth * displayHeight {
		layers = append(layers, pix[i:i+(displayWidth*displayHeight)])
	}

	return layers
}

func getPixelCounts(pix []int) map[int]int {
	vals := make(map[int]int, 0)

	for i := range pix {
		if _, ok := vals[pix[i]]; ok {
			vals[pix[i]]++
		} else {
			vals[pix[i]] = 1
		}
	}
	return vals
}

func getTopVisiblePixel(pos int, layers [][]int) int {

	for index := range layers {
		if layers[index][pos] != 2 {
			return layers[index][pos]
		}
	}

	return 2
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	pix := getData(Entry.PuzzleInput())

	layers := getLayers(pix)

	leastZeros := -1
	leastZeroIndex := 0
	for lIndex, l := range layers {
		counts := getPixelCounts(l)
		if leastZeros == -1 || counts[0] < leastZeros {
			leastZeros = counts[0]
			leastZeroIndex = lIndex
		}
	}

	tlCounts := getPixelCounts(layers[leastZeroIndex])

	return fmt.Sprintf("%v", tlCounts[1]*tlCounts[2])
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	pix := getData(Entry.PuzzleInput())
	layers := getLayers(pix)
	display := ""

	for y := 0; y < displayHeight; y++ {
		for x := 0; x < displayWidth; x++ {

			switch getTopVisiblePixel(x+(y*displayWidth), layers) {
			case 0:
				display += " "
			case 1:
				display += "#"

			}
		}

		display += "\n"
	}

	if updateChan != nil {
		updateChan <- []string{"See output below"}
	}

	return fmt.Sprintf("%v", display)
}
