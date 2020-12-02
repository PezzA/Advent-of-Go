package Day202002

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

func Test_ReadData(t *testing.T) {
	RegisterTestingT(t)

	data := getData(Entry.PuzzleInput())
	Expect(data[0].letter).Should(Equal("f"))
}

func Test_IsPasswordValid(t *testing.T) {
	Expect(isPasswordValid(passwordSpec{
		letter:         "a",
		minOccurrences: 1,
		maxOccurrences: 3,
		password:       "abcde"})).Should(Equal(true))

	Expect(isPasswordValid(passwordSpec{
		letter:         "b",
		minOccurrences: 1,
		maxOccurrences: 3,
		password:       "cdefg"})).Should(Equal(false))

	Expect(isPasswordValid(passwordSpec{
		letter:         "a",
		minOccurrences: 6,
		maxOccurrences: 7,
		password:       "aaaaa"})).Should(Equal(false))

	Expect(isPasswordValid(passwordSpec{
		letter:         "x",
		minOccurrences: 6,
		maxOccurrences: 7,
		password:       "xxxxxxxxxx"})).Should(Equal(false))
}

func Test_IsPositionValid(t *testing.T) {
	RegisterTestingT(t)
	Expect(isPositionValid(passwordSpec{
		letter:         "a",
		minOccurrences: 1,
		maxOccurrences: 3,
		password:       "abcde"})).Should(Equal(true))

	Expect(isPositionValid(passwordSpec{
		letter:         "b",
		minOccurrences: 1,
		maxOccurrences: 3,
		password:       "cdefg"})).Should(Equal(false))

	Expect(isPositionValid(passwordSpec{
		letter:         "c",
		minOccurrences: 2,
		maxOccurrences: 9,
		password:       "ccccccccc"})).Should(Equal(false))
}

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	valid := 0
	data := getData(Entry.PuzzleInput())

	for _, spec := range data {
		if isPasswordValid(spec) {
			valid++
		}
	}

	fmt.Println(valid)

}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

	valid := 0
	data := getData(Entry.PuzzleInput())

	for _, spec := range data {
		if isPositionValid(spec) {
			valid++
		}
	}

	fmt.Println(valid)

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
