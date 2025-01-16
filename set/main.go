package main

import "fmt"

type Set struct {
	elements map[string]int
}

func NewSet() *Set {
	return &Set{elements: make(map[string]int)}
}

// Add добавление элемента в множестве
func (s *Set) Add(key string) {
	s.elements[key] = 1
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
	slice := make([]string, 0, len(s.elements))
	for elem := range s.elements {
		slice = append(slice, elem)
	}
	if len(slice) == 0 {
		return []string{"пустое множество"}
	}
	return slice
}

func main() {
	set1 := NewSet()
	set2 := NewSet()

	set1.Add("a")
	set1.Add("b")
	set1.Add("c")

	set2.Add("b")
	set2.Add("c")
	set2.Add("e")

	unionSet := set1.Union(set2)
	fmt.Println("Объединение:", unionSet.ToSlice())

	subtractionSet := set1.Subtraction(set2)
	fmt.Println("Разность:", subtractionSet.ToSlice())

	intersectionSet := set1.Intersection(set2)
	fmt.Println("Пересечение:", intersectionSet.ToSlice())

}
