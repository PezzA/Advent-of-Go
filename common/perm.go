package common

func permGen(a []int, f func([]int)) {
	perm(a, f, 0)
}

// Permute the values at index i to len(a)-1.
// taken from https://yourbasic.org/golang/generate-permutation-slice-string/
// will roll own https://en.wikipedia.org/wiki/Heap%27s_algorithm later
func perm(a []int, f func([]int), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

// GetPermuations returns all comvination of input int array
func GetPermuations(input []int) [][]int {
	allPerms := [][]int{}
	permGen(input, func(p []int) {
		newArr := []int{}

		for index := range p {
			newArr = append(newArr, p[index])
		}

		allPerms = append(allPerms, newArr)
	})
	return allPerms
}
