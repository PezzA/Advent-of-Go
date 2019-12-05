package Day201904

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

func correctTests(codes []int) []int {

	for index := range codes {
		if codes[index] = 
	}

	return codes
}

func Test_ReadData(t *testing.T) {
	RegisterTestingT(t)
	min, max := getData(Entry.PuzzleInput())

	Expect(min).Should(Equal(271973))
	Expect(max).Should(Equal(785961))
}

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	Expect(isValidPassword(111111)).Should(Equal(true))
	Expect(isValidPassword(223450)).Should(Equal(false))
	Expect(isValidPassword(123789)).Should(Equal(false))
	Expect(isValidPassword(124489)).Should(Equal(true))

	min, max := getData(Entry.PuzzleInput())

	count := 0
	for i := min; i <= max; i++ {
		if isValidPassword(i) {
			count++
		}
	}

	fmt.Println(count)
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

	Expect(getChunks(111111)).Should(Equal([]string{"111111"}))
	Expect(getChunks(111122)).Should(Equal([]string{"1111", "22"}))
	Expect(getChunks(333444)).Should(Equal([]string{"333", "444"}))
	Expect(getChunks(112233)).Should(Equal([]string{"11", "22", "33"}))
	Expect(getChunks(123456)).Should(Equal([]string{"1", "2", "3", "4", "5", "6"}))
	Expect(getChunks(123444)).Should(Equal([]string{"1", "2", "3", "444"}))
	Expect(isValidPassword2(112233)).Should(Equal(true))
	Expect(isValidPassword2(123444)).Should(Equal(false))
	Expect(isValidPassword2(111122)).Should(Equal(true))

	min, max := getData(Entry.PuzzleInput())

	count := 0
	for i := min; i <= max; i++ {
		if isValidPassword2(i) {
			count++
		}
	}

	fmt.Println(count)
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
