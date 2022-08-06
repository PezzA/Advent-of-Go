package Day201512

import (
	"encoding/json"
	"fmt"
)

// Entry holds wraps the data and runner interfaces for this puzzle
var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string, int) {
	return 2015, 12, "JSAbacusFramework.io", 2
}

func parseAndWalk(input interface{}, total int, filter string) int {
	returnValue := 0
	switch input.(type) {
	case map[string]interface{}:
		subItem := input.(map[string]interface{})
		for _, v := range subItem {
			if filter != "" && v == filter {
				return 0
			}
			returnValue += parseAndWalk(v, total, filter)
		}
	case []interface{}:
		subItem := input.([]interface{})
		for _, v := range subItem {
			returnValue += parseAndWalk(v, total, filter)
		}
	case float64:
		returnValue = int(input.(float64))
	}
	return returnValue
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	var customData interface{}
	json.Unmarshal([]byte(inputData), &customData)

	count := parseAndWalk(customData, 0, "")
	return fmt.Sprintf("%v", count)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	var customData interface{}
	json.Unmarshal([]byte(inputData), &customData)

	count := parseAndWalk(customData, 0, "red")
	return fmt.Sprintf("%v", count)
}
