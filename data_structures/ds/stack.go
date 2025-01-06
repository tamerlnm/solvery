package ds

type Stack struct {
	elements []int
}

func (s *Stack) Push(element int) {
	s.elements = append(s.elements, element)
}

func (s *Stack) Pop() (int, bool) {
	if len(s.elements) == 0 {
		return 0, false
	}
	top := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return top, true
}

func (s *Stack) Peek() (int, bool) {
	if len(s.elements) == 0 {
		return 0, false
	}
	return s.elements[len(s.elements)-1], true
}
