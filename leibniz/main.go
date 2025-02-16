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
	n := flag.Int("n", 3, "help message for flag n")
	flag.Parse()
	wg := sync.WaitGroup{}
	results := make(chan float64, *n)
	stop := make(chan struct{})
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		fmt.Println("\nProgram has stopped...")
		close(stop)
	}()

	wg.Add(*n)
	for i := 0; i < *n; i++ {
		go LeibnizFormula(i, *n, &wg, results, stop)
	}

	wg.Wait()
	close(results)

	for r := range results {
		sum += r
	}
	fmt.Println(sum)
}

func LeibnizFormula(id, n int, wg *sync.WaitGroup, results chan float64, stop chan struct{}) {
	defer wg.Done()
	sum := 0.0

	for i := id; ; i += n {
		select {
		case <-stop:
			results <- sum * 4
			return
		default:
			sign := 1.0
			if i%2 != 0 {
				sign = -1.0
			}
			sum += sign / float64(2*i+1)
		}
	}
}
