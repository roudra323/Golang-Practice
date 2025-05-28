package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Millisecond * 500)
	fmt.Printf("Worker %d finished\n", id)
}

// Function to demonstrate counting
func counter(name string, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 5; i++ {
		fmt.Printf("%s: %d\n", name, i)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	fmt.Println("Learning Goroutine")

	fmt.Println("Starting cocurrent workers....")

	var wg sync.WaitGroup

	// Launch 3 worker goroutines

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	// Wait for the workers to finish the work
	wg.Wait()

	// Launch 2 counter goroutines
	wg.Add(2)
	go counter("Counter-A", &wg)
	go counter("Counter-B", &wg)

	// Wait for the counters to finish counting
	wg.Wait()

}
