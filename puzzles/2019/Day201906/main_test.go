package Day201906

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_ReadData(t *testing.T) {
	RegisterTestingT(t)

	o := getData(Entry.PuzzleInput())

	Expect(o[0]).Should(Equal(orbitPair{"KDZ", "KYY"}))
	Expect(o[5]).Should(Equal(orbitPair{"HL3", "X5V"}))
}

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	orbitList := getData(testData)
	on := orbitList.buildOrbitTree()
	Expect(on.countDepths()).Should(Equal(42))
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)
	orbitList := getData(Entry.PuzzleInput())

	on := orbitList.buildOrbitTree()

	me, san := on.findNode("YOU"), on.findNode("SAN")

	closet := 0
	for _, meStep := range me.jfc {
		for _, sanStep := range san.jfc {
			if meStep == sanStep {

				if meStep.distance > closet {
					closet = meStep.distance
				}

			}
		}
	}

	Expect((me.depth - 1 - closet) + (san.depth - 1 - closet)).Should(Equal(286))

}

var testData = `COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L`

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
