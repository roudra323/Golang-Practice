package main

import (
	"fmt"
	"sync"
	"time"
)

type jobs struct {
	id   int
	name string
}

func worker(jobChannel <-chan jobs, resultChannel chan<- jobs, wg *sync.WaitGroup) {

	defer wg.Done()

	for jobs := range jobChannel {
		fmt.Printf("Reading Job id %d\n", jobs.id)
		time.Sleep(time.Millisecond * 100)
		resultChannel <- jobs
	}
}

func main() {

	var wg sync.WaitGroup

	numTasks := 100
	numOfWorkers := 5

	jobChannel := make(chan jobs, numTasks)
	resultChannel := make(chan jobs, numTasks)

	for i := 1; i <= numOfWorkers; i++ {
		wg.Add(1)
		go func() {
			worker(jobChannel, resultChannel, &wg)
		}()
	}

	for i := 0; i < numTasks; i++ {
		job := jobs{
			id:   i,
			name: "worker",
		}
		jobChannel <- job
	}

	close(jobChannel)
	wg.Wait()

	close(resultChannel)

	/*
			    go func() {
		        wg.Wait()
		        close(resultChannel)
		    }()
	*/

	for res := range resultChannel {
		fmt.Println(res)
	}

}
