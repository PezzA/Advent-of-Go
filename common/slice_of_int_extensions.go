package common

import "math"

func Min(in []int) (int, int) {
	smallest := math.MaxInt32
	index := 0
	for i := range in {
		if in[i] < smallest {
			smallest = in[i]
			index = i
		}
	}
	return smallest, index
}

func Max(in []int) (int, int) {
	largest := math.MinInt32
	index := 0

	for i := range in {
		if in[i] > largest {
			largest = in[i]
			index = i
		}
	}
	return largest, index
}

func Same(s []int, t []int) bool {
	if len(s) != len(t) {
		return false
	}

	for i := range s {
		if s[i] != t[i] {
			return false
		}
	}
	return true
}

func ContainsInt(list []int, test int) (bool, int) {
	for i := range list {
		if list[i] == test {
			return true, i
		}
	}
	return false, 0
}

func DeepCopy(in []int) []int {
	out := make([]int, len(in))
	for i := range in {
		out[i] = in[i]
	}
	return out
}
