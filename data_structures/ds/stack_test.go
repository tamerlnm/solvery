package ds

import "testing"

func TestStack_Push(t *testing.T) {
	var s Stack
	s.Push(10)
	s.Push(20)

	if len(s.elements) != 2 {
		t.Errorf("expected stack size 2, got %d", len(s.elements))
	}
	if s.elements[0] != 10 || s.elements[1] != 20 {
		t.Errorf("expected elements [10, 20], got %v", s.elements)
	}
}

func TestStack_Pop(t *testing.T) {
	var s Stack
	s.Push(10)
	s.Push(20)

	val, ok := s.Pop()
	if !ok || val != 20 {
		t.Errorf("expected Pop() to return 20, true, got %d, %v", val, ok)
	}
	if len(s.elements) != 1 {
		t.Errorf("expected stack size 1 after Pop, got %d", len(s.elements))
	}

	val, ok = s.Pop()
	if !ok || val != 10 {
		t.Errorf("expected Pop() to return 10, true, got %d, %v", val, ok)
	}

	val, ok = s.Pop()
	if ok {
		t.Errorf("expected Pop() to return false for an empty stack, got %v", ok)
	}
}

func TestStack_Peek(t *testing.T) {
	var s Stack
	s.Push(10)
	s.Push(20)

	val, ok := s.Peek()
	if !ok || val != 20 {
		t.Errorf("expected Peek() to return 20, true, got %d, %v", val, ok)
	}

	s.Pop()
	val, ok = s.Peek()
	if !ok || val != 10 {
		t.Errorf("expected Peek() to return 10, true, got %d, %v", val, ok)
	}

	s.Pop()
	val, ok = s.Peek()
	if ok {
		t.Errorf("expected Peek() to return false for an empty stack, got %v", ok)
	}
}
