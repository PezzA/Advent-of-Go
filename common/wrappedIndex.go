package common

// GetWrappedIndex will return the next index from an array as if it was a loop
func GetWrappedIndex(currentPos int, loopSize int, modifier int) int {
	if loopSize == 0 {
		return 0
	}
	rawIndex := currentPos + modifier
	return rawIndex % loopSize
}
