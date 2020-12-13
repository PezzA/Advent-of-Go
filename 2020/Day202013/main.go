package Day202013

import (
	"fmt"
	"strconv"
	"strings"
)

type schedule struct {
	active    bool
	timeStamp int
	slot      int
}

func (s schedule) timeToNextBus(offset int) int {
	if !s.active {
		return -1
	}

	return s.timeStamp - (offset % s.timeStamp)
}

func (s schedule) isSync(offset int) bool {
	return offset%s.timeStamp == s.slot
}

func (s schedule) isSyncWithSlot(offset int) bool {
	offset += s.slot
	return offset%s.timeStamp == 0
}

func (s schedule) String() string {
	return fmt.Sprintf("bus %d", s.timeStamp)
}

func getStep(s1 schedule, s2 schedule, offset int, step int, terminal bool) (int, int) {
	i := offset

	match := -1
	for {
		if s1.isSyncWithSlot(i) && s2.isSyncWithSlot(i) {
			if i != 0 {
				if terminal {
					return i, 0
				}
				if match == -1 {
					match = i
				} else {
					return i, i - match
				}

			}
		}

		i += step
	}
}

func getData(input string) (int, []schedule) {
	lines := strings.Split(input, "\n")

	earliestDeparture, _ := strconv.Atoi(lines[0])

	scheduleData := strings.Split(lines[1], ",")

	schedules := make([]schedule, len(scheduleData))

	for i := range scheduleData {
		if scheduleData[i] == "x" {
			schedules[i] = schedule{false, 0, i}
			continue
		}

		ts, _ := strconv.Atoi(scheduleData[i])
		schedules[i] = schedule{
			active:    true,
			timeStamp: ts,
			slot:      i,
		}
	}
	return earliestDeparture, schedules
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	offset, schedules := getData(inputData)

	waitTime := -1
	var targetSchedule = schedule{false, 0, -1}

	for i := range schedules {
		if !schedules[i].active {
			continue

		}
		swt := schedules[i].timeToNextBus(offset)

		if swt < waitTime || waitTime == -1 {
			waitTime = swt
			targetSchedule = schedules[i]
		}
	}

	return fmt.Sprintf("%v", targetSchedule.timeStamp*waitTime)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {

	_, schedules := getData(Entry.PuzzleInput())

	activeS := make([]schedule, 0)

	for i := range schedules {
		if schedules[i].active {
			activeS = append(activeS, schedules[i])
		}
	}

	offset, step := 0, activeS[0].timeStamp

	for i := 0; i < len(activeS)-2; i++ {
		if activeS[i].active {
			offset, step = getStep(activeS[i], activeS[i+1], offset, step, false)
		}
	}

	finalOffset, _ := getStep(activeS[len(activeS)-2], activeS[len(activeS)-1], offset, step, true)

	return fmt.Sprintf("%v", finalOffset)
}
