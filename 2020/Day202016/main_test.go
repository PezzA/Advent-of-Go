package Day202016

import (
	"fmt"
	"strings"
	"testing"

	. "github.com/onsi/gomega"
)

func Test_ReadData(t *testing.T) {
	RegisterTestingT(t)

	fields, mine, others := getData(Entry.PuzzleInput())

	fmt.Println(fields, mine, others)
}

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	fields, _, others := getData(Entry.PuzzleInput())

	tot := 0

	for _, tic := range others {
		_, invalids := tic.invalidValues(fields)

		for i := range invalids {
			tot += invalids[i]
		}
	}

	fmt.Println(tot)
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

	fields, myTicket, others := getData(Entry.PuzzleInput())

	validOthers := make([]ticket, 0)

	for _, other := range others {
		invalid, _ := other.invalidValues(fields)
		if !invalid {
			validOthers = append(validOthers, other)
		}
	}

	foundMatches := make([]match, 0)

	for {
		for _, rule := range fields {

			if isFoundrule(rule.name, foundMatches) {
				continue
			}
			matches := rule.impliedField(validOthers, foundMatches)

			if len(matches) == 1 {
				// hit!
				foundMatches = append(foundMatches, matches[0])
			}
		}

		if len(foundMatches) == len(fields) {
			// they are all allocated
			break
		}
	}

	tot := 1
	for _, fm := range foundMatches {
		if strings.HasPrefix(fm.field, "departure ") {
			tot *= myTicket[fm.col]
		}
	}
	fmt.Println(tot)

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
