package Day201909

import (
	"fmt"
	"strconv"
	"strings"
)

func getListIntData(input string) []int64 {
	retval := []int64{}

	for _, i := range strings.Split(input, ",") {
		newVal, _ := strconv.ParseInt(i, 10, 64)

		retval = append(retval, newVal)
	}

	return retval
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	return fmt.Sprintf("%v", " -- Not Yet Implemented --")
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	return fmt.Sprintf("%v", " -- Not Yet Implemented --")
}
