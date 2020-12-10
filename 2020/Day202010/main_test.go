package Day202010

import (
	"fmt"
	"sort"
	"testing"

	. "github.com/onsi/gomega"
)

func Test_ReadData(t *testing.T) {
	RegisterTestingT(t)

}

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	data := getData(Entry.PuzzleInput())
	sort.Ints(data)

	diffs := make(map[int]int)
	diffs[0] = 0
	diffs[1] = 0
	diffs[2] = 0
	diffs[3] = 0

	for i := range data {
		if i == 0 {
			diffs[data[i]]++
			continue
		}

		diff := data[i] - data[i-1]

		diffs[diff]++
	}

	diffs[3]++
	fmt.Println(diffs[1] * diffs[3])
}

var data1 = `28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3`

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

	data := getData(Entry.PuzzleInput())

	data = append(data, 0)
	sort.Ints(data)

	fmt.Println(data)

	clusters := make([][]int, 0)
	currCluster := make([]int, 0)

	for i := 0; i < len(data); i++ {
		if i == 0 {
			currCluster = append(currCluster, data[i])
			continue
		}

		if data[i]-data[i-1] == 1 {
			currCluster = append(currCluster, data[i])
		} else {
			clusters = append(clusters, currCluster)
			currCluster = []int{data[i]}
		}
	}
	clusters = append(clusters, currCluster)

	tot := 1
	for _, c := range clusters {

		if len(c) == 3 {
			tot *= 2
		}

		if len(c) == 4 {
			tot *= 4
		}

		if len(c) == 5 {
			tot *= 7
		}

	}

	fmt.Println(clusters)

	fmt.Println(tot)
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
