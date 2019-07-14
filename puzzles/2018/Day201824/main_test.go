package Day201824

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_GetData(t *testing.T) {
	RegisterTestingT(t)

	immunes, infections := getData(Entry.PuzzleInput())

	Expect(immunes[0]).Should(Equal(group{1, 4445, 10125, 20, 16, "cold", []string{}, []string{"radiation"}}))
	Expect(immunes[2]).Should(Equal(group{3, 1767, 5757, 27, 4, "radiation", []string{"fire", "radiation"}, []string{}}))
	Expect(immunes[9]).Should(Equal(group{10, 3091, 6684, 17, 19, "cold", []string{"radiation"}, []string{"slashing"}}))

	Expect(infections[0]).Should(Equal(group{1, 1929, 13168, 13, 7, "fire", []string{"bludgeoning"}, []string{}}))
	Expect(infections[2]).Should(Equal(group{3, 1767, 5757, 27, 4, "radiation", []string{"fire", "radiation"}, []string{}}))
	Expect(infections[9]).Should(Equal(group{10, 3091, 6684, 17, 19, "cold", []string{"radiation"}, []string{"slashing"}}))

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
