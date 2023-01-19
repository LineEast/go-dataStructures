package singly

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

// const (
// 	success = "\t\u2713\t"
// 	failed  = "\t\u2717\t"
// )

func compare(a, b int) bool {
	return a > b
}

func newIntOrderdList(len int) *SinglyOrderd[int] {
	l := NewOrderdList(compare)
	for i := 0; i < len; i++ {
		l.Push(i)
	}

	return l
}

func TestOrderdList(t *testing.T) {
	l := newIntOrderdList(5)

	spew.Dump(l.List)
}
