package doubly

type (
	List[T any] struct {
		head, tail *Node[T]

		len uint
	}
)

func New[T any]() *List[T] {
	return &List[T]{}
}

func (l *List[T]) Len() uint {
	return l.len
}

func (l *List[T]) Clear() {
	l.head, l.tail = nil, nil
}

func (l *List[T]) IsEmpty() bool {
	if l.Len() == 0 {
		return true
	}

	return false
}

func (l *List[T]) Front() *Node[T] {
	return l.head
}

func (l *List[T]) Back() *Node[T] {
	return l.tail
}

func (l *List[T]) PushFront(body T) {
	node := NewNode(body)

	if l.IsEmpty() {
		l.head = node
		l.tail = node
	} else {
		node.prev = l.Front()
		l.Front().prev = node
		l.head = node
	}

	l.len++
}

func (l *List[T]) PushNodeFront(node *Node[T]) {
	if l.IsEmpty() {
		l.head = node
		l.tail = node
	} else {
		node.prev = l.Front()
		l.Front().prev = node
		l.head = node
	}

	l.len++
}

func (l *List[T]) PushListFront(list *List[T]) {
	if l.IsEmpty() {
		l.head = list.head
		l.tail = list.tail
	} else {
		l.tail.next = list.head
		l.tail = list.tail
	}

	l.len += list.Len()
}

func (l *List[T]) PushBack(body T) {
	node := NewNode(body)

	if l.IsEmpty() {
		l.head = node
		l.tail = node
	} else {
		node.prev = l.Back()
		l.Back().next = node
		l.tail = node
	}

	l.len++
}

func (l *List[T]) PushNodeBack(node *Node[T]) {
	if l.IsEmpty() {
		l.head = node
		l.tail = node
	} else {
		node.prev = l.Back()
		l.Back().next = node
		l.tail = node
	}

	l.len++
}

func (l *List[T]) PopFront() *Node[T] {
	if l.IsEmpty() {
		return nil
	}

	l.Remove(l.Front())

	return l.Front()
}

func (l *List[T]) PopBack() *Node[T] {
	if l.IsEmpty() {
		return nil
	}

	l.Remove(l.Back())

	return l.Back()
}

func (l *List[T]) Remove(n *Node[T]) {
	n.prev.next = n.next
	n.next.next = n.prev

	n.next = nil
	n.prev = nil

	l.len--
}
