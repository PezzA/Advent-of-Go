package common

// GetWrappedIndex will return the next index from an array as it it was a loop
func GetWrappedIndex(currentPos int, loopsize int, modifier int) int {
	if loopsize == 0 {
		return 0
	}
	rawIndex := currentPos + modifier
	return rawIndex % loopsize
}

// Contains returns true if test string exists in string list
func Contains(list []string, test string) bool {
	for _, val := range list {
		if val == test {
			return true
		}
	}
	return false
}

func Remove(input []string, target string) []string {
	retVal := make([]string, 0)

	for _, val := range input {
		if val != target {
			retVal = append(retVal, val)
		}
	}

	return retVal
}
