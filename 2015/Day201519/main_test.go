package Day201519

import (
	"log"
	"os"
	"regexp"
	"testing"

	. "github.com/onsi/gomega"
)

func getTestData() ([]replacement, string) {
	replacementList := make([]replacement, 0)

	replacementList = append(replacementList, replacement{"e", "H", regexp.MustCompile("e"), regexp.MustCompile("H")})
	replacementList = append(replacementList, replacement{"e", "O", regexp.MustCompile("e"), regexp.MustCompile("O")})
	replacementList = append(replacementList, replacement{"H", "HO", regexp.MustCompile("H"), regexp.MustCompile("HO")})
	replacementList = append(replacementList, replacement{"H", "OH", regexp.MustCompile("H"), regexp.MustCompile("OH")})
	replacementList = append(replacementList, replacement{"O", "HH", regexp.MustCompile("O"), regexp.MustCompile("HH")})

	return replacementList, "HOHOHO"
}

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	replacementList, molecule := getData(Entry.PuzzleInput())

	list := getAllDistinctReplacements(replacementList, molecule)

	Expect(len(list)).Should(Equal(504))
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

	f, err := os.OpenFile("day19.log", os.O_RDWR|os.O_CREATE|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)

	replacementList, molecule := getData(Entry.PuzzleInput())

	trimmedList := stripKey("e", replacementList)

	moveToTargets([]string{"HF", "NAl", "OMg"}, molecule, []string{molecule}, trimmedList, 0)
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
