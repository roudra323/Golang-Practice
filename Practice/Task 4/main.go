/*

Task 2: Channels - Basic Communication
Concept: Using channels for goroutine communication (building on Task 1's goroutine knowledge)
Task: Modify the previous program to use channels instead of WaitGroups. Workers send their results through a channel.

*/

package main

import (
	"fmt"
	"time"
)

func worker(id int, resultChan chan string) {
	for i := 1; i <= 5; i++ {
		message := fmt.Sprintf("Goroutine %d: %d", id, i)
		resultChan <- message // Send to channel
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	// Create a channel for string messages
	resultChan := make(chan string) // Unbuffered channel
	numWorkers := 5

	// Launch workers
	for i := 1; i <= numWorkers; i++ {
		go worker(i, resultChan)
	}

	// Receive messages from channel
	// Each worker sends 5 messages, so we expect 25 total
	for i := 0; i < numWorkers*5; i++ {
		message := <-resultChan // Receive from channel
		fmt.Println(message)
	}

	fmt.Println("All messages received!")
}
