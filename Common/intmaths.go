package common

// Abs Absolute value of integer
func Abs(input int) int {
	if input < 0 {
		return -input
	}
	return input
}
