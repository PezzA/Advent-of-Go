package Day201713

import (
	"bufio"
	"regexp"
	"strconv"
	"strings"
)

// Entry holds wraps the data and runner interfaces for this puzzle
var Entry testDay

type layers map[int]int

var layerRegex = regexp.MustCompile(`([0-9]+): ([0-9]+)`)

type testDay bool

func getLayers(input string) (layers, int) {

	maxdepth := 0
	layerMap := make(layers, 0)
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		matches := layerRegex.FindStringSubmatch(scanner.Text())

		depth, _ := strconv.Atoi(matches[1])

		if depth > maxdepth {
			maxdepth = depth
		}
		layerRange, _ := strconv.Atoi(matches[2])

		layerMap[depth] = layerRange
	}

	return layerMap, maxdepth
}

func getPositionForTime(time int, length int) int {
	if length == 1 {
		return 1
	}
	maxsteps := (2 * length) - 2
	step := time % maxsteps
	if step >= length {
		return maxsteps - step
	}
	return step
}

func (td testDay) PartOne(inputData string) (string, error) {
	layermap, maxDepth := getLayers(inputData)
	picoSeconds := 0
	severity := 0
	for i := 0; i <= maxDepth; i++ {

		if val, ok := layermap[i]; ok {
			if getPositionForTime(picoSeconds, val) == 0 {
				severity += (i * val)
			}
		}

		picoSeconds++
	}
	return strconv.Itoa(severity), nil
}

func (td testDay) PartTwo(inputData string) (string, error) {
	layermap, maxDepth := getLayers(inputData)
	seedPicoseconds := 0
	for {
		picoSeconds := seedPicoseconds
		valid := true
		for i := 0; i <= maxDepth; i++ {

			if val, ok := layermap[i]; ok {
				if getPositionForTime(picoSeconds, val) == 0 {
					valid = false
				}
			}

			if !valid {
				break
			}
			picoSeconds++
		}
		if valid {
			break
		}
		seedPicoseconds++
	}
	return strconv.Itoa(seedPicoseconds), nil
}

func (td testDay) Day() int {
	return 13
}

func (td testDay) Year() int {
	return 2017
}

func (td testDay) Title() string {
	return "Getting the boilerplate in place"
}
