package ds

import "testing"

func TestQueue_Enqueue(t *testing.T) {
	var q Queue
	q.Enqueue(10)
	q.Enqueue(20)

	if len(q.elements) != 2 {
		t.Errorf("expected queue size 2, got %d", len(q.elements))
	}
	if q.elements[0] != 10 || q.elements[1] != 20 {
		t.Errorf("expected elements [10, 20], got %v", q.elements)
	}
}

func TestQueue_Dequeue(t *testing.T) {
	var q Queue
	q.Enqueue(10)
	q.Enqueue(20)

	val, ok := q.Dequeue()
	if !ok || val != 10 {
		t.Errorf("expected Dequeue() to return 10, true, got %d, %v", val, ok)
	}
	if len(q.elements) != 1 {
		t.Errorf("expected queue size 1 after Dequeue, got %d", len(q.elements))
	}

	val, ok = q.Dequeue()
	if !ok || val != 20 {
		t.Errorf("expected Dequeue() to return 20, true, got %d, %v", val, ok)
	}

	val, ok = q.Dequeue()
	if ok {
		t.Errorf("expected Dequeue() to return false for an empty queue, got %v", ok)
	}
}
