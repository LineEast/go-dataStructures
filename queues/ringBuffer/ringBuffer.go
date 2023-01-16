package ringbuffer

import ()

type (
	node struct {
		body any
	}

	RingBuffer struct {
		len   uint64
		nodes []*node
	}
)

func New(size uint64) *RingBuffer {
	return &RingBuffer{
		nodes: make([]*node, roundUp(size)),
	}
}

// Cap returns the capacity of this ring buffer.
func (rb *RingBuffer) Cap() uint64 {
	return uint64(len(rb.nodes))
}

// Len returns the number of items in the queue.
func (rb *RingBuffer) Len() uint64 {
	return rb.len
}

func (rb *RingBuffer) Put(item any) {
	node := &node{
		body: item,
	}

	rb.nodes = append(rb.nodes, node)
}

func (rb *RingBuffer) Get() (item any, err error) {
	item = rb.nodes[0]
	rb.nodes = append(rb.nodes[:0], rb.nodes[1:]...)
	return
}

func roundUp(v uint64) uint64 {
	v--
	v |= v >> 1
	v |= v >> 2
	v |= v >> 4
	v |= v >> 8
	v |= v >> 16
	v |= v >> 32
	v++
	return v
}
