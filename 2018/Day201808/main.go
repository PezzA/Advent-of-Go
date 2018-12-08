package Day201808

import (
	"fmt"
	"strconv"
	"strings"
)

// Entry holds wraps the data and runner interfaces for this puzzle
var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string) {
	return 2018, 8, "Memory Maneuver"
}

//------------------------

type node struct {
	id         string
	childNodes []node
	metadata   []int
}

func getNode(input []int) node {
	children := input[0]
	metadataCount := input[1]

	tail := len(input) - metadataCount
	metadata := input[tail:]

	newInput := input[2:tail]
	node := node{"aa", []node{}, metadata}

	for i := 0; i < children; i++ {
		node.childNodes = append(node.childNodes, getNode(newInput))
	}

	return node
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
	return fmt.Sprintf(" -- Not Yet Implemented --")
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	return fmt.Sprintf(" -- Not Yet Implemented --")
}
