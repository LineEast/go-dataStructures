package queue

import (
	"sync"
)

type (
	Queue struct {
		items []any
		lock  sync.Mutex
	}
)

// New is a constructor for a new threadsafe Queue.
func New(size uint) *Queue {
	return &Queue{
		items: make([]any, 0, size),
	}
}

// Len returns len of q
func (q *Queue) Len() int {
	q.lock.Lock()
	defer q.lock.Unlock()

	return len(q.items)
}

// Put item into q
func (q *Queue) Put(item any) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.items = append(q.items, item)
}

// Get element from queue and save it in q
func (q *Queue) Get() any {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.Empty() {
		return nil
	}

	return q.items[0]
}

// Get element and remove from the queue
func (q *Queue) Pop() (item any) {
	q.lock.Lock()
	defer q.lock.Unlock()

	item = q.items[0]
	q.items = q.items[1:]
	return
}

// Clear cleans the q queue
func (q *Queue) Clear() {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.items = q.items[:0]
}

// Empty returns a bool indicating if this queue is empty.
func (q *Queue) Empty() bool {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.Len() == 0 {
		return true
	}

	return false
}
