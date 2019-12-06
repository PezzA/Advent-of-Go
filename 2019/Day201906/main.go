package Day201906

import (
	"fmt"
	"strings"
)

type orbitPair struct {
	inner string
	outer string
}

type journeyStop struct {
	planet   string
	distance int
}

type orbitList []orbitPair

type orbitNode struct {
	planet   string
	depth    int
	jfc      []journeyStop
	children []*orbitNode
}

func (ol orbitList) buildOrbitTree() *orbitNode {
	tree := &orbitNode{"COM", 0, []journeyStop{}, []*orbitNode{}}
	ol.addTreeNodes(tree)
	return tree
}

func (ol orbitList) addTreeNodes(node *orbitNode) {
	for _, val := range ol.getOuterForInner(node.planet) {
		newNode := &orbitNode{
			val,
			node.depth + 1,
			append(node.jfc, journeyStop{node.planet, node.depth}),
			[]*orbitNode{},
		}
		node.children = append(node.children, newNode)
		ol.addTreeNodes(newNode)
	}
}

func (on *orbitNode) findNode(key string) *orbitNode {
	if on.planet == key {
		return on
	}

	for _, n := range on.children {
		newN := n.findNode(key)

		if newN != nil {
			return newN
		}
	}
	return nil
}

func (on *orbitNode) countDepths() int {
	total := 0

	for _, c := range on.children {
		total += c.countDepths()
	}

	return total + on.depth
}

func getData(input string) orbitList {
	lines := strings.Split(input, "\n")
	orbits := orbitList{}

	for _, line := range lines {
		bits := strings.Split(line, ")")
		orbits = append(orbits, orbitPair{bits[0], bits[1]})
	}

	return orbits
}

func (ol orbitList) getOuterForInner(input string) []string {
	outerWorlds := []string{}
	for _, val := range ol {
		if val.inner == input {
			outerWorlds = append(outerWorlds, val.outer)
		}
	}
	return outerWorlds
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	orbitList := getData(Entry.PuzzleInput())
	on := orbitList.buildOrbitTree()
	return fmt.Sprintf("%v", on.countDepths())
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	orbitList := getData(Entry.PuzzleInput())

	on := orbitList.buildOrbitTree()
	me, san := on.findNode("YOU"), on.findNode("SAN")

	closet := 0
	for _, meStep := range me.jfc {
		for _, sanStep := range san.jfc {
			if meStep == sanStep {

				if meStep.distance > closet {
					closet = meStep.distance
				}

			}
		}
	}
	return fmt.Sprintf("%v", (me.depth-1-closet)+(san.depth-1-closet))
}
