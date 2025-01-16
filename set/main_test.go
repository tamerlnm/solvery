package main

import (
	"testing"
)

func TestUnion(t *testing.T) {
	set1 := NewSet()
	set2 := NewSet()

	set1.Add("a")
	set1.Add("b")
	set1.Add("c")

	set2.Add("b")
	set2.Add("c")
	set2.Add("e")

	expected := []string{"a", "b", "c", "e"}
	result := set1.Union(set2).ToSlice()

	if !compare(result, expected) {
		t.Errorf("Union failed. Expected %v, got %v", expected, result)
	}
}

func TestSubtraction(t *testing.T) {
	set1 := NewSet()
	set2 := NewSet()

	set1.Add("a")
	set1.Add("b")
	set1.Add("c")

	set2.Add("b")
	set2.Add("c")
	set2.Add("e")

	expected := []string{"a"}
	result := set1.Subtraction(set2).ToSlice()

	if !compare(result, expected) {
		t.Errorf("Subtraction failed. Expected %v, got %v", expected, result)
	}
}

func TestIntersection(t *testing.T) {
	set1 := NewSet()
	set2 := NewSet()

	set1.Add("a")
	set1.Add("b")
	set1.Add("c")

	set2.Add("b")
	set2.Add("c")
	set2.Add("e")

	expected := []string{"b", "c"}
	result := set1.Intersection(set2).ToSlice()

	if !compare(result, expected) {
		t.Errorf("Intersection failed. Expected %v, got %v", expected, result)
	}
}

// compare функция для сравнения срезов без учёта порядка
func compare(slice1, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	m := make(map[string]int)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		if m[v] == 0 {
			return false
		}
		m[v]--
	}

	for _, count := range m {
		if count != 0 {
			return false
		}
	}
	return true
}
