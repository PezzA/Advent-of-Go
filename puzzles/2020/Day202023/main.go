package Day202023

import (
	"container/list"
	"fmt"
	"log"
	"math"
	"strconv"
	"time"

	"github.com/pezza/advent-of-code/puzzles/common"
)

type circle []int

func getData(input string) circle {
	cups := make(circle, len(input))

	for i := range input {
		val, _ := strconv.Atoi(string(input[i]))

		cups[i] = val
	}

	return cups
}

func incIndex(index int, length int) int {
	if index == length-1 {
		return 0
	}
	return index + 1
}

func getMinMax(l *list.List) (int, int) {

	startTime := time.Now()
	min := math.MaxInt32
	max := math.MinInt32
	for e := l.Front(); e != nil; e = e.Next() {
		cmp := e.Value.(int)

		if cmp > max {
			max = cmp
		}

		if cmp < min {
			min = cmp
		}

	}

	log.Println(time.Since(startTime).Milliseconds())
	return min, max
}

// getDerivedMax gets the highest value that the circle could have given
// what is in the pickup list.
func getDerivedMax(pickups []int, maxVal int) int {
	hasMax, _ := common.ContainsInt(pickups, maxVal)

	if !hasMax {
		return maxVal
	}

	hasMax1, _ := common.ContainsInt(pickups, maxVal-1)

	if !hasMax1 {
		return maxVal - 1
	}

	hasMax2, _ := common.ContainsInt(pickups, maxVal-2)

	if !hasMax2 {
		return maxVal - 2

	}

	return maxVal - 3

}

// getDerivedMin gets the lowest value that the circle could have given
// what is in the pickup list.
func getDerivedMin(pickups []int) int {
	hasOne, _ := common.ContainsInt(pickups, 1)

	if !hasOne {
		return 1
	}

	hasTwo, _ := common.ContainsInt(pickups, 2)

	if !hasTwo {
		return 2
	}

	hasThree, _ := common.ContainsInt(pickups, 3)

	if !hasThree {
		return 3

	}

	return 4
}

func listRunRound(circle *list.List, targetCup *list.Element, locs map[int]*list.Element) *list.Element {
	circleLen := circle.Len()

	pickUps := make([]int, 3)
	newElement := targetCup
	for i := 0; i < 3; i++ {
		newElement = newElement.Next()

		if newElement == nil {
			newElement = circle.Front()
		}

		pickUps[i] = newElement.Value.(int)
	}

	nextCup := targetCup.Value.(int) - 1

	min, max := getDerivedMin(pickUps), getDerivedMax(pickUps, circleLen)

	for {
		isInPickups, _ := common.ContainsInt(pickUps, nextCup)

		if nextCup < min {
			nextCup = max
			break
		}

		if !isInPickups {
			break
		}

		nextCup--
	}

	circle.MoveAfter(locs[pickUps[2]], locs[nextCup])
	circle.MoveAfter(locs[pickUps[1]], locs[nextCup])
	circle.MoveAfter(locs[pickUps[0]], locs[nextCup])

	if targetCup.Next() != nil {
		return targetCup.Next()

	} else {
		return circle.Front()
	}
}

func (c circle) getDestinationIndex(target int) int {
	for {
		found, index := common.ContainsInt(c, target)
		if found {
			return index
		}

		target--

		smallest, _ := common.Min(c)
		if target < smallest {
			break
		}
	}

	_, i := common.Max(c)

	return i
}

func (c circle) print(targetCup int) string {
	val := ""
	for i := range c {
		if c[i] == targetCup {
			val += fmt.Sprintf("(%d)", c[i])
		} else {
			val += fmt.Sprintf(" %d ", c[i])
		}
	}

	return val
}

func (c circle) runRound(targetCup int) (circle, int) {
	_, currIndex := common.ContainsInt(c, targetCup)

	pickUps := make([]int, 0)
	currIndex = incIndex(currIndex, len(c))
	pickUps = append(pickUps, c[currIndex])
	currIndex = incIndex(currIndex, len(c))
	pickUps = append(pickUps, c[currIndex])
	currIndex = incIndex(currIndex, len(c))
	pickUps = append(pickUps, c[currIndex])

	remainingCircle := make(circle, 0)
	for i := range c {
		found, _ := common.ContainsInt(pickUps, c[i])
		if !found {
			remainingCircle = append(remainingCircle, c[i])
		}
	}

	newIndex := remainingCircle.getDestinationIndex(targetCup - 1)

	newCircle := make(circle, 0)
	newCircle = append(newCircle, remainingCircle[:newIndex+1]...)
	newCircle = append(newCircle, pickUps...)
	newCircle = append(newCircle, remainingCircle[newIndex+1:]...)

	_, newCupIndex := common.ContainsInt(newCircle, targetCup)
	newCupIndex = incIndex(newCupIndex, len(c))
	return newCircle, newCircle[newCupIndex]
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	circ := getData(inputData)

	targetCup := circ[0]

	for count := 0; count <= 99; count++ {
		circ, targetCup = circ.runRound(targetCup)
	}

	_, index := common.ContainsInt(circ, 1)

	val := ""
	for i := 1; i <= len(circ)-1; i++ {
		index = incIndex(index, len(circ))
		val += strconv.Itoa(circ[index])
	}

	return fmt.Sprintf("%v", val)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	circData := getData(inputData)

	circle := list.New()

	for i := range circData {
		circle.PushBack(circData[i])
	}

	// fill in the rest of the circle
	for i := 10; i <= 1000000; i++ {
		circle.PushBack(i)
	}

	// create a lookup list for each element
	locs := make(map[int]*list.Element, 0)
	for e := circle.Front(); e != nil; e = e.Next() {
		locs[e.Value.(int)] = e
	}

	next := circle.Front()
	for i := 0; i < 10000000; i++ {
		next = listRunRound(circle, next, locs)
	}

	// get cup one and the next 2 cups
	cupOne := locs[1]
	nextP1 := cupOne.Next().Value.(int)
	nextP2 := cupOne.Next().Next().Value.(int)

	return fmt.Sprintf("%v", nextP1*nextP2)
}
