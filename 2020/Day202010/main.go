package Day202010

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func getData(input string) []int {
	lines := strings.Split(input, "\n")
	data := make([]int, len(lines))

	for i := range lines {
		val, _ := strconv.Atoi(lines[i])
		data[i] = val
	}

	return data
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
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

	return fmt.Sprintf("%v", diffs[1]*diffs[3])
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	data := getData(Entry.PuzzleInput())

	data = append(data, 0)
	sort.Ints(data)

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
	return fmt.Sprintf("%v", tot)
}
