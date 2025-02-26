package main

import "fmt"

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	if len(stages) == 0 {
		return nil
	}
	safeStage := func(stage Stage, in In) Out {
		out := make(Bi)
		go func() {
			defer close(out)
			// обработка паники
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered from panic:", r)
				}
			}()

			for {
				select {
				case <-done:
					return
				case data, ok := <-in:
					if !ok {
						return
					}
					select {
					case <-done:
						return
					case out <- data:
					}
				}
			}
		}()
		return stage(out)
	}
	pipeline := in
	for _, stage := range stages {
		pipeline = safeStage(stage, pipeline)
	}
	return pipeline
}
