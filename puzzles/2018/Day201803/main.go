package Day201803

import (
	"fmt"
	"strconv"
	"strings"
)

// Entry holds wraps the data and runner interfaces for this puzzle
var Entry dayEntry

type dayEntry bool

type claim struct {
	id         string
	leftOffset int
	topOffset  int
	width      int
	height     int
}

const fabricLength = 1000

func getData(input string) []claim {
	claims := make([]claim, 0)

	for _, line := range strings.Split(input, "\n") {
		bits := strings.Fields(line)

		offSets := strings.Split(strings.Replace(bits[2], ":", "", -1), ",")
		pos := strings.Split(bits[3], "x")

		offsetLeft, _ := strconv.Atoi(offSets[0])
		offsetTop, _ := strconv.Atoi(offSets[1])

		width, _ := strconv.Atoi(pos[0])
		height, _ := strconv.Atoi(pos[1])

		claims = append(claims, claim{bits[0], offsetLeft, offsetTop, width, height})
	}
	return claims
}

func getCloth() [][]int {
	cloth := make([][]int, fabricLength)

	for index := range cloth {
		cloth[index] = make([]int, fabricLength)
	}
	return cloth
}

func getIdCloth() [][][]string {
	cloth := make([][][]string, fabricLength)

	for index := range cloth {
		cloth[index] = make([][]string, fabricLength)
	}
	return cloth
}

func applyClaims(claims []claim) ([][]int, [][][]string) {
	cloth := getCloth()
	idCloth := getIdCloth()

	for _, claim := range claims {
		for x := claim.leftOffset; x < claim.leftOffset+claim.width; x++ {
			for y := claim.topOffset; y < claim.topOffset+claim.height; y++ {
				cloth[y][x]++

				if idCloth[y][x] == nil {
					idCloth[y][x] = make([]string, 0)
				}

				idCloth[y][x] = append(idCloth[y][x], claim.id)
			}
		}
	}

	return cloth, idCloth
}

func countOverlaps(cloth [][]int) int {
	result := 0
	for x, col := range cloth {
		for y := range col {
			if cloth[x][y] > 1 {
				result++
			}
		}
	}
	return result
}

func getNonOverlappedClaim(taggedCloth [][][]string, claims []claim) string {
	candidates := make(map[string]bool)

	for _, claim := range claims {
		candidates[claim.id] = true
	}

	for x, col := range taggedCloth {
		for y := range col {
			if len(taggedCloth[x][y]) > 1 {
				for _, id := range taggedCloth[x][y] {
					delete(candidates, id)
				}

			}
		}
	}

	result := ""
	for k := range candidates {
		result = k
	}
	return result
}

func (td dayEntry) Describe() (int, int, string, int) {
	return 2018, 3, "No Matter How You Slice It"
}
func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	claims := getData(inputData)

	claimed, _ := applyClaims(claims)

	overlaps := countOverlaps(claimed)

	return fmt.Sprintf("%v", overlaps)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	claims := getData(inputData)

	_, tags := applyClaims(claims)

	nonOverlapped := getNonOverlappedClaim(tags, claims)

	return fmt.Sprintf("%v", nonOverlapped)

}
