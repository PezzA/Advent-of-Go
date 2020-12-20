package Day202020

import (
	"fmt"
	"strings"
)

type tile struct {
	id     int
	active [][]bool
}

func getData(input string) []tile {
	lines := strings.Split(input, "\n")
	data := make([]tile, 0)

	tileId := 0
	tileData := make([][]bool, 0)
	for _, line := range lines {
		if line == "" {
			data = append(data, tile{
				tileId, tileData,
			})
			tileData = make([][]bool, 0)
			continue
		}

		if strings.HasPrefix(line, "Tile") {
			fmt.Sscanf(line, "Tile %d:", &tileId)
			continue
		}

		tileLine := make([]bool, 0)
		for _, c := range line {
			tileLine = append(tileLine, c == 35)
		}

		tileData = append(tileData, tileLine)
	}

	return data
}

func createEmptyTile(size int) [][]bool {
	newData := make([][]bool, size)
	for i := range newData {
		newData[i] = make([]bool, size)
	}
	return newData
}

func rotate(data [][]bool) [][]bool {
	size := len(data)
	newData := createEmptyTile(size)

	// row 0 becomes col 9
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			newData[x][size-1-y] = data[y][x]
		}
	}

	return newData
}

func flipTopBottom(data [][]bool) [][]bool {
	size := len(data)
	newData := createEmptyTile(size)

	// row 0 becomes row 9
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			newData[size-1-y][x] = data[y][x]
		}
	}

	return newData
}

func flipLeftRight(data [][]bool) [][]bool {
	size := len(data)
	newData := createEmptyTile(size)

	// col 0 becomes row 9
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			newData[y][size-1-x] = data[y][x]
		}
	}

	return newData
}

func (t tile) printTile() {
	fmt.Printf("Tile %d:\n", t.id)
	printData(t.active)
}

func printData(data [][]bool) {
	for y := 0; y < len(data); y++ {
		for x := 0; x < len(data); x++ {
			if data[y][x] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func (t tile) copyData() [][]bool {
	newData := createEmptyTile(len(t.active))

	for y := 0; y < len(t.active); y++ {
		for x := 0; x < len(t.active); x++ {
			newData[y][x] = t.active[y][x]
		}
	}

	return newData
}

func (t tile) hasBorderMatch(cmp [][]bool) (bool, string) {
	// top to bottom, bottom to top, left to right, right to left
	size := len(t.active)

	topMatch := true
	for i := 0; i < size; i++ {
		if t.active[0][i] != cmp[size-1][i] {
			topMatch = false
			break
		}
	}
	if topMatch {
		return true, "above"
	}

	bottomMatch := true
	for i := 0; i < size; i++ {
		if t.active[size-1][i] != cmp[0][i] {
			bottomMatch = false
			break
		}
	}
	if bottomMatch {
		return true, "below"
	}

	leftMatch := true
	for i := 0; i < size; i++ {
		if t.active[i][0] != cmp[i][size-1] {
			leftMatch = false
			break
		}
	}
	if leftMatch {
		return true, "left"
	}

	rightMatch := true
	for i := 0; i < size; i++ {
		if t.active[i][size-1] != cmp[i][0] {
			rightMatch = false
			break
		}
	}
	if rightMatch {
		return true, "right"
	}

	return false, ""
}

func (t tile) matchingSides(jigsaw []tile) int {
	tot := 0
	for _, piece := range jigsaw {
		if t.id == piece.id {
			continue
		}

		cmp := piece.copyData()
		for i := 0; i < 4; i++ {
			cmp = rotate(cmp)

			match, _ := t.hasBorderMatch(cmp)
			if match {
				tot++
				break
			}

			cmplr := flipLeftRight(cmp)
			match, _ = t.hasBorderMatch(cmplr)
			if match {
				tot++
				break
			}

			cmptb := flipTopBottom(cmp)
			match, _ = t.hasBorderMatch(cmptb)
			if match {
				tot++
				break
			}
		}
	}

	return tot
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	return fmt.Sprintf("%v", " -- Not Yet Implemented --")
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	return fmt.Sprintf("%v", " -- Not Yet Implemented --")
}
