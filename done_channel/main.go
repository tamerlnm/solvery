package main

import (
	"fmt"
	"time"
)

var or func(channels ...<-chan interface{}) <-chan interface{} = func(channels ...<-chan interface{}) <-chan interface{} {
	out := make(chan interface{})
	for _, ch := range channels {
		go func(c <-chan interface{}) {
			<-c
			close(out)
		}(ch)
	}
	return out
}

func sig(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}

func main() {
	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("done after %v\n", time.Since(start))
}
