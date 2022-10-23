package Day201620

import (
	"fmt"
	"sort"
	"testing"

	. "github.com/onsi/gomega"
)

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	data := loadData(Entry.PuzzleInput())

	Expect(data[0].min).Should(Equal(1873924193))
	Expect(data[0].max).Should(Equal(1879728099))
	Expect(data[1074].min).Should(Equal(2076891858))
	Expect(data[1074].max).Should(Equal(2081262092))

	sort.Slice(data, func(i, j int) bool {
		return data[i].min < data[j].min
	})

	fmt.Println("BEFORE", len(data))
	for i := range data {
		fmt.Println(data[i])
	}

	compiledRanges := make(map[iprange]bool, 0)

	for i := range data {
		updateBoundaries(data[i], compiledRanges)
	}

	final := make([]iprange, 0)
	for iprange := range compiledRanges {
		final = append(final, iprange)
	}

	sort.Slice(final, func(i, j int) bool {
		return final[i].min < final[j].min
	})

	fmt.Println("after", len(final))
	count := 0
	fmt.Println(final[0])
	for i := 1; i < len(final); i++ {

		count = count + ((final[i].min - final[i-1].max) - 1)
		fmt.Println(final[i], ((final[i].min - final[i-1].max) - 1), count)
	}

}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)
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
