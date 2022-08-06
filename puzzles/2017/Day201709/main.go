package Day201709

import (
	"fmt"
	"strconv"
	"strings"
)

// Entry holds wraps the data and runner interfaces for this puzzle
var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string, int) {
	return 2017, 9, "Stream Processing", 2
}

type node struct {
	score int
	nodes []node
}

func isSpecialChar(char rune) bool {
	return char == '{' || char == '}' || char == '!' || char == '<' || char == '>'
}

func streamReader(input string, score int, cumlativeScore int, garbageCount int) (node, int, int, int) {
	node := node{score, nil}

	for i := 0; i < len(input); i++ {
		currentRune := rune(input[i])
		if isSpecialChar(currentRune) {
			switch rune(currentRune) {
			case '{':
				childNode, subIndex, newCumlatiuve, newgarbageCount := streamReader(input[i+1:], score+1, cumlativeScore, garbageCount)
				cumlativeScore = newCumlatiuve
				garbageCount = newgarbageCount
				node.nodes = append(node.nodes, childNode)
				i += subIndex
			case '}':
				return node, i + 1, cumlativeScore + node.score, garbageCount
			case '!':
				i++
			case '<':
				terminate := false
				for !terminate {
					i++
					subRune := rune(input[i])
					switch subRune {
					case '!':
						i++
					case '>':
						terminate = true
					default:
						garbageCount++
					}
				}
			}
		}
	}

	return node, 0, cumlativeScore, garbageCount
}

func printTree(n node, depth int) {
	fmt.Println(strings.Repeat(" ", depth*4), "G -", n.score)
	for _, childNode := range n.nodes {
		printTree(childNode, depth+1)
	}
}

func scoreTree(n node, score int) int {
	for _, childNode := range n.nodes {
		score += scoreTree(childNode, score)
	}
	return score + n.score
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	_, _, cumlative, _ := streamReader(inputData, 0, 0, 0)
	return strconv.Itoa(cumlative)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	_, _, _, garbageCount := streamReader(inputData, 0, 0, 0)
	return strconv.Itoa(garbageCount)
}
