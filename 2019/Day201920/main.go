package Day201920

import (
	"fmt"
	"strings"

	"github.com/pezza/advent-of-code/common"
)

type maze [][]rune

type proximity struct {
	name     string
	distance int
}

type portal struct {
	name  string
	start common.Point
	ins   []proximity
	end   common.Point
	outs  []proximity
}

type portalList map[string]portal
type distanceMap map[common.Point]int

const floor = 46
const wall = 35

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

func getData(input string) (maze, portalList) {
	lines := strings.Split(input, "\n")
	newMaze := make(maze, len(lines))

	for y, line := range lines {
		newMaze[y] = make([]rune, len(line))
		for x, char := range line {
			if char == wall || char == floor {
				newMaze[y][x] = char
			}
		}
	}

	portals := make(portalList, 0)

	for y := range newMaze {
		for x := range newMaze[y] {
			if newMaze[y][x] == floor {
				if newMaze[y-1][x] == 0 {
					title := fmt.Sprintf("%v%v",
						string(lines[y-2][x]),
						string(lines[y-1][x]))

					if _, ok := portals[title]; ok {
						portals[title] = portal{
							name:  title,
							start: portals[title].start,
							ins:   []proximity{},
							end:   common.Point{X: x, Y: y},
							outs:  []proximity{},
						}
					} else {
						portals[title] = portal{
							name:  title,
							start: common.Point{X: x, Y: y},
							ins:   []proximity{},
							end:   common.Point{},
							outs:  []proximity{},
						}
					}
				}

				if newMaze[y+1][x] == 0 {
					title := fmt.Sprintf("%v%v",
						string(lines[y+1][x]),
						string(lines[y+2][x]))

					if _, ok := portals[title]; ok {
						portals[title] = portal{
							name:  title,
							start: portals[title].start,
							ins:   []proximity{},
							end:   common.Point{X: x, Y: y},
							outs:  []proximity{},
						}
					} else {
						portals[title] = portal{
							name:  title,
							start: common.Point{X: x, Y: y},
							ins:   []proximity{},
							end:   common.Point{},
							outs:  []proximity{},
						}
					}
				}
				if newMaze[y][x-1] == 0 {
					title := fmt.Sprintf("%v%v",
						string(lines[y][x-2]),
						string(lines[y][x-1]))

					if _, ok := portals[title]; ok {
						portals[title] = portal{
							name:  title,
							start: portals[title].start,
							ins:   []proximity{},
							end:   common.Point{X: x, Y: y},
							outs:  []proximity{},
						}
					} else {
						portals[title] = portal{
							name:  title,
							start: common.Point{X: x, Y: y},
							ins:   []proximity{},
							end:   common.Point{},
							outs:  []proximity{},
						}
					}
				}

				if newMaze[y][x+1] == 0 {
					title := fmt.Sprintf("%v%v",
						string(lines[y][x+1]),
						string(lines[y][x+2]))

					if _, ok := portals[title]; ok {
						portals[title] = portal{
							name:  title,
							start: portals[title].start,
							ins:   []proximity{},
							end:   common.Point{X: x, Y: y},
							outs:  []proximity{},
						}
					} else {
						portals[title] = portal{
							name:  title,
							start: common.Point{X: x, Y: y},
							ins:   []proximity{},
							end:   common.Point{},
							outs:  []proximity{},
						}
					}
				}
			}
		}
	}

	return newMaze, portals
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

func getDistanceMap(p common.Point, m maze) distanceMap {
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

func (p portal) getConnections(pl portalList, m maze) portal {

	for point, distance := range getDistanceMap(p.start, m) {
		if distance > 0 {
			for tPortalName, tPortal := range pl {
				if (point == tPortal.start || point == tPortal.end) && tPortalName != p.name {
					p.ins = append(p.ins, proximity{tPortalName, distance})
				}
			}
		}
	}

	for point, distance := range getDistanceMap(p.end, m) {
		if distance > 0 {
			for tPortalName, tPortal := range pl {
				if (point == tPortal.start || point == tPortal.end) && tPortalName != p.name {
					p.outs = append(p.outs, proximity{tPortalName, distance})
				}
			}
		}
	}

	return p
}

func (pl portalList) getPortal(name string) portal {
	for k, v := range pl {
		if k == name {
			return v
		}
	}
	return portal{}
}

func startWalk(pl portalList) []int {

	start := pl.getPortal("AA")

	return portalWalk(start, true, pl, []string{}, 0, 0)

}

func portalWalk(p portal, in bool, pl portalList, travelled []string, distance int, recursiveDepth int) []int {
	fmt.Println(p, distance)
	if p.name == "ZZ" {
		return []int{distance}
	}

	distances := []int{}

	destList := p.ins
	if !in {
		destList = p.outs
	}

	for _, newP := range destList {
		if recursiveDepth != 0 && (newP.name == "AA" || newP.name == "ZZ") {
			continue
		}

		testPortal := pl.getPortal(newP.name)

		ins := true
		for _, testside := range testPortal.ins {
			if testside.name == p.name {
				ins = false
				break
			}
		}

		found := false

		for _, name := range travelled {
			if name == newP.name {
				found = true
				break
			}
		}

		if found {
			continue
		}
		distances = append(distances, portalWalk(pl.getPortal(newP.name), ins, pl, append(travelled, p.name), distance+newP.distance+1)...)
	}

	return distances
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	return fmt.Sprintf("%v", " -- Not Yet Implemented --")
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	return fmt.Sprintf("%v", " -- Not Yet Implemented --")
}
