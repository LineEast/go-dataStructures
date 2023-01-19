package skip

import (
	"math/rand"
	"time"
)

type (
	Node[T any] struct {
		levels []*Node[T]

		value T
		key   T

		score float64

		prev         *Node[T]
		prevTopLevel *Node[T]
		list         *SkipList[T]
	}

	SkipList[T any] struct {
		levels []*Node[T]

		rand *rand.Rand

		maxLevel uint
		len      uint
		back     *Node[T]
	}
)

func New[T any]() *SkipList[T] {
	source := rand.NewSource(time.Now().UnixNano())
	return &SkipList[T]{
		levels: make([]*Node[T], 48),

		rand: rand.New(source),

		maxLevel: 48,
	}
}
