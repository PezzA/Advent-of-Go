package Day201524

import (
	"fmt"
	"sort"
	"testing"

	. "github.com/onsi/gomega"
)

func getTestPresents() []int {
	return []int{1, 2, 3, 4, 5, 7, 8, 9, 10, 11}
}

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	presents := getTestPresents()

	weight := totalweight(presents)

	combos := combinations(weight/3, presents)

	sort.Slice(combos, func(i, j int) bool {

		if len(combos[i]) == len(combos[j]) {
			return getQuantumEntanglement(combos[i]) < getQuantumEntanglement(combos[j])
		}
		return len(combos[i]) < len(combos[j])
	})

	fmt.Println(fmt.Sprintf("%v", getQuantumEntanglement(combos[0])))
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)
}

func Benchmark_BenchPartOne(b *testing.B) {
	data := Entry.PuzzleInput()
	for n := 0; n < b.N; n++ {
		Entry.PartOne(data, nil)
	}
}

func Benchmark_BenchPartTwo(b *testing.B) {
	data := Entry.PuzzleInput()
	for n := 0; n < b.N; n++ {
		Entry.PartTwo(data, nil)
	}
}
