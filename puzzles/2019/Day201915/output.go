package Day201915

import (
	"fmt"

	"github.com/pezza/advent-of-code/puzzles/common"
)

func drawcounts(maze map[common.Point]int) {
	minX, minY, maxX, maxY := common.MaxInt, common.MaxInt, common.MinInt, common.MinInt

	for k, _ := range maze {
		if k.X < minX {
			minX = k.X
		}

		if k.Y < minY {
			minY = k.Y
		}

		if k.X > maxX {
			maxX = k.X
		}

		if k.Y > maxY {
			maxY = k.Y
		}
	}

	for y := minY; y < maxY+1; y++ {
		for x := minX; x < maxX+1; x++ {

			if val, ok := maze[common.Point{X: x, Y: y}]; ok {
				fmt.Printf("%-4v", val)

			} else {
				fmt.Print("    ")
			}
		}
		fmt.Print("\n")
	}
}

func drawMaze(maze map[common.Point]int, rob common.Point) {
	minX, minY, maxX, maxY := common.MaxInt, common.MaxInt, common.MinInt, common.MinInt

	for k, _ := range maze {
		if k.X < minX {
			minX = k.X
		}

		if k.Y < minY {
			minY = k.Y
		}

		if k.X > maxX {
			maxX = k.X
		}

		if k.Y > maxY {
			maxY = k.Y
		}
	}

	for y := minY; y < maxY+1; y++ {
		for x := minX; x < maxX+1; x++ {

			if x == rob.X && y == rob.Y {
				fmt.Print("D")
				continue
			}

			if val, ok := maze[common.Point{X: x, Y: y}]; ok {
				switch val {
				case 0:
					fmt.Print(" ")
				case 1:
					fmt.Print("#")
				case 2:
					fmt.Print("A")
				}
			} else {
				fmt.Print("_")
			}
		}
		fmt.Print("\n")
	}
}
