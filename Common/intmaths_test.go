package common

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_PartOne(t *testing.T) {

	RegisterTestingT(t)

	Expect(GetWrappedIndex(8, 8, 1)).Should(Equal(1))
	Expect(GetWrappedIndex(0, 8, 1)).Should(Equal(1))
	Expect(GetWrappedIndex(1, 8, 8)).Should(Equal(1))
	Expect(GetWrappedIndex(8, 8, -1)).Should(Equal(7))
}
