package common

// Abs Absolute value of integer
func Abs(input int) int {
	if input < 0 {
		return -input
	}
	return input
}

func Factorial(n int) (result int) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

const MaxInt = int(^uint(0) >> 1)
const MinInt = -MaxInt - 1
