package Day201808

import (
	"fmt"
	"strconv"
	"strings"
)

// Entry holds wraps the data and runner interfaces for this puzzle
var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string, int) {
	return 2018, 8, "Memory Maneuver", 0
}

//------------------------

type node struct {
	id         string
	childNodes []node
	metadata   []int
}

func getNode(input []int) (node, int) {
	childCount := input[0]
	metadataCount := input[1]
	node := node{"aa", make([]node, childCount), nil}

	pos := 2
	for i := 0; i < childCount; i++ {
		var newPos int
		node.childNodes[i], newPos = getNode(input[pos:])
		pos += newPos
	}

	endNode := pos + metadataCount

	node.metadata = input[pos:endNode]
	return node, endNode
}

func countMetaData(n node) int {
	metadataCount := 0

	for _, md := range n.metadata {
		metadataCount += md
	}

	for _, cn := range n.childNodes {
		metadataCount += countMetaData(cn)
	}

	return metadataCount
}

func advancedCount(n node) int {
	value := 0
	if len(n.childNodes) == 0 {
		for _, md := range n.metadata {
			value += md
		}
		return value
	}

	for _, md := range n.metadata {
		if md == 0 || md > len(n.childNodes) {
			continue
		}
		value += advancedCount(n.childNodes[md-1])
	}
	return value
}

func getData(input string) []int {
	vals := make([]int, 0)
	for _, field := range strings.Fields(input) {
		val, _ := strconv.Atoi(field)
		vals = append(vals, val)
	}
	return vals
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	node, _ := getNode(getData(Entry.PuzzleInput()))
	return fmt.Sprintf("%v", countMetaData(node))
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	node, _ := getNode(getData(Entry.PuzzleInput()))
	return fmt.Sprintf("%v", advancedCount(node))
}
