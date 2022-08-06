package common

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_WrappedIndex(t *testing.T) {

	RegisterTestingT(t)

	Expect(GetWrappedIndex(8, 8, 1)).Should(Equal(1))
	Expect(GetWrappedIndex(0, 8, 1)).Should(Equal(1))
	Expect(GetWrappedIndex(1, 8, 8)).Should(Equal(1))
	Expect(GetWrappedIndex(8, 8, -1)).Should(Equal(7))

	Expect(GetWrappedIndex(0, 4, 0)).Should(Equal(0))
	Expect(GetWrappedIndex(1, 4, 0)).Should(Equal(1))
	Expect(GetWrappedIndex(2, 4, 0)).Should(Equal(2))
	Expect(GetWrappedIndex(3, 4, 0)).Should(Equal(3))
	Expect(GetWrappedIndex(4, 4, 0)).Should(Equal(0))
	Expect(GetWrappedIndex(5, 4, 0)).Should(Equal(1))
	Expect(GetWrappedIndex(6, 4, 0)).Should(Equal(2))
	Expect(GetWrappedIndex(7, 4, 0)).Should(Equal(3))
	Expect(GetWrappedIndex(8, 4, 0)).Should(Equal(0))
	Expect(GetWrappedIndex(9, 4, 0)).Should(Equal(1))

}
