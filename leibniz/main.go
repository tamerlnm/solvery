package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	var sum float64
	n := flag.Int("n", 2, "help message for flag n")
	flag.Parse()
	wg := sync.WaitGroup{}
	results := make(chan float64, *n)
	stop := make(chan struct{})
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	wg.Add(*n)
	for i := 0; i < *n; i++ {
		go LeibnizFormula(i, 10000, *n, &wg, results, stop)
	}
	go func() {
		<-sigChan
		close(stop)
	}()

	wg.Wait()
	close(results)
	for r := range results {
		sum += r
	}
	fmt.Println(sum)
}

func LeibnizFormula(id, terms, n int, wg *sync.WaitGroup, results chan float64, stop chan struct{}) {
	defer wg.Done()
	sum := 0.0
	sign := 1.0

	if id%2 != 0 {
		sign = -1.0
	}

	//for i := id; i < terms; i += n {
	//	sum += sign / float64(2*i+1)
	//	sign = -sign
	//}

	for i := id; i < terms; i += n {
		select {
		case <-stop:
			results <- sum * 4
			return
		default:
			sum += sign / float64(2*i+1)
			sign = -sign
		}
	}
	results <- sum * 4
}
