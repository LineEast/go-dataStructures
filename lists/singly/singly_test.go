package singly

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestRemove(t *testing.T) {
	l := New[int]()
	n1 := NewNode(1)
	n2 := NewNode(2)
	n3 := NewNode(3)

	l.PushNodeFront(n3)
	l.PushNodeFront(n2)
	l.PushNodeFront(n1)

	l.Remove(nil)

	l.Clear()

	l.PushNodeBack(n1)
	l.PushNodeBack(n2)
	l.PushNodeBack(n3)

	spew.Dump(l.Front())
}
