package Day201920

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/pezza/advent-of-code/puzzles/common"
)

type maze [][]rune

type proximity struct {
	name     string
	distance int
	inner    bool
}

type portal struct {
	name  string
	inner common.Point
	outer common.Point
}

type portalList map[string]portal
type distanceMap map[common.Point]int

const floor = 46
const wall = 35

func newPortal(name string) portal {
	return portal{
		name:  name,
		inner: common.Point{},
		outer: common.Point{},
	}
}

func getData(input string) (maze, portalList) {
	lines := strings.Split(input, "\n")
	newMaze := make(maze, len(lines))

	tl := common.Point{2, 2}
	//br := common.Point{126, 118}
	br := common.Point{42, 34}
	for y, line := range lines {
		newMaze[y] = make([]rune, len(line))
		for x, char := range line {
			if char == wall || char == floor {
				newMaze[y][x] = char
			}
		}
	}

	pl := make(portalList, 0)

	for y := range newMaze {
		for x := range newMaze[y] {
			if newMaze[y][x] == floor {
				inner := true
				pos := common.Point{X: x, Y: y}
				if x == tl.X || y == tl.Y || x == br.X || y == br.Y {
					inner = false
				}

				title := ""
				if newMaze[y-1][x] == 0 {
					title = fmt.Sprintf("%v%v",
						string(lines[y-2][x]),
						string(lines[y-1][x]))
				}

				if newMaze[y+1][x] == 0 {
					title = fmt.Sprintf("%v%v",
						string(lines[y+1][x]),
						string(lines[y+2][x]))
				}

				if newMaze[y][x-1] == 0 {
					title = fmt.Sprintf("%v%v",
						string(lines[y][x-2]),
						string(lines[y][x-1]))
				}

				if newMaze[y][x+1] == 0 {
					title = fmt.Sprintf("%v%v",
						string(lines[y][x+1]),
						string(lines[y][x+2]))
				}

				if title != "" {
					if _, ok := pl[title]; !ok {
						pl[title] = newPortal(title)
					}

					val := pl[title]
					if inner {
						val.inner = pos
						pl[title] = val
					} else {
						val.outer = pos
						pl[title] = val
					}
				}
			}
		}
	}

	return newMaze, pl
}

func (m maze) getDistanceMap(p common.Point) distanceMap {
	d := make(distanceMap, 0)
	if p.X == 0 && p.Y == 0 {
		return d
	}
	m.distanceCrawl(p, d, 0)
	return d
}

func (m maze) distanceCrawl(p common.Point, d distanceMap, depth int) {
	d[p] = depth
	for _, dir := range common.Cardinal4 {
		newPos := p.Add(dir)
		if m[newPos.Y][newPos.X] != rune(floor) {
			continue
		}
		if val, ok := d[newPos]; ok && val <= depth {
			continue
		} else {
			m.distanceCrawl(newPos, d, depth+1)
		}
	}
}

func (pl portalList) getPortal(name string) (portal, error) {
	for k, v := range pl {
		if k == name {
			return v, nil
		}
	}
	return portal{}, errors.New("Could not find portal")
}

func (m maze) startWalk(pl portalList) []int {
	start, _ := pl.getPortal("AA")

	return m.portalWalk(start.outer, start, pl, 0, 0)
}

var travelled = [][]string{}

func (m maze) portalWalk(pos common.Point, p portal, pl portalList, distance int, recursiveDepth int) []int {
	fmt.Println(p.name, distance, pos, recursiveDepth)
	time.Sleep(time.Millisecond * 100)

	if len(travelled) == recursiveDepth {
		travelled = append(travelled, []string{})
	}

	distances := []int{}
	destinations := []proximity{}

	// build up list of what we can travel too
	for position, distance := range m.getDistanceMap(pos) {
		for _, portal := range pl {
			// normal exclusions
			if portal.name == p.name || common.Contains(travelled[recursiveDepth], portal.name) {
				continue
			}

			if recursiveDepth == 0 {
				// outer most, can only do start and end
				if portal.outer == position && portal.name != "AA" {
					continue
				}
			} else {
				// inner, no start or end
				if portal.name == "AA" || portal.name == "ZZ" {
					continue
				}
			}

			if portal.inner == position || portal.outer == position {
				distance += 1
				destinations = append(destinations, proximity{portal.name, distance, portal.inner == position})
			}
		}
	}

	for _, newP := range destinations {
		if newP.name == "ZZ" {
			return []int{distance - 1} // no step to last portal
		}
		newPortal, _ := pl.getPortal(newP.name)
		travelled[recursiveDepth] = append(travelled[recursiveDepth], p.name)

		var newPos common.Point

		if newP.inner {
			newPos = newPortal.outer
			recursiveDepth++
		} else {
			newPos = newPortal.inner
			recursiveDepth--
		}

		distances = append(distances, m.portalWalk(newPos, newPortal, pl, distance+newP.distance, recursiveDepth)...)
	}

	return distances
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	return fmt.Sprintf("%v", " -- Not Yet Implemented --")
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	return fmt.Sprintf("%v", " -- Not Yet Implemented --")
}

func printMaze(in maze) {
	for _, line := range in {
		for _, cell := range line {
			if cell == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print(string(cell))
			}
		}
		fmt.Println()
	}
}

func getBounds(m map[common.Point]int) (common.Point, common.Point) {
	minX, minY, maxX, maxY := -1, -1, 0, 0
	for k := range m {
		if minX == -1 || k.X < minX {
			minX = k.X
		}
		if minY == -1 || k.Y < minY {
			minY = k.Y
		}

		if k.X > maxX {
			maxX = k.X
		}
		if k.Y > maxY {
			maxY = k.Y
		}
	}
	return common.Point{minX, minY}, common.Point{maxX, maxY}
}

func printDmap(d distanceMap) {
	tl, br := getBounds(d)

	for y := tl.Y; y <= br.Y; y++ {
		for x := tl.X; x <= br.X; x++ {
			if val, ok := d[common.Point{x, y}]; ok {
				fmt.Printf("%4v", val)
			} else {
				fmt.Printf("%4v", "")
			}
		}
		fmt.Println()
	}
}
