package Day201501

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	result, err := Entry.PartOne("(())")
	Expect(err).ShouldNot(HaveOccurred())
	Expect(result).Should(Equal("0"))
}

/*
func (td testDay) GetTestData() ([]string, []string) {
	return []string{
			"(())",
			"()()",
			"(((",
			"(()(()(",
			"))(((((",
			"())",
			"))(",
			")))",
			")())())",
		}, []string{
			")",
			"()())"}
}
*/

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

}

func Benchmark_PartOne(b *testing.B) {
	data := Entry.GetData()
	for n := 0; n < b.N; n++ {
		Entry.PartOne(data)
	}
}

func Benchmark_PartTwo(b *testing.B) {
	data := Entry.GetData()
	for n := 0; n < b.N; n++ {
		Entry.PartTwo(data)
	}
}
