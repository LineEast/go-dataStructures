package singly

type (
	List[T any] struct {
		root *Node[T]

		len uint
	}

	Node[T any] struct {
		body T

		next *Node[T]
	}
)

// New returns an initialized list.
func New[T any]() *List[T] {
	return &List[T]{}
}

// Len returns uint len of the list l
func (l *List[T]) Len() uint {
	return l.len
}

// Init initializes or clears list l.
func (l *List[T]) Clear() {
	go func() {
		last := l.Front()
		for n := last; n.Next() != nil; n = n.Next() {
			last = n.next
			n.next = nil
		}
	}()

	l.root = nil
	l.len = 0
}

// Check if l List is empty
func (l *List[T]) IsEmpty() bool {
	if l.Len() == 0 {
		return true
	}

	return false
}

// Front returns the first element of list l or nil if the list is empty.
func (l *List[T]) Front() *Node[T] {
	return l.root
}

// Back returns the last element of list l or nil if the list is empty.
func (l *List[T]) Back() *Node[T] {
	if l.IsEmpty() {
		return nil
	}

	n := l.root
	for n.Next() != nil {
		n = n.next
	}

	return n
}

// Push to the end of the l list
func (l *List[T]) PushBack(body T) {
	node := NewNode(body)

	if l.IsEmpty() {
		l.root = node
	} else {
		l.Back().next = node
	}

	l.len++
}

// Push n node to the end of the l list
func (l *List[T]) PushNodeBack(n *Node[T]) {
	if l.IsEmpty() {
		l.root = n
	} else {
		l.Back().next = n
	}

	l.len++
}

// Push list List to the back of the l List
func (l *List[T]) PushListBack(list *List[T]) {
	if l.IsEmpty() {
		l.root = list.Front()
	} else {
		l.Back().next = list.Front()
	}

	l.len += list.Len()
}

// Push to the front of the l list
func (l *List[T]) PushFront(body T) {
	node := NewNode(body)

	if l.IsEmpty() {
		l.root = node
	} else {
		node.next = l.root
		l.root = node
	}

	l.len++
}

// Push n node to the front of the l list
func (l *List[T]) PushNodeFront(n *Node[T]) {
	if l.IsEmpty() {
		l.root = n
	} else {
		n.next = l.root
		l.root = n
	}

	l.len++
}

// Push list List to the front of the l List
func (l *List[T]) PushListFront(list *List[T]) {
	if l.IsEmpty() {
		l.root = list.Front()
	} else {
		l.Back().next = list.Front()
	}

	l.len += list.Len()
}

func (l *List[T]) PopFront() (v T) {
	v = l.root.body
	l.Remove(l.Front())
	return
}

// Remove removes e from l if e is an element of list l
func (l *List[T]) Remove(e *Node[T]) {
	if l.IsEmpty() {
		return
	}

	last := l.Front()

	if last == e {
		l.root = l.Front().Next()
		e.next = nil
		l.len--
		return
	}

	for n := l.Front(); n.Next() != e || n.Next() != nil; n = n.Next() {
		if n == e {
			last.next = n.Next()
			e.next = nil
			l.len--
			return
		}

		last = n
	}
}

// Create new node with b body
func NewNode[T any](b T) *Node[T] {
	return &Node[T]{
		body: b,
	}
}

// Create empty node
func NewEmptyNode[T any]() *Node[T] {
	return &Node[T]{}
}

// Set b body to the n Node
func (n *Node[T]) SetBody(b T) {
	n.body = b
}

// Get next Node from n Node
func (n *Node[T]) Next() *Node[T] {
	return n.next
}

func (n *Node[T]) Body() T {
	return n.body
}
