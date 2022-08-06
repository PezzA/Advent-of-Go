package Day201804

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

func (td dayEntry) Describe() (int, int, string, int) {
	return 2018, 4, "Repose Record", 2
}

// Entry holds wraps the data and runner interfaces for this puzzle
var Entry dayEntry

type dayEntry bool

type timedData struct {
	timeStamp time.Time
	data      string
}

func getData(input string) (map[string][]string, time.Time, time.Time) {
	orderedData := make([]timedData, 0)

	var start time.Time
	var end time.Time

	for _, line := range strings.Split(input, "\n") {
		ts, _ := time.Parse("2006-01-02 15:04", line[1:17])

		if start.Year() == 1 || ts.Before(start) {
			start = ts
		}

		if end.Year() == 1 || end.Before(ts) {
			end = ts
		}

		orderedData = append(orderedData, timedData{ts, line})
	}

	sort.Slice(orderedData, func(i, j int) bool { return orderedData[i].timeStamp.Before(orderedData[j].timeStamp) })

	days := make(map[string][]string)

	currGuard, falls := "", 0

	for _, val := range orderedData {
		key := fmt.Sprintf("%v-%v", val.timeStamp.Month(), val.timeStamp.Day())

		if _, ok := days[key]; !ok {
			days[key] = make([]string, 60)
		}

		guardData := strings.Fields(val.data[19:])

		switch guardData[0] {
		case "Guard":
			currGuard = guardData[1]
		case "falls":
			falls = val.timeStamp.Minute()
		case "wakes":
			wakes := val.timeStamp.Minute()

			for mins := falls; mins < wakes; mins++ {

				days[key][mins] = currGuard
			}
		}

	}

	return days, start, end
}

func getSleepyGuard(shifts map[string][]string) string {
	sleeps := make(map[string]int)

	for _, v := range shifts {
		for index := range v {
			if v[index] != "" {
				sleeps[v[index]]++
			}

		}
	}

	max := -1
	maxGuard := ""
	for k, v := range sleeps {
		if max == -1 || v > max {
			max = v
			maxGuard = k
		}
	}
	return maxGuard
}

func getMinSleepyGuardMostAsleep(guard string, shifts map[string][]string) int {
	mins := make(map[int]int)

	for _, v := range shifts {
		for index := range v {
			if v[index] == guard {
				mins[index]++
			}
		}
	}

	max, maxMin := -1, -1

	for k, v := range mins {
		if max == -1 || v > max {
			max = v
			maxMin = k
		}
	}

	return maxMin
}

func getSleepyMinute(shifts map[string][]string) string {

	allSleeps := make(map[string]int)

	for i := 0; i < 60; i++ {
		for _, v := range shifts {
			if v[i] != "" {
				key := fmt.Sprintf("%v-%v", v[i], i)
				allSleeps[key]++
			}
		}
	}

	max, maxKey := -1, ""

	for k, v := range allSleeps {
		if max == -1 || v > max {
			max = v
			maxKey = k
		}
	}

	return maxKey
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	shifts, _, _ := getData(inputData)
	sleepyGuard := getSleepyGuard(shifts)
	minsMostAsleep := getMinSleepyGuardMostAsleep(sleepyGuard, shifts)

	sleepyGuardId, _ := strconv.Atoi(strings.Replace(sleepyGuard, "#", "", -1))
	return fmt.Sprintf("%v", sleepyGuardId*minsMostAsleep)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	shifts, _, _ := getData(inputData)

	sleepyMinuteKey := getSleepyMinute(shifts)

	bits := strings.Split(sleepyMinuteKey, "-")

	sleepyGuardId, _ := strconv.Atoi(strings.Replace(bits[0], "#", "", -1))
	sleepyMinute, _ := strconv.Atoi(bits[1])

	return fmt.Sprintf("%v", sleepyGuardId*sleepyMinute)
}
