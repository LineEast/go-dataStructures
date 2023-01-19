package skip

type (
	NodeC[T any] struct {
		key   T
		value T

		next **NodeC[T]
	}

	SkipListC[T any] struct {
		level uint
		size  uint

		header *NodeC[T]
	}
)
