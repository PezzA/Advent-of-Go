package Day202206

import (
	"fmt"
)

func scanForMarker(input string, length int) int {
	for i := length - 1; i < len(input); i++ {
		chars := make(map[byte]bool, 0)

		for j := 0; j < length; j++ {
			chars[input[i-j]] = true
		}

		if len(chars) == length {
			return i + 1
		}

	}
	return -1
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	return fmt.Sprint(scanForMarker(inputData, 4))
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	return fmt.Sprint(scanForMarker(inputData, 14))
}
