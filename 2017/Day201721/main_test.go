package Day201721

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	initGrid := startGrid()

	Expect(flipX(initGrid)).Should(Equal([]string{".#.", "#..", "###"}))
	Expect(flipY(initGrid)).Should(Equal([]string{"###", "..#", ".#."}))

	grid2 := rotate90(initGrid)
	Expect(grid2).Should(Equal([]string{"#..", "#.#", "##."}))

	Expect(getCell(3, initGrid, 0, 0)).Should(Equal(initGrid))

	bigGrid := []string{"#..#", "....", "....", "#..#"}

	Expect(getCell(2, bigGrid, 0, 0)).Should(Equal([]string{"#.", ".."}))
	Expect(getCell(2, bigGrid, 2, 0)).Should(Equal([]string{".#", ".."}))
	Expect(getCell(2, bigGrid, 0, 2)).Should(Equal([]string{"..", "#."}))
	Expect(getCell(2, bigGrid, 2, 2)).Should(Equal([]string{"..", ".#"}))

	enhance1 := enchanceRule{
		[]string{"..", ".#"},
		[]string{"##.", "#..", "..."},
	}

	enhance2 := enchanceRule{
		[]string{".#.", "..#", "###"},
		[]string{"#..#", "....", "....", "#..#"},
	}

	enhancementRules := []enchanceRule{enhance1, enhance2}

	//	printGrid(initGrid)
	newGrid := loopCells(initGrid, enhancementRules)
	//printGrid(newGrid)
	newGrid = loopCells(newGrid, enhancementRules)
	//printGrid(newGrid)

	enhancementRules = getMatchBook(Entry.PuzzleInput())

	//fmt.Println(getEnhancementRule(initGrid, enhancementRules))

	printGrid(initGrid)
	testGrid := loopCells(initGrid, enhancementRules)
	printGrid(testGrid)
	testGrid = loopCells(testGrid, enhancementRules)
	printGrid(testGrid)
	testGrid = loopCells(testGrid, enhancementRules)
	printGrid(testGrid)
	testGrid = loopCells(testGrid, enhancementRules)
	printGrid(testGrid)
	testGrid = loopCells(testGrid, enhancementRules)
	printGrid(testGrid)

}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

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
