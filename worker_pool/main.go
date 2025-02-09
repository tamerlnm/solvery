package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	numJobs    = 10
	numWorkers = 5
	maxErrors  = 3
)

func main() {
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	var mu sync.Mutex
	errorCount := 0

	var wg sync.WaitGroup

	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &errorCount, &mu, &wg)
	}

	for j := 1; j <= numJobs; j++ {
		mu.Lock()
		if errorCount >= maxErrors {
			mu.Unlock()
			fmt.Println("Error limit exceeded")
			break
		}
		mu.Unlock()
		jobs <- j
	}
	close(jobs)

	wg.Wait()
	close(results)

	fmt.Println("Processing finished")
}

func worker(id int, jobs <-chan int, results chan<- int, errorCount *int, mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		mu.Lock()
		if *errorCount >= maxErrors {
			mu.Unlock()
			return
		}
		mu.Unlock()

		//Симулируем случайную ошибку
		//if j%2 == 0 {
		//	fmt.Println("worker", id, "error occurred on job", j)
		//	mu.Lock()
		//	*errorCount++
		//	mu.Unlock()
		//	continue
		//}

		fmt.Println("worker", id, "finished job", j)
		time.Sleep(time.Second)
		results <- j
	}
}
