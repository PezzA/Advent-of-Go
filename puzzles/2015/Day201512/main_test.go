package Day201512

import (
	"encoding/json"
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

	var customData interface{}
	err := json.Unmarshal([]byte(Entry.PuzzleInput()), &customData)
	if err != nil {
		t.Error(err)
	}

	count := parseAndWalk(customData, 0, "")

	if err != nil {
		t.Fatal(err)
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
