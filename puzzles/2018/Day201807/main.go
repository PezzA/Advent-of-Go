package Day201807

import (
	"fmt"
	"sort"
	"strings"

	"github.com/pezza/advent-of-code/puzzles/Common"
)

var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string) {
	return 2018, 7, "The Sum of Its Parts"
}

type sequence struct {
	item   string
	prereq string
}

type elf struct {
	job       string
	remaining int
}

func getData(input string) []sequence {
	seqs := make([]sequence, 0)

	for _, line := range strings.Split(input, "\n") {
		pre, item := "", ""
		fmt.Sscanf(line, "Step %s must be finished before step %s can begin.", &pre, &item)
		seqs = append(seqs, sequence{item, pre})
	}

	return seqs
}

func getNodeList(input []sequence) []string {
	nodes := make([]string, 0)

	for _, seq := range input {
		if !common.Contains(nodes, seq.item) {
			nodes = append(nodes, seq.item)
		}
		if !common.Contains(nodes, seq.prereq) {
			nodes = append(nodes, seq.prereq)
		}
	}

	return nodes
}

func isStart(test string, seqs []sequence) bool {
	for _, seq := range seqs {
		if seq.item == test {
			return false
		}
	}

	return true
}

func getStarts(seqs []sequence) []string {
	nodes := getNodeList(seqs)

	starts := make([]string, 0)
	for _, node := range nodes {
		if isStart(node, seqs) {
			starts = append(starts, node)
		}
	}

	return starts
}

func getNextCandidates(start string, seqs []sequence) []string {
	nexts := make([]string, 0)

	for _, seq := range seqs {
		if seq.prereq == start {
			nexts = append(nexts, seq.item)
		}
	}

	return nexts
}

func isUnblocked(test string, executed []string, seqs []sequence) bool {

	prereqs := make([]string, 0)

	for _, seq := range seqs {
		if seq.item == test {
			prereqs = append(prereqs, seq.prereq)
		}
	}

	for _, run := range prereqs {
		if !common.Contains(executed, run) {
			return false
		}
	}
	return true
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	seqs := getData(inputData)

	runList := getStarts(seqs)
	sort.Strings(runList)

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

	}
	return strings.Join(executedList, "")
}

func runSim(input string, baseTime int) int {
	elves := make([]elf, 5)

	seqs := getData(input)

	runList := getStarts(seqs)
	sort.Strings(runList)

	executedList := make([]string, 0)
	heldList := make([]string, 0)

	timer := 0
	for len(executedList) != len(getNodeList(seqs)) {

		// decrement each elf's remaining time
		for index := range elves {
			if elves[index].job == "" {
				continue
			}

			elves[index].remaining--

			if elves[index].remaining == 0 {
				executedList = append(executedList, elves[index].job)

				cands := getNextCandidates(elves[index].job, seqs)

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

				elves[index].job = ""
			}
		}

		// try and give each elf a job
		for _, runCand := range runList {
			for index := range elves {
				if elves[index].job == "" {
					elves[index].job = runCand
					elves[index].remaining = baseTime + (int(runCand[0]) - 64)

					runList = runList[1:]
					break
				}
			}
		}

		timer++
	}

	return timer

}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {

	return fmt.Sprintf("%v", runSim(inputData, 60)-1)
}
