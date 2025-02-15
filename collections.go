// collections package implements generic data structures.
package collections

// not intended to be used directly
type collection[T any] struct {
	data []T
}

func (c *collection[T]) Size() int {
	return len(c.data)
}

func (c *collection[T]) IsEmpty() bool {
	return len(c.data) == 0
}

func (b *collection[T]) Clear() {
	b.data = nil
}
