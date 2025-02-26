package main

import (
	"testing"
)

func TestExecutePipeline(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

	in := make(chan interface{})
	go func() {
		for _, num := range []int{1, 2, 3} {
			in <- num
		}
		close(in)
	}()
	result := ExecutePipeline(in, done, Double, AddTen, Square)

	expectedResults := []int{(1*2 + 10) * (1*2 + 10), (2*2 + 10) * (2*2 + 10), (3*2 + 10) * (3*2 + 10)}
	index := 0
	for res := range result {
		if res != expectedResults[index] {
			t.Errorf("Expected %d, got %v", expectedResults[index], res)
		}
		index++
	}
}

func Double(in In) Out {
	out := make(Bi)
	go func() {
		defer close(out)
		for num := range in {
			out <- num.(int) * 2
		}
	}()
	return out
}

func AddTen(in In) Out {
	out := make(Bi)
	go func() {
		defer close(out)
		for num := range in {
			out <- num.(int) + 10
		}
	}()
	return out
}

func Square(in In) Out {
	out := make(Bi)
	go func() {
		defer close(out)
		for num := range in {
			out <- num.(int) * num.(int)
		}
	}()
	return out
}
