package main

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if n <= 0 {
		return errors.New("number of goroutines must be greater than 0")
	}

	var wg sync.WaitGroup
	taskCh := make(chan Task, len(tasks))
	errCh := make(chan struct{}, m)
	done := make(chan struct{})

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case task, ok := <-taskCh:
					if !ok {
						return
					}
					if err := task(); err != nil {
						if len(errCh) < m {
							errCh <- struct{}{}
						}
					}
				case <-done:
					return
				}
			}
		}()
	}

	go func() {
		for _, task := range tasks {
			if len(errCh) >= m {
				close(done)
				break
			}
			taskCh <- task
		}
		close(taskCh)
	}()

	wg.Wait()

	if len(errCh) >= m {
		return ErrErrorsLimitExceeded
	}
	return nil
}
