/* Simple, non concurrent stack implementation
with no guards */
package common

import "errors"

// good candiate for generics?
type Stack struct {
	stack []string
}

func NewStack() *Stack {
	return &Stack{[]string{}}
}

func (s *Stack) Push(input string) {
	s.stack = append(s.stack, input)
}

func (s *Stack) Pop() (string, error) {
	if len(s.stack) == 0 {
		return "", errors.New("Stack Empty")
	}
	retVal := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return retVal, nil
}

func (s *Stack) Peek() string {
	if len(s.stack) == 0 {
		return ""
	}

	return s.stack[len(s.stack)-1]
}

func (s *Stack) Size() int {
	return len(s.stack)
}

// Used for loading data in reverse order
func (s *Stack) AddAtBottom(input string) {
	s.stack = append([]string{input}, s.stack...)
}
