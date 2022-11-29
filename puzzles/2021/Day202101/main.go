package Day202101

import (
    "strings"
    "fmt"
    "strconv"
)

func importData(input string) []int {
    lines:= strings.Split(input, "\n")
    data := make([]int,len(lines))

    for i := range lines {
	val, _ := strconv.Atoi(lines[i])
	data[i] = val
    }

    return data
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
    data := importData(inputData)
    count := 0
    for i := 1; i< len(data); i++ {
	if data[i] > data[i-1] {
	    count++
	}
    }

    return fmt.Sprint(count);
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
    data := importData(inputData)
    count := 0

    for i := 3; i< len(data); i++ {
	if data[i] + data[i-1] + data[i-2] > data[i-1] + data[i-2] + data[i-3] {
	    count++
	}
    }

    return fmt.Sprint(count);
}
