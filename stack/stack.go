package stack

import ()

type (
	stack[T any] struct {
		stack []T
	}
)

// creates a new stack with type from expample
func New[T any](example T) *stack[T] {
	return &stack[T]{}
}

// IsEmpty: check if stack is empty
func (s *stack[T]) IsEmpty() bool {
	return s.Len() == 0
}

// Get len of s stack
func (s *stack[T]) Len() int {
	return len(s.stack)
}

// Clear s stack
func (s *stack[T]) Clear() {
	s.stack = s.stack[:0]
}

// Push a new value onto the stack
func (s *stack[T]) Push(value T) {
	s.stack = append(s.stack, value)
}

// Remove and return top element of stack.
// Return false if stack is empty.
func (s *stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var resutl T
		return resutl, false
	}

	element := s.stack[s.Len()-1]
	s.stack = s.stack[:s.Len()-1]
	return element, true
}

// View the top item on the stack
func (s *stack[T]) Peek() (v T) {
	return s.stack[s.Len()-1]
}
