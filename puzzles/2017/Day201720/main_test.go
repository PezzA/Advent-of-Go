package Day201720

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	//p1 := getParticles(`p=<3,0,0>, v=<2,0,0>, a=<-1,0,0>
	//p=<4,0,0>, v=<0,0,0>, a=<-2,0,0>`)

	p1 := getParticles(Entry.PuzzleInput())

	sortByAcclerationVelocityPosition(p1)

	//fmt.Println(p1)
	Expect(p1[0].index).Should(Equal(243))
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

	particles := getParticles(Entry.PuzzleInput())

	prev, same := 0, 0

	for {
		particles = updateParticles(particles)

		particles = removeCollisions(particles)

		// remove collisions

		if prev == len(particles) {
			same++
		}
		prev = len(particles)

		if same == 20 {
			break
		}
	}

	fmt.Println(prev)
}

func Benchmark_PartOne(b *testing.B) {
	data := Entry.PuzzleInput()
	for n := 0; n < b.N; n++ {
		Entry.PartOne(data, nil)
	}
}

func Benchmark_PartTwo(b *testing.B) {
	data := Entry.PuzzleInput()
	for n := 0; n < b.N; n++ {
		Entry.PartTwo(data, nil)
	}
}
