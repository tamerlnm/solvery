package main

import (
	"fmt"
)

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	for _, stage := range stages {
		in = stage(in)
	}
	return in
}

func generateNumbers(done In) In {
	out := make(Bi)
	go func() {
		defer close(out)
		for i := 1; i <= 5; i++ {
			select {
			case out <- i:
			case <-done:
				return
			}
		}
	}()
	return out
}

func squareNumbers(done In, in In) Out {
	out := make(Bi)
	go func() {
		defer close(out)
		for sq := range in {
			select {
			case out <- sq.(int) * sq.(int):
			case <-done:
				return
			}
		}
	}()
	return out
}

func doubleNumbers(done In, in In) Out {
	out := make(Bi)
	go func() {
		defer close(out)
		for db := range in {
			select {
			case out <- db.(int) * 2:
			case <-done:
				return
			}
		}
	}()
	return out
}

func sumNumbers(done In, in In) Out {
	out := make(Bi)
	go func() {
		defer close(out)
		sum := 0
		for s := range in {
			select {
			case <-done:
				return
			default:
				sum += s.(int)
			}
		}
		out <- sum
	}()
	return out
}

func main() {
	done := make(Bi)
	//close(done)

	out := ExecutePipeline(generateNumbers(done), done, func(in In) Out { return squareNumbers(done, in) }, func(in In) Out { return doubleNumbers(done, in) }, func(in In) Out { return sumNumbers(done, in) })

	for v := range out {
		fmt.Println("Sum of squares doubled:", v)
	}
}
