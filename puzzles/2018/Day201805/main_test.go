package Day201805

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	testPolymer := "dabAcCaCBAcCcaDA"
	final := false
	testPolymer, _ = reactPolymer(testPolymer)
	Expect(testPolymer).Should(Equal("dabAaCBAcCcaDA"))

	testPolymer, _ = reactPolymer(testPolymer)
	Expect(testPolymer).Should(Equal("dabCBAcCcaDA"))

	testPolymer, _ = reactPolymer(testPolymer)
	Expect(testPolymer).Should(Equal("dabCBAcaDA"))

	testPolymer, final = reactPolymer(testPolymer)
	Expect(testPolymer).Should(Equal("dabCBAcaDA"))
	Expect(final).Should(Equal(false))

}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

	Expect(removeUnit(97, "dabAcCaCBAcCcaDA")).Should(Equal("dbcCCBcCcD"))
	Expect(fullyReactPolymer(removeUnit(97, "dabAcCaCBAcCcaDA"))).Should(Equal("dbCBcD"))
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
