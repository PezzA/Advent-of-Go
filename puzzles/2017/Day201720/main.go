package Day201720

import (
	"bufio"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/pezza/advent-of-code/puzzles/common"
)

type dayEntry bool

var Entry dayEntry

func (td dayEntry) Describe() (int, int, string, int) {
	return 2017, 20, "Particle Swarm", 0
}

type vertex struct {
	x int
	y int
	z int
}

type particle struct {
	index int
	p     vertex
	v     vertex
	a     vertex
}

var particleRegex = regexp.MustCompile(`p=<(\S*),(\S*),(\S*)>, v=<(\S*),(\S*),(\S*)>, a=<(\S*),(\S*),(\S*)>`)

func (v vertex) absTotal() int {
	return common.Abs(v.x) + common.Abs(v.y) + common.Abs(v.z)
}

func getParticles(input string) []particle {
	scanner := bufio.NewScanner(strings.NewReader(input))
	parts := make([]particle, 0)

	index := 0
	for scanner.Scan() {
		line := scanner.Text()

		matches := particleRegex.FindStringSubmatch(line)

		px, _ := strconv.Atoi(matches[1])
		py, _ := strconv.Atoi(matches[2])
		pz, _ := strconv.Atoi(matches[3])
		vx, _ := strconv.Atoi(matches[4])
		vy, _ := strconv.Atoi(matches[5])
		vz, _ := strconv.Atoi(matches[6])
		ax, _ := strconv.Atoi(matches[7])
		ay, _ := strconv.Atoi(matches[8])
		az, _ := strconv.Atoi(matches[9])

		part := particle{
			index,
			vertex{
				px,
				py,
				pz,
			},
			vertex{
				vx,
				vy,
				vz,
			},
			vertex{
				ax,
				ay,
				az,
			}}

		parts = append(parts, part)
		index++
	}
	return parts
}

func sortByAcclerationVelocityPosition(pl []particle) []particle {
	sort.Slice(pl, func(i, j int) bool {
		if pl[i].a.absTotal() == pl[j].a.absTotal() {
			if pl[i].v.absTotal() == pl[j].v.absTotal() {
				return pl[i].p.absTotal() < pl[j].p.absTotal()
			}
			return pl[i].v.absTotal() < pl[j].v.absTotal()
		}
		return pl[i].a.absTotal() < pl[j].a.absTotal()
	})

	return pl
}

func (part particle) update() particle {
	part.v.x += part.a.x
	part.v.y += part.a.y
	part.v.z += part.a.z
	part.p.x += part.v.x
	part.p.y += part.v.y
	part.p.z += part.v.z

	return part
}

func (part particle) String() string {
	return fmt.Sprintf("(%v) [%v,%v,%v]", part.index, part.a.absTotal(), part.v.absTotal(), part.p.absTotal())
}

func updateParticles(pl []particle) []particle {
	newpart := make([]particle, 0)
	for index := range pl {
		newpart = append(newpart, pl[index].update())
	}
	return newpart
}

/* aborted trying to determine if movingawayfrom 0, moving away from each other and staying in the same order
func (part particle) manhattan() int {
	return intAbs(part.p.x) + intAbs(part.p.y) + intAbs(part.p.z)
}

func getManhattenorder(list []particle) []particle {
	sort.Slice(list, func(i, j int) bool {
		return list[i].manhattan() < list[j].manhattan()
	})

	return list
}


func isSameOrder(pl1 []particle, pl2 []particle) bool {
	for i := range pl1 {
		if pl1[i].index != pl2[i].index {
			return false
		}
	}
	return true
}

func noDecreasingDistances(pl1 []particle, pl2 []particle) bool {
	for i := 0; i < len(pl1)-1; i++ {
		prevousDistance := pl1[i+1].manhattan() - pl1[i].manhattan()
		currentDistance := pl2[i+1].manhattan() - pl2[i].manhattan()

		if prevousDistance > currentDistance {
			return false
		}
	}
	return true
}


func allMovingAwayFromTheSun(pl1 []particle, pl2 []particle) bool {
	for i := 0; i < len(pl1)-1; i++ {
		if pl1[i].manhattan() >= pl2[i].manhattan() {
			return false
		}
	}
	return true
}
*/

func removeCollisions(part []particle) []particle {

	places := make(map[vertex][]int, 0)

	for _, val := range part {
		if _, ok := places[val.p]; ok {
			places[val.p] = append(places[val.p], val.index)
		} else {
			places[val.p] = []int{val.index}
		}
	}

	if len(places) < len(part) {
		for _, v := range places {
			if len(v) > 1 {
				for _, partIndex := range v {
					for i := 0; i < len(part); i++ {
						if part[i].index == partIndex {
							part = append(part[:i], part[i+1:]...)
						}
					}
				}
			}
		}
	}

	return part
}

// 0, 78, 43, 138
func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	particles := getParticles(Entry.PuzzleInput())

	sortByAcclerationVelocityPosition(particles)

	return fmt.Sprintf("%v", particles[0].index)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	particles := getParticles(Entry.PuzzleInput())

	prev, same := 0, 0

	for {
		particles = removeCollisions(updateParticles(particles))

		if prev == len(particles) {
			same++
		}

		prev = len(particles)

		if same == 20 {
			break
		}
	}

	return fmt.Sprintf("%v", prev)
}
