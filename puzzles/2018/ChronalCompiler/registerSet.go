package ChronalCompiler

import "fmt"

type RegisterSet map[int]int

func (rs RegisterSet) String() string {
	return fmt.Sprintf("[%v %v %v %v]", rs[0], rs[1], rs[2], rs[3])
}

func (rs RegisterSet) DeepCopy() RegisterSet {
	retVal := make(RegisterSet)
	for k, v := range rs {
		retVal[k] = v
	}
	return retVal
}

func (rs RegisterSet) Same(cmp RegisterSet) bool {
	return rs[0] == cmp[0] && rs[1] == cmp[1] && rs[2] == cmp[2] && rs[3] == cmp[3]
}
