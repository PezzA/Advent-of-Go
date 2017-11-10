package Day112016

import "sort"

type facility struct {
	elevator   int
	components componentList
}

func (f facility) equals(compare facility) bool {
	if f.elevator != compare.elevator {
		return false
	}

	sort.Sort(f.components)
	sort.Sort(compare.components)

	for i := 0; i < len(f.components); i++ {
		if !f.components[i].equals(compare.components[i]) {
			return false
		}
	}

	return true
}

type solveState struct {
	f         facility
	subStates []solveState
}
