package model

import "golang.org/x/exp/maps"

type Set struct {
	elements map[string]struct{}
}

func NewSet() *Set {
	return &Set{elements: make(map[string]struct{})}
}

// Add добавление элемента в множестве
func (s *Set) Add(key string) {
	s.elements[key] = struct{}{}
}

// Remove удаление элемента в множестве
func (s *Set) Remove(key string) {
	delete(s.elements, key)
}

// Contains проверка, есть ли элемент в множестве
func (s *Set) Contains(element string) bool {
	_, exists := s.elements[element]
	return exists
}

// Union возвращает объединение двух множеств
func (s *Set) Union(another *Set) *Set {
	unionSet := NewSet()
	for elem := range s.elements {
		unionSet.Add(elem)
	}
	for elem := range another.elements {
		unionSet.Add(elem)
	}
	return unionSet
}

// Subtraction возвращает разность двух множеств
func (s *Set) Subtraction(another *Set) *Set {
	subtractionSet := NewSet()
	for elem := range s.elements {
		if !another.Contains(elem) {
			subtractionSet.Add(elem)
		}
	}

	return subtractionSet
}

// Intersection возвращает пересечение двух множеств
func (s *Set) Intersection(another *Set) *Set {
	intersectionSet := NewSet()
	for elem := range s.elements {
		if another.Contains(elem) {
			intersectionSet.Add(elem)
		}
	}
	return intersectionSet
}
func (s *Set) ToSlice() []string {
	if len(s.elements) == 0 {
		return []string{}
	}

	slice := maps.Keys(s.elements)

	if len(slice) == 0 {
		return []string{}
	}
	return slice
}
