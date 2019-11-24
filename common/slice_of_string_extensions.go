package common

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
