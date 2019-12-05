package intcode

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_ReadData(t *testing.T) {
	RegisterTestingT(t)

	Expect(parseOpCode(1002)).Should(Equal(OpCode{2, 0, 1, 0}))
	Expect(parseOpCode(10001)).Should(Equal(OpCode{1, 0, 0, 1}))
	Expect(parseOpCode(1)).Should(Equal(OpCode{1, 0, 0, 0}))
	Expect(parseOpCode(99)).Should(Equal(OpCode{99, 0, 0, 0}))

}
