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
	n4 := NewNode(4)
	n5 := NewNode(5)

	l.PushNodeFront(n5)
	l.PushNodeFront(n4)
	l.PushNodeFront(n3)
	l.PushNodeFront(n2)
	l.PushNodeFront(n1)

	l.Clear()
	spew.Dump(l.Front())
	t.Log(n1, n2, n3, n4, n5)

	// l.PushNodeBack(n1)
	// l.PushNodeBack(n2)
	// l.PushNodeBack(n3)

}
