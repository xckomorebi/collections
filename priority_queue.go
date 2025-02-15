package collections

type Comparable[T any] interface {
	CompareTo(o T) int
}

type PriorityQueue[T Comparable[T]] struct {
	collection[T]
}

func NewPriorityQueue[T Comparable[T]]() *PriorityQueue[T] {
	return &PriorityQueue[T]{}
}

func (pq *PriorityQueue[T]) swap(i, j int) {
	pq.data[i], pq.data[j] = pq.data[j], pq.data[i]
}

func (pq *PriorityQueue[T]) up(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if pq.data[parent].CompareTo(pq.data[i]) <= 0 {
			break
		}
		pq.swap(parent, i)
		i = parent
	}
}

func (pq *PriorityQueue[T]) Enqueue(val T) {
	pq.data = append(pq.data, val)
	pq.up(len(pq.data) - 1)
}

func (pq *PriorityQueue[T]) down(i, n int) {
	for {
		left := 2*i + 1
		if left >= n {
			break
		}
		j := left
		if right := left + 1; right < n && pq.data[right].CompareTo(pq.data[left]) < 0 {
			j = right
		}
		if pq.data[i].CompareTo(pq.data[j]) <= 0 {
			break
		}
		pq.swap(i, j)
		i = j
	}
}

func (pq *PriorityQueue[T]) Dequeue() (T, bool) {
	if pq.IsEmpty() {
		var zero T
		return zero, false
	}
	n := len(pq.data) - 1
	pq.swap(0, n)
	pq.down(0, n)
	val := pq.data[n]
	pq.data = pq.data[:n]
	return val, true
}

func (pq *PriorityQueue[T]) Peek() (T, bool) {
	if pq.IsEmpty() {
		var zero T
		return zero, false
	}
	return pq.data[0], true
}
