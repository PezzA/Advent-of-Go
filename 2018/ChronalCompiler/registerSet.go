package chronalcompiler

import "fmt"

// RegisterSet is a type to express a map of ints
type RegisterSet map[int]int

// NewRegisterSet returns a new registerset with registers 0-3 set to zero
func NewRegisterSet() RegisterSet {
	newRs := make(RegisterSet, 0)
	newRs[0] = 0
	newRs[1] = 0
	newRs[2] = 0
	newRs[3] = 0

	return newRs
}

func (rs RegisterSet) String() string {
	return fmt.Sprintf("[%v %v %v %v]", rs[0], rs[1], rs[2], rs[3])
}

// DeepCopy will create a new clone of the input registerset
func (rs RegisterSet) DeepCopy() RegisterSet {
	retVal := make(RegisterSet)
	for k, v := range rs {
		retVal[k] = v
	}
	return retVal
}

// Same will compare if 2 register sets have the same register values
func (rs RegisterSet) Same(cmp RegisterSet) bool {
	return rs[0] == cmp[0] && rs[1] == cmp[1] && rs[2] == cmp[2] && rs[3] == cmp[3]
}
