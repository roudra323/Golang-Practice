package main

import (
	"fmt"
	"sync"
)

// sayHi prints a greeting message.
// It uses a WaitGroup to signal its completion.
func sayHi(wg *sync.WaitGroup) {
	defer wg.Done() // Decrement counter when function returns
	fmt.Println("\n-------------Starting Hi Function----------------")
	fmt.Println("Hi Roudra!!")
	fmt.Println("-------------Ending Hi Function----------------\n")
}

// sayBye prints a farewell message.
// It uses a WaitGroup to signal its completion.
func sayBye(wg *sync.WaitGroup) {
	defer wg.Done() // Decrement counter when function returns
	fmt.Println("\n-------------Starting Bye Function----------------")
	fmt.Println("Bye Roudra ...")
	fmt.Println("-------------Ending Bye Function----------------\n")
}

func main() {
	// Declare a WaitGroup to manage goroutines.
	// A WaitGroup waits for a collection of goroutines to finish.
	var wg sync.WaitGroup

	// Increment the WaitGroup counter by 2, for two goroutines.
	// The main goroutine calls Add to set the number of goroutines to wait for.
	wg.Add(2)

	// Launch the sayHi function as a goroutine.
	// Goroutines are functions that run concurrently with other functions.
	go sayHi(&wg)
	// Launch the sayBye function as a goroutine.
	go sayBye(&wg)

	// Wait until the WaitGroup counter becomes 0.
	// This blocks the main goroutine until all goroutines registered with Add have called Done.
	wg.Wait()
	// This message is printed after all goroutines have completed.
	fmt.Println("All goroutines completed")
}
