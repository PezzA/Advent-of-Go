package Day201624

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/pezza/advent-of-code/puzzles/common"
)

const testData string = `###########
#0.1.....2#
#.#######.#
#4.......3#
###########`

func Test_ReadData(t *testing.T) {
	RegisterTestingT(t)

	_, locations := getMaze(testData, true)

	Expect(locations[0]).To(Equal(common.Point{X: 1, Y: 1}))
	Expect(locations[1]).To(Equal(common.Point{X: 3, Y: 1}))
	Expect(locations[2]).To(Equal(common.Point{X: 9, Y: 1}))
	Expect(locations[3]).To(Equal(common.Point{X: 9, Y: 3}))
	Expect(locations[4]).To(Equal(common.Point{X: 1, Y: 3}))
}

func Test_Walk(t *testing.T) {
	RegisterTestingT(t)

	maze, locations := getMaze(testData, true)

	visited := make(map[common.Point]int)

	walk(0, locations[0], locations[4], maze, visited)

	fmt.Println(visited[locations[4]])
	//Expect(walk(0, locations[0], locations[1], maze, []common.Point{})).To(Equal(2))
	//Expect(walk(0, locations[4], locations[1], maze, []common.Point{})).To(Equal(4))
	//Expect(walk(0, locations[0], locations[3], maze, []common.Point{})).To(Equal(10))
}

func Test_GetDistances(t *testing.T) {
	RegisterTestingT(t)

	maze, locations := getMaze(testData, true)
	distanceMap := getDistances(locations, maze)

	Expect(distanceMap[common.NewPoint(0, 4)]).To(Equal(2))
	Expect(distanceMap[common.NewPoint(0, 1)]).To(Equal(2))
	Expect(distanceMap[common.NewPoint(4, 2)]).To(Equal(10))
}

func Test_RouteLength(t *testing.T) {
	RegisterTestingT(t)

	maze, locations := getMaze(testData, true)
	distanceMap := getDistances(locations, maze)

	perm := []int64{4, 1, 2, 3}

	Expect(getRouteLength(perm, distanceMap, false)).To(Equal(14))

}

func Test_GetPerms(t *testing.T) {
	RegisterTestingT(t)

	maze, locations := getMaze(Entry.PuzzleInput(), true)

	distanceMap := getDistances(locations, maze)
	fmt.Println("Got distances")

	vals := make([]int64, 0)

	for k := range locations {
		if k == 0 {
			continue
		}
		vals = append(vals, int64(k))
	}
	perms := common.GetPermuations(vals)

	fmt.Println("Got", len(perms), "perms")

	minVal := -1
	for _, perm := range perms {
		val := getRouteLength(perm, distanceMap, false)

		if minVal == -1 || val < minVal {
			minVal = val
		}

	}
	// 2738 Too High
	fmt.Println(minVal)
}

func Test_PartOne(t *testing.T) {
	// get the maze and all the points.  Figure out the distance between
	// each point to each point

	// work out the shortest path to visit all of them
	RegisterTestingT(t)
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
