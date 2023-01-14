package stack

import ()

type (
	stack struct {
		stack []any
	}
)

func New() *stack {
	return &stack{}
}

// IsEmpty: check if stack is empty
func (s *stack) IsEmpty() bool {
	return len(*&s.stack) == 0
}

// Get len of s stack
func (s *stack) Len() int {
	return len(s.stack)
}

// Clear s stack
func (s *stack) Clear() {
	s.stack = s.stack[:0]
}

// Push a new value onto the stack
func (s *stack) Push(value any) {
	s.stack = append(s.stack, value)
}

// Remove and return top element of stack.
// Return false if stack is empty.
func (s *stack) Pop() (any, bool) {
	if s.IsEmpty() {
		return nil, false
	}

	element := s.stack[s.Len()-1]
	s.stack = s.stack[:s.Len()-1]
	return element, true
}

// View the top item on the stack
func (s *stack) Peek() any {
	return s.stack[:s.Len()-1]
}
