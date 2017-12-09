package Day201709

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	var tree node
	_, _, cumulative, _ := streamReader("{}", 0, 0, 0)
	Expect(cumulative).Should(Equal(1))

	_, _, cumulative, _ = streamReader("{{{}}}", 0, 0, 0)
	Expect(cumulative).Should(Equal(6))

	_, _, cumulative, _ = streamReader("{{},{}}", 0, 0, 0)
	Expect(cumulative).Should(Equal(5))

	_, _, cumulative, _ = streamReader("{{{},{},{{}}}}", 0, 0, 0)
	Expect(cumulative).Should(Equal(16))

	_, _, cumulative, _ = streamReader("{<a>,<a>,<a>,<a>}", 0, 0, 0)
	Expect(cumulative).Should(Equal(1))

	_, _, cumulative, _ = streamReader("{{<ab>},{<ab>},{<ab>},{<ab>}}", 0, 0, 0)
	Expect(cumulative).Should(Equal(9))

	_, _, cumulative, _ = streamReader("{{<!!>},{<!!>},{<!!>},{<!!>}}", 0, 0, 0)
	Expect(cumulative).Should(Equal(9))

	tree, _, cumulative, _ = streamReader("{{<a!>},{<a!>},{<a!>},{<ab>}}", 0, 0, 0)

	printTree(tree, 0)
	Expect(cumulative).Should(Equal(3))

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
