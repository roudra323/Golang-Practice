/*

Task 6: Select Statement and Non-blocking Operations
Concept: Using select for multiple channel operations and timeouts (building on all previous channel knowledge)
Task: Create a monitoring system where:

Multiple worker goroutines send status updates to different channels
A monitor goroutine uses select to handle updates from any worker
Implement a timeout mechanism - if no updates come within 2 seconds, print a timeout message
Allow graceful shutdown when all workers are done
Include a non-blocking channel operation to check system status

*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type WorkerStatus struct {
	WorkerID int
	Message  string
	JobCount int
}

func worker(id int, statusChan chan WorkerStatus, doneChan chan bool, wg *sync.WaitGroup) {
	defer wg.Done()

	jobCount := 0
	// Each worker does 3-7 jobs with random delays
	maxJobs := 3 + rand.Intn(5)

	for i := 0; i < maxJobs; i++ {
		// Simulate work with random delay
		workTime := time.Duration(rand.Intn(1000)) * time.Millisecond
		time.Sleep(workTime)

		jobCount++
		status := WorkerStatus{
			WorkerID: id,
			Message:  fmt.Sprintf("Completed job %d", i+1),
			JobCount: jobCount,
		}

		// Send status update
		statusChan <- status
	}

	// Signal completion
	finalStatus := WorkerStatus{
		WorkerID: id,
		Message:  "Worker finished",
		JobCount: jobCount,
	}
	statusChan <- finalStatus
	doneChan <- true
}

func monitor(statusChan chan WorkerStatus, doneChan chan bool, emergencyChan chan string, totalWorkers int) {
	workersFinished := 0

	for {
		select {
		case status := <-statusChan:
			fmt.Printf("[MONITOR] Worker %d: %s (Total jobs: %d)\n",
				status.WorkerID, status.Message, status.JobCount)

		case <-doneChan:
			workersFinished++
			fmt.Printf("[MONITOR] Worker finished (%d/%d completed)\n", workersFinished, totalWorkers)

			if workersFinished == totalWorkers {
				fmt.Println("[MONITOR] All workers completed!")
				return
			}

		case emergency := <-emergencyChan:
			fmt.Printf("[MONITOR] 🚨 EMERGENCY: %s\n", emergency)

		case <-time.After(2 * time.Second):
			fmt.Println("[MONITOR] ⏰ Timeout: No updates for 2 seconds")

			// Non-blocking check for emergency messages
			select {
			case emergency := <-emergencyChan:
				fmt.Printf("[MONITOR] Found pending emergency: %s\n", emergency)
			default:
				fmt.Println("[MONITOR] No pending emergencies")
			}
		}
	}
}

func emergencyAlertSystem(emergencyChan chan string) {
	time.Sleep(5 * time.Second) // Wait 5 seconds then send alert

	select {
	case emergencyChan <- "System overload detected!":
		fmt.Println("[EMERGENCY] Alert sent successfully")
	case <-time.After(100 * time.Millisecond):
		fmt.Println("[EMERGENCY] Alert channel busy, couldn't send")
	}
}

func main() {

	statusChan := make(chan WorkerStatus, 10) // Buffered for status updates
	doneChan := make(chan bool, 5)            // Buffered for completion signals
	emergencyChan := make(chan string, 1)     // Buffered for emergency alerts

	var wg sync.WaitGroup
	numWorkers := 3

	// Start workers
	fmt.Println("Starting workers...")
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, statusChan, doneChan, &wg)
	}

	// Start emergency alert system
	go emergencyAlertSystem(emergencyChan)

	// Start monitor in a separate goroutine
	go monitor(statusChan, doneChan, emergencyChan, numWorkers)

	// Wait for all workers to complete
	wg.Wait()

	fmt.Println("All systems shut down gracefully!")
}
