package Day202206

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_ReadData(t *testing.T) {
	RegisterTestingT(t)

}

func Test_PartScan(t *testing.T) {
	RegisterTestingT(t)

	Expect(scanForMarker("mjqjpqmgbljsphdztnvjfqwrcgsmlb", 4)).Should(Equal(7))
	Expect(scanForMarker("bvwbjplbgvbhsrlpgdmjqwftvncz", 4)).Should(Equal(5))
	Expect(scanForMarker("nppdvjthqldpwncqszvftbrmjlhg", 4)).Should(Equal(6))
	Expect(scanForMarker("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 4)).Should(Equal(10))
	Expect(scanForMarker("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 4)).Should(Equal(11))
}

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	Expect(Entry.PartOne(Entry.PuzzleInput(), nil)).Should(Equal("1850"))
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

	Expect(Entry.PartTwo(Entry.PuzzleInput(), nil)).Should(Equal("2823"))
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
