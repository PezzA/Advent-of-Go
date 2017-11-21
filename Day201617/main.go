package Day201617

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

// Entry holds wraps the data and runner interfaces for this puzzle
var Entry testDay

type testDay bool

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

func traverse(node *mazeNode) *mazeNode {
	hasher := md5.New()

	hasher.Reset()
	hasher.Write([]byte(node.seed))

	node.hash = hex.EncodeToString(hasher.Sum(nil))

	if node.x == 4 && node.y == 4 {
		fmt.Println(node.seed)
		return node
	}

	// try up
	if isOpen(node.hash[0]) && node.y > 1 {
		upNode := &mazeNode{
			seed: node.seed + "U",
			x:    node.x,
			y:    node.y - 1,
		}
		node.up = traverse(upNode)
	}

	// try down
	if isOpen(node.hash[1]) && node.y < 4 {
		downNode := &mazeNode{
			seed: node.seed + "D",
			x:    node.x,
			y:    node.y + 1,
		}
		node.down = traverse(downNode)
	}

	if isOpen(node.hash[2]) && node.x > 1 {
		leftNode := &mazeNode{
			seed: node.seed + "L",
			x:    node.x - 1,
			y:    node.y,
		}
		node.left = traverse(leftNode)
	}

	if isOpen(node.hash[3]) && node.x < 4 {
		rightNode := &mazeNode{
			seed: node.seed + "R",
			x:    node.x + 1,
			y:    node.y,
		}
		node.right = traverse(rightNode)
	}

	return node
}

func isOpen(input byte) bool {
	return input == 'b' || input == 'c' || input == 'd' || input == 'e' || input == 'f'
}

func (td testDay) PartOne(inputData string) (string, error) {

	rootNode := &mazeNode{
		seed: inputData,
		x:    1,
		y:    1,
	}

	fmt.Println(traverse(rootNode))

	return "", nil

}

func (td testDay) PartTwo(inputData string) (string, error) {
	return "", nil
}

func (td testDay) Day() int {
	return 201617
}

func (td testDay) Heading() string {
	return "--- (2016) Day 17: Two Steps Forward ---"
}

func (td testDay) GetTestData() ([]string, []string) {
	return []string{
			"hijkl",
			//"ihgpwlah",
			//"kglvqrro",
			//"ulqzkmiv",
		},
		[]string{""}
}

func (td testDay) GetData() string {
	return "ioramepc"
}
