package ds

type Queue struct {
	elements []int
}

func (q *Queue) Enqueue(element int) {
	q.elements = append(q.elements, element)
}

func (q *Queue) Dequeue() (int, bool) {
	if len(q.elements) == 0 {
		return 0, false
	}
	front := q.elements[0]
	q.elements = q.elements[1:]
	return front, true
}
