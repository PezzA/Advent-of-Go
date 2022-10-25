package Day201707

import (
	"testing"

	. "github.com/onsi/gomega"
)

func getTestData() string {
	return `pbga (66)
xhth (57)
ebii (61)
havc (66)
ktlj (57)
fwft (72) -> ktlj, cntj, xhth
qoyq (66)
padx (45) -> pbga, havc, qoyq
tknk (41) -> ugml, padx, fwft
jptl (61)
ugml (68) -> gyxo, ebii, jptl
gyxo (61)
cntj (57)`
}

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	Expect(parseNode("flstvi (25) -> hqtwr, uvybtgx, rfguypo")).Should(Equal(node{"flstvi", 25, 0, "hqtwr, uvybtgx, rfguypo", nil}))

	Expect(parseNode("cmgugdz (67)")).Should(Equal(node{"cmgugdz", 67, 0, "", nil}))
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

	nodeList := getNodes(getTestData())

	rootNode := getRootNode(nodeList)

	printNode(buildTree(rootNode, nodeList), 0)

}

func Benchmark_PartOne(b *testing.B) {
	data := Entry.PuzzleInput()
	for n := 0; n < b.N; n++ {
		Entry.PartOne(data, nil)
	}
}

func Benchmark_PartTwo(b *testing.B) {
	data := Entry.PuzzleInput()
	for n := 0; n < b.N; n++ {
		Entry.PartTwo(data, nil)
	}
}
