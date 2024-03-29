package Day201617

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"strings"
)

// Entry holds wraps the data and runner interfaces for this puzzle
var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string, int) {
	return 2016, 17, "Two Steps Forward", 2
}

type mazeNode struct {
	seed  string
	hash  string
	x     int
	y     int
	up    *mazeNode
	down  *mazeNode
	left  *mazeNode
	right *mazeNode
}

type mazeTree struct {
	root mazeNode
}

func traverse(node *mazeNode, matches chan string) *mazeNode {
	hasher := md5.New()

	hasher.Reset()
	hasher.Write([]byte(node.seed))

	node.hash = hex.EncodeToString(hasher.Sum(nil))

	if node.x == 4 && node.y == 4 {
		matches <- node.seed
		return node
	}

	// try up
	if isOpen(node.hash[0]) && node.y > 1 {
		upNode := &mazeNode{
			seed: node.seed + "U",
			x:    node.x,
			y:    node.y - 1,
		}
		node.up = traverse(upNode, matches)
	}

	// try down
	if isOpen(node.hash[1]) && node.y < 4 {
		downNode := &mazeNode{
			seed: node.seed + "D",
			x:    node.x,
			y:    node.y + 1,
		}
		node.down = traverse(downNode, matches)
	}

	if isOpen(node.hash[2]) && node.x > 1 {
		leftNode := &mazeNode{
			seed: node.seed + "L",
			x:    node.x - 1,
			y:    node.y,
		}
		node.left = traverse(leftNode, matches)
	}

	if isOpen(node.hash[3]) && node.x < 4 {
		rightNode := &mazeNode{
			seed: node.seed + "R",
			x:    node.x + 1,
			y:    node.y,
		}
		node.right = traverse(rightNode, matches)
	}

	return node
}

func isOpen(input byte) bool {
	return input == 'b' || input == 'c' || input == 'd' || input == 'e' || input == 'f'
}

func doTraverse(node *mazeNode, matches chan string) {
	traverse(node, matches)
	close(matches)
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	rootNode := &mazeNode{
		seed: inputData,
		x:    1,
		y:    1,
	}

	resultChan := make(chan string, 0)
	go doTraverse(rootNode, resultChan)

	var shortest = ""
	var longest = ""

	for result := range resultChan {
		if shortest == "" || len(result) < len(shortest) {
			shortest = result
		}
		if longest == "" || len(result) > len(longest) {
			longest = result
		}
	}
	return strings.Replace(shortest, inputData, "", 1)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	rootNode := &mazeNode{
		seed: inputData,
		x:    1,
		y:    1,
	}

	resultChan := make(chan string, 0)
	go doTraverse(rootNode, resultChan)

	var shortest = ""
	var longest = ""
	for result := range resultChan {
		if shortest == "" || len(result) < len(shortest) {
			shortest = result
		}
		if longest == "" || len(result) > len(longest) {
			longest = result
		}
	}
	return strconv.Itoa(len(longest) - len(inputData))
}

func (td dayEntry) PuzzleInput() string {
	return "ioramepc"
}
