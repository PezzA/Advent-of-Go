package Day201713

import (
	"bufio"
	"regexp"
	"strconv"
	"strings"
)

type dayEntry bool

var Entry dayEntry

func (td dayEntry) Describe() (int, int, string, int) {
	return 2017, 13, "Packet Scanners", 2
}

type layers map[int]int

var layerRegex = regexp.MustCompile(`([0-9]+): ([0-9]+)`)

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

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
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
	return strconv.Itoa(severity)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
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
	return strconv.Itoa(seedPicoseconds)
}
