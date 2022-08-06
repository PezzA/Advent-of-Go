package Day201518

import (
	"fmt"
	"strings"

	"github.com/pezza/advent-of-code/puzzles/common"
)

// Entry holds wraps the data and runner interfaces for this puzzle
var Entry dayEntry

type dayEntry bool

type display [][]int

func (d display) getCellValue(p common.Point) int {
	maxLen := len(d) - 1
	if p.X < 0 || p.Y < 0 || p.X > maxLen || p.Y > maxLen {
		return 0
	}

	return d[p.Y][p.X]
}

func (d display) printCell(p common.Point) {
	fmt.Println(d.getCellValue(common.Point{p.X - 1, p.Y - 1}), d.getCellValue(common.Point{p.X, p.Y - 1}), d.getCellValue(common.Point{p.X + 1, p.Y - 1}))
	fmt.Println(d.getCellValue(common.Point{p.X - 1, p.Y}), d.getCellValue(common.Point{p.X, p.Y}), d.getCellValue(common.Point{p.X + 1, p.Y}))
	fmt.Println(d.getCellValue(common.Point{p.X - 1, p.Y + 1}), d.getCellValue(common.Point{p.X, p.Y + 1}), d.getCellValue(common.Point{p.X + 1, p.Y + 1}))
	fmt.Println()
}

func (d display) getNeighbourCount(p common.Point) int {
	tl, t, tr := d.getCellValue(common.Point{p.X - 1, p.Y - 1}), d.getCellValue(common.Point{p.X, p.Y - 1}), d.getCellValue(common.Point{p.X + 1, p.Y - 1})
	l, _, r := d.getCellValue(common.Point{p.X - 1, p.Y}), d.getCellValue(common.Point{p.X, p.Y}), d.getCellValue(common.Point{p.X + 1, p.Y})
	bl, b, br := d.getCellValue(common.Point{p.X - 1, p.Y + 1}), d.getCellValue(common.Point{p.X, p.Y + 1}), d.getCellValue(common.Point{p.X + 1, p.Y + 1})
	return tl + t + tr + l + r + bl + b + br
}

func getData(input string) display {
	startingScreen := make(display, 0)
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		screenLine := make([]int, 0)
		for _, char := range line {

			val := 0
			if char == 35 {
				val = 1
			}

			screenLine = append(screenLine, val)
		}
		startingScreen = append(startingScreen, screenLine)
	}

	return startingScreen
}

func (d display) getNewCellValue(p common.Point) int {
	currentState, neighbourCount := d.getCellValue(p), d.getNeighbourCount(p)

	if neighbourCount == 3 || (currentState == 1 && neighbourCount == 2) {
		return 1
	}

	return 0
}

func (d display) pixelOnCount() int {
	count := 0
	for _, line := range d {
		for _, cell := range line {
			count += cell
		}
	}
	return count
}
func (d display) drawFrame() {
	for _, line := range d {
		fmt.Println(line)
	}
}
func (d display) getNewFrame() display {
	newFrame := make(display, len(d))

	for i := range newFrame {
		newFrame[i] = make([]int, len(d[i]))
	}

	for y := 0; y < len(d); y++ {
		for x := 0; x < len(d[y]); x++ {
			newFrame[y][x] = d.getNewCellValue(common.Point{x, y})
		}
	}

	return newFrame
}

func (td dayEntry) Describe() (int, int, string, int) {
	return 2015, 18, "Like a GIF For Your Yard", 2
}
func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {

	startFrame := getData(inputData)

	for i := 0; i < 100; i++ {
		startFrame = startFrame.getNewFrame()
		//updateChan <- []string{fmt.Sprintf("Processing frame %v", i)}
	}

	return fmt.Sprintf("%v", startFrame.pixelOnCount())
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	startFrame := getData(inputData)

	startFrame[0][0] = 1
	startFrame[0][99] = 1
	startFrame[99][0] = 1
	startFrame[99][99] = 1

	for i := 0; i < 100; i++ {
		startFrame = startFrame.getNewFrame()
		startFrame[0][0] = 1
		startFrame[0][99] = 1
		startFrame[99][0] = 1
		startFrame[99][99] = 1

	}

	return fmt.Sprintf("%v", startFrame.pixelOnCount())
}
