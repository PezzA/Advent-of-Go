package common

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_ReadData(t *testing.T) {
	RegisterTestingT(t)

	stack := NewStack()

	stack.Push("A")

	Expect(stack.Size()).Should(Equal(1))

	stack.Push("B")

	Expect(stack.Size()).Should(Equal(2))

	Expect(stack.Peek()).Should(Equal("B"))

	popped := stack.Pop()
	Expect(popped).Should(Equal("B"))
	Expect(stack.Size()).Should(Equal(1))
	Expect(stack.Peek()).Should(Equal("A"))

	popped = stack.Pop()
	Expect(popped).Should(Equal("A"))
	Expect(stack.Size()).Should(Equal(0))

}

func Test_AddToBottom(t *testing.T) {
	RegisterTestingT(t)

	stack := NewStack()

	stack.AddAtBottom("B")
	stack.AddAtBottom("A")

	Expect(stack.Size()).Should(Equal(2))
	Expect(stack.Peek()).Should(Equal("B"))

	stack.AddAtBottom("C")
	Expect(stack.Size()).Should(Equal(3))
	Expect(stack.Peek()).Should(Equal("B"))

}
