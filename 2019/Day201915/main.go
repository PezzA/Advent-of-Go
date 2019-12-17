package Day201915

import (
	"fmt"

	"github.com/pezza/advent-of-code/common"

	"github.com/pezza/advent-of-code/2019/intcode"
)

func botController(input string, in, out chan int64) {
	vm := intcode.New(input)
	vm.RunProgram(nil, nil, in, out, nil)
	close(out)
}

func runProgram(input string) {
	in := make(chan int64, 0)
	out := make(chan int64, 0)

	go botController(input, in, out)

	pos := common.Point{X: 0, Y: 0}
	maze := make(map[common.Point]int, 0)

	maze[pos] = 0
	crawl(pos, maze, in, out)

	for k, v := range maze {
		if v == 2 {
			fmt.Println(k.X, k.Y)
		}
	}

	counts := make(map[common.Point]int, 0)

	pathMaze(maze, common.Point{-17, 12}, counts, 0)

	min := 0
	for _, v := range counts {
		if v > min {
			min = v
		}

	}

	fmt.Println(min)
	// not 391, 394, 395
}

func pathMaze(maze map[common.Point]int, pos common.Point, counts map[common.Point]int, depth int) {
	counts[pos] = depth
	for _, dir := range crawlOrder {
		newpos := pos.Add(dir)
		if maze[newpos] != 1 {
			if val, ok := counts[newpos]; !ok || depth < val {
				pathMaze(maze, newpos, counts, depth+1)
			}
		}
	}
}

var crawlOrder = []common.Point{common.North, common.South, common.West, common.East}

func crawl(pos common.Point, maze map[common.Point]int, in, out chan int64) {
	var crawls []common.Point
	var movs []int
	var revs []int

	for index, dir := range crawlOrder {
		mov, rev := 0, 0
		switch index {
		case 0:
			mov, rev = 1, 2
		case 1:
			mov, rev = 2, 1
		case 2:
			mov, rev = 3, 4
		case 3:
			mov, rev = 4, 3
		}

		newPos := pos.Add(dir)
		if _, ok := maze[newPos]; ok {
			continue
		}

		in <- int64(mov)

		switch <-out {
		case 0:
			maze[newPos] = 1
		case 1:
			maze[pos] = 0
			crawls = append(crawls, newPos)
			movs = append(movs, mov)
			revs = append(revs, rev)
			// move back
			in <- int64(rev)
			<-out
		case 2:
			maze[pos] = 2
			crawls = append(crawls, newPos)
			movs = append(movs, mov)
			revs = append(revs, rev)
			// move back
			in <- int64(rev)
			<-out
		}

	}

	if _, ok := maze[pos]; !ok {
		maze[pos] = 0
	}

	for index := range crawls {
		in <- int64(movs[index])
		<-out
		crawl(crawls[index], maze, in, out)
		in <- int64(revs[index])
		<-out
	}

}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	return fmt.Sprintf("%v", " -- Not Yet Implemented --")
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	return fmt.Sprintf("%v", " -- Not Yet Implemented --")
}
