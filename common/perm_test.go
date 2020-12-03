package common

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

func Test_Perms(t *testing.T) {
	RegisterTestingT(t)

	fmt.Println(GetPermuations([]int64{0, 1, 2, 3, 4}))
}
