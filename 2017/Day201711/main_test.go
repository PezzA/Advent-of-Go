package Day201711

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	distance, _ := travel("ne,ne,ne")
	Expect(distance).Should(Equal(3))

	distance, _ = travel("ne,ne,sw,sw")
	Expect(distance).Should(Equal(0))

	distance, _ = travel("ne,ne,s,s")
	Expect(distance).Should(Equal(2))

	distance, _ = travel("se,sw,se,sw,sw")
	Expect(distance).Should(Equal(3))
}

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
