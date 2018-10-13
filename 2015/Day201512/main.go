package Day201512

import (
	"fmt"
	"strconv"
	"strings"
)

// Entry holds wraps the data and runner interfaces for this puzzle
var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string) {
	return 2015, 12, "JSAbacusFramework.io"
}

func splitAndAdd(input string) int {
	count := 0
	lsb := strings.Split(input, "[")
	for _, innerLsb := range lsb {
		rsb := strings.Split(innerLsb, "]")
		for _, innerRsb := range rsb {
			lfb := strings.Split(innerRsb, "{")
			for _, innerLfb := range lfb {
				rfb := strings.Split(innerLfb, "}")
				for _, innerRfb := range rfb {
					colon := strings.Split(innerRfb, ":")
					for _, innerCol := range colon {
						comma := strings.Split(innerCol, ",")
						for _, innerCom := range comma {
							value, err := strconv.Atoi(innerCom)
							if err == nil {
								count += value
							}
						}
					}
				}
			}
		}
	}
	return count
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	val := splitAndAdd(Entry.PuzzleInput())
	return fmt.Sprintf("%v", val)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	return fmt.Sprintf(" -- Not Yet Implemented --")
}
