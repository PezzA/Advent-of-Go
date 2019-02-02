package Day201818

import (
	"fmt"
	"strings"

	"github.com/pezza/advent-of-code/puzzles/Common"
)

var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string) {
	return 2018, 18, "Settlers of The North Pole"
}

const (
	openArea   = 46
	treeArea   = 124
	lumberYard = 35
)

type acreage [][]int

func getData(input string) acreage {
	land := make(acreage, 0)

	for _, line := range strings.Split(input, "\n") {
		row := make([]int, len(line))

		for index, char := range line {
			row[index] = int(char)
		}

		land = append(land, row)
	}
	return land
}

func (a acreage) draw() {
	for y := range a {
		for x := range a[y] {
			fmt.Print(string(a[y][x]))
		}
		fmt.Print("\n")
	}
}

func (a acreage) getCounts(p common.Point) (int, int, int) {
	open, tree, lumber := 0, 0, 0

	switch a[p.Y][p.X] {
	case openArea:
		open++
	case treeArea:
		tree++
	case lumberYard:
		lumber++

	}

	for _, card := range append(common.Cardinal8) {
		tp := p.Add(card)
		if tp.X < 0 || tp.Y < 0 || tp.X >= len(a) || tp.Y >= len(a) {
			continue
		}

		switch a[tp.Y][tp.X] {
		case openArea:
			open++
		case treeArea:
			tree++
		case lumberYard:
			lumber++

		}
	}
	return open, tree, lumber
}

func (a acreage) transformCell(p common.Point) int {
	_, tree, lumber := a.getCounts(p)

	switch a[p.Y][p.X] {
	case openArea:
		if tree >= 3 {
			return treeArea
		}
	case treeArea:
		if lumber >= 3 {
			return lumberYard
		}
	case lumberYard:
		if lumber >= 2 && tree >= 1 {
			return lumberYard
		}
		return openArea

	}
	return a[p.Y][p.X]
}

func (a acreage) transform() acreage {
	newAcreage := make(acreage, 0)
	for index := range a {
		newAcreage = append(newAcreage, make([]int, len(a[index])))
	}

	for y := range a {
		for x := range a[y] {
			newAcreage[y][x] = a.transformCell(common.Point{x, y})
		}
	}
	return newAcreage
}

func (a acreage) getTypeCount(i int) int {
	count := 0
	for y := range a {
		for x := range a[y] {
			if a[y][x] == i {
				count++
			}
		}
	}
	return count
}

func (a acreage) sameAs(t acreage) bool {
	for index := range a {
		if !common.Same(a[index], t[index]) {
			return false
		}
	}
	return true
}

func contains(s []acreage, t acreage) (bool, int) {
	for index, a := range s {
		if t.sameAs(a) {
			return true, index
		}
	}
	return false, 0
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	land := getData(inputData)
	for i := 0; i < 10; i++ {
		land = land.transform()
	}

	return fmt.Sprintf("%v", land.getTypeCount(treeArea)*land.getTypeCount(lumberYard))
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	land := getData(Entry.PuzzleInput())
	lands := []acreage{land}

	returnVal := 0
	for i := 1; i <= 1000000000; i++ {
		land = land.transform()

		if ok, foundIndex := contains(lands, land); ok {
			loopSpan := 1000000000 - i
			loopLength := i - foundIndex

			returnVal = lands[foundIndex+(loopSpan%loopLength)].getTypeCount(treeArea) * lands[foundIndex+(loopSpan%loopLength)].getTypeCount(lumberYard)

			break
		} else {
			lands = append(lands, land)
		}
	}

	return fmt.Sprintf("%v", returnVal)
}
