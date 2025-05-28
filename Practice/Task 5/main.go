/*
Task 5: Buffered Channels and Channel Closing
Concept: Using buffered channels and proper channel closing (building on Task 2's channel knowledge)
Task: Create a producer-consumer system where:

2 producer goroutines generate jobs (numbers) and send them to a buffered channel
2 consumer goroutines receive jobs from the channel and "process" them
Producers should finish first, then close the channel
Consumers should stop gracefully when the channel is closed
Use buffered channels to allow producers to send multiple jobs without waiting
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(jobsChan chan<- int, wg *sync.WaitGroup, id int) {
	defer wg.Done()

	for i := 1; i <= 3; i++ {
		job := id*10 + i
		fmt.Printf("Producer %d: Creating job %d\n", id, job)
		jobsChan <- job
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Printf("Producer %d: Finished\n", id)

}

func consumers(jobsChan <-chan int, wg *sync.WaitGroup, id int) {
	defer wg.Done()

	for job := range jobsChan { // range automatically handles channel closing
		fmt.Printf("Consumer %d: Processing job %d\n", id, job)
		time.Sleep(200 * time.Millisecond) // Simulate work
	}
	fmt.Printf("Consumer %d: No more jobs, exiting\n", id)
}

func main() {
	fmt.Println("Task 5")

	var jobWg sync.WaitGroup
	var conWg sync.WaitGroup

	jobsChan := make(chan int, 5)

	for i := 0; i < 2; i++ {
		jobWg.Add(1)
		go producer(jobsChan, &jobWg, i)
	}

	for i := 0; i < 2; i++ {
		conWg.Add(1)
		go consumers(jobsChan, &conWg, i)
	}

	go func() {
		jobWg.Wait()
		close(jobsChan)
		fmt.Println("All producers done, channel closed")
	}()

	conWg.Wait()
	fmt.Println("All work completed!")

}
