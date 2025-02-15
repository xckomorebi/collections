package collections_test

import (
	"testing"

	"github.com/xckomorebi/collections"
)

type Int int

func (i Int) CompareTo(o Int) int {
	return int(i) - int(o)
}

func TestPush(t *testing.T) {

	pq := collections.NewPriorityQueue[Int]()

	pq.Enqueue(Int(1))
	pq.Enqueue(Int(3))
	pq.Enqueue(Int(2))

	if pq.Size() != 3 {
		t.Errorf("Expected size 3, got %d", pq.Size())
	}

	val, _ := pq.Dequeue()
	if val != Int(1) {
		t.Errorf("Expected 1, got %d", val)
	}

	val, _ = pq.Dequeue()
	if val != Int(2) {
		t.Errorf("Expected 2, got %d", val)
	}

	val, _ = pq.Dequeue()
	if val != Int(3) {
		t.Errorf("Expected 3, got %d", val)
	}
}
