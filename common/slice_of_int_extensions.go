package common

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

func DeepCopy(in []int) []int {
	out := make([]int, len(in))
	for i := range in {
		out[i] = in[i]
	}
	return out
}
