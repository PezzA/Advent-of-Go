package Day201617

import (
	"testing"
)

func TestTraverse(t *testing.T) {
	rootNode := &mazeNode{
		seed: "kglvqrro",
		x:    1,
		y:    1,
	}

	traverse(rootNode)
}
