package common

func permGen(a []int64, f func([]int64)) {
	perm(a, f, 0)
}

// Permute the values at index i to len(a)-1.
// taken from https://yourbasic.org/golang/generate-permutation-slice-string/
// will roll own https://en.wikipedia.org/wiki/Heap%27s_algorithm later
func perm(a []int64, f func([]int64), i int64) {
	if i > int64(len(a)) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < int64(len(a)); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

// GetPermuations returns all comvination of input int array
func GetPermuations(input []int64) [][]int64 {
	allPerms := [][]int64{}
	permGen(input, func(p []int64) {
		newArr := []int64{}

		for index := range p {
			newArr = append(newArr, p[index])
		}

		allPerms = append(allPerms, newArr)
	})
	return allPerms
}
