package Day202209

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/pezza/advent-of-code/puzzles/common"
)

type instruction struct {
	direction string
	amount    int
}

func getData(input string) []instruction {
	lines := strings.Split(input, "\n")
	retVal := make([]instruction, 0)

	for _, line := range lines {
		bits := strings.Split(line, " ")
		val, _ := strconv.Atoi(bits[1])

		retVal = append(retVal, instruction{bits[0], val})
	}
	return retVal
}

func moveTail(head, tail common.Point) common.Point {
	// update tail
	if head.X == tail.X+2 && tail.Y == head.Y {
		tail.X++
		return tail
	}
	if head.X == tail.X-2 && tail.Y == head.Y {
		tail.X--
		return tail
	}
	if head.Y == tail.Y-2 && tail.X == head.X {
		tail.Y--
		return tail
	}
	if head.Y == tail.Y+2 && tail.X == head.X {
		tail.Y++
		return tail
	}
	// ..H
	// ...
	// T..
	if head.X == tail.X+2 && head.Y == tail.Y+2 {
		tail.X++
		tail.Y++
		return tail
	}
	// H..
	// ...
	// ..T
	if head.X == tail.X-2 && head.Y == tail.Y+2 {
		tail.X--
		tail.Y++
		return tail
	}
	// ..T
	// ...
	// H..
	if head.X == tail.X-2 && head.Y == tail.Y-2 {
		tail.X--
		tail.Y--
		return tail
	}
	// T..
	// ...
	// ..H
	if head.X == tail.X+2 && head.Y == tail.Y-2 {
		tail.X++
		tail.Y--
		return tail
	}

	// ..T
	// H..
	// ...
	if head.X == tail.X-2 && head.Y == tail.Y-1 {
		tail.X--
		tail.Y--
		return tail
	}
	// ...
	// H..
	// ..T
	if head.X == tail.X-2 && head.Y == tail.Y+1 {
		tail.X--
		tail.Y++
		return tail
	}
	// ...
	// ..H
	// T..
	if head.X == tail.X+2 && head.Y == tail.Y+1 {
		tail.X++
		tail.Y++
		return tail
	}
	// T..
	// ..H
	// ...
	if head.X == tail.X+2 && head.Y == tail.Y-1 {
		tail.X++
		tail.Y--
		return tail
	}
	// .H.
	// ...
	// T..
	if head.X == tail.X+1 && head.Y == tail.Y+2 {
		tail.X++
		tail.Y++
		return tail
	}
	// .H.
	// ...
	// ..T
	if head.X == tail.X-1 && head.Y == tail.Y+2 {
		tail.X--
		tail.Y++
		return tail
	}
	// T..
	// ...
	// .H
	if head.X == tail.X+1 && head.Y == tail.Y-2 {
		tail.X++
		tail.Y--
		return tail
	}
	// ..T
	// ...
	// .H
	if head.X == tail.X-1 && head.Y == tail.Y-2 {
		tail.X--
		tail.Y--
		return tail
	}
	return tail
}

func moveHead(direction string, head common.Point) common.Point {
	switch direction {
	case "U":
		head.Y += 1
	case "D":
		head.Y -= 1
	case "L":
		head.X -= 1
	case "R":
		head.X += 1
	}
	return head
}

func processMoves(moves []instruction, print bool) map[string]int {
	head, tail := common.Point{0, 0}, common.Point{0, 0}

	if print {
		printGrid(head, tail, 6, 5)
	}
	tailLocations := make(map[string]int, 0)

	for _, move := range moves {
		for i := 0; i < move.amount; i++ {
			head = moveHead(move.direction, head)
			tail = moveTail(head, tail)

			if _, ok := tailLocations[tail.String()]; ok {
				tailLocations[tail.String()]++
			} else {
				tailLocations[tail.String()] = 1
			}
			if print {
				fmt.Println(move, i)
				printGrid(head, tail, 6, 5)
			}
		}
	}
	return tailLocations
}

func processLongTail(moves []instruction, length int, print bool) map[string]int {

	segments := make([]common.Point, length)

	tailLocations := make(map[string]int, 0)
	for index := range segments {
		segments[index] = common.Point{0, 0}
	}

	for _, move := range moves {
		for i := 0; i < move.amount; i++ {
			for s := 0; s < length; s++ {
				if s == 0 {
					segments[0] = moveHead(move.direction, segments[0])
					continue
				}

				segments[s] = moveTail(segments[s-1], segments[s])
			}

			if _, ok := tailLocations[segments[length-1].String()]; ok {
				tailLocations[segments[length-1].String()]++
			} else {
				tailLocations[segments[length-1].String()] = 1
			}
			if print {
				printGridSegments(segments, 26, 26)
			}
		}
	}
	return tailLocations
}

func printGridSegments(segments []common.Point, maxX int, maxY int) {

	fmt.Println(segments)
	//	start := common.Point{0, 0}
	grid := make([]string, maxY+1)

	for y := 0; y <= maxY; y++ {
		// we start off drawing the top most line and work back to 0
		transY := maxY - y
		tmpString := ""
		for x := 0; x <= maxX; x++ {

			drawn := false
			for s := 0; s < len(segments); s++ {
				if segments[s].X == x && segments[s].Y == transY {
					if s == 0 {
						tmpString += "H"
					} else {
						tmpString += fmt.Sprint(s)
					}
					drawn = true
					break
				}
			}

			if !drawn {
				if 0 == x && 0 == transY {
					tmpString += "S"
				} else {
					tmpString += "."
				}
			}
		}
		grid[y] = tmpString
	}
	for _, line := range grid {
		fmt.Println(line)
	}
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func printGrid(head, tail common.Point, maxX int, maxY int) {
	start := common.Point{0, 0}
	grid := make([]string, maxX)

	for y := 0; y <= maxY; y++ {
		// we start off drawing the top most line and work back to 0
		transY := maxY - y
		tmpString := ""
		for x := 0; x <= maxX; x++ {
			if head.X == x && head.Y == transY {
				tmpString += "H"
				continue
			}
			if tail.X == x && tail.Y == transY {
				tmpString += "T"
				continue
			}
			if start.X == x && start.Y == transY {
				tmpString += "S"
				continue
			}
			tmpString += "."
		}

		grid[y] = tmpString
	}

	for _, line := range grid {
		fmt.Println(line)
	}
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	moves := getData(inputData)
	tailLocations := processMoves(moves, false)
	return fmt.Sprint(len(tailLocations))
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	moves := getData(inputData)
	tailLocations := processLongTail(moves, 10, false)
	return fmt.Sprint(len(tailLocations))
}
