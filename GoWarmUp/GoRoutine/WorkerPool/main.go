package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID   int
	Data string
}

type Result struct {
	Job    Job
	Output string
	Error  error
}

func worker(jobsChan <-chan Job, resultChan chan<- Result, wg *sync.WaitGroup, id int) {
	defer wg.Done()

	for job := range jobsChan {
		fmt.Printf("Worker %d processing job %d\n", id, job.ID)
		time.Sleep(time.Second)

		resultChan <- Result{
			Job:    job,
			Output: fmt.Sprintf("Processed %s by worker %d", job.Data, id),
		}
	}
}

func main() {

	numWorkers := 3
	numJobs := 5

	jobsChan := make(chan Job, numJobs)
	result := make(chan Result, numJobs)

	wg := sync.WaitGroup{}

	// start wokers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(jobsChan, result, &wg, i)
	}

	// send jobs to channel
	for i := 0; i < numJobs; i++ {
		jobsChan <- Job{ID: i, Data: fmt.Sprintf("data-%d", i)}
	}

	close(jobsChan)

	// Close result channel when all workers finish
	go func() {
		wg.Wait()
		close(result)
	}()

	// Collect results
	for res := range result {
		if res.Error != nil {
			fmt.Printf("Error: %v\n", res.Error)
		} else {
			fmt.Printf("Job %d: %s\n", res.Job.ID, res.Output)
		}
	}

}
