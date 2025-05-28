/*
Task 1: Basic Goroutines and Wait Groups
Concept: Understanding goroutines and synchronization with WaitGroups
Task: Create a program that launches 5 goroutines, each printing numbers 1-5 with their goroutine ID, and ensure the main function waits for all to complete.
*/

package main

import (
	"fmt"
	"sync"
)

func printNumber(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		fmt.Printf("ID: %d\t%d\n", id, i)
	}
}

func main() {
	fmt.Println("Task 3")

	var wg sync.WaitGroup

	for i := 1; i < 5; i++ {
		wg.Add(1)
		go printNumber(i, &wg)

	}

	wg.Wait()

}
