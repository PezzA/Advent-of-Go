package Day202011

import (
	"fmt"
	"strings"
)

type seatPlan [][]rune

const emptySeat = 76
const floor = 46
const occupiedSeat = 35

func getData(input string) seatPlan {
	lines := strings.Split(input, "\n")
	var grid = make(seatPlan, len(lines))

	for l, line := range lines {
		gridRow := make([]rune, len(line))

		for i, r := range line {
			gridRow[i] = r
		}

		grid[l] = gridRow
	}
	return grid
}

func (sp seatPlan) adjacentOccupiedSeats(x int, y int) int {
	tot := 0
	//N
	if y > 0 {
		if sp[y-1][x] == occupiedSeat {
			tot++
		}
	}
	//NE
	if y > 0 && x < len(sp[y])-1 {
		if sp[y-1][x+1] == occupiedSeat {
			tot++
		}
	}
	//E
	if x < len(sp[y])-1 {
		if sp[y][x+1] == occupiedSeat {
			tot++
		}
	}
	//SE
	if x < len(sp[y])-1 && y < len(sp)-1 {
		if sp[y+1][x+1] == occupiedSeat {
			tot++
		}
	}

	//S
	if y < len(sp)-1 {
		if sp[y+1][x] == occupiedSeat {
			tot++
		}
	}

	//SW
	if y < len(sp)-1 && x > 0 {
		if sp[y+1][x-1] == occupiedSeat {
			tot++
		}
	}

	//W
	if x > 0 {
		if sp[y][x-1] == occupiedSeat {
			tot++
		}
	}

	//NW
	if x > 0 && y > 0 {
		if sp[y-1][x-1] == occupiedSeat {
			tot++
		}
	}

	return tot
}
func (sp seatPlan) maxY() int {
	return len(sp)
}

func (sp seatPlan) maxX() int {
	return len(sp[0])
}

func (sp seatPlan) extendedOccupiedSeats(startX int, startY int) int {
	tot := 0
	//N
	x, y := startX, startY-1
	for {
		if y < 0 {
			break
		}
		if sp[y][x] == emptySeat {
			break
		}
		if sp[y][x] == occupiedSeat {
			tot++
			break
		}
		y--
	}

	//NE
	x, y = startX+1, startY-1
	for {
		if y < 0 || x == sp.maxX() {
			break
		}

		if sp[y][x] == emptySeat {
			break
		}
		if sp[y][x] == occupiedSeat {
			tot++
			break
		}
		y--
		x++
	}

	//E
	x, y = startX+1, startY
	for {
		if x == sp.maxX() {
			break
		}
		if sp[y][x] == emptySeat {
			break
		}
		if sp[y][x] == occupiedSeat {
			tot++
			break
		}
		x++
	}

	//SE
	x, y = startX+1, startY+1
	for {
		if x == sp.maxX() || y == sp.maxY() {
			break
		}
		if sp[y][x] == emptySeat {
			break
		}
		if sp[y][x] == occupiedSeat {
			tot++
			break
		}
		y++
		x++
	}

	//S
	x, y = startX, startY+1
	for {
		if y == sp.maxY() {
			break
		}
		if sp[y][x] == emptySeat {
			break
		}
		if sp[y][x] == occupiedSeat {
			tot++
			break
		}
		y++
	}

	//SW
	x, y = startX-1, startY+1
	for {
		if y == sp.maxY() || x < 0 {
			break
		}
		if sp[y][x] == emptySeat {
			break
		}
		if sp[y][x] == occupiedSeat {
			tot++
			break
		}
		y++
		x--
	}

	//W
	x, y = startX-1, startY
	for {
		if x < 0 {
			break
		}
		if sp[y][x] == emptySeat {
			break
		}
		if sp[y][x] == occupiedSeat {
			tot++
			break
		}
		x--
	}

	//NW
	x, y = startX-1, startY-1
	for {
		if x < 0 || y < 0 {
			break
		}
		if sp[y][x] == emptySeat {
			break
		}
		if sp[y][x] == occupiedSeat {
			tot++
			break
		}
		y--
		x--
	}

	return tot
}

func (sp seatPlan) equals(cmp seatPlan) bool {
	for y := range sp {
		for x := range sp[y] {
			if sp[y][x] != cmp[y][x] {
				return false
			}
		}
	}
	return true
}

func (sp seatPlan) countOccupiedSeats() int {
	tot := 0
	for y := range sp {
		for x := range sp[y] {
			if sp[y][x] == occupiedSeat {
				tot++
			}
		}
	}
	return tot
}

func (sp seatPlan) nextState(tolerance int, extendedCheck bool) seatPlan {
	newPlan := make(seatPlan, len(sp))
	for i := range sp {
		newPlan[i] = make([]rune, len(sp[i]))
	}

	for y, row := range sp {
		for x := range row {
			if sp[y][x] == floor {
				newPlan[y][x] = floor
				continue
			}

			occ := 0
			if extendedCheck {
				occ = sp.extendedOccupiedSeats(x, y)
			} else {
				occ = sp.adjacentOccupiedSeats(x, y)
			}

			if sp[y][x] == emptySeat && occ == 0 {
				newPlan[y][x] = occupiedSeat
				continue
			}

			if sp[y][x] == occupiedSeat && occ >= tolerance {
				newPlan[y][x] = emptySeat
				continue
			}

			newPlan[y][x] = sp[y][x]
		}
	}
	return newPlan
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	data := getData(inputData)

	occ := 0

	newState := data.nextState(4, false)

	for {
		newState2 := newState.nextState(4, false)

		if newState2.equals(newState) {
			occ = newState.countOccupiedSeats()
			break
		}

		newState = newState2
	}
	return fmt.Sprintf("%v", occ)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	data := getData(inputData)

	occ := 0

	newState := data.nextState(5, true)

	for {
		newState2 := newState.nextState(5, true)

		if newState2.equals(newState) {
			occ = newState.countOccupiedSeats()
			break
		}

		newState = newState2
	}
	return fmt.Sprintf("%v", occ)
}
