package main

import (
	"fmt"
	"sync"
	"time"
)

func getData(c chan string) {
	c <- "data"
}

func getAnotherData(c chan string) {
	c <- "anotherData"
}

func basic() {
	myChannel := make(chan string)
	anotherChannel := make(chan string)

	go getData(myChannel)
	go getAnotherData(anotherChannel)

	select {
	case msg1 := <-myChannel:
		fmt.Println("myChannel:", msg1)
	case msg2 := <-anotherChannel:
		fmt.Println("anotherChannel:", msg2)
	}
}

// For Select loop
func forSelect() {
	charChannel := make(chan string, 3)
	chars := []string{"a", "b", "c"}

	for _, s := range chars {
		charChannel <- s
	}

	close(charChannel)

	for result := range charChannel {
		fmt.Println(result)
	}
}

func doWork(done <-chan bool) {

	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("Doing Work")
		}
	}
}

func doneChannel() {
	done := make(chan bool)

	go doWork(done)

	time.Sleep(time.Second * 3)

	close(done)
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Mark this worker as done when the function ends

	fmt.Printf("Worker %d: started\n", id)
	fmt.Printf("Worker %d: finished\n", id)
}

func testWaitGroup() {
	var wg sync.WaitGroup // Create a WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)         // Add one task to the WaitGroup
		go worker(i, &wg) // Start worker as a goroutine
	}

	wg.Wait() // Wait for all goroutines to call Done()
	fmt.Println("All workers completed")
}

func main() {
	// basic()
	forSelect()
	// doneChannel()
	// testWaitGroup()
}
