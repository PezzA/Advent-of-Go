package Day201707

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type dayEntry bool

var Entry dayEntry

func (td dayEntry) Describe() (int, int, string) {
	return 2017, 7, "Recursive Circus"
}

var nodeRegex = regexp.MustCompile(`([a-z]+) \(([0-9]+)\)( -> ((\w*(,| )*)*))*`)

type node struct {
	title           string
	weight          int
	cumlativeWeight int
	nodeList        string
	children        []*node
}

func parseNode(input string) node {
	match := nodeRegex.FindStringSubmatch(input)

	title := match[1]

	weight, _ := strconv.Atoi(match[2])

	nodeList := match[4]

	return node{
		title:    title,
		weight:   weight,
		nodeList: nodeList,
	}
}

func getNodes(input string) []node {
	scanner := bufio.NewScanner(strings.NewReader(input))

	nodeList := make([]node, 0)

	for scanner.Scan() {
		nodeList = append(nodeList, parseNode(scanner.Text()))
	}

	return nodeList
}

func getRootNode(nl []node) string {
	targetNode := ""

	// well, if it works, it works
	for _, node := range nl {
		match := true
		for _, checkNode := range nl {
			if strings.Contains(checkNode.nodeList, node.title) {
				match = false
			}
		}

		if match {
			targetNode = node.title
		}
	}

	return targetNode
}

func getNode(input string, nl []node) *node {
	for _, node := range nl {
		if node.title == input {
			return &node
		}
	}

	return &node{}
}

func addNodes(node *node, nl []node) *node {
	node.cumlativeWeight = node.weight
	if node.nodeList == "" {
		return node
	}

	nodeKeyList := strings.Split(node.nodeList, ",")

	for _, nodekey := range nodeKeyList {
		actKey := strings.TrimSpace(nodekey)
		leafNode := getNode(actKey, nl)
		node.children = append(node.children, addNodes(leafNode, nl))
	}

	weights := make(map[int]int, 0)

	for _, child := range node.children {

		if _, ok := weights[child.cumlativeWeight]; ok {
			weights[child.cumlativeWeight]++
		} else {
			weights[child.cumlativeWeight] = 1
		}

		node.cumlativeWeight += child.cumlativeWeight
	}

	return node
}

func printNode(node *node, depth int) {
	if len(node.children) > 0 {
		fmt.Println(strings.Repeat("-", depth*4), node.cumlativeWeight, node.weight)
	}
	for _, child := range node.children {
		printNode(child, depth+1)
	}
}

func buildTree(rootNodeKey string, nl []node) *node {
	return addNodes(getNode(rootNodeKey, nl), nl)
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	return getRootNode(getNodes(inputData))
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	nodeList := getNodes(inputData)

	rootNode := getRootNode(nodeList)

	printNode(buildTree(rootNode, nodeList), 0)
	return fmt.Sprintf(" -- Not Yet Implemented --")
}
