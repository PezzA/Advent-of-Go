package Day201912

import (
	"fmt"
	"strings"

	"github.com/pezza/advent-of-code/common"
)

type planet struct {
	position vector3d
	velocity vector3d
}

type vector3d struct {
	x int
	y int
	z int
}

type planetList []planet

func (p planet) getEnergy() int {
	return (common.Abs(p.position.x) + common.Abs(p.position.y) + common.Abs(p.position.z)) * (common.Abs(p.velocity.x) + common.Abs(p.velocity.y) + common.Abs(p.velocity.z))
}

func (pl planetList) applyGravity(pairs [][]int) {
	for _, pair := range pairs {
		if pl[pair[0]].position.x > pl[pair[1]].position.x {
			pl[pair[0]].velocity.x--
			pl[pair[1]].velocity.x++
		} else if pl[pair[0]].position.x < pl[pair[1]].position.x {
			pl[pair[0]].velocity.x++
			pl[pair[1]].velocity.x--
		}
		if pl[pair[0]].position.y > pl[pair[1]].position.y {
			pl[pair[0]].velocity.y--
			pl[pair[1]].velocity.y++
		} else if pl[pair[0]].position.y < pl[pair[1]].position.y {
			pl[pair[0]].velocity.y++
			pl[pair[1]].velocity.y--
		}
		if pl[pair[0]].position.z > pl[pair[1]].position.z {
			pl[pair[0]].velocity.z--
			pl[pair[1]].velocity.z++
		} else if pl[pair[0]].position.z < pl[pair[1]].position.z {
			pl[pair[0]].velocity.z++
			pl[pair[1]].velocity.z--
		}
	}
}

func (pl planetList) applyVelocity() {
	for index := range pl {
		pl[index].position.x += pl[index].velocity.x
		pl[index].position.y += pl[index].velocity.y
		pl[index].position.z += pl[index].velocity.z
	}
}

func (pl planetList) totalEnergy() int {
	energy := 0
	for index := range pl {
		energy += pl[index].getEnergy()
	}
	return energy
}

func containsPair(check []int, list [][]int) bool {
	for _, item := range list {
		if (check[0] == item[0] && check[1] == item[1]) || (check[1] == item[0] && check[0] == item[1]) {
			return true
		}
	}
	return false
}

func getPairs(size int) [][]int {
	pairs := [][]int{}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if !containsPair([]int{i, j}, pairs) {
				pairs = append(pairs, []int{i, j})
			}
		}
	}

	return pairs
}

func getData(input string) planetList {
	lines := strings.Split(input, "\n")

	bodies := planetList{}

	for _, line := range lines {
		var x, y, z int
		fmt.Sscanf(line, "<x=%d, y=%d, z=%d>", &x, &y, &z)
		bodies = append(bodies, planet{vector3d{x, y, z}, vector3d{0, 0, 0}})
	}
	return bodies
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	bodies := getData(inputData)

	pairs := getPairs(4)
	for i := 0; i < 1000; i++ {
		bodies.applyGravity(pairs)
		bodies.applyVelocity()
	}

	return fmt.Sprintf("%v", bodies.totalEnergy())
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	return fmt.Sprintf("%v", " -- Not Yet Implemented --")
}
