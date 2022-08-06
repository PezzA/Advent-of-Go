package Day202020

import (
	"fmt"
	"math"
	"strings"

	"github.com/pezza/advent-of-code/puzzles/common"
)

type grid [][]bool

type tile struct {
	id int
	grid
}

type placement struct {
	side int
	tile
}

const monsterMask = `                  # 
#    ##    ##    ###
 #  #  #  #  #  #   `

func getMonsterMask() grid {
	lines := strings.Split(monsterMask, "\n")

	mask := make(grid, len(lines))

	for i := range mask {
		mask[i] = make([]bool, len(lines[i]))
	}

	for i := range lines {
		for j := range lines[i] {
			mask[i][j] = lines[i][j] == 35
		}
	}

	return mask
}

func getData(input string) []tile {
	lines := strings.Split(input, "\n")
	data := make([]tile, 0)

	tileId := 0
	tileData := make(grid, 0)
	for _, line := range lines {
		if line == "" {
			data = append(data, tile{
				tileId, tileData,
			})
			tileData = make(grid, 0)
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

func createEmptyTile(size int) grid {
	newData := make(grid, size)
	for i := range newData {
		newData[i] = make([]bool, size)
	}
	return newData
}

func (g grid) rotate() grid {
	size := len(g)
	newData := createEmptyTile(size)

	// row 0 becomes col 9
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			newData[x][size-1-y] = g[y][x]
		}
	}

	return newData
}

func (g grid) flipTopBottom() grid {
	size := len(g)
	newData := createEmptyTile(size)

	// row 0 becomes row 9
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			newData[size-1-y][x] = g[y][x]
		}
	}

	return newData
}

func (g grid) flipLeftRight() grid {
	size := len(g)
	newData := createEmptyTile(size)

	// col 0 becomes row 9
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			newData[y][size-1-x] = g[y][x]
		}
	}

	return newData
}

func (t tile) printTile() {
	fmt.Printf("Tile %d:\n", t.id)
	t.grid.printData()
}

func (g grid) printData() {
	for y := 0; y < len(g); y++ {
		for x := 0; x < len(g); x++ {
			if g[y][x] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func (g grid) copyData() grid {
	newData := createEmptyTile(len(g))

	for y := 0; y < len(g); y++ {
		for x := 0; x < len(g); x++ {
			newData[y][x] = g[y][x]
		}
	}

	return newData
}

func (t tile) hasBorderMatch(cmp grid) (bool, int) {
	// top to bottom, bottom to top, left to right, right to left
	size := len(t.grid)

	topMatch := true
	for i := 0; i < size; i++ {
		if t.grid[0][i] != cmp[size-1][i] {
			topMatch = false
			break
		}
	}
	if topMatch {
		return true, 1
	}

	bottomMatch := true
	for i := 0; i < size; i++ {
		if t.grid[size-1][i] != cmp[0][i] {
			bottomMatch = false
			break
		}
	}
	if bottomMatch {
		return true, 4
	}

	leftMatch := true
	for i := 0; i < size; i++ {
		if t.grid[i][0] != cmp[i][size-1] {
			leftMatch = false
			break
		}
	}
	if leftMatch {
		return true, 8
	}

	rightMatch := true
	for i := 0; i < size; i++ {
		if t.grid[i][size-1] != cmp[i][0] {
			rightMatch = false
			break
		}
	}
	if rightMatch {
		return true, 2
	}

	return false, 0
}

func (t tile) matchingSides(jigsaw []tile, completed [][]tile) []placement {
	placements := make([]placement, 0)

	for _, piece := range jigsaw {
		if t.id == piece.id {
			continue
		}

		found := false
		for _, compRow := range completed {
			for _, compT := range compRow {
				if compT.id == piece.id {
					found = true
				}
			}
		}

		if found {
			continue
		}

		cmp := piece.copyData()
		for i := 0; i < 4; i++ {
			cmp = cmp.rotate()

			match, side := t.hasBorderMatch(cmp)
			if match {
				placements = append(placements, placement{side, tile{piece.id, cmp}})
				break
			}

			cmpLR := cmp.flipLeftRight()
			match, side = t.hasBorderMatch(cmpLR)
			if match {
				placements = append(placements, placement{side, tile{piece.id, cmpLR}})
				break
			}

			cmpTB := cmp.flipTopBottom()
			match, side = t.hasBorderMatch(cmpTB)
			if match {
				placements = append(placements, placement{side, tile{piece.id, cmpTB}})
				break
			}
		}
	}

	return placements
}

func getTopLeft(tiles []tile) (bool, tile) {
	for _, t := range tiles {
		matches := t.matchingSides(tiles, [][]tile{})
		if len(matches) == 2 {
			if (matches[0].side == 2 && matches[1].side == 4) || (matches[0].side == 4 && matches[1].side == 2) {
				return true, t
			}
		}
	}

	return false, tile{}
}

func (g grid) stripBorder() grid {
	newG := make(grid, len(g)-2)

	for i := range newG {
		newG[i] = make([]bool, len(g)-2)
	}

	for y := 1; y < len(g)-1; y++ {
		for x := 1; x < len(g)-1; x++ {
			newG[y-1][x-1] = g[y][x]
		}
	}

	return newG
}

// find corner piece that goes in the top left and then work it out in rows.
func solveJigSaw(tiles []tile) [][]tile {
	sideSize := int(math.Sqrt(float64(len(tiles))))
	_, topLeftTile := getTopLeft(tiles)

	completedJig := make([][]tile, sideSize)

	for y := 0; y < sideSize; y++ {
		completedJig[y] = make([]tile, sideSize)
	}

	completedJig[0][0] = topLeftTile

	for y := 0; y < sideSize; y++ {
		for x := 0; x < sideSize; x++ {
			matches := completedJig[y][x].matchingSides(tiles, completedJig)

			// we are going top to bottom, left to right
			for _, m := range matches {
				switch m.side {
				case 2:
					completedJig[y][x+1] = m.tile
				case 4:
					completedJig[y+1][x] = m.tile
				}
			}

		}
	}

	return completedJig
}

func stitchTogether(tiles [][]tile) grid {
	strippedTileLength := len(tiles[0][0].grid) - 2
	masterSize := len(tiles) * strippedTileLength

	masterGrid := make(grid, masterSize)
	for i := range masterGrid {
		masterGrid[i] = make([]bool, masterSize)
	}

	for y, row := range tiles {
		for x, t := range row {
			strippedTile := t.grid.stripBorder()

			for sy := range strippedTile {
				for sx := range strippedTile[sy] {
					masterGrid[(y*strippedTileLength)+sy][(x*strippedTileLength)+sx] = strippedTile[sy][sx]
				}
			}

		}
	}

	return masterGrid
}

func (g grid) getAllHash() map[common.Point]bool {
	points := make(map[common.Point]bool, 0)
	for y := 0; y < len(g); y++ {
		for x := 0; x < len(g); x++ {
			if g[y][x] {
				points[common.Point{X: x, Y: y}] = true
			}
		}
	}

	return points
}
func searchGrid(mask grid, masterGrid grid) (int, int) {
	masterSize := len(masterGrid)
	maskHeight, maskLength := len(mask), len(mask[0])

	hashMap := masterGrid.getAllHash()

	tot := 0
	for y := 0; y < masterSize-maskHeight; y++ {
		for x := 0; x < masterSize-maskLength; x++ {
			found := true

			monsterPieces := make([]common.Point, 0)

			for my := range mask {
				for mx := range mask[my] {
					if mask[my][mx] && masterGrid[y+my][x+mx] {
						monsterPieces = append(monsterPieces, common.Point{X: x + mx, Y: y + my})
					}

					if mask[my][mx] && !masterGrid[y+my][x+mx] {
						found = false
						break
					}

				}

				if !found {
					break
				}
			}

			if found {
				for _, mp := range monsterPieces {
					delete(hashMap, mp)
				}
				tot++
			}
		}
	}

	return tot, len(hashMap)
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	tiles := getData(inputData)
	tot := 1
	// don't have to put it all together, just need the corners
	for _, t := range tiles {
		if len(t.matchingSides(tiles, [][]tile{})) == 2 {
			tot *= t.id
		}
	}
	return fmt.Sprintf("%v", tot)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	tiles := getData(inputData)
	completedJig := solveJigSaw(tiles)
	grid := stitchTogether(completedJig)
	mask := getMonsterMask()
	cmp := grid.copyData()

	found, roughness := 0, 0

	for i := 0; i < 4; i++ {
		cmp = cmp.rotate()

		found, roughness = searchGrid(mask, cmp)

		if found > 0 {
			break
		}

		cmpLR := cmp.flipLeftRight()

		found, roughness = searchGrid(mask, cmpLR)

		if found > 0 {
			break
		}

		cmpTB := cmp.flipTopBottom()
		found, roughness = searchGrid(mask, cmpTB)

		if found > 0 {
			break
		}

	}

	return fmt.Sprintf("%v", roughness)
}
