package common

func Same(s []int, t []int) bool {
	for i := range s {
		if s[i] != t[i] {
			return false
		}
	}
	return true
}
