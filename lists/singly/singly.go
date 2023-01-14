package singly

// New returns an initialized list.
func New() *List {
	return &List{}
}

// Len returns uint len of the list l
func (l *List) Len() uint {
	return l.len
}

// Init initializes or clears list l.
func (l *List) Clear() {
	l.root = nil
	l.len = 0
}

// Front returns the first element of list l or nil if the list is empty.
func (l *List) Front() *node {
	if l.Len() == 0 {
		return nil
	}

	return l.root
}

// Back returns the last element of list l or nil if the list is empty.
func (l *List) Back() *node {
	if l.Len() == 0 {
		return nil
	}

	n := l.root
	for n.next != nil {
		n = n.next
	}

	return n
}

func (l *List) PushBack(body any) {
	node := &node{
		body: body,
	}

	if l.len == 0 {
		l.root = node
	} else {
		n := l.root
		for n.next != nil {
			n = n.next
		}

		n.next = node
	}

	l.len++
}

func (l *List) PushFront(body *any) {
	node := &node{
		body: body,
		next: l.root,
	}

	l.root = node

	l.len++
}

// Remove removes e from l if e is an element of list l.
// It returns the element value e.body. The element must not be nil.
func (l *List) Remove(e *node) any {
	if l.root == e {
		l.root = e.next
		return e.body
	}

	n := l.root
	index := 0
	for n.next != e {
		n = n.next
		index++
	}

	last := l.root
	for i := 0; i < index; i++ {
		last = last.next
	}

	last.next = n.next
	return e.body
}

// // Returns body of first element
// func (l *LinkedList) Front() any {
// 	l.rLock()
// 	defer l.rUnlock()

// 	return l.root.body
// }

// // Returns body of last element
// func (l *LinkedList) Back() any {
// 	l.rLock()
// 	defer l.rUnlock()

// 	n := l.root
// 	for n.next != nil {
// 		n = n.next
// 	}

// 	return n.body
// }

// func (l *LinkedList) Len() uint {
// 	l.rLock()
// 	defer l.rUnlock()

// 	return l.len
// }

// func (l *LinkedList) lock() {
// 	l.mtx.Lock()
// }

// func (l *LinkedList) unlock() {
// 	l.mtx.Unlock()
// }

// func (l *LinkedList) rLock() {
// 	l.mtx.RLock()
// }

// func (l *LinkedList) rUnlock() {
// 	l.mtx.RUnlock()
// }
