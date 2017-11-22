package Day201617

import (
	"fmt"
	"testing"
)

func TestTraverse(t *testing.T) {
	rootNode := &mazeNode{
		seed: "ulqzkmiv",
		x:    1,
		y:    1,
	}

	resultChan := make(chan string, 0)
	go doTraverse(rootNode, resultChan)

	var shortest = ""
	for result := range resultChan {
		if shortest == "" || len(result) < len(shortest) {

			shortest = result
			fmt.Println("New shortest ->", shortest)
		}
	}
}
