package singly

import (
	"sync"
)

type (
	// L interface {
	// 	New() *List

	// 	Back() *node
	// 	Front() *node
	// 	Init() *List
	// 	InsertAfter(v any, mark *node) *node
	// 	InsertBefore(v any, mark *node) *node
	// 	Len() int
	// 	MoveAfter(e, mark *node)
	// 	MoveBefore(e, mark *node)
	// 	MoveToBack(e *node)
	// 	MoveToFront(e *node)
	// 	PushBack(v any) *node
	// 	PushBackList(other *List)
	// 	PushFront(v any) *node
	// 	PushFrontList(other *List)
	// 	Remove(e *node) any
	// }
	List struct {
		root *node

		mtx sync.RWMutex

		len uint
	}

	node struct {
		body any

		next *node
	}
)
