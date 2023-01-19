package singly

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

const (
	success = "\t\u2713\t"
	failed  = "\t\u2717\t"
)

func newIntList(len uint) *List[int] {
	l := New[int]()
	for i := 0; i < int(len); i++ {
		n := NewNode(i)
		l.PushNodeBack(n)
	}

	return l
}

func TestIsEmpty(t *testing.T) {
	l := newIntList(5)

	if l.IsEmpty() {
		t.Logf("%s", failed)
	} else {
		t.Logf("%s", success)
	}
}

func TestClear(t *testing.T) {
	l := newIntList(5)

	l.Clear()

	if l.IsEmpty() {
		t.Logf("%s", success)
	} else {
		spew.Dump(l)
		t.Logf("%s", failed)
	}
}
