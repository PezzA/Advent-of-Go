package Day201902

import (
	"fmt"

	"github.com/pezza/advent-of-code/2019/intcode"
)

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	vc := intcode.New(inputData)

	init := make(map[int64]int64, 0)
	init[1] = 12
	init[2] = 2

	vc.RunProgram(init, []int64{1}, nil, nil, nil)

	return fmt.Sprintf("%v", vc.GetValue(0))
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	vc := intcode.New(inputData)

	done := false
	actNoun, actVerb := 0, 0

	for noun := 0; noun <= 100; noun++ {
		for verb := 0; verb <= 100; verb++ {

			init := make(map[int64]int64, 0)
			init[1] = int64(noun)
			init[2] = int64(verb)

			vc.RunProgram(init, []int64{}, nil, nil, nil)

			if vc.GetValue(0) == 19690720 {
				actNoun = noun
				actVerb = verb
				done = true
			}

			if done {
				break
			}
		}
		if done {
			break
		}
	}

	return fmt.Sprintf("%v", 100*actNoun+actVerb)
}
