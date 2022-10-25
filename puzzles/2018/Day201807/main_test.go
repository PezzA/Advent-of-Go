package Day201807

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/pezza/advent-of-code/puzzles/common"
)

var testData = `Step C must be finished before step A can begin.
Step C must be finished before step F can begin.
Step A must be finished before step B can begin.
Step A must be finished before step D can begin.
Step B must be finished before step E can begin.
Step D must be finished before step E can begin.
Step F must be finished before step E can begin.`

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	seqs := getData(testData)
	Expect(seqs[0]).Should(Equal(sequence{"A", "C"}))

	start := getStarts(seqs)

	runList := start

	executedList := make([]string, 0)
	heldList := make([]string, 0)

	for len(runList) > 0 {
		next := runList[0]
		runList = runList[1:]
		executedList = append(executedList, next)
		cands := getNextCandidates(next, seqs)

		for _, cand := range cands {
			if isUnblocked(cand, executedList, seqs) {
				if !common.Contains(runList, cand) {
					runList = append(runList, cand)
				}
			} else {
				if !common.Contains(heldList, cand) {
					heldList = append(heldList, cand)
				}
			}

		}

		newHeldList := make([]string, 0)
		for _, check := range heldList {
			if isUnblocked(check, executedList, seqs) {
				if !common.Contains(runList, check) {
					runList = append(runList, check)
				}
			} else {
				newHeldList = append(newHeldList, check)
			}
		}
		heldList = newHeldList

		sort.Strings(runList)

		fmt.Println(next, runList, heldList, executedList)
	}

	fmt.Println(strings.Join(executedList, ""))
}

func Test_PuzzleData(t *testing.T) {
	RegisterTestingT(t)
	seqs := getData(Entry.PuzzleInput())
	Expect(seqs[0]).Should(Equal(sequence{"G", "W"}))

	start := getStarts(seqs)
	sort.Strings(start)
	runList := start

	executedList := make([]string, 0)
	heldList := make([]string, 0)

	for len(executedList) != len(getNodeList(seqs)) {
		next := runList[0]
		runList = runList[1:]
		executedList = append(executedList, next)
		cands := getNextCandidates(next, seqs)

		// get newly available candidates
		for _, cand := range cands {
			if isUnblocked(cand, executedList, seqs) {
				if !common.Contains(runList, cand) {
					runList = append(runList, cand)
				}
			} else {
				if !common.Contains(heldList, cand) {
					heldList = append(heldList, cand)
				}
			}

		}
		// check that anything previously held is now available
		newHeldList := make([]string, 0)
		for _, check := range heldList {
			if isUnblocked(check, executedList, seqs) {
				if !common.Contains(runList, check) {
					runList = append(runList, check)
				}
			} else {
				newHeldList = append(newHeldList, check)
			}
		}
		heldList = newHeldList

		// sort the list alphabetically
		sort.Strings(runList)

		fmt.Println(next, runList, heldList, executedList)
	}

	fmt.Println(strings.Join(executedList, ""))
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

	runSim(Entry.PuzzleInput(), 60)
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
