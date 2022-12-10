package Day202208

import (
	"fmt"
	"strings"

	"github.com/pezza/advent-of-code/puzzles/common"
)

// assumes will always contains digits 0-9
func getData(input string) ([][]int, int, int) {
	lines := strings.Split(input, "\n")

	grid := make([][]int, len(lines))

	for lineIndex, line := range lines {
		grid[lineIndex] = make([]int, 0)

		for _, char := range line {
			// convert ascii code for number to number itself
			grid[lineIndex] = append(grid[lineIndex], int(char)-48)
		}
	}

	return grid, len(grid[0]), len(grid)
}

func isTreeVisible(grid [][]int, point common.Point) bool {
	if point.X == 0 || point.Y == 0 || point.X == len(grid)-1 || point.Y == len(grid[0])-1 {
		return true
	}

	if canViewFromUp(grid, point) ||
		canViewFromDown(grid, point) ||
		canViewFromLeft(grid, point) ||
		canViewFromRight(grid, point) {
		return true
	}

	return false
}

func getScenicScore(grid [][]int, point common.Point) int {
	height := grid[point.Y][point.X]
	seeUp, seeDown, seeLeft, seeRight := 0, 0, 0, 0
	// how many up
	for i := point.Y - 1; i >= 0; i-- {
		seeUp++
		if grid[i][point.X] >= height {
			break
		}
	}

	// how many down
	for i := point.Y + 1; i < len(grid); i++ {
		seeDown++
		if grid[i][point.X] >= height {
			break
		}
	}

	// how many left
	for i := point.X - 1; i >= 0; i-- {
		seeLeft++
		if grid[point.Y][i] >= height {
			break
		}
	}
	// how many right
	for i := point.X + 1; i < len(grid[0]); i++ {
		seeRight++
		if grid[point.Y][i] >= height {
			break
		}
	}
	return seeUp * seeDown * seeLeft * seeRight
}

func canViewFromUp(grid [][]int, point common.Point) bool {
	height := grid[point.Y][point.X]

	for i := point.Y - 1; i >= 0; i-- {
		if grid[i][point.X] >= height {
			return false
		}
	}
	return true
}

func canViewFromDown(grid [][]int, point common.Point) bool {
	height := grid[point.Y][point.X]

	for i := point.Y + 1; i < len(grid); i++ {
		if grid[i][point.X] >= height {
			return false
		}
	}
	return true
}

func canViewFromLeft(grid [][]int, point common.Point) bool {
	height := grid[point.Y][point.X]

	for i := point.X - 1; i >= 0; i-- {
		if grid[point.Y][i] >= height {
			return false
		}
	}
	return true
}

func canViewFromRight(grid [][]int, point common.Point) bool {
	height := grid[point.Y][point.X]

	for i := point.X + 1; i < len(grid[0]); i++ {
		if grid[point.Y][i] >= height {
			return false
		}
	}
	return true
}

func getTreeCount(grid [][]int) int {
	count := 0
	for x := 0; x < len(grid[0]); x++ {
		for y := 0; y < len(grid); y++ {
			if isTreeVisible(grid, common.Point{x, y}) {
				count++
			}
		}
	}
	return count
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	grid, _, _ := getData(inputData)
	return fmt.Sprint(getTreeCount(grid))
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	grid, _, _ := getData(inputData)

	bestScenicScore := 0

	for x := 0; x < len(grid[0]); x++ {
		for y := 0; y < len(grid); y++ {
			score := getScenicScore(grid, common.Point{x, y})

			if score > bestScenicScore {
				bestScenicScore = score
			}
		}
	}

	return fmt.Sprint(bestScenicScore)
}
