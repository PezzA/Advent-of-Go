package Day201812

import (
	"fmt"
	"strings"
)

var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string, int) {
	return 2018, 12, "Subterranean Sustainability", 0
}

type transform struct {
	pattern       []bool
	producesPlant bool
}

func getData(input string) (map[int]bool, []transform) {
	data := strings.Split(input, "\n")

	plantlist := strings.Fields(data[0])[2]

	pl := make(map[int]bool)

	for index, val := range plantlist {
		pl[index] = val == 35
	}
	transforms := make([]transform, 0)

	for i := 2; i < len(data); i++ {
		bits := strings.Fields(data[i])
		transforms = append(transforms, transform{
			[]bool{
				bits[0][0] == 35,
				bits[0][1] == 35,
				bits[0][2] == 35,
				bits[0][3] == 35,
				bits[0][4] == 35},
			bits[2] == "#"})
	}

	return pl, transforms
}

func same(source []bool, target []bool) bool {
	for index := range source {
		if source[index] != target[index] {
			return false
		}
	}
	return true
}

func getBounds(input map[int]bool) (int, int) {

	testMin, testMax := 0, 0

	first := true
	for k, _ := range input {

		if first || k < testMin {
			testMin = k
		}

		if first || k > testMax {
			testMax = k
		}
		first = false
	}

	return testMin, testMax
}

func regenerate(pl map[int]bool, tl []transform) map[int]bool {

	min, max := getBounds(pl)

	newPl := make(map[int]bool)
	for i := min - 2; i < max+2; i++ {
		testList := []bool{pl[i-2], pl[i-1], pl[i], pl[i+1], pl[i+2]}
		newPl[i] = false
		for _, val := range tl {
			if val.producesPlant {
				if same(testList, val.pattern) {
					newPl[i] = true
				}
			}
		}
	}

	for i := min - 2; i < max; i++ {
		if newPl[i] {
			break
		}

		delete(newPl, i)
	}

	for i := max + 2; i == 0; i-- {
		if newPl[i] {
			break
		}

		delete(newPl, i)
	}
	return newPl
}

func drawPlants(input map[int]bool) {
	_, max := getBounds(input)
	for i := -10; i < max; i++ {
		if input[i] {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
	}
	fmt.Print("\n")
}

func isShifted(source map[int]bool, target map[int]bool) bool {
	for k, _ := range source {
		if source[k] != target[k+1] {
			return false
		}
	}

	return true
}
func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {

	pl, tl := getData(inputData)

	for i := 0; i < 20; i++ {
		pl = regenerate(pl, tl)
	}

	count := 0

	for k, v := range pl {
		if v {
			count += k
		}
	}
	return fmt.Sprintf("%v", count)
}

func deepCopy(input map[int]bool) map[int]bool {
	newCopy := make(map[int]bool)

	for k, v := range input {
		newCopy[k] = v
	}

	return newCopy
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {

	opl, tl := getData(Entry.PuzzleInput())

	count := 50000000000

	inLoop := false
	for !inLoop {
		pl := deepCopy(opl)
		pl = regenerate(pl, tl)

		inLoop = isShifted(opl, pl)

		opl = pl
		count--
	}

	finalcount := 0
	for k, _ := range opl {
		if opl[k] {
			finalcount += k + count
		}
	}

	// high 2100000000470
	return fmt.Sprintf("%v", finalcount)
}
