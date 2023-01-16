package stack

import (
	s "github.com/golang-collections/collections/stack"
	"testing"
)

func BenchmarkPush(b *testing.B) {
	s := NewWithSize[int](100)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
}

func BenchmarkStandartPush(b *testing.B) {
	s := s.New()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
}
