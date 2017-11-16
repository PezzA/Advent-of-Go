package Day201611

import "strings"

func getStartingFacility(s string) facility {
	componentList := make([]component, 0)

	// find all the microchips in play
	for index, line := range strings.Split(s, "\n") {
		for _, word := range strings.Fields(line) {
			if strings.Index(word, "-") > 0 {
				element := strings.Split(word, "-")
				componentList = append(componentList, component{index, element[0], true})
			}
		}
	}

	// then go and find where the reactors are
	for index, line := range strings.Split(s, "\n") {
		for _, chip := range componentList {
			if strings.Index(line, chip.element+" generator") > 0 {
				componentList = append(componentList, component{index, chip.element, false})
			}
		}
	}

	return facility{0, componentList}
}
